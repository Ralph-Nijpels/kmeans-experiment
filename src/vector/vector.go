package vector

import (
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

// Vector implements a simple mathematical vector
type Vector struct {
	cells []float64
}

// Make creates a vector based on a list of values
func Make(f []float64) Vector {
	var v Vector

	v.cells = make([]float64, len(f))
	for i := range f {
		v.cells[i] = f[i]
	}

	return v
}

// Zero creates a vector of the requested size set to the origin
func Zero(dimension int) Vector {
	var v Vector

	v.cells = make([]float64, dimension)

	return v
}

// Rand creates a vector of the requested size set to a random location
func Rand(dimension int) Vector {
	var v Vector

	v.cells = make([]float64, dimension)
	for i := range v.cells {
		v.cells[i] = rand.Float64()
	}

	return v
}

// Abs provides the euclidian lenth of a vector
// That's how mathematicians specify the absolute value of a vector
func (v Vector) Abs() float64 {

	l := 0.0
	for i := range v.cells {
		l += math.Pow(v.cells[i], 2)
	}

	return math.Sqrt(l)
}

// Cbd provides the city-block-distance length of a vector
func (v Vector) Cbd() float64 {
	var l float64

	for i := range v.cells {
		l += math.Abs(v.cells[i])
	}

	return l
}

// Add substracts one vector from another
func (v Vector) Add(w Vector) Vector {
	if len(v.cells) != len(w.cells) {
		log.Fatalf("Vector.Add: dimensions of vectors must be the same")
	}

	var r Vector
	r.cells = make([]float64, len(v.cells))
	for i := range r.cells {
		r.cells[i] = v.cells[i] + w.cells[i]
	}
	return r
}

// Sub substracts one vector from another
func (v Vector) Sub(w Vector) Vector {
	if len(v.cells) != len(w.cells) {
		log.Fatalf("Vector.Sub: dimensions of vectors must be the same")
	}

	var r Vector
	r.cells = make([]float64, len(v.cells))
	for i := range r.cells {
		r.cells[i] = v.cells[i] - w.cells[i]
	}
	return r
}

// Min provides a vector containing the smallest elements of both vectors
func (v Vector) Min(w Vector) Vector {
	if len(v.cells) != len(w.cells) {
		log.Fatalf("Vector.Min: dimensions of vectors must be the same")
	}

	var r Vector

	r.cells = make([]float64, len(v.cells))
	for i := range r.cells {
		r.cells[i] = math.Min(v.cells[i], w.cells[i])
	}

	return r
}

// MinD provides the index of the smallest value in the vector
func (v Vector) MinD() int {
	r := 0

	for i := range v.cells {
		if v.cells[i] < v.cells[r] {
			r = i
		}
	}

	return r
}

// Max set every element of the resulting vector to the highest option
// It provides the first occurence if there are more dimensions with this value
func (v Vector) Max(w Vector) Vector {
	if len(v.cells) != len(w.cells) {
		log.Fatalf("Vector.Max: dimensions of vectors must be the same")
	}

	var r Vector

	r.cells = make([]float64, len(v.cells))
	for i := range r.cells {
		r.cells[i] = math.Max(v.cells[i], w.cells[i])
	}

	return r
}

// Unit provides a vector with length 1 in the direction of the vector
func (v Vector) Unit() Vector {
	return v.Divs(v.Abs())
}

// MaxD provides the index of the largest value in the vector
// It provides the first occurence if there are more dimensions with this value
func (v Vector) MaxD() int {
	r := 0

	for i := range v.cells {
		if v.cells[i] > v.cells[r] {
			r = i
		}
	}

	return r
}

// Muls multiplies a vector by a scalar.
func (v Vector) Muls(s float64) Vector {
	var r Vector
	r.cells = make([]float64, len(v.cells))
	for i := range r.cells {
		r.cells[i] = v.cells[i] * s
	}
	return r
}

// Divs divides a vector by a scalar.
func (v Vector) Divs(s float64) Vector {
	var r Vector

	r.cells = make([]float64, len(v.cells))
	for i := range r.cells {
		r.cells[i] = v.cells[i] / s
	}

	return r
}

// Raw retrieves the values of the vector as a slice
func (v Vector) Raw() []float64 {
	return v.cells
}

// Get retrieves the value of a single cell
func (v Vector) Get(i int) float64 {
	if (i < 0) || (i >= len(v.cells)) {
		log.Fatalf("Vector.Get: Index out of bounds")
	}

	return v.cells[i]
}

// Set changes the value of a single cell
func (v Vector) Set(i int, f float64) {
	if (i < 0) || (i >= len(v.cells)) {
		log.Fatalf("Vector.Get: Index out of bounds")
	}

	v.cells[i] = f
}

// String() implements the Stringer interface
func (v Vector) String() string {
	var s strings.Builder

	s.WriteString("[")
	for i, f := range v.cells {
		if i > 0 {
			s.WriteString(", ")
		}
		s.WriteString(strconv.FormatFloat(f, 'f', 3, 64))
	}
	s.WriteString("]")

	return s.String()
}
