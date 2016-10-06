package cmd

import (
	"github.com/intercity/intercity-server/utils"
	"github.com/spf13/cobra"
)

var (
	emailAddress string
)

// enable-sslCmd represents the enable-ssl command
var enableSSLCmd = &cobra.Command{
	Use:   "enable-ssl [email-address]",
	Short: "Enable SSL using Let's Encrypt",
	Run:   EnableSSL,
}

func init() {
	RootCmd.AddCommand(enableSSLCmd)
}

func EnableSSL(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		utils.LogError("Email address is required", nil)
	} else {
		emailAddress = args[0]
	}

	if utils.IntercityInstalled() {
		if !utils.CheckValidEmail(emailAddress) {
			utils.LogError("Valid email address is required", nil)
		}

		checkRequirements()
		configureSSL()
		utils.BuildIntercity()
		utils.RestartIntercity()
	} else {
		utils.PrintInstallationInstructions()
	}
}

func configureSSL() {
	configFile := "/var/intercity/containers/app.yml"
	utils.ReplaceData(configFile, "#- \"templates/web.ssl.template.yml\"",
		"- \"templates/web.ssl.template.yml\"")

	utils.ReplaceData(configFile, "#- \"templates/web.letsencrypt.ssl.template.yml\"",
		"- \"templates/web.letsencrypt.ssl.template.yml\"")

	utils.ReplaceData(configFile, "LETSENCRYPT_ACCOUNT_EMAIL: \"example@example.com\"",
		"LETSENCRYPT_ACCOUNT_EMAIL: \""+emailAddress+"\"")
}

func checkRequirements() {
	if _, err := utils.RunCommand("cat /var/intercity/containers/app.yml | grep 8880:80"); err == nil {
		utils.LogError("You can't enable SSL while running Intercity on a custom port", nil)
	}
}
