package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Storage       Storage    `yaml:"storage"`
	DatabasesNode yaml.Node  `yaml:"databases"`
	Databases     []Database `yaml:"-"`
	Key           string     `yaml:"key"`
}

func Parse(pathBase ...string) (*Config, error) {
	path := "config.yml"
	if pathBase != nil {
		path = pathBase[0]
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	if err = yaml.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	if err = parseStorage(&config.Storage); err != nil {
		return nil, err
	}

	databases, err := parseDatabase(config.DatabasesNode)
	if err != nil {
		return nil, err
	}
	config.Databases = databases

	return &config, nil
}
