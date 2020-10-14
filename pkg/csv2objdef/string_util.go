package csv2objdef

import (
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

var plu = pluralize.NewClient()

func ToCamelCase(str string) string {
	return strcase.ToLowerCamel(str)
}

func ToUpperCamelCase(str string) string {
	return strcase.ToCamel(str)
}

func Plural(str string) string {
	return plu.Plural(str)
}

func Singular(str string) string {
	return plu.Singular(str)
}
