package csv2objdef

import (
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

var plu = pluralize.NewClient()

func SnakeToCamelCase(str string) string {
	return strcase.ToLowerCamel(str)
}

func Plural(str string) string {
	return plu.Plural(str)
}
