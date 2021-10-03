package coordinate_supplier

import (
	"fmt"
	"math/rand"
	"sync"
)

// CoordinateSupplier provides XY coordinates in a 2x2 grid and is safe for concurrent usage
// ... NextMode determines the order that coordinates will be handed out (Asc, Desc, Random)
// ... repeat determines if each coordinate should be handed out once, or if iterating through should loop indefinitely
type coordinateSupplier struct {
	coordinates []coordinate
	at          int
	repeat      bool
	mode        NextMode

	rw sync.RWMutex
}

func NewCoordinateSupplier(width, height int, mode NextMode, repeat bool) (*coordinateSupplier, error) {
	if width < 1 {
		return nil, fmt.Errorf("minimum width is 1")
	}
	if height < 1 {
		return nil, fmt.Errorf("minimum height is 1")
	}

	cs := &coordinateSupplier{
		repeat: repeat,
		rw:     sync.RWMutex{},
	}

	switch mode {
	case Asc:
		cs.coordinates = makeAscCoordinates(width, height)
	case Random:
		cs.coordinates = makeAscCoordinates(width, height)
		rand.Shuffle(len(cs.coordinates), func(i, j int) { cs.coordinates[i], cs.coordinates[j] = cs.coordinates[j], cs.coordinates[i] })
	case Desc:
		cs.coordinates = makeAscCoordinates(width, height)
		reverseCoordinates(cs.coordinates)
	default:
		return nil, fmt.Errorf("unknown mode specified")
	}

	return cs, nil
}

// Next iterates through each pair of coordinates
// If done is false, the returned coordinates should be used, they are valid.
// If done is true, the returned coordinates should be discarded, there were none left to use.
func (c *coordinateSupplier) Next() (x, y int, done bool) {
	c.rw.Lock()
	defer c.rw.Unlock()

	if c.at >= len(c.coordinates) {
		if c.repeat {
			c.at = 0
		} else {
			return 0, 0, true
		}
	}

	defer func() { c.at++ }()
	return c.coordinates[c.at].x, c.coordinates[c.at].y, false
}
