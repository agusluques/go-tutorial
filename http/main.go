package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type myWriter struct{}

func main() {
	resp, err := http.Get("http://agustinluques.com.ar")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// br := make([]byte, 9999999)
	// resp.Body.Read(br)
	// fmt.Println(string(br))

	// io.Copy(os.Stdout, resp.Body)

	mw := myWriter{}
	io.Copy(mw, resp.Body)
}

func (myWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
