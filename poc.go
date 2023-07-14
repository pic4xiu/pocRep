package main

import (
	"fmt"
	"image"
	"os"
	"runtime"

	"github.com/disintegration/imaging"
)

func main() {
	runtime.GOMAXPROCS(1)
	file, _ := os.Open("output.tiff")
	src, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	imaging.Grayscale(src)
	// src = imaging.CropAnchor(src, 30, 30, imaging.Center)

}
