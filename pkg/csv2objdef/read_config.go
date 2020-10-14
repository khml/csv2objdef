package csv2objdef

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Dtype struct {
	Type  string
	Stype string
}

type Setting struct {
	Dtypes []Dtype
	Header struct {
		Table   int
		Column  int
		Logical int
		Dtype   int
	}
}

func ReadSetting(filename string) (Setting, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return Setting{}, err
	}

	return readSettingFromYaml(buf)
}

func readSettingFromYaml(fileBuffer []byte) (Setting, error) {
	setting := Setting{}

	err := yaml.Unmarshal(fileBuffer, &setting)
	if err != nil {
		fmt.Println(err)
		return Setting{}, err
	}
	return setting, nil
}

type DtypeMap = map[string]string

func MakeDtypeMap(setting *Setting) DtypeMap {
	var dtypeMap = make(DtypeMap)
	for _, dtype := range setting.Dtypes {
		dtypeMap[dtype.Type] = dtype.Stype
	}
	return dtypeMap
}
