package main

import (
	"log"
	"os"
	"os/user"

	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/utils"

	"github.com/sevlyar/go-daemon"
	"golang.org/x/exp/slices"
)

const PidFile = "/run/ahenk-go.pid"
const ExecutablePath = "/usr/bin/ahenk-go"
const DataDir = "/etc/ahenk-go/"
const LogFile = DataDir + "ahenk.log"
const LibDir = "/usr/share/ahenk-go/"
const PluginDir = LibDir + "/plugins/"

// Main Function that starts daemon and controls arguments
func main() {
	if len(os.Args) == 2 && slices.Contains([]string{"start", "stop", "restart", "nodaemon", "tmptest"}, os.Args[1]) {
		switch os.Args[1] {
		case "start":
			utils.CreatePath(DataDir)
			cntxt := &daemon.Context{
				PidFileName: PidFile,
				PidFilePerm: 0644,
				LogFileName: LogFile,
				LogFilePerm: 0640,
				WorkDir:     LibDir,
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
		case "restart":
			i, _ := daemon.ReadPidFile(PidFile)
			Restart(i, 15)
		case "nodaemon":
			log.Print("STARTED AS NO-DAEMON")

			f := utils.OpenLogFile(LogFile)
			defer f.Close()
			log.SetOutput(f)
			// TODO Also pipe fmt.Print* commands

		case "tmptest":
			log.Print("TEMPORARY TEST STARTED, log files are NOT redirecting!")
			PluginManager("tmptest")
		}
	} else {
		panic("Please enter a valid option !")
	}

	current, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	if current.Uid != "0" {
		log.Fatal("Ahenk-go requires superuser privilege")
	}

	PluginManager()
	// NEXT Make PluginManager async !
}
