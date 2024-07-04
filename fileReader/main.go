package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	filename := os.Args[1]
	fmt.Println(filename)

	f, _ := os.Open(filename)

	io.Copy(logWriter{}, f)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
