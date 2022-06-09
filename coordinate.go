package coordinate_supplier

import (
	"fmt"
	"math/rand"
)

type Coordinate struct {
	X int
	Y int
	Z int
}

// MakeCoordinateList returns a slice of Coordinate, with each item representing one cell in the XY grid.
// The Order determines the ordering of the coordinates in the slice.
func MakeCoordinateList(width, height, depth int, order Order) (cs []Coordinate, err error) {
	switch order {
	case Asc:
		cs = makeAscCoordinates(width, height, depth)
	case Random:
		cs = makeAscCoordinates(width, height, depth)
		shuffleCoordinates(cs)
	case Desc:
		cs = makeAscCoordinates(width, height, depth)
		reverseCoordinates(cs)
	default:
		err = fmt.Errorf("unknown order specified")
	}
	return
}

func makeAscCoordinates(width, height, depth int) []Coordinate {
	coordinates := make([]Coordinate, 0, width*height*depth)
	var atX, atY, atZ int
	for {
		coordinates = append(coordinates, Coordinate{
			X: atX,
			Y: atY,
			Z: atZ,
		})

		atX++
		if atX >= width {
			atX = 0
			atY++
		}
		if atY >= height {
			atY = 0
			atZ++
		}
		if atZ >= depth {
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
