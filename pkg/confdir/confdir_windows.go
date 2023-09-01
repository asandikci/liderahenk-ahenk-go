//go:build windows

package confdir

import "os"

func getPaths() Path {
	var path Path
	path.Program = "TODO"
	path.UserConfig, _ = os.UserConfigDir()
	path.Data = "TODO"
	path.Plugins = path.Data + "plugins/"
	path.Logs = path.Data + "logs/"
	path.Pid = "TODO"
	return path
}
