package fcompute

import "testing"

func TestNetReturnsData_for_loss(t *testing.T) {
  data := ComputedFData{
    NetDailyAssetValue: []float64 {5,6,1},
  }

  expected := 0.0
  got := data.NetReturns()

  if got != expected {
    t.Errorf("Test failed for Net returns on loss, expected = %f, got = %f", expected, got)
  }
}

func TestNetReturnsData_for_profit(t *testing.T) {
  data := ComputedFData{
    NetDailyAssetValue: []float64 {5,6,1,8},
  }

  expected := ((8-5)/5) * 100.0
  got := data.NetReturns()

  if got != expected {
    t.Errorf("Test failed for Net returns on loss, expected = %f, got = %f", expected, got)
  }
}
