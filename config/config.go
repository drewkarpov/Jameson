package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
		Host     string `yaml:"host"`
		DbName   string `yaml:"db_name"`
	} `yaml:"database"`
}

func InitConfig(filename string) Config {
	var path = "./config/" + filename + ".yaml"
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	var dbConfig Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&dbConfig)
	if err != nil {
		logrus.Fatalf("cannot read config file from path " + path)
	}
	var currentConfig = Config{Database: dbConfig.Database}

	return currentConfig
}
