package cmd

import (
	"os"
	"path/filepath"

	"github.com/intercity/intercity-server/utils"
	"github.com/spf13/cobra"
)

var (
	hostname      string
	useCustomPort bool
)

var installCmd = &cobra.Command{
	Use:   "install [hostname]",
	Short: "Install your very own Intercity instance",
	Run:   Install,
}

func init() {
	RootCmd.AddCommand(installCmd)

	installCmd.Flags().Bool("use-custom-port", false, "Use Intercity with a custom port")
}

func Install(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		utils.LogError("hostname needs to be provided", nil)
	} else {
		hostname = args[0]
	}

	useCustomPort, _ = cmd.Flags().GetBool("use-custom-port")

	if utils.IntercityInstalled() {
		utils.PrintUpdateInstructions()
	} else {
		if !utils.CheckValidDomain(hostname) {
			utils.LogError("valid hostname needs to be provided", nil)
		}

		installDocker()
		downloadIntercity()
		configureIntercity()
		buildIntercity()
		startIntercity()
	}
}

func installDocker() {
	utils.LogCommand("Installing Docker")

	if _, err := utils.RunCommand("which docker"); err != nil {
		if _, err := utils.RunCommand("wget -nv -O - https://get.docker.com | sh"); err != nil {
			utils.LogError("Could not install Docker", err)
		} else {
			utils.LogSuccess()
		}
	} else {
		println("     Docker is already installed. Let's continue")
	}
}

func downloadIntercity() {
	utils.LogCommand("Downloading Intercity")

	path := filepath.Join("/var", "intercity")
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		utils.LogError("Could not create Intercity installation directory.", err)
	}

	cmd := "git clone https://github.com/intercity/intercity-docker.git -b 0-4-stable /var/intercity"
	if _, err := utils.RunCommand(cmd); err != nil {
		utils.LogError("Could not download Intercity", err)
	}
	utils.LogSuccess()
}

func configureIntercity() {
	utils.LogCommand("Configuring Intercity")

	if _, err := utils.RunCommand("cp /var/intercity/samples/app.yml /var/intercity/containers/"); err != nil {
		utils.LogError("Could not copy configuration file", err)
	}

	configFile := "/var/intercity/containers/app.yml"
	utils.ReplaceData(configFile, "intercity.example.com", hostname)

	if useCustomPort {
		utils.ReplaceData(configFile, "80:80", "8880:80")
		utils.ReplaceData(configFile, "443:443", "8443:443")
	}

	utils.LogSuccess()
}

func buildIntercity() {
	utils.LogCommand("Building Intercity")
	if _, err := utils.RunCommand("/var/intercity/launcher bootstrap app"); err != nil {
		utils.LogError("Could not bootstrap Intercity", err)
	}

	utils.LogSuccess()
}

func startIntercity() {
	utils.StartIntercity()
}
