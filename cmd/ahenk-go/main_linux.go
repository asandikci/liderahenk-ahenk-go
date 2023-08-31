//go:build linux && !windows

package main

import (
	"log"
	"os"
	"syscall"

	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/utils"
)

// FIXME there isn't any difference with Stop() function
// TODO There can be a Start() function in it but start function doesnt work properly right now
// Restart ahenk daemon with a specific PID (running from second copy)
func Restart(pid, signal int) {
	Stop(pid, signal)
}

// Stop ahenk daemon with a specific PID (running from second copy)
func Stop(pid, signal int) {
	log.Println("Stop Signal Caught")

	// FILLME what you want to do before stopping daemon?

	if err := syscall.Kill(pid, syscall.Signal(signal)); err == nil {
		log.Printf("Ahenk Daemon with pid number %v Successfully stopped", pid)
		f := utils.OpenLogFile(LogFile)
		defer f.Close()
		log.SetOutput(f)
		log.Printf("Ahenk Daemon with pid number %v Successfully stopped", pid)
	} else {
		log.Fatal(err)
	}
	os.Exit(0)
}
