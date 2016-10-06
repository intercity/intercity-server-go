package utils

import (
	"log"
	"os"
)

func LogCommand(command string) {
	println("---- " + command)
}

func LogSuccess() {
	println("    Done")
}

func LogError(message string, err error) {
	println(message)
	if err != nil {
		log.Fatal(err)
	} else {
		os.Exit(1)
	}
}
