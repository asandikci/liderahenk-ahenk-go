//go:build !safeplugins

package plugin // pluginmanager

type Unloadable interface {
	Unload() error
}

var TmpTestModule Unloadable
var ResourcesModule Unloadable

// Connect plugins to main code
//
// Gets plugin names as string slice
func connectPlugins(params []string) {

	conf := getJITConf()

	var pluginM PluginModule

	for _, v := range params {
		pluginM = LoadPlugin(v, conf)
		plug := pluginM.plugin

		switch v {
		case "tmptest":
			Tmptest = plug.(TmpTestI)
			TmpTestModule = pluginM.module
			checkPlugin(Tmptest)
		case "resources":
			Resources = plug.(ResourcesI)
			ResourcesModule = pluginM.module
			checkPlugin(Resources)

			// FILLME[epic=new-plugin-template]
			// case "pluginname":
			//   Pluginname = plug.(NewPluginInterfaceI)
			// 	 checkPlugin(Pluginname, ok)
		}
	}
}
