//go:build windows

package main

import (
	"log"
	"os"
)

func Restart(pid, signal int) {
	Stop(pid, signal)
}

func Stop(pid, signal int) {
	log.Println("Stop Signal Caught")
	log.Println("TODO, STOP FUNCTION IS NOT FUNCTIONAL IN WINDOWS YET!")
	os.Exit(0)
}
