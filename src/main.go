package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"./nodemap"
	"./palette"
)

var nodeTable *palette.Palette
var nodeAssignments *nodemap.Nodemap

// The nodeTable is initialize to a set of random colors
func initProcess(img *image.RGBA) {
	nodeTable = palette.New(img)
	nodeAssignments = nodemap.New(img)
}

func setToNearest(img *image.RGBA) {
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			// Extract pixel
			pOffset := img.PixOffset(x, y)
			rImg := float64(img.Pix[pOffset])
			gImg := float64(img.Pix[pOffset+1])
			bImg := float64(img.Pix[pOffset+2])
			// Calculate the closest color for nmap pixel
			iClose := nodeTable.Append(rImg, gImg, bImg)
			// Assign nmap pixel to the closest
			nodeAssignments.SetGroup(x, y, iClose)
		}
	}
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
	fmt.Println("round 1 - run")
	setToNearest(imgTmp)
	fmt.Println("round 1 - shift")
	fmt.Printf("%v \n", nodeTable)
	nodeTable.Shift()
	fmt.Println("round 1 - split")
	nodeTable.Split()
	fmt.Println("round 1 - reset")
	nodeTable.Reset()

	fmt.Println("round 2 - run")
	setToNearest(imgTmp)
	fmt.Println("round 2 - shift")
	nodeTable.Shift()
	fmt.Printf("%v \n", nodeTable)
	fmt.Println("round 2 - split")
	nodeTable.Split()
	fmt.Println("round 2 - reset")
	nodeTable.Reset()

	fmt.Println("round 3 - run")
	setToNearest(imgTmp)
	fmt.Println("round 3 - shift")
	nodeTable.Shift()
	fmt.Printf("%v \n", nodeTable)
	fmt.Println("round 3 - split")
	nodeTable.Split()
	fmt.Println("round 3 - reset")
	nodeTable.Reset()

	fmt.Println("round 4 - run")
	setToNearest(imgTmp)
	fmt.Println("round 4 - shift")
	nodeTable.Shift()
	fmt.Printf("%v \n", nodeTable)
	fmt.Println("round 4 - split")
	nodeTable.Split()
	fmt.Println("round 4 - reset")
	nodeTable.Reset()

	fmt.Println("round 5 - run")
	setToNearest(imgTmp)
	fmt.Println("round 5 - shift")
	nodeTable.Shift()
	// No Split!!
	fmt.Println("round 5 - reset")
	nodeTable.Reset()

	// Let's set every pixel to the average of it's neighbours
	fmt.Println("write")
	imgOut := image.NewRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Take color from the colortable
			pColor := nodeTable.Get(nodeAssignments.GetGroup(x, y))
			pNew := color.NRGBA{uint8(pColor.Get(0)), uint8(pColor.Get(1)), uint8(pColor.Get(2)), 0xFF}
			imgOut.Set(x, y, pNew)
		}
	}

	// Encode as PNG.
	outFile, _ := os.Create("image.png")
	png.Encode(outFile, imgOut)
}
