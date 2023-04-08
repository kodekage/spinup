/*
Copyright © 2023 Prosper Opara oparaprosper79@gmail.com
*/

package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/kodekage/spinup/app"
	"github.com/kodekage/spinup/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// reactCmd represents the React command
var reactCmd = &cobra.Command{
	Use:   "react",
	Short: "spinup react web application",
	Long:  "spinup react web application",
	Run: func(cmd *cobra.Command, args []string) {
		app := app.New()
		appType := viper.GetString("type")

		switch appType {
		case "nextjs":
			bootstrapNextApp(app)
		default:
			app.CreateProject("npx", "create-react-app", ".", "--template typescript")
		}

	},
}

func init() {
	rootCmd.AddCommand(reactCmd)
}

// bootstrapNextApp creates a barebones nextjs project
func bootstrapNextApp(app app.SpinupApp) {
	execTime := util.ExecTime()
	defer execTime()

	path := util.CreateDirectory(app.Name, app.Directory)
	templateJson := "./templates/react/next.package.json"
	destPackageJson := fmt.Sprintf("%s/package.json", path)

	file, err := os.Open(templateJson)
	if err != nil {
		log.Fatal("Open error ", err)
	}

	dst, err := os.Create(destPackageJson)
	if err != nil {
		log.Fatal("Create error ", err)
	}

	_, err = io.Copy(dst, file)
	if err != nil {
		log.Fatal("Copy error ", err)
	}

	isValid := util.ValidateCommand("npm")

	if isValid {
		fmt.Println("=> Bootstrapping Application...💨")

		cmd := exec.Command("npm", "install", "next", "react", "react-dom")

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
