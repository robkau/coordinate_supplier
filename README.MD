# coordinate_supplier

[![Go Reference](https://pkg.go.dev/badge/github.com/robkau/coordinate_supplier.svg)](https://pkg.go.dev/github.com/robkau/coordinate_supplier)

coordinate_supplier is a utility Go package to supply coordinates in a XY grid.

```
opts := CoordinateSupplierOptions{Width: 10, Height: 10, Order: Asc, Repeat: false}
cs, err := NewCoordinateSupplier(opts)
if err != nil {
  // handle me
}

for x, y, done := cs.Next(); !done; x, y, done = cs.Next() {
    fmt.Println("The next coordinate is", x, y)
}
```


----
## Features:
 - Hand out each coordinate exactly once, or loop through all coordinates repeatedly
 - Hand out coordinates in ascending order, descending order, or random order
 - Fast and mostly concurrent-safe (no data races) implementation via atomic.AddUint64 
 - Strictly concurrent-safe (guaranteed in order) implementation via sync.RWMutex
----