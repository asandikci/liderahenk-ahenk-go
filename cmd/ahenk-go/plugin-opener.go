package main

// // Load Plugin with plugin name and function name
// func LoadPlugin(plugName, funcName string) {
// 	plug, err := plugin.Open("../../plugins/resources/main.so")
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	symGreeter, err := plug.Lookup("Greeter")
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	var greeter Greeter
// 	greeter, ok := symGreeter.(Greeter)
// 	if !ok {
// 		fmt.Println("unexpected type from module symbol")
// 		os.Exit(1)
// 	}
// }

// // NEXT move plugin-manager.go main here !
