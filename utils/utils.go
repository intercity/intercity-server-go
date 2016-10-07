package utils

import (
	"os"
	"regexp"
)

func IntercityInstalled() bool {
	if _, err := os.Stat("/var/intercity"); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func PrintInstallationInstructions() {
	println("Intercity is not installed.")
	println("To install Intercity, run:")
	println("    intercity-server install [hostname]")
}

func PrintUpdateInstructions() {
	println("Intercity is already installed.")
	println("To update Intercity, run:")
	println("    intercity-server update")
}

func CheckValidDomain(domain string) bool {
	var validID = regexp.MustCompile(`^([a-z0-9\*-]+\.)*[a-z0-9\*-]+$`)
	return validID.MatchString(domain)
}

func CheckValidEmail(email string) bool {
	var validID = regexp.MustCompile(`([^@]+)@([^\.]+)`)
	return validID.MatchString(email)
}
