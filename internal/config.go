package internal

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type ServerConfig struct {
	Port int32
	Name string
}

type UserConfig map[string]interface{}

type SysConfig struct {
	Server *ServerConfig
	Config UserConfig
}

func NewSysConfig() *SysConfig {
	return &SysConfig{Server: &ServerConfig{Port: 8080, Name: "tgin"}}
}

func InitConfig() *SysConfig {
	config := NewSysConfig()
	if b := LoadConfigFile(); b != nil {
		err := yaml.Unmarshal(b, config)
		if err != nil {
			panic(err)
		}
	}
	return config
}

func LoadConfigFile() []byte {
	dir, _ := os.Getwd()
	file := dir + "/application.yaml"
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return b
}
