package nodemap

import "image"

// Nodemap contains the assignment of each pixel to a color in the colortable
type Nodemap struct {
	Assignments []uint16
	Top         int
	Left        int
	Bottom      int
	Right       int
}

// New creates a nodemap for all pixels in your image
func New(img *image.RGBA) *Nodemap {
	top := img.Bounds().Min.Y
	left := img.Bounds().Min.X
	bottom := img.Bounds().Max.Y
	right := img.Bounds().Max.X
	return &Nodemap{make([]uint16, (bottom-top)*(right-left)), top, left, bottom, right}
}

// GetGroup provides the assigned group number of a pixel at (x,y)
func (nmap *Nodemap) GetGroup(x, y int) int {
	return int(nmap.Assignments[(y-nmap.Top)*(nmap.Right-nmap.Left)+x])
}

// SetGroup assigns group v to a pixel at (x,y)
func (nmap *Nodemap) SetGroup(x, y int, v int) {
	nmap.Assignments[(y-nmap.Top)*(nmap.Right-nmap.Left)+x] = uint16(v)
}
