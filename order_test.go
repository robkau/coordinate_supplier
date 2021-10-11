package coordinate_supplier

import (
	"fmt"
	"testing"
)

func TestOrderToString(t *testing.T) {

	tests := []struct {
		o    Order
		want string
	}{
		{Asc, "Asc"},
		{Desc, "Desc"},
		{Random, "Random"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("want-%s", tt.want), func(t *testing.T) {
			if got := OrderToString(tt.o); got != tt.want {
				t.Errorf("OrderToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
