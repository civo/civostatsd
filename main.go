package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/civo/civostatsd/config"
	"github.com/civo/civostatsd/gather"
	"github.com/civo/civostatsd/send"
)

func main() {
	var configFile string
	var testMode bool

	flag.StringVar(&configFile, "config", "/etc/civostatsd", "Set the location of the configuration file")
	flag.BoolVar(&testMode, "test", false, "run a single iteration as a test")
	flag.Parse()

	configuration := config.Load(configFile)

	ticker := time.NewTicker(60 * time.Second)
	death := make(chan os.Signal, 1)
	signal.Notify(death, os.Interrupt, os.Kill)

	fmt.Println("Civostatsd v2.0")
	fmt.Println("Using API server: " + configuration.Server)
	fmt.Println("Using Token:      " + configuration.Token)

	stats := gather.All()
	send.ToAPI(configuration, stats)

	if testMode {
		fmt.Println(stats.String())
		os.Exit(0)
	}

	for {
		select {
		case <-ticker.C:
			stats := gather.All()
			fmt.Println(stats.String())
			send.ToAPI(configuration, stats)
		case <-death:
			os.Exit(0)
		}
	}
}
