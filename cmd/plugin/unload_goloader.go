//go:build !safeplugins

package plugin

import "log"

func unloadPlugins(params []string) {

	for _, v := range params {
		switch v {
		case "tmptest":
			err := TmpTestModule.Unload()
			unloaded, ok := Tmptest.(PlugUnloaded)
			if ok {
				log.Printf("Succesfully unloaded tmptest: %v - %v", unloaded, err)
			} else {
				log.Fatalf("There was an error while unloading tmptest: %v - %v", unloaded, err)
			}
		case "resources":
			err := ResourcesModule.Unload()
			unloaded, ok := Resources.(PlugUnloaded)
			if ok {
				log.Printf("Succesfully unloaded resources: %v - %v", unloaded, err)
			} else {
				log.Fatalf("There was an error while unloading resources: %v - %v", unloaded, err)
			}

			// FILLME[epic=new-plugin-template]
			// case "pluginname":
			//   PluginnameModule.Unload()
			// TODO proper template
		}
	}
}
