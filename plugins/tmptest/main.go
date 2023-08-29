package main

import (
	"fmt"
	"time"
)

type plug string

// exported plugin Symbol
var TmptestConnect plug

func (p plug) TmpTest() {
	fmt.Println()
	fmt.Println("----- Entered Test Area -----")
	fmt.Println()
	fmt.Println()

	// THIS FILE WON'T UPDATE ANYMORE (because of .gitignore)

	fmt.Println()
	fmt.Println()
	fmt.Println("----- End of Test Area -----")
	fmt.Println()
	time.Sleep(5 * time.Second)
}

func (p plug) Info() map[string]string {
	inf := make(map[string]string)
	inf["name"] = "tmptest"
	inf["version"] = "0.0.1"
	inf["support"] = "debian"
	inf["description"] = "Temporary testing"
	inf["developer"] = "asandikci@aliberksandikci.com.tr"
	return inf
}
