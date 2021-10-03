package coordinate_supplier

import (
	"fmt"
	"math/rand"
)

type Coordinate struct {
	X int
	Y int
}

// MakeCoordinateList returns a slice of Coordinate, with each item representing one cell in the XY grid.
// The NextMode determines the ordering of the slice.
func MakeCoordinateList(width, height int, mode NextMode) (cs []Coordinate, err error) {
	switch mode {
	case Asc:
		cs = makeAscCoordinates(width, height)
	case Random:
		cs = makeAscCoordinates(width, height)
		shuffleCoordinates(cs)
	case Desc:
		cs = makeAscCoordinates(width, height)
		reverseCoordinates(cs)
	default:
		err = fmt.Errorf("unknown mode specified")
	}
	return
}

func makeAscCoordinates(width, height int) []Coordinate {
	coordinates := make([]Coordinate, 0, width*height)
	var atX, atY int
	for {
		coordinates = append(coordinates, Coordinate{
			X: atX,
			Y: atY,
		})

		atX++
		if atX >= width {
			atX = 0
			atY++
		}
		if atY >= height {
			break
		}
	}
	return coordinates
}

func reverseCoordinates(cs []Coordinate) {
	i := 0
	j := len(cs) - 1
	for i < j {
		cs[i], cs[j] = cs[j], cs[i]
		i++
		j--
	}
}

func shuffleCoordinates(cs []Coordinate) {
	rand.Shuffle(len(cs), func(i, j int) { cs[i], cs[j] = cs[j], cs[i] })
}
