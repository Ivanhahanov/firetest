package fire

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Target  Target    `yaml:"target"`
	Actions []*Action `yaml:"actions"`
}

type Target struct {
	Type    string `yaml:"type"`
	Address string `yaml:"address"`
	Users   int    `yaml:"users"`
}

func (c *Config) Parce(filepath string) {
	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
