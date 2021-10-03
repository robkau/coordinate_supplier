# coordinate_supplier

[![Go Reference](https://pkg.go.dev/badge/github.com/robkau/coordinate_supplier.svg)](https://pkg.go.dev/github.com/robkau/coordinate_supplier)

coordinate_supplier is a utility Go package to supply coordinates in a XY grid of any size.

```
// first initialize a coordinate supplier with given size and options
func NewCoordinateSupplier(width, height int, mode NextMode, repeat bool) (*coordinateSupplier, error)

// then iterate through the provided coordinates until done
func (c *coordinateSupplier) Next() (x, y int, done bool) {
```

##Features:
 - hand out each coordinate exactly once, or loop through all coordinates repeatedly
 - hand out coordinates in ascending order, descending order, or random order
 - safe for concurrent use between goroutines


  
See unit tests for examples.