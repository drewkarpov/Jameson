package config

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
		Host     string
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
	dbConfig.Database.Host = os.Getenv("HOST")
	var currentConfig = Config{Database: dbConfig.Database}

	return currentConfig
}
