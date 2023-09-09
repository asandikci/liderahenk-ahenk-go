//go:build windows

package confdir

import "os"

func getPaths() Path {
	var path Path
	path.Program = "C:\\Program Files\\ahenk-go\\ahenk-go.exe"
	path.UserConfig, _ = os.UserConfigDir()
	path.Data = "C:\\ProgramData\\ahenk-go\\"
	path.Plugins = path.Data + "plugins\\"
	path.Logs = path.Data + "logs\\"
	path.Version = path.Data + "version"
	return path
}
