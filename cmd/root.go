/*
Copyright © 2023 Prosper Opara oparaprosper79@gmail.com
*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmdArgs struct {
	outputDir string
	appName   string
	appType   string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "spinup",
	Short: "Quickly 'spinup' web applications",
	Long: `Spinup a developer tool for quickly bootstrapping web applications
in any choice framework.
`,
	Version: "0.1.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&rootCmdArgs.outputDir, "output", "o", "", "directory to setup application (default is $HOME/)")
	rootCmd.PersistentFlags().StringVarP(&rootCmdArgs.appName, "name", "n", "", "bootstrapped application name and directory name (required)")
	rootCmd.PersistentFlags().StringVarP(&rootCmdArgs.appType, "type", "t", "next", "application framework name")

	viper.BindPFlag("name", rootCmd.PersistentFlags().Lookup("name"))
	viper.BindPFlag("type", rootCmd.PersistentFlags().Lookup("type"))
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))

	rootCmd.MarkPersistentFlagRequired("name")
}
