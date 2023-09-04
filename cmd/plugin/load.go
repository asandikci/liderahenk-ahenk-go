package plugin

import (
	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/confdir"
	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/utils"
	"github.com/eh-steve/goloader/jit"
)

// Load Plugin placed in confdir.Paths.Plugins, returns empty interface with channel
//
// Do not forget to cast in plugin manager
func LoadPlugin(plugName string, conf jit.BuildConfig, chn chan interface{}) {

	// TODO if error caugth try without relative path, this will be good for local testing
	loadable, err := jit.BuildGoPackage(conf, confdir.Paths.Plugins+plugName+"/")
	utils.Check(err)

	module, err := loadable.Load()
	utils.Check(err)
	plugSymbols := module.SymbolsByPkg[loadable.ImportPath]

	// TODO also allow lookup another symbol other than PlugnameConnect
	plugOut := plugSymbols[(utils.FirstUpperEN(plugName) + "Connect")]
	if plugOut == nil {
		panic("There is no exported symbol in " + plugName)
	}

	chn <- plugOut
}
