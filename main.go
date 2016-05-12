package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"git.absolutedevops.io/internal/civostatsd/gather"
	"git.absolutedevops.io/internal/civostatsd/send"
)

var (
	Version   string
	BuildTime string
)

func main() {
	ticker := time.NewTicker(5 * time.Second)
	death := make(chan os.Signal, 1)
	signal.Notify(death, os.Interrupt, os.Kill)

	fmt.Println("Civostatsd -- v" + Version + " (" + BuildTime + ")")

	boolPtr := flag.Bool("test", false, "run a single iteration as a test")
	flag.Parse()

	stats := gather.All()
	send.ToAPI(stats)

	if *boolPtr {
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
