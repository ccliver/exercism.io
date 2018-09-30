package triangle

import "math"

// Kind defines the type of a triangle.
type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides determines what kind of triangle was passed in from the length of the sides.
func KindFromSides(a, b, c float64) Kind {
	switch {
	case math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) ||
		math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0):
		return NaT
	case a <= 0 || b <= 0 || c <= 0:
		return NaT
	case ((a+b < c) || (b+c < a) || (a+c < b)):
		return NaT
	case a == b && b == c:
		return Equ
	case a == b || b == c || a == c:
		return Iso
	case a != b && b != c:
		return Sca
	}

	return NaT
}
