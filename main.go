package main

import (
	"fmt"
	// "log"
	// "net"
	"os"
	// "path/filepath"
	// "strings"
	"sync"

	"github.com/leangeder/chatops/lib/bot"

	yaml "gopkg.in/yaml.v2"


	// for test
	"flag"
	//"plugin"
	"io/ioutil"
)

type Plugin interface {

}

func main() {
	configFile := flag.String("config", "chatops.yml", "Path of the configuration file")
	flag.Parse()

	c := loadConfig(*configFile)

	fmt.Println("Error ", c["username"])

	var wg sync.WaitGroup
	for {
		wg.Add(1)
		go func() {
			bot.New(c).Process()
		}()
	}
	wg.Wait()
}


func loadConfig(path string) map[string]interface{} {

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	var c map[string]interface{}

	buff, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	yaml.Unmarshal(buff, &c)

	return c
}

// func loadPlugins() {
//
// 	for currentPlugin in (cat plugins)
// 		plugin, err := plugin.Open(currentPlugin)
// 		if err != nil {
// 			panic(fmt.Println("Unable to load plugin %s: %s", currentPlugin, err)
// 		}
//
// 		config, err := plugin.Lookup("loadConfig")
// 		if err != nil {
// 			panic(fmt.Println("Unable to load configuration for the plugin %s: %s", config, err))
// 		}
// 	}
// }
