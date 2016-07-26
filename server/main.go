// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server3 is an "echo" server that displays request parameters.
package main

import (
	//"fmt"
	"log"
	"net/http"
	"strconv"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"io"
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
	//set up our vars num=iterations
	var  num int 
	for k, v := range r.Form {
		if k!="num" { continue }
		n,err := strconv.Atoi(v[0])
		if err != nil {
			log.Print(err)
			continue
		}
		switch k {
			case "num": num = n
		}
		//fmt.Fprintf(w, "%q plus 1 = %d\n", k,n+1)
	}
	//fmt.Fprintf(w, "%d\n%d\n",cycles,other)
	mandelb(w, num)
}

//!-handler

func mandelb(w io.Writer, n int) {
	const (
		//iterations = 631*600// 600*600
		//x0, y0, side = -0.8, 0.15, 0.1
		x0, y0, side = -0.75, 0.21, 0.04
		width, height = 1024, 1024
	)
	if n==0 {n=1} // need some number of iterations  ...
	iterations := n*600

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		//y := float64(py)/height*(ymax-ymin) + ymin
		y := y0 + float64(py)/height*side
		for px := 0; px < width; px++ {
			//x := float64(px)/width*(xmax-xmin) + xmin
			x := x0 +float64(px)/width*side
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, height-py, mandelbrot(z,iterations))
		}
	}
	//png.Encode(os.Stdout, img) // NOTE: ignoring errors
	png.Encode(w, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128, iterations int) color.Color {
	const contrast = 15
	var c uint64
	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			//return color.Gray{255 - contrast*n}
			c = coloRatio(n,iterations)
			w := getColors(c)
			return color.RGBA{w[0],w[1],w[2],255}
		}
	}
	return color.Black
}

func coloRatio(z, max int) uint64{
	const T = 8355771 //codeColors([3]int{127,127,127})
	const B	 = float64(1256) 
	x := float64(z)/float64(max)
	y := (1 - B*x - (1-B*float64(max))*x*x)*float64(T)
	return uint64(math.Floor(y))
}


func codeColors(c [3]int) int{
	const B = 1<<8
	return ((c[0]*B) +c[1])*B + c[2] 
}

func getColors(v uint64) [3] uint8 {
	const B = 1<<8
	var w [3] uint8 
	for i:=0; i<3; i++ {
		w[2-i] = uint8(v%B)
		v = v/B
	}
	return w
}
