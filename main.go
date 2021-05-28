package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	/*
		if resp.StatusCode == 200 {
			//bs := []byte{}
			bs := make([]byte, 32x1024)
			resp.Body.Read(bs)
			fmt.Println(string(bs))
		}
	*/
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Wrote bytes:", len(bs))

	return len(bs), nil
}
