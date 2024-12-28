package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Postgres struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Url      string `yaml:"url"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
}

type HTTP struct {
	Port string `yaml:"port"`
}

type Config struct {
	Postgres `yaml:"postgres"`
	HTTP     `yaml:"http"`
}

func Load(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	d := yaml.NewDecoder(file)
	if err = d.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
