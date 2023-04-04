package app

import (
	"fmt"
	"github.com/kodekage/spinup/util"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
)

type App interface {
	CreateProject(cmd string, args ...string) bool
	RemoveProject(path string) bool
}

type spinupApp struct {
	name      string
	directory string
}

func New() spinupApp {
	app := spinupApp{
		name:      viper.GetString("name"),
		directory: viper.GetString("output"),
	}

	return app
}

func (a spinupApp) CreateProject(cmd string, args ...string) bool {
	execTime := util.ExecTime()
	homeDir := util.HomeDir()
	directory := homeDir

	// use user specified directory
	if len(a.directory) != 0 {
		directory = fmt.Sprintf("%s/%s", homeDir, a.directory)
	}

	isValid := util.ValidateCommand("npx")

	if isValid {
		path := util.CreateDirectory(a.name, directory)

		fmt.Println("=> Bootstrapping Application...💨")

		cmd := exec.Command(cmd, args...)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Dir = path

		err := cmd.Run()

		if err != nil {
			log.Fatalf("there was an error bootstrapping app %s", err)
		}
		fmt.Println("=> Application Created ✅✅ ")
		execTime()
	}

	return true
}
