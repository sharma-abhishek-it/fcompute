package fcompute

import "testing"

func testing_data()(
  products_fdata [][]float64,
  products_weights []float64,
  investment float64) {

  products_fdata = [][]float64 {{5,4,6},{4,3,3},{6,4,7}}
  products_weights = []float64{0.2,0.3,0.5}

  return products_fdata, products_weights, 100
}

func TestPreComputeWithRebalancedEndDay(t *testing.T) {
  products_fdata,weights,investment := testing_data()
  expected_daily_asset_value,expected_net_asset := [3]int {20,31,52}, 104

  daily_asset_value, net_daily_asset_value :=
    PreCompute(products_fdata,weights,3,investment)

  for i,p_data := range daily_asset_value  {
    if int(p_data[2]) != expected_daily_asset_value[i] {
      t.Errorf("Expected output mismatched for product #%d | Expected = %d, Got = %f",
      i, expected_daily_asset_value[i], daily_asset_value[i][2])
    }
  }

  if int(net_daily_asset_value[2]) != expected_net_asset {
    t.Errorf("Expected output mismatched for net_daily_asset | Expected = %d, Got = %f",
    expected_net_asset, net_daily_asset_value[2])
  }
}

func TestPreComputeWithUnbalancedEndDay(t *testing.T) {
  products_fdata,weights,investment := testing_data()
  expected_daily_asset_value, expected_net_asset := [3]int {21,21,62}, 105

  daily_asset_value, net_daily_asset_value :=
    PreCompute(products_fdata,weights,2,investment)

  for i,p_data := range daily_asset_value  {
    if int(p_data[2]) != expected_daily_asset_value[i] {
      t.Errorf("Expected output mismatched for product #%d | Expected = %d, Got = %f",
      i, expected_daily_asset_value[i], daily_asset_value[i][2])
    }
  }

  if int(net_daily_asset_value[2]) != expected_net_asset {
    t.Errorf("Expected output mismatched for net_daily_asset | Expected = %d, Got = %f",
    expected_net_asset, net_daily_asset_value[2])
  }
}
