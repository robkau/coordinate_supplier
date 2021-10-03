package coordinate_supplier

import (
	"context"
	"fmt"
)

const chanBufferSize = 10

// NewCoordinateSupplierChan performs the same as CoordinateSupplier interface with a few differences.
// 1: To get the supply of coordinates, you should range over the returned channel. Each returned value is a valid coordinate to use.
// 2: You should either fully consume the channel or cancel the context when done reading values from the channel.
func NewCoordinateSupplierChan(ctx context.Context, opts CoordinateSupplierOptions) (<-chan Coordinate, error) {
	if opts.Width < 1 {
		return nil, fmt.Errorf("minimum width is 1")
	}
	if opts.Height < 1 {
		return nil, fmt.Errorf("minimum height is 1")
	}
	coordinates, err := MakeCoordinateList(opts.Width, opts.Height, opts.Mode)
	if err != nil {
		return nil, fmt.Errorf("failed make coordinate list: %w", err)
	}

	coordChan := make(chan Coordinate, chanBufferSize)
	go func() {
		at := 0
		defer close(coordChan)
		for {
			// get current coordinate
			if !opts.Repeat && at >= len(coordinates) {
				// finished all values and not on repeat
				return
			}
			atClamped := at % len(coordinates)

			// send current coordinate
			select {
			case coordChan <- coordinates[atClamped]:
				// sent
				at++
			case <-ctx.Done():
				// cancelled
				return
			}
		}
	}()

	return coordChan, nil
}
