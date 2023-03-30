package app

import (
	"fmt"
	"github.com/kodekage/spinup/util"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
	"time"
)

type Application struct {
	Name      string
	Directory string
	Type      string
}

func New() *Application {
	app := Application{
		Name:      viper.GetString("name"),
		Directory: viper.GetString("output"),
		Type:      viper.GetString("type"),
	}

	return &app
}

func (app Application) CreateReactApp() {
	start := time.Now()
	homeDir := util.HomeDir()
	directory := homeDir

	// use user specified directory
	if len(app.Directory) != 0 {
		directory = fmt.Sprintf("%s/%s", homeDir, app.Directory)
	}

	isValid := util.ValidateCommand("npx")

	if isValid {
		path := util.CreateDirectory(app.Name, directory)

		fmt.Println("=> Bootstrapping React Application...💨")

		cmd := exec.Command("npx", "create-react-app", ".", "--template typescript --eslint --src-dir --import-alias '\"@/*\"'")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Dir = path

		err := cmd.Run()

		if err != nil {
			log.Fatalf("there was an error bootstrapping app %s", err)
		}
		fmt.Println("=> Application Created ✅✅ ")
		elapsed := time.Now().Sub(start)
		fmt.Printf("Time Ellapsed %s", elapsed)
	}
}
