//go:build !safeplugins

package plugin

import (
	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/confdir"
	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/utils"
	"github.com/eh-steve/goloader"
	"github.com/eh-steve/goloader/jit"
)

type PluginModule struct {
	plugin interface{}
	module *goloader.CodeModule
}

// Load Plugin placed in confdir.Paths.Plugins, returns interface with channel
//
// Do not forget to cast in plugin manager
func LoadPlugin(plugName string, conf jit.BuildConfig) PluginModule {

	// TODO if error caugth try without relative path, this will be good for local testing
	loadable, err := jit.BuildGoPackage(conf, confdir.Paths.Plugins+plugName)
	utils.Check(err)

	module, err := loadable.Load()
	utils.Check(err)
	plugSymbols := module.SymbolsByPkg[loadable.ImportPath]

	// TODO also allow lookup another symbol other than PlugnameConnect
	plugOut := plugSymbols[(utils.FirstUpperEN(plugName) + "Connect")]
	if plugOut == nil {
		panic("There is no exported symbol in " + plugName)
	}

	myPlugin := PluginModule{plugin: plugOut, module: module}
	return myPlugin
}
