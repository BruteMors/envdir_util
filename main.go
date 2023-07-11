package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	//args := []string{"./go-envdir", "testdata/env", "/bin/bash", "testdata/echo.sh", "arg1=1", "arg2=2"}
	//fmt.Println(args)

	envs, err := ReadDir(args[1])
	if err != nil {
		fmt.Println(err)
	}

	exitCode := RunCmd(args[2:], envs)
	os.Exit(exitCode)
	//fmt.Println(envs)
}
