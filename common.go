package main

import "os"

func getInput(filename string) (result string) {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	result = string(b)
	return
}
