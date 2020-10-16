package csv2objdef

import (
	"fmt"
	"strings"
)

type TblAttr struct {
	Table string
	Col   string
	Lgcl  string
	Dtype string
}

func (t *TblAttr) FormatAttr(dtypeMap DtypeMap) {
	t.Table = Plural(t.Table)
	t.Col = ToCamelCase(t.Col)

	s, ok := dtypeMap[strings.TrimSpace(t.Dtype)]
	if ok {
		t.Dtype = s
	}
}

func (t *TblAttr) AttrFormat(indent int) string {
	var result string
	space := strings.Repeat(" ", indent)
	result += space + "// " + t.Lgcl + ";\n"
	result += space + "private " + t.Dtype + " " + ToCamelCase(t.Col) + ";\n"
	return result
}

type TblDef struct {
	Name string
	Attr []TblAttr
}

func (t *TblDef) Append(a *TblAttr) {
	t.Attr = append(t.Attr, *a)
}

func (t TblDef) Show() {
	for _, attr := range t.Attr {
		fmt.Println(attr)
	}
}

func (t TblDef) AttrFormat(indent int, classFormat string) string {
	var result string
	for _, attr := range t.Attr {
		result += attr.AttrFormat(indent) + "\n"
	}
	return fmt.Sprintf(classFormat, t.Name, Singular(ToUpperCamelCase(t.Name)), result)
}

func ConvTblAttr(csvRecords *CsvRecords, tbl int, col int, logical int, dtype int) []TblAttr {
	var attrs []TblAttr
	for _, r := range *csvRecords {
		attrs = append(attrs,
			TblAttr{Table: r[tbl], Col: r[col], Lgcl: r[logical], Dtype: r[dtype]},
		)
	}
	return attrs
}

type TblMap = map[string]TblDef

func GenTblMap(tblAttrs *[]TblAttr) TblMap {
	var tableMap = make(map[string]*TblDef)

	for _, attr := range *tblAttrs {
		_, ok := tableMap[attr.Table]
		if !ok {
			tableMap[attr.Table] = &TblDef{Name: attr.Table}
		}
		def := tableMap[attr.Table]
		def.Append(&attr)
	}

	var tblMap = make(map[string]TblDef)
	for key, def := range tableMap {
		tblMap[key] = *def
	}

	return tblMap
}
