package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWritter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	//Acordding to documentarion in go, the Http doesn't return the response body directly in text
	// you have to pass to a reader to read the bytes to some functtion that implements the Reader interface
	// to write this value in the screen, you have to overload the write Function and use some function
	// in this example io.Copy implements the write interface

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWritter{}

	io.Copy(lw, resp.Body)

}

func (logWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes", len(bs))
	return len(bs), nil
}
