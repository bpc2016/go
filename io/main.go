package main

import (
	"os"
	//"encoding/binary"
	"fmt"
	//"time"
	"bytes"
	"io"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	//"math/rand"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	//t := time.Now().Nanosecond()
	//rand.Seed(int64(t))

	fp,err := os.Create("./test.png")
	if err !=nil {
		fmt.Printf("File creation error\n")
		return
	}

	buf := new(bytes.Buffer)
	w := io.Writer(buf)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors, was os.Stdout
	
	//write to file
	fp.Write(buf.Bytes())


	fmt.Println("goodbye")
	fp.Close()
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
