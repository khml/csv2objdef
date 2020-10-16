package csv2objdef

import (
	"encoding/csv"
	"io/ioutil"
	"os"
)

func ReadTxtFile(txtPath *string) (string, error) {
	bytes, err := ioutil.ReadFile(*txtPath)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func WriteTxtFile(txtPath string, content string) error {
	file, err := os.Create(txtPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

type CsvHeader = []string
type CsvRecords = [][]string

type CsvData struct {
	Header  CsvHeader
	Records CsvRecords
}

func ReadCsv(filepath string, skipRowNum int) (CsvData, error) {
	var records CsvRecords

	f, err := os.Open(filepath)
	if err != nil {
		return CsvData{}, err
	}
	defer f.Close()

	reader := csv.NewReader(f)

	// skip rows
	for i := skipRowNum; i > 0; i-- {
		_, err := reader.Read()
		if err != nil {
			return CsvData{}, err
		}
	}

	csvHeader, err := reader.Read()
	if err != nil {
		return CsvData{}, err
	}

	for {
		r, err := reader.Read()
		if err != nil {
			break
		}
		records = append(records, r)
	}
	return CsvData{Header: csvHeader, Records: records}, nil
}

func CreateDir(dirname string) error {
	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		err = os.Mkdir(dirname, 0777)
		return err
	}
	return nil
}
