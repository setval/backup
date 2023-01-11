package config

import (
	"gopkg.in/yaml.v3"
)

const (
	DatabaseTypeMySQL = "mysql"
)

type Database struct {
	Name     string   `yaml:"name"`
	Login    string   `yaml:"login"`
	Password string   `yaml:"password"`
	Host     string   `yaml:"host"`
	Port     int      `yaml:"port"`
	Driver   string   `yaml:"driver"`
	Tables   []string `yaml:"tables"`
}

func parseDatabase(nodes yaml.Node) ([]Database, error) {
	var databases []Database
	for _, content := range nodes.Content {
		var db Database
		if err := content.Decode(&db); err != nil {
			return nil, err
		}
		databases = append(databases, db)
	}
	return databases, nil
}
