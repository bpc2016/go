package main

import (
	"log"
	"net/http"
	//"os"
	"fmt"
	"bytes"
	"io"
	"image"
	"image/draw"
	"image/color"
	"image/png"
	"math/cmplx"
	//"time"
	//"math/rand"
	"encoding/base64"
)

func main() {
	http.HandleFunc("/", handler_init) // each request calls handler
	http.HandleFunc("/image", handler_image) // each request calls handler
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir(".")) ) ) 
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// display the base html
func handler_init(w http.ResponseWriter, r *http.Request) {
	//t := time.Now().Nanosecond()
	//rand.Seed(int64(t))

	setCoords()
 	// note the use of backtick!!
	w.Write([]byte(`
<html>
<head>
<title>Mandelbrot</title>
<script type="text/javascript" src="static/jquery-1.9.1.min.js"></script>
<script type="text/javascript" src="static/mandelbrot.js"></script>
</head>
<body><h2>Mandelbrot</h2>
<div id="imgs" style="position:relative">
</div>
</body>
</html>`))}

// return a mandelbrot image
func handler_image(w http.ResponseWriter, r *http.Request) {
	getImage(w) // fetch the base64 encoded png image
}

var (
	width = 1024
	height = 1024
	x [1024*1024] int
	y [1024*1024] int
	position = 0
	array_set bool
)

func setCoords(){
	//fill our coordinate arrays
	k := 0;
	for iy := 0; iy<height; iy++{
		for ix := 0; ix<width; ix++{
			x[k] = ix
			y[k] = iy
			k++	
		}
	}		
	array_set = true
	position = 0
}

func getImage(w io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height, N          = 1024, 1024, 1024*1024
	)

	fmt.Println("position is: ",position)

	if position == N {  // we are done
		io.WriteString(w,"0")
		return 
	}
	lim := position + 1024*100 //1024*256
	if lim > N { lim=N }

	var px,py int	//pixel positions

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), image.Transparent, image.ZP, draw.Src)

	buf := new(bytes.Buffer)
	for k := position; k<lim; k++{ // N
		px = x[k]; py = y[k]
		rx := float64(px)/width*(xmax-xmin) + xmin
		ry := float64(py)/height*(ymax-ymin) + ymin
		z := complex(rx, ry)
		// Image point (px, py) represents complex value z.
		img.Set(px, py, mandelbrot(z))
		position = k+1;
	}

	png.Encode(io.Writer(buf), img) // NOTE: ignoring errors, to an io.Writer

	// convert to base64
	encoder := base64.NewEncoder(base64.StdEncoding, w) // send to target w
	encoder.Write(buf.Bytes())
	encoder.Close()
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
