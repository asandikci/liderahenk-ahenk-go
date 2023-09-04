//go:build windows

package plugin

import "github.com/eh-steve/goloader/jit"

func getJITConf() jit.BuildConfig {
	return jit.BuildConfig{
		KeepTempFiles:   true,                     // Files are copied/written to a temp dir to ensure it is writable. This retains the temporary copies
		ExtraBuildFlags: []string{"-x"},           // Flags passed to go build command
		BuildEnv:        []string{"GOOS=windows"}, // Env vars to set for go build toolchain
		TmpDir:          "",                       // To control where temporary files are copied
		DebugLog:        true,                     //
	}
}
