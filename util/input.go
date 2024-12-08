package util

import (
	"bufio"
	"os"
)

func ReadInput(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	contents := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	return contents
}
