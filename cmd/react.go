/*
Copyright © 2023 Prosper Opara oparaprosper79@gmail.com
*/

package cmd

import (
	"github.com/kodekage/spinup/app"
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
			app.CreateProject("npx", "create-next-app", ".", "--typescript --eslint --src-dir --import-alias '\"@/*\"'")
		default:
			app.CreateProject("npx", "create-react-app", ".", "--template typescript")
		}

	},
}

func init() {
	rootCmd.AddCommand(reactCmd)
}
