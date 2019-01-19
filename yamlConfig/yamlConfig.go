package yamlConfig

import (
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

func ReadConfig(file_name string , type_defined interface{}) {
	yamlFile, err := ioutil.ReadFile(file_name)
	failOnError(err, "配置文件读取失败")
	err = yaml.Unmarshal(yamlFile, type_defined)
	return
}

func failOnError(err error, msg string)  {
	if err != nil {
		log.Fatalf("%s: %s",msg, err)
	}
}
