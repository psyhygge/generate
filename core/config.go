package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Database struct {
		Type string `yaml:"type"`
		DSN  string `yaml:"dsn"`
	} `yaml:"database"`
	Output struct {
		Language        string   `yaml:"language"`
		ModelsDir       string   `yaml:"models_dir"`
		NamingStyle     string   `yaml:"naming_style"`
		FileNamingStyle string   `yaml:"file_naming_style"`
		Tables          []string `yaml:"tables"`
	} `yaml:"output"`
}

func LoadConfig() *Config {
	configFile, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}

	var config Config
	if err := yaml.Unmarshal(configFile, &config); err != nil {
		panic(fmt.Errorf("failed to parse config: %w", err))
	}
	return &config
}
