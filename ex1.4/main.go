// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main
// modified to print names of files where dups occur

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	where := make(map[[2]string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts,"STDIN",where)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts,arg,where)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			for pair,_ := range where {
				if pair[0]==line{
					fmt.Printf("%s\t%s\n", pair[1], line)
				}
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, filename string,where map[[2]string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		t := input.Text()
		counts[t]++
		s := [2]string{t,filename}
		where[s]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
