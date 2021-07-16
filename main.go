package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

var myPallette *Palette

// The nodeTable is initialize to a set of random colors
func initProcess(img *image.RGBA) {
	myPallette = CreatePalette(64)
}

func main() {
	fmt.Println("We've started")

	inFile, err := os.Open("foto.jpg")
	if err != nil {
		log.Fatal(err)
	}

	img, err := jpeg.Decode(inFile)
	if err != nil {
		log.Fatal(err)
	}
	bounds := img.Bounds()

	// Copy the picture into a structure we can manipulate
	imgTmp := image.NewRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			imgTmp.Set(x, y, img.At(x, y))
		}
	}

	// Initialize my pallete
	fmt.Println("init")
	initProcess(imgTmp)

	myPallette.SetFromImage(imgTmp)
	myNodeMap := CreateNodemap(imgTmp, myPallette)

	// Let's set every pixel to the average of it's neighbours
	fmt.Println("write")
	imgOut := image.NewRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Take color from the colortable
			pColor := myPallette.Get(myNodeMap.GetGroup(x, y))
			pNew := color.NRGBA{uint8(pColor.Get(0)), uint8(pColor.Get(1)), uint8(pColor.Get(2)), 0xFF}
			imgOut.Set(x, y, pNew)
		}
	}

	// Encode as PNG.
	outFile, _ := os.Create("image.png")
	png.Encode(outFile, imgOut)
}
