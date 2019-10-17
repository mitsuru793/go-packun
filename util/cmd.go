package util

import (
	"os/exec"
	"os"
	"github.com/mattn/go-shellwords"
)

func RunCommand(cmdStr string) error {
	words, err := shellwords.Parse(cmdStr)
	if err != nil {
		return err
	}

	var cmd *exec.Cmd
	switch len(words) {
	case 0:
		// empty
		return nil
	case 1:
		// only command
		cmd = exec.Command(words[0])
	default:
		// command + option
		cmd = exec.Command(words[0], words[1:]...)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
