package coordinate_supplier

import (
	"github.com/stretchr/testify/require"
	"sync"
	"sync/atomic"
	"testing"
)

func Test_Coordinate_Supplier_Asc_10x1(t *testing.T) {
	cs, err := NewCoordinateSupplier(10, 1, Asc, false)
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
}

func Test_Coordinate_Supplier_Asc_1x10(t *testing.T) {
	cs, err := NewCoordinateSupplier(1, 10, Asc, false)
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
}

func Test_Coordinate_Supplier_Asc_2x2(t *testing.T) {
	cs, err := NewCoordinateSupplier(2, 2, Asc, false)
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
}

func Test_Coordinate_Supplier_Desc_2x2(t *testing.T) {
	cs, err := NewCoordinateSupplier(2, 2, Desc, false)
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
}

func Test_Coordinate_Supplier_Asc_3x2_Repeat(t *testing.T) {
	cs, err := NewCoordinateSupplier(3, 2, Asc, true)
	require.NoError(t, err)

	xPattern := []int{0, 1, 2, 0, 1, 2}
	yPattern := []int{0, 0, 0, 1, 1, 1}
	for seen := 0; seen < 1000; seen++ {
		x, y, done := cs.Next()
		require.False(t, done)
		require.Equal(t, xPattern[seen%len(xPattern)], x)
		require.Equal(t, yPattern[seen%len(yPattern)], y)
	}
}

func Test_Coordinate_Supplier_Desc_3x2_Repeat(t *testing.T) {
	cs, err := NewCoordinateSupplier(3, 2, Desc, true)
	require.NoError(t, err)

	xPattern := []int{2, 1, 0, 2, 1, 0}
	yPattern := []int{1, 1, 1, 0, 0, 0}
	for seen := 0; seen < 1000; seen++ {
		x, y, done := cs.Next()
		require.False(t, done)
		require.Equal(t, xPattern[seen%len(xPattern)], x)
		require.Equal(t, yPattern[seen%len(yPattern)], y)
	}
}

func Test_Coordinate_Supplier_Desc_2x2_Repeat(t *testing.T) {
	cs, err := NewCoordinateSupplier(2, 2, Desc, true)
	require.NoError(t, err)

	xPattern := []int{1, 0, 1, 0}
	yPattern := []int{1, 1, 0, 0}
	for seen := 0; seen < 1000; seen++ {
		x, y, done := cs.Next()
		require.False(t, done)
		require.Equal(t, xPattern[seen%len(xPattern)], x)
		require.Equal(t, yPattern[seen%len(yPattern)], y)
	}
}

func runBenchmarkCoordinateSupplier(w, h int, next NextMode, numConsumers int, maxConsumed uint64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		// make supplier
		var repeat bool
		var consumed uint64
		if maxConsumed > 0 {
			// instead of consuming once, will repeat until maxConsumed
			repeat = true
		}
		cs, err := NewCoordinateSupplier(w, h, next, repeat)
		if err != nil {
			b.Fatalf("failed create supplier: %s", err)
		}

		// run consumers to get all coordinates
		wg := sync.WaitGroup{}
		for i := 0; i < numConsumers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for _, _, done := cs.Next(); !done && atomic.LoadUint64(&consumed) <= maxConsumed; _, _, done = cs.Next() {
					if repeat {
						atomic.AddUint64(&consumed, 1)
					}
				}
			}()
		}
		wg.Wait()
	}
}

func Benchmark_3x3_1_ConsumeOnce(b *testing.B)  { runBenchmarkCoordinateSupplier(3, 3, Asc, 1, 0, b) }
func Benchmark_3x3_3_ConsumeOnce(b *testing.B)  { runBenchmarkCoordinateSupplier(3, 3, Asc, 3, 0, b) }
func Benchmark_3x3_30_ConsumeOnce(b *testing.B) { runBenchmarkCoordinateSupplier(3, 3, Asc, 30, 0, b) }
func Benchmark_3x3_300_ConsumeOnce(b *testing.B) {
	runBenchmarkCoordinateSupplier(3, 3, Asc, 300, 0, b)
}
func Benchmark_3x3_3000_ConsumeOnce(b *testing.B) {
	runBenchmarkCoordinateSupplier(3, 3, Asc, 3000, 0, b)
}

