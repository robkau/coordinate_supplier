package coordinate_supplier

import (
	"fmt"
	"math/rand"
	"sync/atomic"
)

// CoordinateSupplier provides XY coordinates in a XY grid and is safe for concurrent usage
// ... NextMode determines the order that coordinates will be handed out (Asc, Desc, Random)
// ... repeat determines if each coordinate should be handed out once, or if iterating through should loop indefinitely
type CoordinateSupplier struct {
	coordinates []coordinate
	at          uint64
	repeat      bool
	mode        NextMode
}

func NewCoordinateSupplier(width, height int, mode NextMode, repeat bool) (*CoordinateSupplier, error) {
	if width < 1 {
		return nil, fmt.Errorf("minimum width is 1")
	}
	if height < 1 {
		return nil, fmt.Errorf("minimum height is 1")
	}

	cs := &CoordinateSupplier{
		repeat: repeat,
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
func (c *CoordinateSupplier) Next() (x, y int, done bool) {
	atNow := atomic.AddUint64(&c.at, 1) - 1

	if !c.repeat && atNow >= uint64(len(c.coordinates)) {
		return 0, 0, true
	}

	atNowClamped := atNow % uint64(len(c.coordinates))
	return c.coordinates[atNowClamped].x, c.coordinates[atNowClamped].y, false
}
