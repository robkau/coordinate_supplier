package coordinate_supplier

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"sync"
	"sync/atomic"
	"testing"
)

var suppliersToTest = []struct {
	name string
	new  func(options CoordinateSupplierOptions) (CoordinateSupplier, error)
}{
	{"atomic", NewCoordinateSupplierAtomic},
	{"rw", NewCoordinateSupplierRWMutex},
}

func Test_Coordinate_Supplier_Asc_10x1(t *testing.T) {
	testOpts := CoordinateSupplierOptions{10, 1, Asc, false}
	// test the ones behind CoordinateSupplier interface
	for _, supplier := range suppliersToTest {
		t.Run(supplier.name, func(t *testing.T) {
			cs, err := supplier.new(testOpts)
			require.NoError(t, err)

			seen := 0
			last := -1
			for x, y, done := cs.Next(); !done; x, y, done = cs.Next() {
				require.Equal(t, 0, y)
				require.Greater(t, x, last)
				last = x
				seen++
			}
			require.Equal(t, 10, seen)
		})
	}
}

func Test_Coordinate_Supplier_Asc_1x10(t *testing.T) {
	testOpts := CoordinateSupplierOptions{1, 10, Asc, false}
	// test the ones behind CoordinateSupplier interface
	for _, supplier := range suppliersToTest {
		t.Run(supplier.name, func(t *testing.T) {
			cs, err := supplier.new(testOpts)
			require.NoError(t, err)

			seen := 0
			last := -1
			for x, y, done := cs.Next(); !done; x, y, done = cs.Next() {
				require.Equal(t, 0, x)
				require.Greater(t, y, last)
				last = y
				seen++
			}
			require.Equal(t, 10, seen)
		})
	}
}

func Test_Coordinate_Supplier_Asc_2x2(t *testing.T) {
	testOpts := CoordinateSupplierOptions{2, 2, Asc, false}
	// test the ones behind CoordinateSupplier interface
	for _, supplier := range suppliersToTest {
		t.Run(supplier.name, func(t *testing.T) {
			cs, err := supplier.new(testOpts)
			require.NoError(t, err)

			x1, y1, done1 := cs.Next()
			x2, y2, done2 := cs.Next()
			x3, y3, done3 := cs.Next()
			x4, y4, done4 := cs.Next()
			_, _, done5 := cs.Next()

			require.Equal(t, 0, x1)
			require.Equal(t, 1, x2)
			require.Equal(t, 0, x3)
			require.Equal(t, 1, x4)

			require.Equal(t, 0, y1)
			require.Equal(t, 0, y2)
			require.Equal(t, 1, y3)
			require.Equal(t, 1, y4)

			require.False(t, done1)
			require.False(t, done2)
			require.False(t, done3)
			require.False(t, done4)
			require.True(t, done5)
		})
	}
}

func Test_Coordinate_Supplier_Desc_2x2(t *testing.T) {
	testOpts := CoordinateSupplierOptions{2, 2, Desc, false}
	// test the ones behind CoordinateSupplier interface
	for _, supplier := range suppliersToTest {
		t.Run(supplier.name, func(t *testing.T) {
			cs, err := supplier.new(testOpts)
			require.NoError(t, err)

			x1, y1, done1 := cs.Next()
			x2, y2, done2 := cs.Next()
			x3, y3, done3 := cs.Next()
			x4, y4, done4 := cs.Next()
			_, _, done5 := cs.Next()

			require.Equal(t, 1, x1)
			require.Equal(t, 0, x2)
			require.Equal(t, 1, x3)
			require.Equal(t, 0, x4)

			require.Equal(t, 1, y1)
			require.Equal(t, 1, y2)
			require.Equal(t, 0, y3)
			require.Equal(t, 0, y4)

			require.False(t, done1)
			require.False(t, done2)
			require.False(t, done3)
			require.False(t, done4)
			require.True(t, done5)
		})
	}
}

func Test_Coordinate_Supplier_Asc_3x2_Repeat(t *testing.T) {
	testOpts := CoordinateSupplierOptions{3, 2, Asc, true}
	xPattern := []int{0, 1, 2, 0, 1, 2}
	yPattern := []int{0, 0, 0, 1, 1, 1}
	// test the ones behind CoordinateSupplier interface
	for _, supplier := range suppliersToTest {
		t.Run(supplier.name, func(t *testing.T) {
			cs, err := supplier.new(testOpts)
			require.NoError(t, err)
			for seen := 0; seen < 1000; seen++ {
				x, y, done := cs.Next()
				require.False(t, done)
				require.Equal(t, xPattern[seen%len(xPattern)], x)
				require.Equal(t, yPattern[seen%len(yPattern)], y)
			}
		})
	}
}

