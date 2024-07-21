//nolint:lll,gocognit // Code generator, verbosity is expected here.
package generator

import (
	"fmt"
	"io"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/hoomy-official/go-sonos/specs"
)

func GenerateService(w io.Writer, api specs.APISpec, service specs.Service, d specs.Documentation) error {
	packageName := api.Model
	file := jen.NewFile(packageName)

	file.HeaderComment("CODE GENERATED, DO NOT EDIT")

	file.
		Const().
		Defs(
			jen.Comment(fmt.Sprintf("namespace%s is the UPnP namespace for the Sonos service.", service.ServiceName)).
				Line().
				Id(fmt.Sprintf("namespace%s", service.ServiceName)).Op("=").Lit(service.ServiceType),

			jen.Line().Comment(fmt.Sprintf("event%s is the event path for the Sonos event service.", service.ServiceName)).
				Line().
				Id(fmt.Sprintf("event%s", service.ServiceName)).Op("=").Lit(service.EventSubURL),

			jen.Line().Comment(fmt.Sprintf("control%s is the action path for the Sonos service.", service.ServiceName)).
				Line().
				Id(fmt.Sprintf("control%s", service.ServiceName)).Op("=").Lit(service.ControlURL),
		).
		Line().
		Comment(fmt.Sprintf("%s is based on UPnP and wraps a SOAP client for communicating with Sonos services.", service.ServiceName)).
		Line().
		Type().
		Id(service.ServiceName).
		Struct(
			jen.Id("doer").Qual("github.com/hoomy-official/go-shared/pkg/net/do", "Doer"),
		).
		Line().
		Comment(fmt.Sprintf("New%s initializes a new %s with the provided Doer.", service.ServiceName, service.ServiceName)).
		Line().
		Func().
		Id(fmt.Sprintf("New%s", service.ServiceName)).
		Params(
			jen.Id("doer").Qual("github.com/hoomy-official/go-shared/pkg/net/do", "Doer"),
		).
		Op("*").Id(service.ServiceName).
		Block(
			jen.Return(
				jen.Op("&").
					Id(service.ServiceName).
					Values(
						jen.Dict{
							jen.Id("doer"): jen.Id("doer"),
						},
					),
			),
		).
		Line().
		Comment("SubscribeEvent subscribes for events from the Sonos Event service using the provided callback URL.").
		Line().
		Func().
		Params(jen.Id(strings.ToLower(service.ServiceName[:1])).Id(service.ServiceName)).
		Id("SubscribeEvent").
		Params(
			jen.Id("ctx").Qual("context", "Context"),
			jen.Id("callbackURL").Op("*").Qual("net/url", "URL"),
		).
		Error().
		Block(
			jen.Return(
				jen.Id(strings.ToLower(service.ServiceName[:1])).Dot("doer").Dot("Do").Call(
					jen.Id("ctx"),
					jen.Id("do").Dot("WithPath").Call(jen.Id(fmt.Sprintf("event%s", service.ServiceName))),
					jen.Id("do").Dot("WithMethod").Call(jen.Lit("SUBSCRIBE")),
					jen.Id("do").Dot("WithExtraHeaderf").Call(jen.Lit("NT"), jen.Lit("upnp:event")),
					jen.Id("do").Dot("WithExtraHeaderf").Call(jen.Lit("CALLBACK"), jen.Qual("fmt", "Sprintf").Call(jen.Lit("<%s>"), jen.Id("callbackURL").Dot("String").Call())),
					jen.Id("do").Dot("WithExtraHeaderf").Call(jen.Lit("TIMEOUT"), jen.Lit("Second-1800")),
				),
			),
		).
		Line().
		Line().
		Do(servideDef(service, d))

	return file.Render(w)
}

