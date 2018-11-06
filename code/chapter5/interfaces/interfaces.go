// Sample program to show how to write a simple version of curl using
// the io.Reader and io.Writer interface support.
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

// init is called before main.
func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./example2 <url>")
		os.Exit(-1)
	}
}

func curl() {
	// Get a response from the web server.
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Copies from the Body to Stdout.
	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}

func bufferz() {
	var b bytes.Buffer

	// Write a string to the buffer.
	b.Write([]byte("Hello"))

	// Use Fprintf to concatenate a string to the Buffer.
	fmt.Fprintf(&b, "World!")

	// Write the content of the Buffer to stdout.
	io.Copy(os.Stdout, &b)
}

// main is the entry point for the application.
func main() {
	// curl()
	bufferz()
}
