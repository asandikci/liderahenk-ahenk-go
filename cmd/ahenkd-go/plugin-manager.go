package main

import (
	"ahenk-go/pkg/plugins/resources"
	"fmt"
	"log"
	"time"
)

func PluginManager() {
	log.Print("Plugin Manager Started")

	for {
		logPlugin(resources.AgentInfo())
		time.Sleep(30 * time.Second)
	}

}

func logPlugin(mp map[string]interface{}) {
	fmt.Printf("\nOs Info:\n")
	for i, v := range mp {
		fmt.Printf("%v: %v\n", i, v)
	}
}

// TODO response to Lider
// func createResponse() {

// }
