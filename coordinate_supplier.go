package coordinate_supplier

// CoordinateSupplier provides coordinates in a XYZ grid
type CoordinateSupplier interface {
	// Next should be called repeatedly to iterate through each pair of coordinates.
	// If done is false, the returned coordinates should be used, they are valid.
	// If done is true, the returned coordinates should be discarded and Next should not be called any longer.
	Next() (x, y, z int, done bool)
}

// CoordinateSupplierOptions control the way coordinates are handed out.
type CoordinateSupplierOptions struct {
	Width  int   // x-width of Coordinate grid
	Height int   // y-height of Coordinate grid
	Depth  int   // z-depth of Coordinate grid
	Order  Order // order that coordinates will be handed out (Asc, Desc, Random)
	Repeat bool  // if each Coordinate should be handed out exactly once, or if iterating should loop through indefinitely
}

// NewCoordinateSupplier returns the default CoordinateSupplier implementation.
func NewCoordinateSupplier(opts CoordinateSupplierOptions) (CoordinateSupplier, error) {
	return NewCoordinateSupplierAtomic(opts)
}
