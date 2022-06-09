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
	//{"rw", NewCoordinateSupplierRWMutex},
}

func Test_Coordinate_Supplier_Asc_10x1x1(t *testing.T) {
	testOpts := CoordinateSupplierOptions{10, 1, 1, Asc, false}
	// test the ones behind CoordinateSupplier interface
	for _, supplier := range suppliersToTest {
		t.Run(supplier.name, func(t *testing.T) {
			cs, err := supplier.new(testOpts)
			require.NoError(t, err)

			seen := 0
			last := -1
			for x, y, z, done := cs.Next(); !done; x, y, z, done = cs.Next() {
				require.Equal(t, 0, y)
				require.Equal(t, 0, z)
				require.Greater(t, x, last)
				last = x
				seen++
			}
			require.Equal(t, 10, seen)
		})
	}
}

func Test_Coordinate_Supplier_Asc_1x10x1(t *testing.T) {
	testOpts := CoordinateSupplierOptions{1, 10, 1, Asc, false}
	// test the ones behind CoordinateSupplier interface
	for _, supplier := range suppliersToTest {
		t.Run(supplier.name, func(t *testing.T) {
			cs, err := supplier.new(testOpts)
			require.NoError(t, err)

			seen := 0
			last := -1
			for x, y, z, done := cs.Next(); !done; x, y, z, done = cs.Next() {
				require.Equal(t, 0, x)
				require.Equal(t, 0, z)
				require.Greater(t, y, last)
				last = y
				seen++
			}
			require.Equal(t, 10, seen)
		})
	}
}

func Test_Coordinate_Supplier_Asc_1x1x10x(t *testing.T) {
	testOpts := CoordinateSupplierOptions{1, 1, 10, Asc, false}
	// test the ones behind CoordinateSupplier interface
	for _, supplier := range suppliersToTest {
		t.Run(supplier.name, func(t *testing.T) {
			cs, err := supplier.new(testOpts)
			require.NoError(t, err)

			seen := 0
			last := -1
			for x, y, z, done := cs.Next(); !done; x, y, z, done = cs.Next() {
				require.Equal(t, 0, x)
				require.Equal(t, 0, y)
				require.Greater(t, z, last)
				last = y
				seen++
			}
			require.Equal(t, 10, seen)
		})
	}
}

func Test_Coordinate_Supplier_Asc_2x2x2(t *testing.T) {
	testOpts := CoordinateSupplierOptions{2, 2, 2, Asc, false}
	// test the ones behind CoordinateSupplier interface
	for _, supplier := range suppliersToTest {
		t.Run(supplier.name, func(t *testing.T) {
			cs, err := supplier.new(testOpts)
			require.NoError(t, err)

			x1, y1, z1, done1 := cs.Next()
			x2, y2, z2, done2 := cs.Next()
			x3, y3, z3, done3 := cs.Next()
			x4, y4, z4, done4 := cs.Next()
			x5, y5, z5, done5 := cs.Next()
			x6, y6, z6, done6 := cs.Next()
			x7, y7, z7, done7 := cs.Next()
			x8, y8, z8, done8 := cs.Next()
			_, _, _, done9 := cs.Next()

			require.Equal(t, 0, x1)
			require.Equal(t, 1, x2)
			require.Equal(t, 0, x3)
			require.Equal(t, 1, x4)
			require.Equal(t, 0, x5)
			require.Equal(t, 1, x6)
			require.Equal(t, 0, x7)
			require.Equal(t, 1, x8)

			require.Equal(t, 0, y1)
			require.Equal(t, 0, y2)
			require.Equal(t, 1, y3)
			require.Equal(t, 1, y4)
			require.Equal(t, 0, y5)
			require.Equal(t, 0, y6)
			require.Equal(t, 1, y7)
			require.Equal(t, 1, y8)

			require.Equal(t, 0, z1)
			require.Equal(t, 0, z2)
			require.Equal(t, 0, z3)
			require.Equal(t, 0, z4)
			require.Equal(t, 1, z5)
			require.Equal(t, 1, z6)
			require.Equal(t, 1, z7)
			require.Equal(t, 1, z8)

			require.False(t, done1)
			require.False(t, done2)
			require.False(t, done3)
			require.False(t, done4)
			require.False(t, done5)
			require.False(t, done6)
			require.False(t, done7)
			require.False(t, done8)
			require.True(t, done9)
		})
	}
}

