package config

import (
	"bufio"
	"go/types"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type ServerProperties struct {
	Bind           string `cfg:"bind"`
	Port           int    `cfg:"port"`
	AppendOnly     bool   `cfg:"appendonly"`
	AppendFilename string `cfg:"appendFilename"`
	MaxClients     string `cfg:"maxClients"`
	Databases      string `cfg:"databases"`
	Peers          string `cfg:"peers"`
	Self           string `cfg:"self"`
}

var Properties *ServerProperties

var a types.Slice

func init() {
	//默认值
	Properties = &ServerProperties{
		Bind:       "127.0.0.1",
		Port:       6379,
		AppendOnly: false,
	}
}

func parseConfig(src io.Reader) *ServerProperties {
	config := &ServerProperties{}

	rawMap := make(map[string]string)
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := scanner.Text()
		//跳过无效和注释行
		if len(line) > 0 && line[0] == '#' {
			continue
		}
		index := strings.Index(line, " ")
		if 0 < index && index < len(line)-1 {
			key := line[0:index]
			value := strings.TrimSpace(line[index+1:])
			rawMap[strings.ToLower(key)] = value
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	t := reflect.TypeOf(config)
	v := reflect.ValueOf(config)
	for i := 0; i < t.Elem().NumField(); i++ {
		field := t.Elem().Field(i)
		fieldVal := v.Elem().Field(i)
		key, ok := field.Tag.Lookup("cfg")
		if !ok {
			key = field.Name
		}
		value, ok := rawMap[strings.ToLower(key)]
		if ok {
			//结构体字段必须.Type.Kind()
			switch field.Type.Kind() {
			case reflect.String:
				fieldVal.SetString(value)
			case reflect.Int:
				intV, err := strconv.ParseInt(value, 10, 64)
				if err == nil {
					fieldVal.SetInt(intV)
				}
			case reflect.Bool:
				boolV := value == "yes"
				fieldVal.SetBool(boolV)
			case reflect.Slice:
				if field.Type.Elem().Kind() == reflect.String {
					sliceV := strings.Split(value, ",")
					fieldVal.Set(reflect.ValueOf(sliceV))
				}
			}

		}
	}
	return config

}

func SetupConfig(configFilename string) {
	file, err := os.Open(configFilename)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	Properties = parseConfig(file)
}
