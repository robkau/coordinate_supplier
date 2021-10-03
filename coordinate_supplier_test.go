package coordinate_supplier

import (
	"github.com/stretchr/testify/require"
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