func Test_Coordinate_Supplier_Desc_2x2x2(t *testing.T) {
	testOpts := CoordinateSupplierOptions{2, 2, 2, Desc, false}
	// test the ones behind CoordinateSupplier interface
	for _, supplier := range suppliersToTest {
		t.Run(supplier.name, func(t *testing.T) {
			cs, err := supplier.new(testOpts)
			require.NoError(t, err)

			x1, y1, z1, done1 := cs.Next()
			x2, y2, z2, done2 := cs.Next()
			x3, y3, z3, done3 := cs.Next()
			x4, y4, z4, done4 := cs.Next()
			x5, y5, z5, done5 := cs.Next()
			x6, y6, z6, done6 := cs.Next()
			x7, y7, z7, done7 := cs.Next()
			x8, y8, z8, done8 := cs.Next()
			_, _, _, done9 := cs.Next()

			require.Equal(t, 1, x1)
			require.Equal(t, 0, x2)
			require.Equal(t, 1, x3)
			require.Equal(t, 0, x4)
			require.Equal(t, 1, x5)
			require.Equal(t, 0, x6)
			require.Equal(t, 1, x7)
			require.Equal(t, 0, x8)

			require.Equal(t, 1, y1)
			require.Equal(t, 1, y2)
			require.Equal(t, 0, y3)
			require.Equal(t, 0, y4)
			require.Equal(t, 1, y5)
			require.Equal(t, 1, y6)
			require.Equal(t, 0, y7)
			require.Equal(t, 0, y8)

			require.Equal(t, 1, z1)
			require.Equal(t, 1, z2)
			require.Equal(t, 1, z3)
			require.Equal(t, 1, z4)
			require.Equal(t, 0, z5)
			require.Equal(t, 0, z6)
			require.Equal(t, 0, z7)
			require.Equal(t, 0, z8)

			require.False(t, done1)
			require.False(t, done2)
			require.False(t, done3)
			require.False(t, done4)
			require.False(t, done5)
			require.False(t, done6)
			require.False(t, done7)
			require.False(t, done8)
			require.True(t, done9)
		})
	}
}

func Test_Coordinate_Supplier_Asc_3x2x1_Repeat(t *testing.T) {
	testOpts := CoordinateSupplierOptions{3, 2, 1, Asc, true}
	xPattern := []int{0, 1, 2, 0, 1, 2}
	yPattern := []int{0, 0, 0, 1, 1, 1}
	zPattern := []int{0, 0, 0, 0, 0, 0}
	// test the ones behind CoordinateSupplier interface
	for _, supplier := range suppliersToTest {
		t.Run(supplier.name, func(t *testing.T) {
			cs, err := supplier.new(testOpts)
			require.NoError(t, err)
			for seen := 0; seen < 1000; seen++ {
				x, y, z, done := cs.Next()
				require.False(t, done)
				require.Equal(t, xPattern[seen%len(xPattern)], x)
				require.Equal(t, yPattern[seen%len(yPattern)], y)
				require.Equal(t, zPattern[seen%len(zPattern)], z)
			}
		})
	}
}

func Test_Coordinate_Supplier_Desc_3x2x1_Repeat(t *testing.T) {
	testOpts := CoordinateSupplierOptions{3, 2, 1, Desc, true}
	xPattern := []int{2, 1, 0, 2, 1, 0}
	yPattern := []int{1, 1, 1, 0, 0, 0}
	zPattern := []int{0, 0, 0, 0, 0, 0}
	// test the ones behind CoordinateSupplier interface
	for _, supplier := range suppliersToTest {
		t.Run(supplier.name, func(t *testing.T) {
			cs, err := supplier.new(testOpts)
			require.NoError(t, err)
			for seen := 0; seen < 1000; seen++ {
				x, y, z, done := cs.Next()
				require.False(t, done)
				require.Equal(t, xPattern[seen%len(xPattern)], x)
				require.Equal(t, yPattern[seen%len(yPattern)], y)
				require.Equal(t, zPattern[seen%len(zPattern)], z)
			}
		})
	}
}

