package helpers

import (
	"os"
)

func getWorkingDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return wd
}

func ReadFileHelper(filename string) string {
	content, err := os.ReadFile(getWorkingDirectory() + filename)
	if err != nil {
		panic(err)
	}

	return string(content)
}
