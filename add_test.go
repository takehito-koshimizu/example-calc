package calc

import "testing"

func TestAdd(t *testing.T) {
  const x, y, expected = 1, 2, 3
  result := Add(x, y)
  if result != expected {
    t.Errorf("Add(%v, %v) = %v, want: %v", x, y, result, expected)
  }
}
