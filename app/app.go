package app

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/kodekage/spinup/util"
	"github.com/spf13/viper"
)

type App interface {
	CreateProject(cmd string, args ...string) interface{}
	RemoveProject(path string) bool
}

type SpinupApp struct {
	Name      string
	Directory string
}

func New() SpinupApp {
	app := SpinupApp{
		Name:      viper.GetString("name"),
		Directory: viper.GetString("output"),
	}

	return app
}

func (a SpinupApp) CreateProject(cmd string, args ...string) {
	execTime := util.ExecTime()
	defer execTime()

	isValid := util.ValidateCommand(cmd)

	if isValid {
		path := util.CreateDirectory(a.Name, a.Directory)

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
	}
}
