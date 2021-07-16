package main

import (
	"image"

	"github.com/ralph-nijpels/vector"
)

// Nodemap contains the assignment of each pixel to a color in the colortable
type Nodemap struct {
	bounds      image.Rectangle
	assignments []int
}

// New creates a nodemap for all pixels in your image
func CreateNodemap(img *image.RGBA, p *Palette) *Nodemap {

	bounds := img.Bounds()
	nmap := Nodemap{bounds, make([]int, bounds.Dy()*bounds.Dx())}

	for py := 0; py < bounds.Dy(); py++ {
		for px := 0; px < bounds.Dx(); px++ {

			pOffset := img.PixOffset(
				bounds.Min.X+px,
				bounds.Min.Y+py)

			pVector := vector.Make([]float64{
				float64(img.Pix[pOffset]),
				float64(img.Pix[pOffset+1]),
				float64(img.Pix[pOffset+2])})

			iClose := 0
			dClose := p.Get(0).Sub(pVector).Abs()
			for c := 1; c < p.Len(); c++ {
				distance := p.Get(c).Sub(pVector).Abs()
				if distance < dClose {
					dClose = distance
					iClose = c
				}
			}

			nmap.assignments[py*bounds.Dx()+px] = iClose
		}
	}

	return &nmap
}

// GetGroup provides the assigned group number of a pixel at (x,y)
func (nmap *Nodemap) GetGroup(x, y int) int {
	iWidth := nmap.bounds.Dx()
	return nmap.assignments[y*iWidth+x]
}
