package main

import (
	"fmt"
	"image"
	"sort"
	"strings"

	"github.com/ralph-nijpels/vector"
)

// node represents an item in the Palette.
type node struct {
	center vector.Vector // center is the current color
	min    vector.Vector // Miniumum values of assigned pixels
	max    vector.Vector // Maximum values of assigned pixels
	total  vector.Vector // Running total of assigned pixel
	count  int
}

// Pallette nodes can be sorted by count
type byCount []node

func (n byCount) Len() int {
	return len(n)
}

func (n byCount) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n byCount) Less(i, j int) bool {
	return n[i].count < n[j].count
}

// Palette represents the group of colors
type Palette struct {
	nodes []node
}

// New creates a Palette for you
func CreatePalette(size int) *Palette {
	var p Palette

	p.nodes = make([]node, size)
	for i := range p.nodes {
		// No seed used, but that is OK for now
		p.nodes[i].center = vector.Rand(3).Muls(256.0)
		p.nodes[i].total = vector.Zero(3)
		p.nodes[i].min = vector.Zero(3)
		p.nodes[i].max = vector.Zero(3)
		p.nodes[i].count = 0
	}

	return &p
}

// reset clears the administrative part of the palette
func (p *Palette) reset() {
	// Administration set to 0
	for i := range p.nodes {
		p.nodes[i].min = vector.Zero(3)
		p.nodes[i].max = vector.Zero(3)
		p.nodes[i].total = vector.Zero(3)
		p.nodes[i].count = 0
	}
}

// appendPixel adds a pixel from the picture to a pallete node.
func (p *Palette) analysePixel(pix vector.Vector) {

	// Find the closest palette entry
	iClose := 0
	dClose := p.nodes[iClose].center.Sub(pix).Abs()
	for i := range p.nodes {
		d := p.nodes[i].center.Sub(pix).Abs()
		if d < dClose {
			iClose = i
			dClose = d
		}
	}

	// Keep a running bounded box of the color area covered
	// to speed up the 'split' operations
	if p.nodes[iClose].count > 0 {
		p.nodes[iClose].min = p.nodes[iClose].min.Min(pix)
		p.nodes[iClose].max = p.nodes[iClose].max.Max(pix)
	} else {
		p.nodes[iClose].min = pix
		p.nodes[iClose].max = pix
	}

	// Keep a running total for all pixels added to
	// speed up the 'shift' operation
	p.nodes[iClose].total = p.nodes[iClose].total.Add(pix)
	p.nodes[iClose].count++
}

// analyse maps the pixels of the image to the palette entries
func (p *Palette) analyse(img *image.RGBA) {
	p.reset()
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			pOffset := img.PixOffset(x, y)
			pVector := vector.Make([]float64{
				float64(img.Pix[pOffset]),
				float64(img.Pix[pOffset+1]),
				float64(img.Pix[pOffset+2])})
			p.analysePixel(pVector)
		}
	}
}

// Shift moves the colors to the center of their group
// There is a magic number, currently set at '1.0', to determine what counts as a shift. This to prevent the
// algorithm to start 'hunting' when pixels shift from one side to the other and back (again...)
func (p *Palette) shift() int {
	shifts := 0
	for i := range p.nodes {
		if p.nodes[i].count > 0 {
			newCenter := p.nodes[i].total.Divs(float64(p.nodes[i].count))
			if p.nodes[i].center.Sub(newCenter).Abs() > 1.0 {
				p.nodes[i].center = newCenter
				shifts++
			}
		}
	}
	return shifts
}

// Split splits the overpopulated nodes and splits them
// Perform a kind of butterfly scan over the nodes. Split by a relatively small value (1% of the colorspace) to prevent
// stealing from other groups. If the colorspace doesn't leave room for splitting, we are going for a next kandidate.
// The minimum size of the colorspace needed is set to '3.0'
func (p *Palette) split() int {

	splits := 0

	sort.Sort(byCount(p.nodes))
	b, t := 0, len(p.nodes)-1
	for p.nodes[b].count*10 < p.nodes[t].count {
		colorSpace := p.nodes[t].max.Sub(p.nodes[t].min)
		if colorSpace.Abs() > 3.0 {
			splitVector := colorSpace.Unit().Muls(0.01)
			p.nodes[b].center = p.nodes[t].center.Sub(splitVector)
			p.nodes[t].center = p.nodes[t].center.Add(splitVector)
			splits++
			b++
		}
		t--
	}

	return splits
}

// SetFromImage determines the colors needed to best represent the picture by
// applying a k-nearest algorithm where k is determined by the size of the palette.
// It will run until no more shifts and splits are needed.
func (p *Palette) SetFromImage(img *image.RGBA) error {

	round := 0
	p.analyse(img)
	shifts := p.shift()
	splits := p.split()
	for (shifts > 0 || splits > 0) && (round < 5) {
		p.analyse(img)
		shifts = p.shift()
		splits = p.split()
		round++
	}

	return nil
}

// String provides a printable version of the current content of the palette
func (p *Palette) String() string {
	var s strings.Builder

	for i := range p.nodes {
		s.WriteString(fmt.Sprintf("%02d: %06d %v\n", i, p.nodes[i].count, p.nodes[i].center))
	}

	return s.String()
}

// Len provides the number of colors in the palette
func (p *Palette) Len() int {
	return len(p.nodes)
}

// Get provides the current center color in the palette
func (p *Palette) Get(i int) vector.Vector {
	return p.nodes[i].center
}
