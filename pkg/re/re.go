package re

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func Re(command string) {
	cmd := exec.Command("bash", "-c", command)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slurpOut, _ := io.ReadAll(stdout) // TODO: fix this
	fmt.Printf("%s\n", slurpOut)

	slurpErr, _ := io.ReadAll(stderr) // TODO: fix this
	fmt.Printf("%s\n", slurpErr)

	if err := cmd.Wait(); err != nil {
		log.Println(err)

		log.Println("Restarting...\n\n ")
		Re(command)
	}
}
