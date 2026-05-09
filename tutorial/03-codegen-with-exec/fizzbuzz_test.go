package fizzbuzz

import "testing"

// TestFizzBuzz exercises the contract:
//
//   - n divisible by 3 and 5 → "FizzBuzz"
//   - n divisible by 3 only  → "Fizz"
//   - n divisible by 5 only  → "Buzz"
//   - otherwise              → strconv.Itoa(n)
//
// The agent's job is to create fizzbuzz.go in this package with a
// FizzBuzz(n int) string function that satisfies these cases.
func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		in   int
		want string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, "7"},
		{9, "Fizz"},
		{10, "Buzz"},
		{15, "FizzBuzz"},
		{30, "FizzBuzz"},
		{45, "FizzBuzz"},
		{98, "98"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := FizzBuzz(tt.in); got != tt.want {
				t.Errorf("FizzBuzz(%d) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
