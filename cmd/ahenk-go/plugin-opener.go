package main

import (
	"fmt"
	"os"
	"plugin"

	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/utils"
)

// Load Plugin placed in PluginDir, returns empty interface.
// Do not forget to cast in plugin manager
//
// Give Plugin Name as argument and be sure you compiled plugins with `-buildmode=plugin` to PluginDir as `pluginname.so`
func LoadPlugin(plugName string) interface{} {

	// TODO if error caugth try without relative path, this will be good for local testing
	plug, err := plugin.Open(PluginDir + plugName + ".so")
	utils.Check(err)

	// TODO also allow lookup another symbol other than PlugnameConnect
	symPlug, err := plug.Lookup(utils.FirstUpperEN(plugName) + "Connect")
	utils.Check(err)

	var plugOut interface{}
	plugOut, ok := symPlug.(interface{})
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}
	return plugOut
}
