package csv2objdef

import (
	"fmt"
	"strings"
)

func AttrFormat(indent int, lgcl string, dtype string, col string) string {
	var result string
	space := strings.Repeat(" ", indent)
	result += space + "// " + lgcl + ";\n"
	result += space + "private " + dtype + " " + ToCamelCase(col) + ";\n"
	return result
}

func TblFormat(indent int, classFormat string, setting *Setting, name string, records *CsvRecords) string {
	var result string

	h := setting.Header
	for _, r := range *records {
		result += AttrFormat(indent, r[h.Logical], r[h.Dtype], r[h.Column]) + "\n"
	}
	return fmt.Sprintf(classFormat, name, Singular(ToUpperCamelCase(name)), result)
}
