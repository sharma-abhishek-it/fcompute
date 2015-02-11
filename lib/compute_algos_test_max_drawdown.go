package fcompute

import "testing"

func TestMaxDrawdownLogic_Descending_Series(t *testing.T) {
  data := ComputedFData {
    NetDailyAssetValue: []float64 {5,4,3,2,1},
  }

  expected := (5-1)/5.0
  got := data.MaximumDrawdown()

  if got != expected {
    t.Errorf("Expected output mismatched for max drawdown | Expected = %d, Got = %d", expected, got)
  }
}

func TestMaxDrawdownLogic_Ascending_Series(t *testing.T) {
  data := ComputedFData {
    NetDailyAssetValue: []float64 {1,2,3,4,5},
  }

  expected := 0.0
  got := data.MaximumDrawdown()

  if got != expected {
    t.Errorf("Expected output mismatched for max drawdown | Expected = %d, Got = %d", expected, got)
  }
}

func TestMaxDrawdownLogic_Only_One_Peak(t *testing.T) {
  data := ComputedFData {
    NetDailyAssetValue: []float64 {1,2,5,4,3},
  }

  expected := (5-3)/5.0
  got := data.MaximumDrawdown()

  if got != expected {
    t.Errorf("Expected output mismatched for max drawdown | Expected = %d, Got = %d", expected, got)
  }
}

func TestMaxDrawdownLogic_Actual_Peak_on_left(t *testing.T) {
  data := ComputedFData {
    NetDailyAssetValue: []float64 {1,5,2,4,3},
  }

  expected := (5-2)/5.0
  got := data.MaximumDrawdown()

  if got != expected {
    t.Errorf("Expected output mismatched for max drawdown | Expected = %d, Got = %d", expected, got)
  }
}

func TestMaxDrawdownLogic_Actual_Peak_on_right(t *testing.T) {
  data := ComputedFData {
    NetDailyAssetValue: []float64 {1,4,3,5,2},
  }

  expected := (5-2)/5.0
  got := data.MaximumDrawdown()

  if got != expected {
    t.Errorf("Expected output mismatched for max drawdown | Expected = %d, Got = %d", expected, got)
  }
}

func TestMaxDrawdownLogic_Many_Smaller_Peaks(t *testing.T) {
  data := ComputedFData {
    NetDailyAssetValue: []float64 {1,2,10,7,6,8,2,1,4,3},
  }

  expected := (10-1)/10.0
  got := data.MaximumDrawdown()

  if got != expected {
    t.Errorf("Expected output mismatched for max drawdown | Expected = %d, Got = %d", expected, got)
  }
}

func TestMaxDrawdownLogic_Complex_Case(t *testing.T) {
  data := ComputedFData {
    NetDailyAssetValue: []float64 {1,2,10,7,6,8,2,1,4,3,24,10,9,11,2,18},
  }

  expected := (24-2)/24.0
  got := data.MaximumDrawdown()

  if got != expected {
    t.Errorf("Expected output mismatched for max drawdown | Expected = %d, Got = %d", expected, got)
  }
}

func TestMaxDrawdownLogic_New_peak_at_last(t *testing.T) {
  data := ComputedFData {
    NetDailyAssetValue: []float64 {1,2,10,7,6,8,2,1,4,3,24},
  }

  expected := (10-1)/10.0
  got := data.MaximumDrawdown()

  if got != expected {
    t.Errorf("Expected output mismatched for max drawdown | Expected = %d, Got = %d", expected, got)
  }
}
