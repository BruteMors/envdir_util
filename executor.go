package main

import (
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for name, value := range env {
		if value.NeedRemove {
			err := os.Unsetenv(name)
			if err != nil {
				fmt.Printf("cant unset env var: %v /n", err)
			}
		} else {
			err := os.Setenv(name, value.Value)
			if err != nil {
				fmt.Printf("cant set env var: %v /n", err)
			}
		}
	}
	app := exec.Cmd{
		Path:   cmd[0],
		Args:   cmd,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	if err := app.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode()
		}
	}
	return
}
