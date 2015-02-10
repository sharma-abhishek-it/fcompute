package fcompute

import (
  "testing"
  "encoding/json"
  "time"
  "reflect"
)
func TestJSONUnmarshallingUserPrefs(t *testing.T) {
  var input = `
  {
    "start_date": "2014-12-30",
    "end_date":   "2014-12-31",
    "frequency":  7,
    "weights":    [3.1, 4.1, 5.1],
    "investment": 1000.9
  }
  `
  expected := UserPrefs{
    StartDate: FDataDate(time.Date(2014, time.December, 30, 0,0,0,0,time.UTC)),
    EndDate:   FDataDate(time.Date(2014, time.December, 31, 0,0,0,0,time.UTC)),
    RebalanceFrequency: 7,
    Weights:   []float64{3.1, 4.1, 5.1},
    Investment: 1000.9,
  }
  got := UserPrefs{}
  json.Unmarshal([]byte(input), &got)


  if !(reflect.DeepEqual(got,expected)) {
    t.Error("Unmarshalled UserPrefs are different. They are equal on weights, start, end, freq",
      reflect.DeepEqual(got.Weights,expected.Weights),
      got.StartDate == expected.StartDate,
      got.EndDate   == expected.EndDate,
      got.RebalanceFrequency == expected.RebalanceFrequency,
    )
  }
}

func TestUserPrefsDateDifference_ToBe_Inclusive_Of_BothDates(t *testing.T) {
  t1 := FDataDate(time.Date(2014, time.December, 30, 0,0,0,0,time.UTC))
  t2 := FDataDate(time.Date(2014, time.December, 31, 0,0,0,0,time.UTC))
  expected := 2

  if t1.DaysSince(t2) != expected {
    t.Error("Difference in days is not expected", t1.DaysSince(t2), expected)
  }
}

func TestUserPrefsDateDifference_ToCalculate_Absolute_Difference_Before_or_After(t *testing.T) {
  t1 := FDataDate(time.Date(2014, time.December, 31, 0,0,0,0,time.UTC))
  t2 := FDataDate(time.Date(2014, time.December, 30, 0,0,0,0,time.UTC))
  expected := 2

  if t1.DaysSince(t2) != expected {
    t.Error("Difference in days is not expected", t1.DaysSince(t2), expected)
  }
}
