package utils

func StartIntercity() {
	LogCommand("Starting Intercity")

	if _, err := RunCommand("/var/intercity/launcher start app"); err != nil {
		LogError("Could not start Intercity", err)
	}

	LogSuccess()
}

func RestartIntercity() {
	LogCommand("Restarting Intercity")

	if _, err := RunCommand("/var/intercity/launcher restart app"); err != nil {
		LogError("Could not restart Intercity", err)
	}
	LogSuccess()
}

func BuildIntercity() {
	LogCommand("Building Intercity")
	if _, err := RunCommand("/var/intercity/launcher rebuild app"); err != nil {
		LogError("Could not build Intercity", err)
	}

	LogSuccess()
}

func BootstrapIntercity() {
	LogCommand("Bootstrapping Intercity")
	if _, err := RunCommand("/var/intercity/launcher bootstrap app"); err != nil {
		LogError("Could not bootstrap Intercity", err)
	}

	LogSuccess()
}
