package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ApiPort                     string `yaml:"api_port"`
	ContainerInitServiceAddress string `yaml:"container_init_service_address"`
}

func LoadConfig(filePath string) (*Config, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var conf Config
	if err := yaml.Unmarshal(data, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
