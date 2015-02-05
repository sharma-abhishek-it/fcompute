package fcompute

func PreCompute(
  products_fdata [][]float64,
  products_weights []float64,
  rebalance_frequency int, investment float64)(daily_asset_value [][]float64,net_daily_asset_value []float64) {


  daily_asset_value          = make([][]float64, 0)
  net_daily_asset_value      = []float64{investment}
  total_days,total_products := len(products_fdata[0]),len(products_fdata)
  last_rebalanced_day       := 0


  for _,weight := range products_weights {
    daily_asset_value        = append(daily_asset_value, []float64{weight * investment})
  }

  for day:=1; day<total_days; day++ {
    net_daily_value := float64(0)

    for product_i:=0; product_i<total_products; product_i++ {
      day_closing_rate,last_rebalanced_rate := products_fdata[product_i][day],products_fdata[product_i][last_rebalanced_day]
      day_investment_value := daily_asset_value[product_i][last_rebalanced_day] * day_closing_rate / last_rebalanced_rate

      daily_asset_value[product_i] = append(daily_asset_value[product_i], day_investment_value)
      net_daily_value += day_investment_value
    }
    net_daily_asset_value = append(net_daily_asset_value, net_daily_value)


    if ((day+1) % rebalance_frequency == 0) {
      net_assets := float64(0)
      for product_i:=0; product_i<total_products; product_i++ {
        net_assets += daily_asset_value[product_i][day]
      }
      for product_i,weight := range products_weights {
        last_rebalanced_day               = day
        daily_asset_value[product_i][day] = weight * net_assets
      }
    }
  }

  return daily_asset_value, net_daily_asset_value
}
