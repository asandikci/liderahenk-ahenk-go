package main

import (
	"log"
	"os"
	"os/user"
	"runtime"
	"time"

	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/cmd/plugin"
	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/confdir"
	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/utils"

	"github.com/sevlyar/go-daemon"
	"golang.org/x/exp/slices"
)

// Main Function that starts daemon and controls arguments
func main() {
	confdir.GetPaths()
	if len(os.Args) == 2 && slices.Contains([]string{"start", "stop", "restart", "nodaemon", "tmptest", "winsafe"}, os.Args[1]) {
		switch os.Args[1] {
		case "start":
			if runtime.GOOS == "windows" {
				panic("start argument is not compatible with windows right now") // TODO
			}
			utils.CreatePath(confdir.Paths.Data)
			utils.CreatePath(confdir.Paths.Logs)
			cntxt := &daemon.Context{
				PidFileName: confdir.Paths.Pid,
				PidFilePerm: 0644,
				LogFileName: confdir.Paths.Logs + "main.log",
				LogFilePerm: 0640,
				WorkDir:     confdir.Paths.Data,
				Umask:       027,
				Args:        []string{confdir.Paths.Program, "start"},
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
			i, _ := daemon.ReadPidFile(confdir.Paths.Pid)
			Stop(i, 15)
		case "restart":
			i, _ := daemon.ReadPidFile(confdir.Paths.Pid)
			Restart(i, 15)
		case "nodaemon":
			log.Print("STARTED AS NO-DAEMON")

			f := utils.OpenLogFile(confdir.Paths.Logs + "main.log")
			defer f.Close()
			log.SetOutput(f)
			// TODO Also pipe fmt.Print* commands

		case "tmptest":
			log.Print("TEMPORARY TEST STARTED, log files are NOT redirecting!")
			plugNames := []string{"tmptest"}
			plugin.ConnectPlugins(plugNames)

			plugin.Tmptest.TmpTest()
			// if a, b := plugin.Resources.AgentInfo(); b == nil {
			// 	plugin.LogPlugin("Agent Info", a, true)
			// }
			// if a, b := plugin.Resources.ResourceUsage(); b == nil {
			// 	plugin.LogPlugin("Agent Info", a, true)
			// }

			log.Println("waiting...")
			time.Sleep(5 * time.Second)

			plugin.UnloadPlugins(plugNames)
			log.Println("Plugins unloaded succesfully, following lines should be give error!")
			// err := plugin.Tmptest.TmpTest()
			// if err != nil {
			// 	log.Println("Plugin was not loaded or already unloaded !: ", err)
			// }
			// if a, err := plugin.Resources.AgentInfo(); err == nil {
			// 	plugin.LogPlugin("Agent Info", a, true)
			// } else {
			// 	log.Println("Plugin was not loaded or already unloaded !: ", err)
			// }
			// if a, err := plugin.Resources.ResourceUsage(); err == nil {
			// 	plugin.LogPlugin("Agent Info", a, true)
			// } else {
			// 	log.Println("Plugin was not loaded or already unloaded !: ", err)
			// }

			// comment these if you want to run usual code after tmptest
			log.Println("exiting before making usual code")
			os.Exit(0)
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
	time.Sleep(5 * time.Second)
	log.Println("bye")
}
