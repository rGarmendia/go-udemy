package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// println(string(bs))
	// that is equals to

	lw := logWriter{}
	// takes first argument a writer type and as second a reader type
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
