package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get the version of the Intercity-Server cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version: 0.3.0")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
