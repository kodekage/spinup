/*
Copyright © 2023 Prosper Opara oparaprosper79@gmail.com
*/

package cmd

import (
	"github.com/kodekage/spinup/app"
	"github.com/spf13/cobra"
)

// reactCmd represents the React command
var reactCmd = &cobra.Command{
	Use:   "react",
	Short: "spinup react web application",
	Long:  "spinup react web application",
	Run: func(cmd *cobra.Command, args []string) {
		app := app.New()

		app.CreateReactApp()
	},
}

func init() {
	rootCmd.AddCommand(reactCmd)
}
