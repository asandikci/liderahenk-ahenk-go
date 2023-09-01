package pluginmanager

import "fmt"

// General Functions/Methods that each plugin has
type PlugGeneral interface {
	Info() (interface{}, error)
}

// plugins/resources
type Resources interface {
	AgentInfo() (interface{}, error)
	ResourceUsage() (interface{}, error)
}

type TmpTest interface {
	TmpTest() (interface{}, error)
}

type anny any

// Start All plugins
func StartPlugins(params map[string][]string) {
	// GeneralFunctions := []string{"Info"}

	conf := getJITConf()
	chanPlug := make(chan PluginStruct)

	// var myint TmpTest

	// for i, v := range params {
	// 	v := append(v, GeneralFunctions...)
	// 	go LoadPlugin2(i, v, conf, chanPlug)
	// }

	go LoadPlugin3("tmptest", conf, chanPlug)
	a, ok := <-chanPlug
	var an anny = a
	fmt.Println(a, ok)
	fmt.Println(an)
	var aa []func() (interface{}, error)
	for _, v := range a.Functions {
		aa = append(aa, v)
	}
	fmt.Println(aa)
	var loo interface{} = aa
	fmt.Println(loo.(TmpTest))
}

//
// -------------------------
// OLD PLUGIN IMPLEMENTATION
// -------------------------
//

// // FILLME creating new plugin interface, template
// // type NewPluginInterface interface {
// // 	PluginMethod() returnType
// // }

// Loads Plugins and runs them concurrently.
// When you create a new plugin create a new interface and call this plugin in this function
// func PluginManager(params ...string) {
// 	chanPlug := make(chan interface{})

// 	go LoadPlugin("resources", chanPlug)
// 	res, ok := <-chanPlug
// 	var resources Resources = res.(Resources)
// 	checkPlugin(resources, ok)

// 	if len(params) > 0 && params[0] == "tmptest" {
// 		go LoadPlugin("tmptest", chanPlug)
// 		res, ok := <-chanPlug
// 		var tmptest TmpTest = res.(TmpTest)
// 		checkPlugin(res, ok)
// 		tmptest.TmpTest()
// 	}

// 	// FILLME Loading new plugin, template
// 	// go LoadPlugin("pluginName", chanPlug)
// 	// res, ok = <-chanPlug
// 	// var pluginName NewPluginInterface = res.(NewPluginInterface)
// 	// checkPlugin(res, ok)

// 	// Run plugins concurrently and log out
// 	go logPlugin("AgentInfo", resources.AgentInfo(), true)
// 	for {
// 		go logPlugin("ResourceUsage", resources.ResourceUsage(), true)

// 		// FILLME Running/Log out a plugin, template
// 		// go logPlugin("InfoAboutFunction", pluginName.Function(), true)

//		time.Sleep(30 * time.Second)
//	}
// }

// // Logs plugin outputs.
// func logPlugin(title string, mp map[string]interface{}, toJson bool) {
// 	log.Printf("\n----- %v -----\n", title)
// 	if toJson {
// 		data, err := json.MarshalIndent(&mp, "", " ")
// 		utils.Check(err)
// 		log.Println(string(data))
// 	} else {
// 		for i, v := range mp {
// 			log.Printf("%v: %v\n", i, v)
// 		}
// 	}
// }

// // Checks plugin status
// func checkPlugin(plugVal interface{}, status bool) {
// 	if status {
// 		if plugVal == nil {
// 			log.Fatal("Plugin loaded but there is no value!")
// 		} else {
// 			plugInfo, _ := plugVal.(PlugGeneral).Info()
// 			log.Printf("Plugin \"%v\" loaded and ready to use, version \"%v\" ", plugInfo["name"], plugInfo["version"])
// 		}
// 	} else {
// 		if plugVal == nil {
// 			log.Fatal("Plugin closed and there is no value! ")
// 		} else {
// 			log.Fatal("Plugin closed or there is an error")
// 		}
// 	}
// }

// // // TODO response to Lider
// // // func createResponse() {

// // // }
