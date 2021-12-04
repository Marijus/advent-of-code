package common

import (
	"os"
	"path"
)

func GetInput(filename string) (result string) {
	filePath := path.Join("../inputs", filename)
	b, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	result = string(b)
	return
}
