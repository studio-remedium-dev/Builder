package directory

import (
	"Builder/utils/log"
	"os"
	"os/user"
	"runtime"
)

func BuilderDir(path string) (bool, error) {
	//check if file path exists, returns err = nil if file exists
	_, err := os.Stat(path)

	if err == nil {
		log.Error(".builder dir already exists")
	}

	// should return true if dir doesn't exist
	if os.IsNotExist(err) {

		errDir := os.Mkdir(path, 0755)
		//should return nil once directory is made, if not, throw err
		if errDir != nil {
			log.Fatal("failed to make directory", path, err)
		}
	}

	return true, err
}

// MakeBuilderDir does...
func MakeBuilderDir() {
	if runtime.GOOS != "windows" {
		user, _ := user.Current()
		homeDir := user.HomeDir

		builderPath := homeDir + "/.builder"

		BuilderDir(builderPath)
	} else {
		appDataDir := os.Getenv("LOCALAPPDATA")
		if appDataDir == "" {
			appDataDir = os.Getenv("APPDATA")
		}

		builderPath := appDataDir + "/Builder"

		BuilderDir(builderPath)
	}
}
