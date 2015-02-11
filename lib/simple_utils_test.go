package fcompute

import(
  "testing"
)

func TestEven_with_even(t *testing.T) {
  if (even(6) != true) {
    t.Error("even failed on 6")
  }
}
func TestEven_with_odd(t *testing.T) {
  if (even(11) != false) {
    t.Error("even failed on 11")
  }
}
func TestOdd_with_even(t *testing.T) {
  if (odd(8) != false) {
    t.Error("odd failed on 8")
  }
}
func TestOdd_with_odd(t *testing.T) {
  if (odd(13) != true) {
    t.Error("odd failed on 13")
  }
}

func TestMean(t *testing.T) {
  input := []float64{2,5,11}

  expected := 6.0
  got := mean(input)

  if got != expected {
    t.Errorf("TestMean failed expected - %f, got - %f", expected, got)
  }
}
