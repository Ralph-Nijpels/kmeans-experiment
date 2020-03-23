package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
)

type KNode struct {
	RCenter, GCenter, BCenter      float64 // Center
	RTotal, GTotal, BTotal, NTotal float64 // Running total of assigned pixel
}

// Nodemap contains the assignment of each pixel to a color in the colortable
type nodemap struct {
	Assignments []uint16
	Top         int
	Left        int
	Bottom      int
	Right       int
}

// NewNodemap creates a nodemap for all pixels in your image
func NewNodemap(img *image.RGBA) *nodemap {
	top := img.Bounds().Min.Y
	left := img.Bounds().Min.X
	bottom := img.Bounds().Max.Y
	right := img.Bounds().Max.X
	return &nodemap{make([]uint16, (bottom-top)*(right-left)), top, left, bottom, right}
}

func (nmap *nodemap) GetGroup(x, y int) uint16 {
	return nmap.Assignments[(y-nmap.Top)*(nmap.Right-nmap.Left)+x]
}

func (nmap *nodemap) SetGroup(x, y int, v uint16) {
	nmap.Assignments[(y-nmap.Top)*(nmap.Right-nmap.Left)+x] = v
}

var nodeTable [64]KNode
var nodeAssignments *nodemap

// The nodeTable is initialize to a set of random colors
func initPalette(img *image.RGBA) {

	for i := 0; i < len(nodeTable); i++ {
		// No seed used, but that is OK for now
		nodeTable[i].RCenter = rand.Float64() * 256
		nodeTable[i].GCenter = rand.Float64() * 256
		nodeTable[i].BCenter = rand.Float64() * 256
		// Administration set to 0
		nodeTable[i].RTotal = 0
		nodeTable[i].GTotal = 0
		nodeTable[i].BTotal = 0
		nodeTable[i].NTotal = 0
	}

	nodeAssignments = NewNodemap(img)
}

func colorDistance(rFrom, gFrom, bFrom, rTo, gTo, bTo float64) float64 {
	// Euclidian distance
	// return math.Sqrt(math.Pow(rTo-rFrom, 2) + math.Pow(gTo-gFrom, 2) + math.Pow(bTo-bFrom, 2))
	// City Block Distance works just as well and is a lot quicker
	return math.Abs(rTo-rFrom) + math.Abs(gTo-gFrom) + math.Abs(bTo-bFrom)
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
			iClose := uint16(0)
			rTbl := nodeTable[iClose].RCenter
			gTbl := nodeTable[iClose].GCenter
			bTbl := nodeTable[iClose].BCenter
			dClose := colorDistance(rTbl, gTbl, bTbl, rImg, gImg, bImg)
			for i := 1; i < len(nodeTable); i++ {
				rTbl = nodeTable[i].RCenter
				gTbl = nodeTable[i].GCenter
				bTbl = nodeTable[i].BCenter
				dNode := colorDistance(rTbl, gTbl, bTbl, rImg, gImg, bImg)
				if dNode < dClose {
					iClose = uint16(i)
					dClose = dNode
				}
			}
			// Assign nmap pixel to the closest
			nodeAssignments.SetGroup(x, y, iClose)
			// Update the closest node
			nodeTable[iClose].RTotal += rImg
			nodeTable[iClose].GTotal += gImg
			nodeTable[iClose].BTotal += bImg
			nodeTable[iClose].NTotal++
		}
	}
}

func readjustPallette() float64 {
	var totalChange float64
	// Change all centers to their average
	for i := 0; i < len(nodeTable); i++ {
		// No seed used, but that is OK for now
		if nodeTable[i].NTotal > 0 {
			rNew := nodeTable[i].RTotal / nodeTable[i].NTotal
			gNew := nodeTable[i].GTotal / nodeTable[i].NTotal
			bNew := nodeTable[i].BTotal / nodeTable[i].NTotal
			totalChange += math.Abs(nodeTable[i].RCenter - rNew)
			totalChange += math.Abs(nodeTable[i].GCenter - gNew)
			totalChange += math.Abs(nodeTable[i].BCenter - bNew)
			nodeTable[i].RCenter = rNew
			nodeTable[i].GCenter = gNew
			nodeTable[i].BCenter = bNew
		}
		// Administration set to 0
		nodeTable[i].RTotal = 0
		nodeTable[i].GTotal = 0
		nodeTable[i].BTotal = 0
		nodeTable[i].NTotal = 0
	}
	return totalChange
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
	initPalette(imgTmp)
	setToNearest(imgTmp)
	fmt.Println(readjustPallette())
	setToNearest(imgTmp)
	fmt.Println(readjustPallette())
	setToNearest(imgTmp)
	fmt.Println(readjustPallette())
	setToNearest(imgTmp)
	fmt.Println(readjustPallette())
	setToNearest(imgTmp)
	fmt.Println(readjustPallette())
	setToNearest(imgTmp)
	fmt.Println(readjustPallette())
	setToNearest(imgTmp)
	fmt.Println(readjustPallette())
	setToNearest(imgTmp)
	fmt.Println(readjustPallette())

	// Let's set every pixel to the average of it's neighbours
	imgOut := image.NewRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Take color from the colortable
			pNode := nodeAssignments.GetGroup(x, y)
			rPixel := uint8(nodeTable[pNode].RCenter)
			gPixel := uint8(nodeTable[pNode].GCenter)
			bPixel := uint8(nodeTable[pNode].BCenter)
			pNew := color.NRGBA{rPixel, gPixel, bPixel, 0xFF}
			imgOut.Set(x, y, pNew)
		}
	}

	// Encode as PNG.
	outFile, _ := os.Create("image.png")
	png.Encode(outFile, imgOut)
}
