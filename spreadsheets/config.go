package spreadsheets

import (
	configApp "github.com/jasiu001/maestro/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const CONFIG_FILE = "config.yaml"

type Config struct {
	Spreadsheet Spreadsheet
}

type Spreadsheet struct {
	Id    string `yaml:"id"`
	Date  string `yaml:"date"`
	Scope string `yaml:"range"`
}

func InitConfig() *Config {
	yamlFile, err := ioutil.ReadFile(configApp.GetFile(CONFIG_FILE))
	if err != nil {
		log.Fatalf("Unable to read config file: %v", err)
	}

	config := Config{}
	err = yaml.Unmarshal([]byte(yamlFile), &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &config
}
