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

	_ = csv2objdef.CreateDir(setting.Result.Dir)

	dtypeMap := csv2objdef.MakeDtypeMap(&setting)
	err = csvData.ReplaceDtype(setting.Header.Dtype, &dtypeMap)
	if err != nil {
		log.Fatal(err)
	}

	tblMap, err := csvData.ToTableMap(setting.Header.Table)
	if err != nil {
		log.Fatal(err)
	}

	for name, def := range tblMap {
		outputPath := createFilePath(name, &setting)
		err = csv2objdef.WriteTxtFile(outputPath, csv2objdef.TblFormat(4, clsFormat, &setting, name, &def))
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
