package configuration

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// type Configuration interface {
// 	Pseudo  string `yaml:"pseudo"`
// 	Plugins PluginConfiguration
// 	Plugins []PlugingConfiguration
// 	// Chats  []struct {
// 	// 	Name  string `yaml:"name"`
// 	// 	Token string `yaml:"token"`
// 	// 	Url   string `yaml:"url"`
// 	// } `yaml:"chats"`
// }

type Configuration struct {
	Pseudo  string `yaml:"pseudo"`
	Plugins PluginConfiguration
	// Chats  []struct {
	// 	Name  string `yaml:"name"`
	// 	Token string `yaml:"token"`
	// 	Url   string `yaml:"url"`
	// } `yaml:"chats"`
}

type PluginConfiguration struct {
	Name    string ""
	Options map[string]interface{}
}

func NewConfiguration(configurationPath string) (*Configuration, error) {

	bytes, err := ioutil.ReadFile(configurationPath)
	if err != nil {
		return nil, err
	}

	c := Configuration{}
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		log.Printf("Failed to validate syntax: %s", configurationPath)
		return nil, err
	}

	return &c, nil
}
