package main

import (
	"flag"

	"github.com/leangeder/chatops/lib/plugins"
)

func main() {
	configFile := flag.String("config", "chatops.yml", "Path of the configuration file")
	flag.Parse()

	if *configFile == "" {
		return
	}

	plugins.Load(*configFile)

	<-make(chan struct{})
	return
}
