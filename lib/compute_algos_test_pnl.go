package fcompute

import (
  "testing"
  "reflect"
)

func TestPNLData_For_Only_Positive_Vals(t *testing.T) {
  data := ComputedFData {
    NetDailyAssetValue: []float64 {1,2,3,4,5},
  }
  expected := []float64{0,1,2,3,4}

  got := data.PNLData()

  if !reflect.DeepEqual(got, expected) {
    t.Error("Expected output mismatched for pnl data - expected, got - ", expected, ",", got)
  }
}

func TestPNLData_For_Only_Negative_Vals(t *testing.T) {
  data := ComputedFData {
    NetDailyAssetValue: []float64 {5,4,3,2,1},
  }
  expected := []float64{0,-1,-2,-3,-4}

  got := data.PNLData()

  if !reflect.DeepEqual(got, expected) {
    t.Error("Expected output mismatched for pnl data - expected, got - ", expected, ",", got)
  }
}

func TestPNLData_For_Complex_Input(t *testing.T) {
  data := ComputedFData {
    NetDailyAssetValue: []float64 {10,8,12,11,10,9,7,10,10,12,14},
  }
  expected := []float64{0,-2,2,1,0,-1,-3,0,0,2,4}

  got := data.PNLData()

  if !reflect.DeepEqual(got, expected) {
    t.Error("Expected output mismatched for pnl data - expected, got - ", expected, ",", got)
  }
}