func Test_Coordinate_Supplier_Asc_3x3x3_Random(t *testing.T) {
	testOpts := CoordinateSupplierOptions{3, 3, 3, Random, false}
	// test the ones behind CoordinateSupplier interface
	for _, supplier := range suppliersToTest {
		t.Run(supplier.name, func(t *testing.T) {
			cs, err := supplier.new(testOpts)
			require.NoError(t, err)
			type xyz struct {
				x int
				y int
				z int
			}
			seen := map[xyz]bool{}
			for x, y, z, done := cs.Next(); !done; x, y, z, done = cs.Next() {
				seen[xyz{x, y, z}] = true
				require.True(t, x >= 0 && x <= 2)
				require.True(t, y >= 0 && y <= 2)
				require.True(t, z >= 0 && z <= 2)
			}
			require.Len(t, seen, 27)
		})
	}
}

func Test_Coordinate_Supplier_Desc_2x2x2_Repeat(t *testing.T) {
	testOpts := CoordinateSupplierOptions{2, 2, 2, Desc, true}
	xPattern := []int{1, 0, 1, 0, 1, 0, 1, 0}
	yPattern := []int{1, 1, 0, 0, 1, 1, 0, 0}
	zPattern := []int{1, 1, 1, 1, 0, 0, 0, 0}
	// test the ones behind CoordinateSupplier interface
	for _, supplier := range suppliersToTest {
		t.Run(supplier.name, func(t *testing.T) {
			cs, err := supplier.new(testOpts)
			require.NoError(t, err)
			for seen := 0; seen < 1000; seen++ {
				x, y, z, done := cs.Next()
				require.False(t, done)
				require.Equal(t, xPattern[seen%len(xPattern)], x)
				require.Equal(t, yPattern[seen%len(yPattern)], y)
				require.Equal(t, zPattern[seen%len(zPattern)], z)
			}
		})
	}
}

func Test_Coordinate_Supplier_Asc_1000x1000x1_Concurrent(t *testing.T) {
	testOpts := CoordinateSupplierOptions{1000, 1000, 1, Asc, false}
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
	testOpts := CoordinateSupplierOptions{100, 100, 2, Asc, false}
	// test the ones behind CoordinateSupplier interface
	for _, supplier := range suppliersToTest {
		t.Run(supplier.name, func(t *testing.T) {
			// consume the coordinates
			cs, err := supplier.new(testOpts)
			require.NoError(t, err)
			for _, _, _, done := cs.Next(); !done; _, _, _, done = cs.Next() {
				require.False(t, done)
			}

			// get a bunch past the end... it should still be done
			extras := 0
			for {
				if extras > 1000000 {
					break
				}
				extras++

				_, _, _, done := cs.Next()
				require.True(t, done)
			}
		})
	}
}

func Test_Readme_Example(t *testing.T) {
	opts := CoordinateSupplierOptions{Width: 10, Height: 10, Depth: 1, Order: Asc, Repeat: false}
	cs, err := NewCoordinateSupplier(opts)
	if err != nil {
		require.NoError(t, err)
	}

	for x, y, z, done := cs.Next(); !done; x, y, z, done = cs.Next() {
		fmt.Println("The next coordinate is", x, y, z)
	}
}

func BenchmarkCoordinateSuppliers(b *testing.B) {
	upToWidth := 1000
	upToHeight := 1000
	depth := 5
	upToConsumers := 1000
	upToConsumed := 1000000

	for width := 1; width <= upToWidth; width *= 30 {
		for height := 1; height <= upToHeight; height *= 30 {
			for consume := 1; consume <= upToConsumed; consume *= 1000000 {
				for consumers := 1; consumers <= upToConsumers; consumers *= 10 {
					// run CoordinateSuppliers
					for _, supplier := range suppliersToTest {
						b.Run(fmt.Sprintf("%s-%dw-%dh-%dd-%dconsumers-consume%d", supplier.name, width, height, depth, consumers, consume), func(b *testing.B) {
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
								cs, err := supplier.new(CoordinateSupplierOptions{width, height, depth, Asc, repeat})
								require.NoError(b, err)

								count := runCoordinateSupplier(cs, consumers, uint64(useConsume))
								if useConsume == 0 {
									require.Equal(b, uint64(width*height*depth), count)
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
			for _, _, _, done := cs.Next(); !done; _, _, _, done = cs.Next() {
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
