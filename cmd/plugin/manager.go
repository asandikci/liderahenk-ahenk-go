package plugin // pluginmanager

import (
	"encoding/json"
	"log"

	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/utils"
)

// General Functions/Methods that each plugin has
type PlugGeneral interface {
	// Get Plugin Info
	Info() map[string]interface{}
}

type PlugUnloaded interface {
	error
}

// plugins/resources
type ResourcesI interface {
	// Get Agent Info, see plugins/resources for more information
	AgentInfo() (map[string]interface{}, error) // LINK plugins/resources/main.go#AgentInfo

	// Get Resource Usage, see plugins/resources for more information
	ResourceUsage() (map[string]interface{}, error) // LINK plugins/resources/main.go#ResourceUsage
}

// plugins/tmptest
type TmpTestI interface {
	// Run temporary tests
	TmpTest() error // LINK plugins/tmptest/main.go#TmpTest
}

// FILLME[epic=new-plugin-template]
// type NewPluginInterfaceI interface {
// 	// explanation
// 	PluginFunction() returnType // link plugins/pluginname/file.go#PluginFunction
// }

var Tmptest TmpTestI
var Resources ResourcesI

// FILLME[epic=new-plugin-template]
// var pluginname NewPluginInterfaceI

// Connect plugins to main code
//
// Gets plugin names as string slice
func ConnectPlugins(params []string) {
	connectPlugins(params)
}

func UnloadPlugins(params []string) {
	unloadPlugins(params)
}

// Logs plugin outputs.
func LogPlugin(title string, mp map[string]interface{}, toJson bool) {
	log.Printf("\n----- %v -----\n", title)
	if toJson {
		data, err := json.MarshalIndent(&mp, "", " ")
		utils.Check(err)
		log.Println(string(data))
	} else {
		for i, v := range mp {
			log.Printf("%v: %v\n", i, v)
		}
	}
}

// // TODO response to Lider
// func createResponse() {

// }
