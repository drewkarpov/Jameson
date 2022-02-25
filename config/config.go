package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type Config struct {
	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
		Host     string
		DbName   string
	} `yaml:"database"`
}

func InitConfig(filename string) Config {
	var dbConfig Config

	if os.Getenv("MONGO_CREDENTIALS") != "" {
		credentials := strings.Split(os.Getenv("MONGO_CREDENTIALS"), ":")
		dbConfig.Database.Username = credentials[0]
		dbConfig.Database.Password = credentials[1]
	} else {
		var path = "./config/" + filename + ".yaml"
		f, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		decoder := yaml.NewDecoder(f)
		err = decoder.Decode(&dbConfig)
		if err != nil {
			logrus.Fatalf("cannot read config file from path " + path)
		}

	}

	dbConfig.Database.Host = "127.0.0.1"
	dbConfig.Database.DbName = "jameson"
	var currentConfig = Config{Database: dbConfig.Database}

	return currentConfig
}
