package io

import (
	"bufio"
	"os"
)

func FileReadAll[T string | []byte](path string) T {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return T(data)
}

func FileReadAllLines(path string, f func(string)) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		f(scanner.Text())
	}
}