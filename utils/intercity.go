package utils

import "log"

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
		log.Fatal(err)
	}
	LogSuccess()
}
