package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func Read(configPath string) (c Config) {
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		fmt.Printf("Error parsing config file: %s\n", err)
		panic(err)
	}
	return
}

type Config struct {
	Server struct {
		Port         string `yaml:"port"`
		ReadTimeout  int    `yaml:"readTimeout"`
		WriteTimeout int    `yaml:"writeTimeout"`
	} `yaml:"server"`
	PSQL struct {
		DSN string `yaml:"dsn"`
	} `yaml:"psql"`
	Bucket struct {
		IPLimit       int `yaml:"ipLimit"`
		LoginLimit    int `yaml:"loginLimit"`
		PasswordLimit int `yaml:"passwordLimit"`
		BucketTTL     int `yaml:"bucketTtl"`
	} `yaml:"bucket"`
}
