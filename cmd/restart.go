package cmd

import (
	"os"

	"github.com/intercity/intercity-server/utils"
	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart your Intercity instance",
	Run:   Restart,
}

func init() {
	RootCmd.AddCommand(restartCmd)
}

func Restart(cmd *cobra.Command, args []string) {
	if utils.IntercityInstalled() {
		utils.RestartIntercity()
	} else {
		utils.PrintInstallationInstructions()
		os.Exit(1)
	}
}
