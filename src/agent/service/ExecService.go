package service

import (
	"agent/utill"
	"log"
	"os/exec"
	"strings"
)

func ExecCommand(commandType string, command string) {
	var cmd *exec.Cmd

	commandType = strings.ToUpper(commandType)

	switch commandType {
	case "BASH":
		cmd = exec.Command("sh", "-c", command)
	case "FILE":

		if strings.HasSuffix(command, ".py") {
			cmd = exec.Command("python", command)
		} else if strings.HasSuffix(command, ".sh") {

			if strings.HasPrefix(command, "http") {

				fileBody := utill.GetRemoteFileBody(command)

				cmd = exec.Command("sh", "-c", fileBody)

			} else {
				cmd = exec.Command("sh", "-c", command)

			}

		} else {
			panic("Unknown file type! " + command)
		}

	default:
		cmd = exec.Command("sh", "-c", "echo 'loop once, no command found!'")
	}

	out, err := cmd.Output()

	if err != nil {
		log.Print(err.Error())
	}

	log.Printf("%s", string(out))
}
