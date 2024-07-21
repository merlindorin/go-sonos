//nolint:lll // Code generator, verbosity is expected here.
package generator

import (
	"fmt"
	"io"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/hoomy-official/go-sonos/specs"
)

func GenerateServices(w io.Writer, api specs.APISpec, doc specs.Documentation) error {
	packageName := api.Model
	serviceOutputFile := jen.NewFile(packageName)

	serviceOutputFile.HeaderComment("CODE GENERATED, DO NOT EDIT")
	serviceOutputFile.PackageComment(fmt.Sprintf("Package %s contains the implementation for the %s (Model %s) services.", packageName, api.ModelDescription, api.Model))

	serviceOutputFile.
		Comment("Services struct contains service clients for different functionalities of a Sonos device.").
		Line().
		Type().
		Id("Services").
		StructFunc(func(group *jen.Group) {
			for _, service := range api.Services {
				s := doc.Service(service.ServiceName)

				if s != nil {
					group.Add(
						jen.Comment(s.Description),
					)
				}

				group.Add(
					jen.Id(service.Name).Op("*").Id(service.ServiceName),
				)

				if s != nil {
					group.Add(jen.Line())
				}
			}
		}).
		Line().
		Comment("NewService creates and returns a new Services struct initialized with base URLs for Sonos services available.").
		Line().
		Func().Id("NewService").Params(jen.Id("doer").Qual("github.com/hoomy-official/go-shared/pkg/net/do", "Doer")).Op("*").Id("Services").Block(
		jen.Return(
			jen.Op("&").Id("Services").Values(
				jen.DictFunc(
					func(dict jen.Dict) {
						for _, service := range api.Services {
							dict[jen.Id(service.Name)] = jen.Id(fmt.Sprintf("New%s", service.ServiceName)).Call(jen.Id("doer"))
						}
					},
				),
			),
		)).
		Line().
		Comment("Model returns the model identifier of the Sonos device.").
		Line().
		Func().Params(jen.Id("s").Id("Services")).Id(fmt.Sprintf("%s%s", strings.ToUpper("model"[:1]), "model"[1:])).Params().String().Block(jen.Return(jen.Lit(api.Model))).
		Line().
		Comment("ModelDescription returns a human-readable description of the Sonos model.").
		Line().
		Func().Params(jen.Id("s").Id("Services")).Id(fmt.Sprintf("%s%s", strings.ToUpper("modelDescription"[:1]), "modelDescription"[1:])).Params().String().Block(jen.Return(jen.Lit(api.ModelDescription))).
		Line().
		Comment("SoftwareGeneration returns the generation of the software the services are running on.").
		Line().
		Func().Params(jen.Id("s").Id("Services")).Id(fmt.Sprintf("%s%s", strings.ToUpper("softwareGeneration"[:1]), "softwareGeneration"[1:])).Params().Int().Block(jen.Return(jen.Lit(api.SoftwareGeneration))).
		Line().
		Comment("SoftwareVersion returns the current software version installed on the Sonos device.").
		Line().
		Func().Params(jen.Id("s").Id("Services")).Id(fmt.Sprintf("%s%s", strings.ToUpper("softwareVersion"[:1]), "softwareVersion"[1:])).Params().String().Block(jen.Return(jen.Lit(api.SoftwareVersion))).
		Line().
		Comment("DiscoveryDate returns a timestamp (in Unix epoch time) representing when the services were discovered.").
		Line().
		Func().Params(jen.Id("s").Id("Services")).Id(fmt.Sprintf("%s%s", strings.ToUpper("discoveryDate"[:1]), "discoveryDate"[1:])).Params().Int().Block(jen.Return(jen.Lit(int(api.DiscoveryDate.Unix()))))

	return serviceOutputFile.Render(w)
}
