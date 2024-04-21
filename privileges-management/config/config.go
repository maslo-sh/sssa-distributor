package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

const ConfigFile = "config.yaml"

type Config struct {
	DomainController struct {
		Port              int    `yaml:"port"`
		Host              string `yaml:"host"`
		BaseDN            string `yaml:"base-dn"`
		AccessCredentials struct {
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		} `yaml:"access-credentials"`
	} `yaml:"domain-controller"`

	Kafka struct {
		Port        int    `yaml:"port"`
		Host        string `yaml:"host"`
		WriterTopic string `yaml:"writer-topic"`
	} `yaml:"kafka"`
}

func GetConfig() Config {
	f, err := os.Open(ConfigFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
