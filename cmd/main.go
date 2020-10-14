package main

import (
	"csv2objdef/pkg/csv2objdef"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) <= 2 {
		log.Fatalf("Usage: %s path/to/csv path/to/format.txt\n", os.Args[0])
	}

	csvPath := os.Args[1]
	formatPath := os.Args[2]
	data, err := csv2objdef.ReadCsv(&csvPath, 1)
	if err != nil {
		log.Fatal(err)
	}

	const resultDir = "results"
	csv2objdef.CreateDir(resultDir)

	clsFormat, err := csv2objdef.ReadTxtFile(&formatPath)

	setting, err := csv2objdef.ReadSetting("config.yml")
	if err != nil {
		log.Fatalf("Read config.yml error")
	}

	fmt.Println(setting)

	attrs := csv2objdef.ConvTblAttr(data,
		setting.Header.Table,
		setting.Header.Column,
		setting.Header.Logical,
		setting.Header.Dtype)

	dtypeMap := csv2objdef.MakeDtypeMap(&setting)
	tblMap := csv2objdef.GenTblMap(attrs, &dtypeMap)

	for _, def := range tblMap {
		outputPath := csv2objdef.ToUpperCamelCase(csv2objdef.Singular(def.Name)) + "Dto.java"
		outputPath = filepath.Join(resultDir, outputPath)
		fmt.Println(outputPath)
		_ = csv2objdef.WriteTxtFile(outputPath, def.AttrFormat(4, clsFormat))
	}
}
