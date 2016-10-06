package cmd

import (
	"log"
	"os"

	"github.com/intercity/intercity-server/utils"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update your Intercity instance",
	Run:   update,
}

func init() {
	RootCmd.AddCommand(updateCmd)
}

func update(cmd *cobra.Command, args []string) {
	if utils.IntercityInstalled() {
		utils.LogCommand("Updating Intercity")

		if _, err := utils.RunCommand("/var/intercity/launcher rebuild app"); err != nil {
			log.Fatal(err)
		}
		utils.LogSuccess()
	} else {
		utils.PrintInstallationInstructions()
		os.Exit(1)
	}
}
