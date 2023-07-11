package main

import (
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	envValues := make(Environment)

	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fileInfo, err := entry.Info()
			if err != nil {
				return nil, err
			}
			nameFile := fileInfo.Name()
			if fileInfo.Size() == 0 {
				envValues[nameFile] = EnvValue{
					Value:      "",
					NeedRemove: true,
				}
				continue
			}

			if strings.Contains(nameFile, "=") {
				continue
			}

			valueFile, err := os.ReadFile(dir + "/" + nameFile)
			if err != nil {
				return nil, err
			}

			stringsOfFile := strings.Split(string(valueFile), "\n")
			firstStringOfFileWithoutZeroEsc := strings.Replace(stringsOfFile[0], "\000", "\n", -1)
			if strings.TrimSpace(firstStringOfFileWithoutZeroEsc) == "" {
				envValues[nameFile] = EnvValue{
					Value:      "",
					NeedRemove: true,
				}
				continue
			}

			envValues[nameFile] = EnvValue{
				Value:      firstStringOfFileWithoutZeroEsc,
				NeedRemove: false,
			}
		}
	}
	return envValues, nil
}
