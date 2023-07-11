package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	envs, err := ReadDir(args[1])
	if err != nil {
		fmt.Println(err)
	}

	exitCode := RunCmd(args[2:], envs)
	os.Exit(exitCode)
}
