package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Service struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"service"`
}

func GetConfig(configPath string) Config {
	file, err := os.Open(configPath)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Can't close file config.yaml: %s\n", err)
		}
	}(file)
	if err != nil {
		log.Fatalf("Can't open file config.yaml: %s\n", err)
	}

	var serviceConfig Config
	if err = yaml.NewDecoder(file).Decode(&serviceConfig); err != nil {
		log.Fatalf("Can't decode config.yaml: %s\n", err)
	}
	return serviceConfig
}
