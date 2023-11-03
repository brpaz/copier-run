package copier

import (
	"os"
	"os/exec"
)

// Runcopier-runerator runs the copier copier-runerator with the given git url and destination
func RunGenerator(gitUrl string, destination string) error {
	cmd := exec.Command("copier", "copy", gitUrl, destination, "--trust")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
