package main

import (
	"fmt"
	"log"
	"time"
)

// plugins/resources
type Resources interface {
	AgentInfo() map[string]interface{}
}

// FILLME creating new plugin interface template
// type NewPluginInterface interface {
// 	PluginMethod() returnType
// }

// Loads Plugins and runs them concurrently.
// When you create a new plugin create a new interface and call this plugin in this function
func PluginManager() {
	chanPlug := make(chan interface{})

	go LoadPlugin("resources", chanPlug)
	res, ok := <-chanPlug
	var resources Resources = res.(Resources)
	checkPlugin(res, ok)

	// FILLME Loading new plugin template
	// go LoadPlugin("pluginName", chanPlug)
	// res, ok = <-chanPlug
	// var pluginName NewPluginInterface = res.(NewPluginInterface)
	// checkPlugin(res, ok)

	// Run plugins concurrently and log out
	for {
		go logPlugin("AgentInfo", resources.AgentInfo())

		// FILLME Running/Logout a plugin template
		// go logPlugin("InfoAboutFunction", pluginName.Function() )

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

// Checks plugin status
func checkPlugin(plugVal interface{}, status bool) {
	if status {
		if plugVal == nil {
			log.Fatal("Plugin loaded but there is no value!")
		} else {
			log.Println("Plugin loaded and ready to use")
		}
	} else {
		if plugVal == nil {
			log.Fatal("Plugin closed and there is no value! ")
		} else {
			log.Fatal("Plugin closed or there is an error")
		}
	}
}

// // TODO response to Lider
// // func createResponse() {

// // }
