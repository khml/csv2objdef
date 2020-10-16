package csv2objdef

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CsvHeader = []string
type CsvRecords = [][]string
type TableMap = map[string]CsvRecords

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

func (c *CsvData) ToTableMap(keyIdx int) (TableMap, error) {
	if keyIdx >= len(c.Header) {
		return TableMap{}, fmt.Errorf("keyIdx = %d is Out of range", keyIdx)
	}

	var tableMap = make(map[string]CsvRecords)

	for _, record := range c.Records {
		key := record[keyIdx]
		_, ok := tableMap[key]
		if !ok {
			tableMap[key] = CsvRecords{}
		}
		tableMap[key] = append(tableMap[key], record)
	}

	return tableMap, nil
}
