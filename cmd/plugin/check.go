package plugin

import "log"

// Checks plugin status
func checkPlugin(plugVal interface{}) {
	if plugVal == nil {
		log.Fatal("Plugin loaded but there is no value!")
	} else {
		plugInfo := plugVal.(PlugGeneral).Info()
		log.Printf("Plugin \"%v\" loaded and ready to use with version \"%v\" ", plugInfo["name"], plugInfo["version"])
	}
}

func checkExist() {

}
