package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DB       string `yaml:"db"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"database"`
}

func NewConfig() *Config {
	return &Config{}
}

func MustLoad() *Config {
	config := NewConfig()

	file, err := os.Open("/Users/muradkadymov/Documents/freshtrack/internal/config/config.yaml")
	if err != nil {
		log.Fatalf("error loading config %v: ", err)
	}

	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		log.Fatalf("error loading config %v: ", err)
	}

	return config
}
