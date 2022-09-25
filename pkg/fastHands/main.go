package fastHands

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Data struct {
	FastHands map[string]map[string]interface{} `yaml:"FastHands"`
	MetaData  map[string]string                 `yaml:"MetaData"`
	Weights   map[string]map[string]interface{} `yaml:"Weights"`
}

func Run() Data {
	var data Data
	yamlFile, err := ioutil.ReadFile("/etc/configs/drills.yaml")
	if err != nil {
		log.Printf("Error reading yaml file %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		log.Fatalf("Cant unmarshal yaml %v", err)
	}

	return data

}
