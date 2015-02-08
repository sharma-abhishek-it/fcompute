package fcompute

func (fData *ComputedFData) PreCompute(prefs UserPrefs) {

  fData.DailyAssetValue      = make([][]float64, 0)
  fData.NetDailyAssetValue   = []float64{prefs.Investment}
  total_days,total_products := len(fData.OriginalData[0]),len(fData.OriginalData)
  last_rebalanced_day       := 0


  for _,weight := range prefs.Weights {
    fData.DailyAssetValue    = append(fData.DailyAssetValue, []float64{weight * prefs.Investment})
  }

  for day:=1; day<total_days; day++ {
    todays_net_value := float64(0)

    for product_i:=0; product_i < total_products; product_i++ {
      rate_of_today                := fData.OriginalData[product_i][day]
      rate_of_last_rebalancing_day := fData.OriginalData[product_i][last_rebalanced_day]
      asset_value_on_last_rebalancing_day := fData.DailyAssetValue[product_i][last_rebalanced_day]

      todays_investment_value  := asset_value_on_last_rebalancing_day * rate_of_today / rate_of_last_rebalancing_day
      todays_net_value         += todays_investment_value

      fData.DailyAssetValue[product_i] = append(fData.DailyAssetValue[product_i], todays_investment_value)
    }
    fData.NetDailyAssetValue = append(fData.NetDailyAssetValue, todays_net_value)


    if ((day+1) % prefs.RebalanceFrequency == 0) {
      last_rebalanced_day   = day
      for product_i,weight := range prefs.Weights {
        fData.DailyAssetValue[product_i][day] = weight * todays_net_value
      }
    }
  }
}
