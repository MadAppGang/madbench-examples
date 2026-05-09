package calc

import "testing"

// TestAdd pins the contract: Add returns the sum of its two arguments.
func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{3, 4, 7},
		{0, 0, 0},
		{-5, 5, 0},
		{10, 20, 30},
		{-1, -1, -2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Add(tt.a, tt.b); got != tt.want {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
