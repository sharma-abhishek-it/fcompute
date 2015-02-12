package fcompute

import (
  "os"
  "time"
  "math"
)

const ENV_START_DATE_KEY string = "FDATA_STARTING_DATE"
const ENV_END_DATE_KEY string = "FDATA_ENDING_DATE"

// Slice original data based on dates.
// Currently stores the beginning date for all data in environment variable
// o/w a proper sync mechanism to redis when sidekiq pushes data will be needed
// Ending Date is assumed that it will be within range else take full data for ending
func (fData ComputedFData) getSlicingIndexes(prefs UserPrefs)(int, int) {
  starting_index, ending_index := 0, len(fData.OriginalData[0]) - 1

  starting_date_string := os.Getenv(ENV_START_DATE_KEY)
  ending_date_string   := os.Getenv(ENV_END_DATE_KEY)
  starting_date, _     := time.Parse(ShortTimeFormat, starting_date_string)
  ending_date, _       := time.Parse(ShortTimeFormat, ending_date_string)

  if( !starting_date.IsZero() &&
      !time.Time(prefs.StartDate).IsZero() &&
      starting_date.Before(time.Time(prefs.StartDate)) ) {

    starting_index = prefs.StartDate.DaysSince(FDataDate(starting_date))
    starting_index -= 1 // 0 based array indexes
    if (starting_index >= ending_index) {
      starting_index = 0
    }
  }

  if( !ending_date.IsZero() &&
      !time.Time(prefs.EndDate).IsZero() &&
      ending_date.After(time.Time(prefs.EndDate)) ) {

    ending_index -= prefs.EndDate.DaysSince(FDataDate(ending_date))
    ending_index += 1 // slicing is exclusive of end index hence the increment
    if (ending_index <= starting_index) {
      ending_index = len(fData.OriginalData[0]) - 1
    }
  }

  return starting_index, ending_index
}

func (fData *ComputedFData) PreCompute(prefs UserPrefs) {

  start,end                 := fData.getSlicingIndexes(prefs)
  for i,_ := range fData.OriginalData {
    fData.OriginalData[i] = fData.OriginalData[i][start:end+1]
  }

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

func (fData ComputedFData) PNLData() (pnl []float64) {
  series := fData.NetDailyAssetValue
  pnl = series
  investment := series[0]

  for i, val := range series { pnl[i] = val - investment }

  return pnl
}

func (fData ComputedFData) NetReturns() float64 {
  series := fData.NetDailyAssetValue
  last   := len(series) - 1

  net_returns := series[last] - series[0]

  if net_returns < 0 {
    return 0
  } else {
    return (net_returns / series[0]) * 100
  }
}

func (fData ComputedFData) AnnualizedReturns() float64 {
  series := fData.NetDailyAssetValue

  annualized_returns := math.Exp(252*mean(series)-1) * 100

  return annualized_returns
}

func (fData ComputedFData) MaximumDrawdown() float64 {
  calc_drawdown := func(peak, low float64) float64 { return (peak-low) / peak }

  series := fData.NetDailyAssetValue
  last_peak_value, max_drawdown := series[0], float64(0)

  for i,length := 1,len(series); i < length; i++ {
    prev, now := series[i-1], series[i]

    if prev < now && now > last_peak_value {
      last_peak_value = now
    } else if prev > now && calc_drawdown(last_peak_value, now) > max_drawdown {
      max_drawdown = calc_drawdown(last_peak_value, now)
    }

  }

  return max_drawdown

}