func Test_Coordinate_Supplier_Desc_3x2_Repeat(t *testing.T) {
	testOpts := CoordinateSupplierOptions{3, 2, Desc, true}
	xPattern := []int{2, 1, 0, 2, 1, 0}
	yPattern := []int{1, 1, 1, 0, 0, 0}
	// test the ones behind CoordinateSupplier interface
	for _, supplier := range suppliersToTest {
		t.Run(supplier.name, func(t *testing.T) {
			cs, err := supplier.new(testOpts)
			require.NoError(t, err)
			for seen := 0; seen < 1000; seen++ {
				x, y, done := cs.Next()
				require.False(t, done)
				require.Equal(t, xPattern[seen%len(xPattern)], x)
				require.Equal(t, yPattern[seen%len(yPattern)], y)
			}
		})
	}
}

func Test_Coordinate_Supplier_Desc_2x2_Repeat(t *testing.T) {
	testOpts := CoordinateSupplierOptions{2, 2, Desc, true}
	xPattern := []int{1, 0, 1, 0}
	yPattern := []int{1, 1, 0, 0}
	// test the ones behind CoordinateSupplier interface
	for _, supplier := range suppliersToTest {
		t.Run(supplier.name, func(t *testing.T) {
			cs, err := supplier.new(testOpts)
			require.NoError(t, err)
			for seen := 0; seen < 1000; seen++ {
				x, y, done := cs.Next()
				require.False(t, done)
				require.Equal(t, xPattern[seen%len(xPattern)], x)
				require.Equal(t, yPattern[seen%len(yPattern)], y)
			}
		})
	}
}

func Test_Coordinate_Supplier_Asc_1000x1000_Concurrent(t *testing.T) {
	testOpts := CoordinateSupplierOptions{1000, 1000, Asc, false}
	// test the ones behind CoordinateSupplier interface
	for _, supplier := range suppliersToTest {
		t.Run(supplier.name, func(t *testing.T) {
			cs, err := supplier.new(testOpts)
			require.NoError(t, err)
			consumed := runCoordinateSupplier(cs, 10, 0)
			require.Equal(t, uint64(testOpts.Width*testOpts.Height), consumed)
		})
	}
}

func Test_ConsumePastEnd(t *testing.T) {
	testOpts := CoordinateSupplierOptions{100, 100, Asc, false}
	// test the ones behind CoordinateSupplier interface
	for _, supplier := range suppliersToTest {
		t.Run(supplier.name, func(t *testing.T) {
			// consume the coordinates
			cs, err := supplier.new(testOpts)
			require.NoError(t, err)
			for _, _, done := cs.Next(); !done; _, _, done = cs.Next() {
				require.False(t, done)
			}

			// get a bunch past the end... it should still be done
			extras := 0
			for {
				if extras > 1000000 {
					break
				}
				extras++

				_, _, done := cs.Next()
				require.True(t, done)
			}
		})
	}
}

func Test_Readme_Example(t *testing.T) {
	opts := CoordinateSupplierOptions{Width: 10, Height: 10, Order: Asc, Repeat: false}
	cs, err := NewCoordinateSupplier(opts)
	if err != nil {
		require.NoError(t, err)
	}

	for x, y, done := cs.Next(); !done; x, y, done = cs.Next() {
		fmt.Println("The next coordinate is", x, y)
	}
}

func BenchmarkCoordinateSuppliers(b *testing.B) {
	upToWidth := 1000
	upToHeight := 1000
	upToConsumers := 1000
	upToConsumed := 1000000

	for width := 1; width <= upToWidth; width *= 30 {
		for height := 1; height <= upToHeight; height *= 30 {
			for consume := 1; consume <= upToConsumed; consume *= 1000000 {
				for consumers := 1; consumers <= upToConsumers; consumers *= 10 {
					// run CoordinateSuppliers
					for _, supplier := range suppliersToTest {
						b.Run(fmt.Sprintf("%s-%dw-%dh-%dconsumers-consume%d", supplier.name, width, height, consumers, consume), func(b *testing.B) {
							for i := 0; i < b.N; i++ {
								// special case if consume == 1, then consume all coordinates once
								// otherwise, loop through coordinates on repeat until upToConsumed
								useConsume := consume
								if consume == 1 {
									useConsume = 0
								}
								var repeat bool
								if useConsume > 0 {
									// instead of consuming once, will loop until upToConsumed
									repeat = true
								}
								cs, err := supplier.new(CoordinateSupplierOptions{width, height, Asc, repeat})
								require.NoError(b, err)

								count := runCoordinateSupplier(cs, consumers, uint64(useConsume))
								if useConsume == 0 {
									require.Equal(b, uint64(width*height), count)
								} else {
									require.Equal(b, uint64(useConsume), count)
								}
							}
						})
					}
				}
			}
		}
	}
}

func runCoordinateSupplier(cs CoordinateSupplier, numConsumers int, maxConsumed uint64) (consumed uint64) {
	// run consumers to get all coordinates
	wg := sync.WaitGroup{}
	var requested uint64
	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, _, done := cs.Next(); !done; _, _, done = cs.Next() {
				// if on repeat, break when reach max consumed limit
				if maxConsumed != 0 {
					now := atomic.AddUint64(&requested, 1)
					if now > maxConsumed {
						return
					}
				}
				atomic.AddUint64(&consumed, 1)
			}
		}()
	}
	wg.Wait()
	return
}
