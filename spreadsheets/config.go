package spreadsheets

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path"
	"runtime"
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
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("Runtime cannot read current file")
	}
	dir := path.Join(path.Dir(filename))

	yamlFile, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", dir, CONFIG_FILE))
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
