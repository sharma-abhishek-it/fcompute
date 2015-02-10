package fcompute

import "testing"

func testing_data()(ComputedFData, UserPrefs) {
  data  := ComputedFData {
    OriginalData: [][]float64 {{5,4,6},{4,3,3},{6,4,7}},
  }
  prefs := UserPrefs{
    Investment:         100,
    Weights:            []float64{0.2,0.3,0.5},
  }

  return data, prefs
}

func TestPreComputeWithRebalancedEndDay(t *testing.T) {
  data,prefs := testing_data()
  prefs.RebalanceFrequency = 3
  expected_daily_asset_value,expected_net_asset := [3]int {20,31,52}, 104

  data.PreCompute(prefs)

  for i,p_data := range data.DailyAssetValue  {
    if int(p_data[2]) != expected_daily_asset_value[i] {
      t.Errorf("Expected output mismatched for product #%d | Expected = %d, Got = %f",
      i, expected_daily_asset_value[i], data.DailyAssetValue[i][2])
    }
  }

  if int(data.NetDailyAssetValue[2]) != expected_net_asset {
    t.Errorf("Expected output mismatched for net_daily_asset | Expected = %d, Got = %f",
    expected_net_asset, data.NetDailyAssetValue[2])
  }
}

func TestPreComputeWithUnbalancedEndDay(t *testing.T) {
  data,prefs := testing_data()
  prefs.RebalanceFrequency = 2

  expected_daily_asset_value,expected_net_asset := [3]int {21,21,62}, 105

  data.PreCompute(prefs)

  for i,p_data := range data.DailyAssetValue  {
    if int(p_data[2]) != expected_daily_asset_value[i] {
      t.Errorf("Expected output mismatched for product #%d | Expected = %d, Got = %f",
      i, expected_daily_asset_value[i], data.DailyAssetValue[i][2])
    }
  }

  if int(data.NetDailyAssetValue[2]) != expected_net_asset {
    t.Errorf("Expected output mismatched for net_daily_asset | Expected = %d, Got = %f",
    expected_net_asset, data.NetDailyAssetValue[2])
  }
}
