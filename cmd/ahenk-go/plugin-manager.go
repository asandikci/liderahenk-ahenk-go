package main

import (
	"fmt"
	"time"
)

// plugins/resources
type Resources interface {
	AgentInfo() map[string]interface{}
}

// Loads Plugins and runs them.
// When you create a new plugin create a new interface and call this plugin in this function
func PluginManager() {
	var resources Resources = LoadPlugin("resources").(Resources)
	for {
		logPlugin("AgentInfo", resources.AgentInfo())
		time.Sleep(30 * time.Second)
	}
}

// Logs plugin outputs.
func logPlugin(title string, mp map[string]interface{}) {
	fmt.Printf("\n----- %v -----\n", title)
	for i, v := range mp {
		fmt.Printf("%v: %v\n", i, v)
	}
}

// // TODO response to Lider
// // func createResponse() {

// // }
