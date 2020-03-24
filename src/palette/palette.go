package palette

import (
	"image"

	"../vector"
)

// paletteNode represents an item in the Palette.
type paletteNode struct {
	center vector.Vector
	min    vector.Vector // Miniumum values of assigned pixels
	max    vector.Vector // Maximum values of assigned pixels
	total  vector.Vector // Running total of assigned pixel
	count  int
}

// Palette represents the group of colors
type Palette struct {
	nodes []paletteNode
}

// New creates a Palette for you
func New(img *image.RGBA) *Palette {
	var p Palette

	p.nodes = make([]paletteNode, 64)
	for i := range p.nodes {
		// No seed used, but that is OK for now
		p.nodes[i].center = vector.Rand(3).Muls(256.0)
		p.nodes[i].total = vector.Zero(3)
		p.nodes[i].count = 0
	}

	return &p
}

// Append adds a pixel from the picture to a pallete node.
func (p *Palette) Append(rImg float64, gImg float64, bImg float64) int {
	iClose := 0

	img := vector.Set([]float64{rImg, gImg, bImg})
	dClose := p.nodes[iClose].center.Sub(img).Len()
	for i := range p.nodes {
		d := p.nodes[i].center.Sub(img).Len()
		if d < dClose {
			iClose = i
			dClose = d
		}
	}

	p.nodes[iClose].total = p.nodes[iClose].total.Add(img)
	if p.nodes[iClose].count > 0 {
		p.nodes[iClose].min = p.nodes[iClose].min.Min(img)
		p.nodes[iClose].max = p.nodes[iClose].max.Max(img)
	} else {
		p.nodes[iClose].min = img
		p.nodes[iClose].max = img
	}
	p.nodes[iClose].count++

	return iClose
}

// Reset moves the colors to the center of their group
func (p *Palette) Reset() {
	// Change all centers to their average
	for i := range p.nodes {
		if p.nodes[i].count > 0 {
			p.nodes[i].center = p.nodes[i].total.Divs(float64(p.nodes[i].count))
		}
	}
	// Now we need to split 
	// Administration set to 0
	for i := range p.nodes {
		p.nodes[i].total = vector.Zero(3)
		p.nodes[i].count = 0
	}
}

// Get provides the current center color in the table
func (p *Palette) Get(i int) vector.Vector {
	return p.nodes[i].center
}
