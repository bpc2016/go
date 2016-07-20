// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.

// 1.7 this exercise replaces loading a buffer with direct copy to Stdout
// somehow wanted the second error placed in err2 - complained that
// err was already used??

package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err2 := io.Copy(os.Stdout,resp.Body)
		resp.Body.Close()
		if  err2 != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err2)
			os.Exit(1)
		}
	}
}

//!-
