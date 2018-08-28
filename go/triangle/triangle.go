package triangle

type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

func KindFromSides(a, b, c float64) Kind {
	var k Kind

	if (a <= 0 || b <= 0 || c <= 0) || (a + b < c) || (b + c < a) || (a + c < b) {
		k = NaT
	} else if a == b && b == c {
		k = Equ
	} else if a == b || b == c || a == c {
		k = Iso
	} else if a != b && b != c {
    k = Sca
  }

	return k
}