func servideDef(service specs.Service, doc specs.Documentation) func(statement *jen.Statement) {
	return func(statement *jen.Statement) {
		for _, action := range service.Actions {
			statement.Line().
				Do(ServiceResType(action, service)).
				Do(func(statement *jen.Statement) {
					if s := doc.Service(service.ServiceName); s != nil {
						if a := s.Action(action.Name); a != nil {
							statement.Comment(fmt.Sprintf("%s %s", action.Name, a.Description))
						}
					}
				}).Line().
				Func().
				Params(jen.Id(strings.ToLower(service.ServiceName[:1])).Id(service.ServiceName)).
				Id(action.Name).
				ParamsFunc(func(g *jen.Group) {
					g.Id("ctx").Qual("context", "Context")

					for _, input := range action.Inputs {
						var s specs.StateVariable
						service.StateVariables.Get(input.RelatedStateVariableName, &s)

						if input.Name == "InstanceID" {
							continue
						}

						g.Id(lowerCamelCase(input.Name)).Id(sanitizeTypeName(input.RelatedStateVariableName))
					}

					if action.Outputs != nil && len(action.Outputs) > 0 {
						g.Id(lowerCamelCase(action.Name)).Op("*").Id(fmt.Sprintf("%s%sRes", service.Name, action.Name))
					}
				}).
				Error().
				Block(
					jen.Type().Id(fmt.Sprintf("%sReq", action.Name)).StructFunc(func(g *jen.Group) {
						for _, input := range action.Inputs {
							var s specs.StateVariable
							service.StateVariables.Get(input.RelatedStateVariableName, &s)

							g.Id(input.Name).Id(sanitizeTypeName(input.RelatedStateVariableName))
						}
					}),
					jen.Line(),
					jen.Type().Id(fmt.Sprintf("%sNode", action.Name)).Struct(jen.Id("Xmlns").String().Tag(map[string]string{"xml": "xmlns:u,attr"}), jen.Id(fmt.Sprintf("%sReq", action.Name))),
					jen.Line(),
					jen.Type().Id(fmt.Sprintf("%sBody", action.Name)).Struct(
						jen.Id("XMLName").Qual("encoding/xml", "Name").Tag(map[string]string{"xml": "s:Body"}),
						jen.Id(fmt.Sprintf("%sNode", action.Name)).Id(fmt.Sprintf("%sNode", action.Name)).Tag(map[string]string{"xml": fmt.Sprintf("u:%s,omitempty", action.Name)}),
					),
					jen.Line(),
					jen.Id("body").Op(":=").Id(fmt.Sprintf("%sBody", action.Name)).Values(
						jen.Dict{
							jen.Id(fmt.Sprintf("%sNode", action.Name)): jen.Id(fmt.Sprintf("%sNode", action.Name)).Values(
								jen.Dict{
									jen.Id("Xmlns"): jen.Id(fmt.Sprintf("namespace%s", service.ServiceName)),
									jen.Id(fmt.Sprintf("%sReq", action.Name)): jen.Id(fmt.Sprintf("%sReq", action.Name)).Values(jen.DictFunc(func(dict jen.Dict) {
										for _, input := range action.Inputs {
											if input.Name == "InstanceID" {
												dict[jen.Id(input.Name)] = jen.Lit(0)
												continue
											}

											var s specs.StateVariable
											service.StateVariables.Get(input.RelatedStateVariableName, &s)
											dict[jen.Id(input.Name)] = jen.Id(lowerCamelCase(input.Name))
										}
									})),
								},
							)},
					),
					jen.Line(),
					jen.Return(
						jen.Id(strings.ToLower(service.ServiceName[:1])).
							Dot("doer").
							Dot("Do").
							Call(
								jen.Line().
									List(
										jen.Id("ctx"),
										jen.Qual("github.com/hoomy-official/go-sonos/opts", "WithSonosUPnPOptions").
											CallFunc(func(group *jen.Group) {
												group.Id(fmt.Sprintf("control%s", service.ServiceName))
												group.Id(fmt.Sprintf("namespace%s", service.ServiceName))
												group.Lit(action.Name)
												group.Id("body")

												if action.Outputs == nil || len(action.Outputs) == 0 {
													group.Nil()
												} else {
													group.Id(lowerCamelCase(action.Name))
												}
											},
											),
									).Op("..."),
							),
					),
				).Line()
		}
	}
}

func ServiceResType(action specs.Actions, service specs.Service) func(statement *jen.Statement) {
	return func(statement *jen.Statement) {
		if action.Outputs == nil || len(action.Outputs) == 0 {
			return
		}
		statement.
			Type().
			Id(fmt.Sprintf("%s%sRes", service.Name, action.Name)).
			StructFunc(func(g *jen.Group) {
				for _, output := range action.Outputs {
					g.Id(output.Name).Id(sanitizeTypeName(output.RelatedStateVariableName))
				}
			}).
			Line()
	}
}
