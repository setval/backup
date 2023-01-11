package config

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

const (
	TypeStorageLocal = "local"
	TypeStorageS3    = "s3"
)

type Storage struct {
	Type        string    `yaml:"type"`
	Auth        yaml.Node `yaml:"auth"`
	S3Config    *S3Config
	LocalConfig *LocalConfig
}

type S3Config struct {
	Endpoint  string `yaml:"endpoint"`
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
	SSL       bool   `yaml:"ssl"`
}

type LocalConfig struct {
	Directory string `yaml:"directory"`
}

func parseStorage(s *Storage) error {
	switch s.Type {
	case TypeStorageLocal:
		var cfg LocalConfig
		if err := s.Auth.Decode(&cfg); err != nil {
			return err
		}
		s.LocalConfig = &cfg
	case TypeStorageS3:
		var cfg S3Config
		if err := s.Auth.Decode(&cfg); err != nil {
			return err
		}
		s.S3Config = &cfg
	default:
		return fmt.Errorf("unknown storage type: %s", s.Type)
	}
	return nil
}
