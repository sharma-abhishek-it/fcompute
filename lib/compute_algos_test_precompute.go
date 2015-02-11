package fcompute

import (
  "testing"
  "os"
)

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

func TestPreCompute_First_NetDailyAssetValue_Equals_Initial_Investment(t *testing.T) {
  data,prefs := testing_data()
  prefs.RebalanceFrequency = 2

  expected := prefs.Investment

  data.PreCompute(prefs)

  if data.NetDailyAssetValue[0] != expected {
    t.Errorf("Expected output mismatched for net_daily_asset | Expected = %f, Got = %f",
    expected, data.NetDailyAssetValue[0])
  }
}

func TestPreCompute_With_Custom_Starting_Ending_Dates_Data_Length(t *testing.T) {
  data,prefs := testing_data()
  data.OriginalData = [][]float64 {{3,5,4,6,7},{2,4,3,3,8},{5,6,4,7,9}}

  os.Setenv(ENV_START_DATE_KEY, "2014-01-22")
  os.Setenv(ENV_END_DATE_KEY, "2014-01-26")

  prefs.fillDates("2014-01-23", "2014-01-25")

  prefs.RebalanceFrequency = 2

  expected_length := 2

  data.PreCompute(prefs)

  if len(data.NetDailyAssetValue) != expected_length {
    t.Errorf("Expected length mismatched for net_daily_assets | Expected = %d, Got = %f",
    expected_length, len(data.NetDailyAssetValue))
  }
}


func TestPreCompute_With_Custom_Starting_Ending_Dates_Exact_Data(t *testing.T) {
  data,prefs := testing_data()
  data.OriginalData = [][]float64 {{3,5,4,6,7},{2,4,3,3,8},{5,6,4,7,9}}

  os.Setenv(ENV_START_DATE_KEY, "2014-01-22")
  os.Setenv(ENV_END_DATE_KEY, "2014-01-26")

  prefs.fillDates("2014-01-23", "2014-01-25")

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
