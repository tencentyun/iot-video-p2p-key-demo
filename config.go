package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config service config
type Config struct {
	Redis struct {
		Addr   string `yaml:"addr"`
		Pass   string `yaml:"pass"`
		Expire int64  `yaml:"expire"`
	} `yaml:"redis"`
	Service struct {
		Addr string `yaml:"addr"`
	} `yaml:"service"`
}

// NewConfig returns a Config instance
func NewConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
