package main

import (
	"csv2objdef/pkg/csv2objdef"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) <= 2 {
		log.Fatalf("Usage: %s path/to/csvRecords path/to/format.txt\n", os.Args[0])
	}

	csvPath := os.Args[1]
	formatPath := os.Args[2]
	csvData, err := csv2objdef.ReadCsv(csvPath, 1)
	if err != nil {
		log.Fatal(err)
	}

	clsFormat, err := csv2objdef.ReadTxtFile(&formatPath)

	setting, err := csv2objdef.ReadSetting("config.yml")
	if err != nil {
		log.Fatalf("Read config.yml error")
	}

	fmt.Println(setting)

	_ = csv2objdef.CreateDir(setting.Result.Dir)
	attrs := createAttrs(&csvData.Records, &setting)
	dtypeMap := csv2objdef.MakeDtypeMap(&setting)
	attrs = replaceDtypes(&attrs, dtypeMap)
	tblMap := csv2objdef.GenTblMap(&attrs)

	for _, def := range tblMap {
		outputPath := createFilePath(def.Name, &setting)
		err = csv2objdef.WriteTxtFile(outputPath, def.AttrFormat(4, clsFormat))
		if err != nil {
			_ = fmt.Errorf("output error. file = %s\n", outputPath)
		} else {
			fmt.Println(outputPath)
		}
	}
}

func createFilePath(baseName string, setting *csv2objdef.Setting) string {
	outputPath := setting.Result.Prefix + csv2objdef.ToUpperCamelCase(csv2objdef.Singular(baseName)) + setting.Result.Suffix
	outputPath = filepath.Join(setting.Result.Dir, outputPath)
	return outputPath
}

func createAttrs(data *csv2objdef.CsvRecords, setting *csv2objdef.Setting) []csv2objdef.TblAttr {
	attrs := csv2objdef.ConvTblAttr(data,
		setting.Header.Table,
		setting.Header.Column,
		setting.Header.Logical,
		setting.Header.Dtype)
	return attrs
}

func replaceDtypes(attrs *[]csv2objdef.TblAttr, dtypeMap csv2objdef.DtypeMap) []csv2objdef.TblAttr {
	var newAttar []csv2objdef.TblAttr
	for _, attr := range *attrs {
		s, ok := dtypeMap[strings.TrimSpace(attr.Dtype)]
		if ok {
			attr.Dtype = s
		}
		newAttar = append(newAttar, attr)
	}
	return newAttar
}
