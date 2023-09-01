package pluginmanager

import (
	"fmt"
	"os"
	"plugin"

	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/confdir"
	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/utils"
	"github.com/eh-steve/goloader/jit"
)

// type mytype func() (interface{}, error)

type PluginStruct struct {
	PluginFunction map[string]interface{}
	Functions      map[string](func() (interface{}, error))
}

func LoadPlugin3(plugName string, conf jit.BuildConfig, chn chan PluginStruct) {
	loadable, err := jit.BuildGoPackage(conf, confdir.Paths.Plugins+plugName+"/")
	utils.Check(err)
	module, err := loadable.Load()
	utils.Check(err)
	plugSymbols := module.SymbolsByPkg[loadable.ImportPath]

	// fmt.Println(*deneme, *deneme2)

	// var empt interface{} = plugSymbols
	// a = empt.(TmpTest)

	// var a TmpTest
	// jit.RegisterTypes(plugSymbols["TmpTest"])
	var plugMan PluginStruct
	plugMan.PluginFunction = plugSymbols
	fmt.Println(plugMan.PluginFunction)
	plugMan.Functions = make(map[string](func() (interface{}, error)))
	for i, v := range plugMan.PluginFunction {
		switch f := v.(type) {
		case func([]byte) (interface{}, error):
			result, err := f([]byte(`{"k":"v"}`))
			if err != nil {
				panic(err)
			}
			fmt.Println(result)
		case func() (interface{}, error):
			fmt.Println(i, 1)
			plugMan.Functions[i] = f
		default:
			fmt.Println(f)
			panic("Function signature was not what was expected")
		}
	}
	fmt.Println(plugMan.Functions)
	for i, o := range plugMan.Functions {
		// fmt.Println(i, o)
		fmt.Println(i, &o)
		// fmt.Println(i, *o)
	}
	// plugMan.Functions["TmpTest"]()
	fmt.Println(1)
	chn <- plugMan
	// var empt mytype = plugSymbols["TmpTest"].(mytype)

	// a := empt.(TmpTest)
	// fmt.Println(empt)

	// switch f := plugSymbols["MyFunc"].(type) {
	// case func([]byte) (interface{}, error):
	// 	result, err := f([]byte(`{"k":"v"}`))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(result)
	// default:
	// 	fmt.Println(f)
	// 	panic("Function signature was not what was expected")
	// }

	// chn <- deneme
}

// func LoadPlugin2(plugName string, plugFunction []string, conf jit.BuildConfig, chn chan interface{}) {

// 	loadable, err := jit.BuildGoPackage(conf, confdir.Paths.Plugins+plugName+"/")
// 	utils.Check(err)

// 	module, err := loadable.Load()

// 	plugSymbols := module.SymbolsByPkg[loadable.ImportPath]
// 	utils.Check(err)

// 	for i, v := range plugFunction {
// 		var f interface{} = plugSymbols

// 	}

// 	// switch f := plugSymbols["MyFunc"].(type) {
// 	// case func([]byte) (interface{}, error):
// 	// 	result, err := f([]byte(`{"k":"v"}`))
// 	// 	if err != nil {
// 	// 		panic(err)
// 	// 	}
// 	// 	fmt.Println(result)
// 	// default:
// 	// 	fmt.Println(f)
// 	// 	panic("Function signature was not what was expected")
// 	// }

// 	// chn <-
// }

//
// -------------------------
// OLD PLUGIN IMPLEMENTATION
// -------------------------
//

// Load Plugin placed in PluginDir, returns empty interface with channel
// Do not forget to cast in plugin manager
//
// Give Plugin Name as argument and be sure you compiled plugins with `-buildmode=plugin` to PluginDir as `pluginname.so`
func LoadPlugin(plugName string, chn chan interface{}) {

	// TODO if error caugth try without relative path, this will be good for local testing
	plug, err := plugin.Open(confdir.Paths.Plugins + plugName + ".so")
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
	chn <- plugOut
}
