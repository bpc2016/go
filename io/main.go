package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	http.Handle("/images/",http.StripPrefix("/images/", http.FileServer(http.Dir(".")) ) ) 
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<html><title>Mandelbrot</title><body>Superimposed Mandelbrot<div style=\"position:relative\"><div style=\"position:absolute;left:0;top:0;\"><img src=\"http://localhost:8000/images/1.png\" width=1024 height=1024/></div></div></body></html>"))
}
