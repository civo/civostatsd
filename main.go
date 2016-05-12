package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/absolutedevops/civostatsd/gather"
	"github.com/absolutedevops/civostatsd/send"
)

var (
	Version   string
	BuildTime string
)

type Config struct {
	Server string
	Token  string
}

func main() {
	var configFile string
	var testMode bool

	flag.StringVar(&configFile, "config", "/etc/civostatsd", "Set the location of the configuration file")
	flag.BoolVar(&testMode, "test", false, "run a single iteration as a test")
	flag.Parse()

	fmt.Println(configFile)
	_, err := os.Stat(configFile)
	if err != nil {
		log.Fatal("Config file is missing: ", configFile)
	}

	var config Config
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(5 * time.Second)
	death := make(chan os.Signal, 1)
	signal.Notify(death, os.Interrupt, os.Kill)

	fmt.Println("Civostatsd")
	fmt.Println("Version/Build   : v" + Version + " (" + BuildTime + ")")
	fmt.Println("Using API server: " + config.Server)
	fmt.Println("Using Token:      " + config.Token)

	stats := gather.All()
	send.ToAPI(stats)

	if testMode {
		fmt.Println(stats.String())
		os.Exit(0)
	}

	for {
		select {
		case <-ticker.C:
			stats := gather.All()
			fmt.Println(stats.String())
			send.ToAPI(stats)
		case <-death:
			os.Exit(0)
		}
	}
}
