package coordinate_supplier

import (
	"fmt"
	"sync/atomic"
)

type coordinateSupplierAtomic struct {
	coordinates []Coordinate
	at          uint64
	repeat      bool
	mode        NextMode
}

// NewCoordinateSupplierAtomic returns a CoordinateSupplier synchronized with atomic.AddUint64.
// It blocks less and is up to 10x faster than NewCoordinateSupplierRWMutex in concurrent usage.
// But it may be possible to receive some coordinates slightly out of order in concurrent usage.
func NewCoordinateSupplierAtomic(opts CoordinateSupplierOptions) (CoordinateSupplier, error) {
	if opts.Width < 1 {
		return nil, fmt.Errorf("minimum width is 1")
	}
	if opts.Height < 1 {
		return nil, fmt.Errorf("minimum height is 1")
	}
	coords, err := MakeCoordinateList(opts.Width, opts.Height, opts.Mode)
	if err != nil {
		return nil, fmt.Errorf("failed make coordinate list: %w", err)
	}

	cs := &coordinateSupplierAtomic{
		repeat:      opts.Repeat,
		coordinates: coords,
	}

	return cs, nil
}

func (c *coordinateSupplierAtomic) Next() (x, y int, done bool) {
	// concurrent-safe get the next value
	atNow := atomic.AddUint64(&c.at, 1) - 1

	if !c.repeat && atNow >= uint64(len(c.coordinates)) {
		return 0, 0, true
	}

	atNowClamped := atNow % uint64(len(c.coordinates))

	// by now no longer concurrent safe - but should usually be in order
	return c.coordinates[atNowClamped].X, c.coordinates[atNowClamped].Y, false
}
