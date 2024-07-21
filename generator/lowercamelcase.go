package generator

import (
	"github.com/iancoleman/strcase"
)

func lowerCamelCase(s string) string {
	if s == "type" || s == "Type" {
		s = "t"
	}

	return strcase.ToLowerCamel(s)
}
