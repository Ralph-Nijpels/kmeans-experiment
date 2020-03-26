package palette

import (
	"fmt"
	"image"
	"log"
	"math"
	"sort"
	"strings"

	"../vector"
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
func New(img *image.RGBA) *Palette {
	var p Palette

	p.nodes = make([]node, 64)
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

// Append adds a pixel from the picture to a pallete node.
func (p *Palette) Append(rImg float64, gImg float64, bImg float64) int {
	iClose := 0

	img := vector.Make([]float64{rImg, gImg, bImg})
	dClose := p.nodes[iClose].center.Sub(img).Abs()
	if math.IsNaN(dClose) {
		log.Panicf("Abs of %v - %v is NaN", p.nodes[iClose].center, img)
	}
	for i := range p.nodes {
		d := p.nodes[i].center.Sub(img).Abs()
		if math.IsNaN(d) {
			log.Panicf("Abs of %v - %v is NaN", p.nodes[i].center, img)
		}
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

// Shift moves the colors to the center of their group
func (p *Palette) Shift() {
	for i := range p.nodes {
		if p.nodes[i].count > 0 {
			p.nodes[i].center = p.nodes[i].total.Divs(float64(p.nodes[i].count))
		}
	}
}

// Split splits the overpopulated nodes and splits them
func (p *Palette) Split() {
	sort.Sort(byCount(p.nodes))

	// Perform a kind of butterfly scan over the nodes
	// Split by a relatively small value (1% of the colorspace) to prevent stealing from other groups
	// If the colorspace doesn't leave room for splitting, we are going for a next kandidate.
	b, t := 0, len(p.nodes)-1
	for p.nodes[b].count*3 < p.nodes[t].count {
		colorSpace := p.nodes[t].max.Sub(p.nodes[t].min)
		if colorSpace.Abs() != 0.0 {
			splitVector := colorSpace.Unit().Muls(0.01)
			if math.IsNaN(splitVector.Abs()) {
				log.Panicf("SplitVector %v in %v(%v,%v) leads to NaN", splitVector, colorSpace, p.nodes[t].min, p.nodes[t].max)
			}
			p.nodes[b].center = p.nodes[t].center.Sub(splitVector)
			p.nodes[t].center = p.nodes[t].center.Add(splitVector)
			b++
		}
		t--
	}
}

// Reset moves the colors to the center of their group
func (p *Palette) Reset() {
	// Administration set to 0
	for i := range p.nodes {
		p.nodes[i].total = vector.Zero(3)
		p.nodes[i].count = 0
	}
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
