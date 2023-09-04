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

// plugins/resources
type ResourcesI interface {
	// Get Agent Info, see plugins/resources for more information
	AgentInfo() map[string]interface{} // LINK plugins/resources/main.go#AgentInfo

	// Get Resource Usage, see plugins/resources for more information
	ResourceUsage() map[string]interface{} // LINK plugins/resources/main.go#ResourceUsage
}

// plugins/tmptest
type TmpTestI interface {
	// Run temporary tests
	TmpTest() // LINK plugins/tmptest/main.go#TmpTest
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

	conf := getJITConf()
	chanPlug := make(chan interface{})

	for _, v := range params {
		go LoadPlugin(v, conf, chanPlug)
		plug, ok := <-chanPlug

		switch v {
		case "tmptest":
			Tmptest = plug.(TmpTestI)
			checkPlugin(Tmptest, ok)
		case "resources":
			Resources = plug.(ResourcesI)
			checkPlugin(Resources, ok)

			// FILLME[epic=new-plugin-template]
			// case "pluginname":
			//   Pluginname = plug.(NewPluginInterfaceI)
			// 	 checkPlugin(Pluginname, ok)
		}
	}
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

// Checks plugin status
func checkPlugin(plugVal interface{}, status bool) {
	if status {
		if plugVal == nil {
			log.Fatal("Plugin loaded but there is no value!")
		} else {
			plugInfo := plugVal.(PlugGeneral).Info()
			log.Printf("Plugin \"%v\" loaded and ready to use with version \"%v\" ", plugInfo["name"], plugInfo["version"])
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
// func createResponse() {

// }
