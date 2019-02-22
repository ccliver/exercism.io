package robotname

import (
	"math/rand"
	"strconv"
	"time"
)

// Robot defines a robot with a name.
type Robot struct {
	name string
}

// Name generates a random name for a Robot.
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		bytes := make([]byte, 2)
		ints := ""
		rand.Seed(time.Now().UTC().UnixNano())

		for i := 0; i < 2; i++ {
			bytes[i] = byte(65 + rand.Intn(25)) //A=65 and Z = 65+25
		}
		for i := 0; i < 3; i++ {
			ints += strconv.Itoa(rand.Intn(9))
		}

		name := string(bytes) + ints
		r.name = name
	}

	return r.name, nil
}

// Reset blanks a Robot's name.
func (r *Robot) Reset() {
    r.name = ""
}
