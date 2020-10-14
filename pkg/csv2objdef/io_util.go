package csv2objdef

import (
	"encoding/csv"
	"io/ioutil"
	"log"
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

func ReadCsv(filepath *string, skipRowNum int) ([][]string, error) {
	var records [][]string

	f, err := os.Open(*filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)

	for {
		l, err := reader.Read()
		if skipRowNum > 0 {
			skipRowNum--
			continue
		}

		if err != nil {
			break
		}
		records = append(records, l)
	}
	return records, nil
}

func CreateDir(dirname string) {
	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		err = os.Mkdir(dirname, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}
}
