package config

import (
	"fmt"
	"os"

	"github.com/go-yaml/yaml"
)

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

type application_config struct {
	Environment string `yaml:"environment"`
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
}

type database_config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Config struct {
	Application application_config `yaml:"application"`
	Database    database_config    `yaml:"database"`
}

func NewDefault() *Config {
	return &Config{
		Application: application_config{
			Environment: "local",
			Host:        "localhost",
			Port:        8080,
		},
		Database: database_config{
			Host:     "localhost",
			Port:     5432,
			Username: "postgres",
			Password: "postgres",
			Name:     "postgres",
		},
	}
}

func (c *Config) LoadFromYaml(path string) error {
	const op = "config.Config.LoadFromYaml"

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("%s: %v", op, err)
	}
	defer file.Close()

	if err := yaml.NewDecoder(file).Decode(c); err != nil {
		return fmt.Errorf("%s: %v", op, err)
	}

	return nil
}
