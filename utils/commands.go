package utils

import "os/exec"

func RunCommand(cmd string) (output string, err error) {
	out, err := exec.Command("bash", "-c", cmd).Output()
	return string(out), err
}