func Benchmark_3x3_1_ConsumeTenMillionItems(b *testing.B) {
	runBenchmarkCoordinateSupplier(3, 3, Asc, 1, 10000000, b)
}
func Benchmark_3x3_3_ConsumeTenMillionItems(b *testing.B) {
	runBenchmarkCoordinateSupplier(3, 3, Asc, 3, 10000000, b)
}
func Benchmark_3x3_30_ConsumeTenMillionItems(b *testing.B) {
	runBenchmarkCoordinateSupplier(3, 3, Asc, 30, 10000000, b)
}
func Benchmark_3x3_300_ConsumeTenMillionItems(b *testing.B) {
	runBenchmarkCoordinateSupplier(3, 3, Asc, 300, 10000000, b)
}
func Benchmark_3x3_3000_ConsumeTenMillionItems(b *testing.B) {
	runBenchmarkCoordinateSupplier(3, 3, Asc, 3000, 10000000, b)
}

func Benchmark_21x21_1_ConsumeOnce(b *testing.B) {
	runBenchmarkCoordinateSupplier(21, 21, Asc, 1, 0, b)
}
func Benchmark_21x21_3_ConsumeOnce(b *testing.B) {
	runBenchmarkCoordinateSupplier(21, 21, Asc, 3, 0, b)
}
func Benchmark_21x21_30_ConsumeOnce(b *testing.B) {
	runBenchmarkCoordinateSupplier(21, 21, Asc, 30, 0, b)
}
func Benchmark_21x21_300_ConsumeOnce(b *testing.B) {
	runBenchmarkCoordinateSupplier(21, 21, Asc, 300, 0, b)
}
func Benchmark_21x21_3000_ConsumeOnce(b *testing.B) {
	runBenchmarkCoordinateSupplier(21, 21, Asc, 3000, 0, b)
}

func Benchmark_21x21_1_ConsumeTenMillionItems(b *testing.B) {
	runBenchmarkCoordinateSupplier(21, 21, Asc, 1, 10000000, b)
}
func Benchmark_21x21_3_ConsumeTenMillionItems(b *testing.B) {
	runBenchmarkCoordinateSupplier(21, 21, Asc, 3, 10000000, b)
}
func Benchmark_21x21_30_ConsumeTenMillionItems(b *testing.B) {
	runBenchmarkCoordinateSupplier(21, 21, Asc, 30, 10000000, b)
}
func Benchmark_21x21_300_ConsumeTenMillionItems(b *testing.B) {
	runBenchmarkCoordinateSupplier(21, 21, Asc, 300, 10000000, b)
}
func Benchmark_21x21_3000_ConsumeTenMillionItems(b *testing.B) {
	runBenchmarkCoordinateSupplier(21, 21, Asc, 3000, 10000000, b)
}

func Benchmark_210x210_1_ConsumeOnce(b *testing.B) {
	runBenchmarkCoordinateSupplier(210, 210, Asc, 1, 0, b)
}
func Benchmark_210x210_3_ConsumeOnce(b *testing.B) {
	runBenchmarkCoordinateSupplier(210, 210, Asc, 3, 0, b)
}
func Benchmark_210x210_30_ConsumeOnce(b *testing.B) {
	runBenchmarkCoordinateSupplier(210, 210, Asc, 30, 0, b)
}
func Benchmark_210x210_300_ConsumeOnce(b *testing.B) {
	runBenchmarkCoordinateSupplier(210, 210, Asc, 300, 0, b)
}
func Benchmark_210x210_3000_ConsumeOnce(b *testing.B) {
	runBenchmarkCoordinateSupplier(210, 210, Asc, 3000, 0, b)
}

func Benchmark_210x210_1_ConsumeTenMillionItems(b *testing.B) {
	runBenchmarkCoordinateSupplier(210, 210, Asc, 1, 10000000, b)
}
func Benchmark_210x210_3_ConsumeTenMillionItems(b *testing.B) {
	runBenchmarkCoordinateSupplier(210, 210, Asc, 3, 10000000, b)
}
func Benchmark_210x210_30_ConsumeTenMillionItems(b *testing.B) {
	runBenchmarkCoordinateSupplier(210, 210, Asc, 30, 10000000, b)
}
func Benchmark_210x210_300_ConsumeTenMillionItems(b *testing.B) {
	runBenchmarkCoordinateSupplier(210, 210, Asc, 300, 10000000, b)
}
func Benchmark_210x210_3000_ConsumeTenMillionItems(b *testing.B) {
	runBenchmarkCoordinateSupplier(210, 210, Asc, 3000, 10000000, b)
}
