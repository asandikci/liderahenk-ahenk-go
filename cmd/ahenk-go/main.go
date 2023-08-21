package main

import (
	"ahenk-go/pkg/utils"
	"log"
	"os"
	"syscall"

	"github.com/sevlyar/go-daemon"
	"golang.org/x/exp/slices"
)

const PidFile = "/run/ahenkd-go.pid"
const ExecutablePath = "/usr/bin/ahenkd-go"
const ConfDir = "/etc/ahenk-go/"

// FIXME there isn't any difference with Stop() function
// TODO There can be a Start() function in it but start function doesnt work properly right now
func Restart(pid, signal int) {
	Stop(pid, signal)
}

// Stop Ahenkd Daemon with a specific PID (running from second copy)
// do not forget to use also os.Exit() when using in ExecStop
func Stop(pid, signal int) {
	log.Println("Stop Signal Caught")

	// FILLME what you want to do before stopping daemon?

	if err := syscall.Kill(pid, syscall.Signal(signal)); err == nil {
		log.Printf("Ahenk Daemon with pid number %v Successfully stopped", pid) // TODO Also log to /etc/ahenk-go/ahenkd.log
	} else {
		log.Fatal(err)
	}
}

// Main Function that starts daemon and controls arguments
func main() {
	if len(os.Args) == 2 && slices.Contains([]string{"start", "stop", "restart", "nodaemon"}, os.Args[1]) {
		switch os.Args[1] {
		case "start":
			utils.CreatePath(ConfDir)
			cntxt := &daemon.Context{
				PidFileName: PidFile,
				PidFilePerm: 0644,
				LogFileName: ConfDir + "ahenkd.log",
				LogFilePerm: 0640,
				WorkDir:     ConfDir,
				Umask:       027,
				Args:        []string{ExecutablePath, "start"},
			}
			d, err := cntxt.Reborn()
			if err != nil {
				log.Fatal("Daemon Reborn ERROR:", err)
			}
			if d != nil {
				return
			}
			defer cntxt.Release()
			log.Println("- - - - - - - - - - - - - - - - - -")
			log.Print("Daemon Succesfully Started")
		case "stop":
			i, _ := daemon.ReadPidFile(PidFile)
			Stop(i, 15)
			os.Exit(0)
		case "restart":
			i, _ := daemon.ReadPidFile(PidFile)
			Restart(i, 15)
			os.Exit(0)
		case "nodaemon":
			log.Print("STARTED AS NO-DAEMON")
		}
	} else {
		panic("Please enter a valid option !")
	}
	PluginManager()
	log.Print("Plugin Manager Started Succesfully")
	// NEXT Make PluginManager async !
}
