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
	//s
	// THIS FILE WON'T UPDATE ANYMORE (bsecause of .gitignore)

	fmt.Println()
	fmt.Println()
	fmt.Println("----- End of Test Area -----")
	fmt.Println()
	time.Sleep(5 * time.Second)
}
