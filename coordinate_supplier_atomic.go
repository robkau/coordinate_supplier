package coordinate_supplier

import (
	"fmt"
	"math/rand"
	"sync/atomic"
)

type coordinateSupplierAtomic struct {
	at      uint64
	width   uint64
	height  uint64
	depth   uint64
	size    uint64
	coprime uint64
	offset  uint64
	done    uint64
	repeat  bool
	order   Order
}

// NewCoordinateSupplierAtomic returns a CoordinateSupplier synchronized with atomic.AddUint64.
// It is the fastest implementation but some coordinates could be received slightly out-of-order when called concurrently.
func NewCoordinateSupplierAtomic(opts CoordinateSupplierOptions) (CoordinateSupplier, error) {
	if opts.Width < 1 {
		return nil, fmt.Errorf("minimum width is 1")
	}
	if opts.Height < 1 {
		return nil, fmt.Errorf("minimum height is 1")
	}
	if opts.Depth < 1 {
		return nil, fmt.Errorf("minimum depth is 1")
	}

	size := opts.Width * opts.Height * opts.Depth
	cs := &coordinateSupplierAtomic{
		repeat:  opts.Repeat,
		order:   opts.Order,
		width:   uint64(opts.Width),
		height:  uint64(opts.Height),
		depth:   uint64(opts.Depth),
		size:    uint64(size),
		coprime: selectCoprime(size/2, size),
		offset:  uint64(rand.Intn(size)),
	}

	return cs, nil
}

func selectCoprime(min, target int) uint64 {
	maxCount := 100000
	count := 0
	selected := 0

	for val := min; val < target; val++ {
		if coprime(val, target) {
			count += 1
			if count == 1 || rand.Intn(count) < 1 {
				selected = val
			}
		}
		if count == maxCount {
			return uint64(val)
		}
	}
	return uint64(selected)
}

func coprime(a, b int) bool {
	return gcd(a, b) == 1
}

func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	}
	return a
}

// Next returns the next coordinate to be supplied.
// It may be possible to receive some coordinates slightly out of order when called concurrently.
// https://lemire.me/blog/2017/09/18/visiting-all-values-in-an-array-exactly-once-in-random-order/
func (c *coordinateSupplierAtomic) Next() (x, y, z int, done bool) {

	if atomic.LoadUint64(&c.done) > 0 {
		// already done
		return 0, 0, 0, true
	}

	// concurrent-safe and in-order get the next index
	atNow := atomic.AddUint64(&c.at, 1) - 1

	// check if now done
	if atNow >= c.size {
		if !c.repeat {
			// mark as done
			atomic.AddUint64(&c.done, 1)
			return 0, 0, 0, true
		}
		// keep repeating
		atNow = atNow % c.size
	}

	switch c.order {
	case Desc:
		atNow = c.size - atNow - 1
	case Random:
		atNow = (atNow*c.coprime + c.offset) % c.size
	}

	return int(atNow % c.width), int((atNow / c.width) % c.height), int(atNow / (c.width * c.height)), false
}
