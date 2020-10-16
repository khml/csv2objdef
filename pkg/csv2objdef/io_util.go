package csv2objdef

import (
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

func CreateDir(dirname string) error {
	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		err = os.Mkdir(dirname, 0777)
		return err
	}
	return nil
}
