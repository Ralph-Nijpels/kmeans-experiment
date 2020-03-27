package vector

import (
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

// Vector3d implements the direct version of the vector interface for 3D vectors
type vector3d [3]float64

func make3D(f []float64) Vector {
	var r vector3d

	for i := 0; i < r.Len() && i < len(f); i++ {
		r[i] = f[i]
	}

	return r
}

func zero3D() Vector {
	var r vector3d

	return r
}

// Internal 'rand' function
func rand3D() Vector {
	var v vector3d

	for i := 0; i < v.Len(); i++ {
		v = v.Set(i, rand.Float64()).(vector3d)
	}

	return v
}

// Unit provides a vector with length 1 in the direction of the vector
func (v vector3d) Unit() Vector {
	return v.Divs(v.Abs())
}

// Abs provides the euclidian lenth of a vector
// That's how mathematicians specify the absolute value of a vector
func (v vector3d) Abs() float64 {

	l := 0.0
	for i := 0; i < v.Len(); i++ {
		l += math.Pow(v.Get(i), 2)
	}

	return math.Sqrt(l)
}

// Cbd provides the city-block-distance length of a vector
func (v vector3d) Cbd() float64 {
	var l float64

	for i := 0; i < v.Len(); i++ {
		l += math.Abs(v.Get(i))
	}

	return l
}

// Add substracts one vector from another
func (v vector3d) Add(w Vector) Vector {
	if v.Len() != w.Len() {
		log.Fatalf("Vector.Add: dimensions of vectors must be the same")
	}

	var r vector3d
	for i := 0; i < r.Len(); i++ {
		r = r.Set(i, v.Get(i)+w.Get(i)).(vector3d)
	}

	return r
}

// Sub substracts one vector from another
func (v vector3d) Sub(w Vector) Vector {
	if v.Len() != w.Len() {
		log.Fatalf("Vector.Sub: dimensions of vectors must be the same")
	}

	var r vector3d
	for i := 0; i < r.Len(); i++ {
		r = r.Set(i, v.Get(i)-w.Get(i)).(vector3d)
	}

	return r
}

// Min provides a vector containing the smallest elements of both vectors
func (v vector3d) Min(w Vector) Vector {
	if v.Len() != w.Len() {
		log.Fatalf("Vector.Min: dimensions of vectors must be the same")
	}

	var r vector3d
	for i := 0; i < r.Len(); i++ {
		r = r.Set(i, math.Min(v.Get(i), w.Get(i))).(vector3d)
	}

	return r
}

// MinD provides the index of the smallest value in the vector
func (v vector3d) MinD() int {

	r := 0
	for i := 0; i < v.Len(); i++ {
		if v.Get(i) < v.Get(r) {
			r = i
		}
	}

	return r
}

// Max set every element of the resulting vector to the highest option
// It provides the first occurence if there are more dimensions with this value
func (v vector3d) Max(w Vector) Vector {
	if v.Len() != w.Len() {
		log.Fatalf("Vector.Max: dimensions of vectors must be the same")
	}

	var r vector3d
	for i := 0; i < r.Len(); i++ {
		r = r.Set(i, math.Max(v.Get(i), w.Get(i))).(vector3d)
	}

	return r
}

// MaxD provides the index of the largest value in the vector
// It provides the first occurence if there are more dimensions with this value
func (v vector3d) MaxD() int {

	r := 0
	for i := 0; i < v.Len(); i++ {
		if v.Get(i) > v.Get(r) {
			r = i
		}
	}

	return r
}

// Muls multiplies a vector by a scalar.
func (v vector3d) Muls(s float64) Vector {

	var r vector3d
	for i := 0; i < r.Len(); i++ {
		r = r.Set(i, v.Get(i)*s).(vector3d)
	}

	return r
}

// Divs divides a vector by a scalar.
func (v vector3d) Divs(s float64) Vector {

	var r vector3d
	for i := 0; i < r.Len(); i++ {
		r = r.Set(i, v.Get(i)/s).(vector3d)
	}

	return r
}

// Len retrieves the number of dimensions
func (v vector3d) Len() int {
	return 3
}

// Raw retrieves the values of the vector as a slice
func (v vector3d) Raw() []float64 {
	// silly?
	r := make([]float64, v.Len())
	for i := 0; i < v.Len(); i++ {
		r[i] = v.Get(i)
	}

	return r
}

// Get retrieves the value of a single cell
func (v vector3d) Get(i int) float64 {
	if (i < 0) || (i >= 3) {
		log.Fatalf("Vector.Get: Index out of bounds")
	}

	return v[i]
}

// Set changes the value of a single cell
func (v vector3d) Set(i int, f float64) Vector {
	if (i < 0) || (i >= 3) {
		log.Fatalf("Vector.Set: Index out of bounds")
	}
	v[i] = f
	return v
}

// String() implements the Stringer interface
func (v vector3d) String() string {
	var s strings.Builder

	s.WriteString("[")
	for i := 0; i < v.Len(); i++ {
		if i > 0 {
			s.WriteString(", ")
		}
		s.WriteString(strconv.FormatFloat(v.Get(i), 'f', 3, 64))
	}
	s.WriteString("]")

	return s.String()
}
