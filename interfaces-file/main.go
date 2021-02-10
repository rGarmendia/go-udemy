package main

import (
	"fmt"
	"io"
	"os"
)

type file struct{}

func main() {

	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}

func (f file) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	fmt.Println("Just wrote this many bytes:", len(p))
	return len(p), nil
}
