package server

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

var ConfInfo *Conf

func init() {
	ConfInfo, _ = ReadYaml()
}

func ReadYaml() (*Conf, error) {
	buf, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}
	var conf Conf
	err = yaml.Unmarshal(buf, &conf)
	if err != nil {
		fmt.Println("读取yaml出错！", err.Error())
		return nil, err
	}

	return &conf, nil
}

type Conf struct {
	Pass string `yaml:"pass"`
	Port string `yaml:"port"`
}
