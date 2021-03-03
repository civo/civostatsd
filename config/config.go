package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

//Config is the struct to save the config from the file
type Config struct {
	Server string
	Token  string
	Region string
}

// Load is the function to load the config and put into the struct
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
