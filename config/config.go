package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server string
	Token  string
}

func Load(configFile string) Config {
	_, err := os.Stat(configFile)
	if err != nil {
		log.Fatal("Config file is missing: ", configFile)
	}

	var config Config
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal(err)
	}
	return config
}
