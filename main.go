package main

import (

	// "net"

	"os"

	// "path/filepath"
	// "strings"

	"github.com/leangeder/chatops/lib/configuration"
	"github.com/leangeder/chatops/plugins"

	// for test
	"flag"
	//"plugin"
)

func main() {
	configFile := flag.String("config", "chatops.yml", "Path of the configuration file")
	flag.Parse()

	if *configFile == "" {
		return
	}

	config, err := configuration.NewConfiguration(*configFile)
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	plugins.New(config)

	<-make(chan struct{})
	return

	// c := loadConfig(*configFile)

	// // go func() {
	// // 	plugins.New(c)
	// // }()

	// fmt.Println("Error ", c["username"])
	// fmt.Println("Error ", c["slack"])
	//	plugins.New(config)

	// var wg sync.WaitGroup
	// for {
	// 	wg.Add(1)
	// 	go func() {
	// 		bot.New(c).Process()
	// 	}()
	// }
	// wg.Wait()
}

// func loadConfig(path string) map[string]interface{} {
//
// 	f, err := os.Open(path)
// 	defer f.Close()
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	var c map[string]interface{}
//
// 	buff, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	yaml.Unmarshal(buff, &c)
//
// 	return c
// }

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
