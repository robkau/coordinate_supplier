# coordinate_supplier

[![Go Reference](https://pkg.go.dev/badge/github.com/robkau/coordinate_supplier.svg)](https://pkg.go.dev/github.com/robkau/coordinate_supplier)

coordinate_supplier is a utility Go package to supply coordinates in a XY grid of any size.

----
## Features:
 - hand out each coordinate exactly once, or loop through all coordinates repeatedly
 - hand out coordinates in ascending order, descending order, or random order
 - Concurrent-safe implementations via sync.RWMutex, atomic.AddUint64, or channel
----