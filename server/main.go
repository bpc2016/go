// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server3 is an "echo" server that displays request parameters.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!+handler
// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	//set up our vars
	M := make(map[string]int)
	M["cycles"]=0;
	var  cycles, other int 
	for k, v := range r.Form {
		//fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
		//convert this to an actual integer
		if k!="cycles" && k!="other" { continue }
		n,err := strconv.Atoi(v[0])
		if err != nil {
			log.Print(err)
			continue
		}
		M[k] = n
		switch k {
			case "cycles": cycles=n
			case "other": other=n
		}
		fmt.Fprintf(w, "%q plus 1 = %d\n", k,n+1)
	}
	fmt.Fprintf(w, "%d\n%d\n%d\n",M["cycles"],M["other"],cycles+other)
}

//!-handler
