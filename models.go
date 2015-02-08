package fcompute

import (
  "time"
)

const shortTimeFormat = "2006-01-02"

type FDataDate time.Time

func (t *FDataDate) UnmarshalJSON(b []byte) error {
  v, _ := time.Parse(shortTimeFormat, string(b[1:len(b)-1]))
  *t = FDataDate(v)
  return nil
}

func (t1 FDataDate) DaysSince(t2 FDataDate) int {
  this, that := time.Time(t1), time.Time(t2)
  return int(that.Sub(this).Hours() / 24) + 1
}

type UserPrefs struct {
  StartDate            FDataDate `json:"start_date"`
  EndDate              FDataDate `json:"end_date"`
  RebalanceFrequency   int       `json:"frequency"`
  Investment           float64   `json:"investment"`
  Weights              []float64 `json:"weights"`
}

type ComputedFData struct {
  OriginalData        [][]float64
  DailyAssetValue     [][]float64
  NetDailyAssetValue  []float64
  ProductNames        []string
  SectorToProducts    map[string] []string
}
