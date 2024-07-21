//nolint:lll,gocognit // Code generator, verbosity is expected here.
package generator

import (
	"fmt"
	"io"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/hoomy-official/go-sonos/specs"
	"github.com/iancoleman/strcase"
)

//nolint:gochecknoinits // generator, this is an easy and enough way
func init() {
	strcase.ConfigureAcronym("IRRepeater", "IRRepeater")
}

func GenerateTypes(w io.Writer, api specs.APISpec) error {
	packageName := api.Model
	serviceFile := jen.NewFile(packageName)

	serviceFile.HeaderComment("CODE GENERATED, DO NOT EDIT")
	serviceFile.PackageComment(fmt.Sprintf("Package %s contains the implementation for the %s (Model %s) services.", packageName, api.ModelDescription, api.Model))

	serviceFile.
		Comment("Services struct contains service clients for different functionalities of a Sonos device.")

	variables := map[string]specs.StateVariable{}

	for _, service := range api.Services {
		for _, variable := range service.StateVariables {
			variables[sanitizeTypeName(variable.Name)] = variable
		}
	}

	serviceFile.Do(func(statement *jen.Statement) {
		for _, variable := range variables {
			statement.
				Do(func(statement *jen.Statement) {
					if variable.AllowedValues != nil {
						statement.Const().DefsFunc(func(group *jen.Group) {
							if allowedValues, ok := variable.AllowedValues.([]interface{}); ok {
								for _, value := range allowedValues {
									switch value.(type) {
									case string:
										group.Id(fmt.Sprintf("%s%s", strcase.ToCamel(fmt.Sprintf("%s", value)), sanitizeTypeName(variable.Name))).Id(sanitizeTypeName(variable.Name)).Op("=").Lit(value)
									default:
										panic(value)
									}
								}
							}
						}).Line().Line()
					}
				}).
				Type().
				Id(sanitizeTypeName(variable.Name)).
				Do(func(statement *jen.Statement) {
					switch variable.DataType {
					case "string":
						statement.String()
					case "boolean":
						statement.Bool()
					case "ui2", "ui4":
						statement.Uint()
					case "i2", "i4":
						statement.Int()
					default:
						panic(variable.DataType)
					}
				}).
				Line().
				Line()
		}
	})

	return serviceFile.Render(w)
}

func sanitizeTypeName(name string) string {
	name = strings.ReplaceAll(name, "A_ARG_TYPE_", "")
	name = strings.ReplaceAll(name, "Next", "")
	return strings.ReplaceAll(name, "Current", "")
}
