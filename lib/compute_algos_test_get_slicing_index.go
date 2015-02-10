package fcompute

import (
  "testing"
  "time"
  "os"
)

func (prefs *UserPrefs) fillDates(start_date string, end_date string) {
  s, _ := time.Parse(ShortTimeFormat, start_date)
  e, _ := time.Parse(ShortTimeFormat, end_date)
  prefs.StartDate = FDataDate(s)
  prefs.EndDate   = FDataDate(e)
}

var computedDummyWithOriginalDataOnly = ComputedFData {
  OriginalData: [][]float64 {{1,2,3,4,5},{2,3,4,5,6},{3,4,5,6,7},{4,5,6,7,8}},
}

func TestSlicingIndexLogic_Behavior_When_ENV_Not_Set(t *testing.T) {
  data  := computedDummyWithOriginalDataOnly
  prefs := UserPrefs{}
  prefs.fillDates("2014-01-24", "2014-01-25")

  expected_start, expected_end := 0, 4
  got_start, got_end := data.getSlicingIndexes(prefs)

  if got_start != expected_start {
    t.Errorf("Expected output mismatched for start index | Expected = %d, Got = %d", expected_start, got_start)
  }
  if got_end != expected_end {
    t.Errorf("Expected output mismatched for end index | Expected = %d, Got = %d", expected_end, got_end)
  }
}

func TestSlicingIndexLogic_Behavior_When_ENV_Set(t *testing.T) {
  data  := computedDummyWithOriginalDataOnly
  prefs := UserPrefs{}
  os.Setenv(ENV_START_DATE_KEY, "2014-01-22")
  os.Setenv(ENV_END_DATE_KEY, "2014-01-26")

  prefs.fillDates("2014-01-24", "2014-01-25")

  expected_start, expected_end := 2, 3
  got_start, got_end := data.getSlicingIndexes(prefs)

  if got_start != expected_start {
    t.Errorf("Expected output mismatched for start index | Expected = %d, Got = %d", expected_start, got_start)
  }
  if got_end != expected_end {
    t.Errorf("Expected output mismatched for end index | Expected = %d, Got = %d", expected_end, got_end)
  }
}

func TestSlicingIndexLogic_Behavior_Corner_Case_Prefs_Start_Eqls_Global_Start(t *testing.T) {
  data  := computedDummyWithOriginalDataOnly
  prefs := UserPrefs{}
  os.Setenv(ENV_START_DATE_KEY, "2014-01-22")
  os.Setenv(ENV_END_DATE_KEY, "2014-01-26")

  prefs.fillDates("2014-01-22", "2014-01-24")

  expected_start, expected_end := 0, 2
  got_start, got_end := data.getSlicingIndexes(prefs)

  if got_start != expected_start {
    t.Errorf("Expected output mismatched for start index | Expected = %d, Got = %d", expected_start, got_start)
  }
  if got_end != expected_end {
    t.Errorf("Expected output mismatched for end index | Expected = %d, Got = %d", expected_end, got_end)
  }
}

func TestSlicingIndexLogic_Behavior_Corner_Case_Prefs_End_1GreaterThan_Global_End(t *testing.T) {
  data  := computedDummyWithOriginalDataOnly
  prefs := UserPrefs{}
  os.Setenv(ENV_START_DATE_KEY, "2014-01-22")
  os.Setenv(ENV_END_DATE_KEY, "2014-01-26")

  prefs.fillDates("2014-01-23", "2014-01-24")

  expected_start, expected_end := 1, 2
  got_start, got_end := data.getSlicingIndexes(prefs)

  if got_start != expected_start {
    t.Errorf("Expected output mismatched for start index | Expected = %d, Got = %d", expected_start, got_start)
  }
  if got_end != expected_end {
    t.Errorf("Expected output mismatched for end index | Expected = %d, Got = %d", expected_end, got_end)
  }
}


func TestSlicingIndexLogic_Behavior_Corner_Case_Prefs_End_Eqls_Global_End(t *testing.T) {
  data  := computedDummyWithOriginalDataOnly
  prefs := UserPrefs{}
  os.Setenv(ENV_START_DATE_KEY, "2014-01-22")
  os.Setenv(ENV_END_DATE_KEY, "2014-01-26")

  prefs.fillDates("2014-01-23", "2014-01-26")

  expected_start, expected_end := 1, 4
  got_start, got_end := data.getSlicingIndexes(prefs)

  if got_start != expected_start {
    t.Errorf("Expected output mismatched for start index | Expected = %d, Got = %d", expected_start, got_start)
  }
  if got_end != expected_end {
    t.Errorf("Expected output mismatched for end index | Expected = %d, Got = %d", expected_end, got_end)
  }
}

func TestSlicingIndexLogic_Behavior_Corner_Case_Prefs_End_1LessThan_Global_End(t *testing.T) {
  data  := computedDummyWithOriginalDataOnly
  prefs := UserPrefs{}
  os.Setenv(ENV_START_DATE_KEY, "2014-01-22")
  os.Setenv(ENV_END_DATE_KEY, "2014-01-26")

  prefs.fillDates("2014-01-23", "2014-01-25")

  expected_start, expected_end := 1, 3
  got_start, got_end := data.getSlicingIndexes(prefs)

  if got_start != expected_start {
    t.Errorf("Expected output mismatched for start index | Expected = %d, Got = %d", expected_start, got_start)
  }
  if got_end != expected_end {
    t.Errorf("Expected output mismatched for end index | Expected = %d, Got = %d", expected_end, got_end)
  }
}


func TestSlicingIndexLogic_Behavior_Corner_Case_Start_Date_Greater_Than_Global_End(t *testing.T) {
  data  := computedDummyWithOriginalDataOnly
  prefs := UserPrefs{}
  os.Setenv(ENV_START_DATE_KEY, "2014-01-22")
  os.Setenv(ENV_END_DATE_KEY, "2014-01-26")

  prefs.fillDates("2014-01-27", "2014-01-24")

  expected_start, expected_end := 0, 2
  got_start, got_end := data.getSlicingIndexes(prefs)

  if got_start != expected_start {
    t.Errorf("Expected output mismatched for start index | Expected = %d, Got = %d", expected_start, got_start)
  }
  if got_end != expected_end {
    t.Errorf("Expected output mismatched for end index | Expected = %d, Got = %d", expected_end, got_end)
  }
}

func TestSlicingIndexLogic_Behavior_Corner_Case_Start_Date_Greater_Than_Prefs_End(t *testing.T) {
  data  := computedDummyWithOriginalDataOnly
  prefs := UserPrefs{}
  os.Setenv(ENV_START_DATE_KEY, "2014-01-22")
  os.Setenv(ENV_END_DATE_KEY, "2014-01-26")

  prefs.fillDates("2014-01-25", "2014-01-24")

  expected_start, expected_end := 3, 4
  got_start, got_end := data.getSlicingIndexes(prefs)

  if got_start != expected_start {
    t.Errorf("Expected output mismatched for start index | Expected = %d, Got = %d", expected_start, got_start)
  }
  if got_end != expected_end {
    t.Errorf("Expected output mismatched for end index | Expected = %d, Got = %d", expected_end, got_end)
  }
}

func TestSlicingIndexLogic_Behavior_Corner_Case_End_Date_Lesser_Than_Global_Start(t *testing.T) {
  data  := computedDummyWithOriginalDataOnly
  prefs := UserPrefs{}
  os.Setenv(ENV_START_DATE_KEY, "2014-01-22")
  os.Setenv(ENV_END_DATE_KEY, "2014-01-26")

  prefs.fillDates("2014-01-24", "2014-01-21")

  expected_start, expected_end := 2, 4
  got_start, got_end := data.getSlicingIndexes(prefs)

  if got_start != expected_start {
    t.Errorf("Expected output mismatched for start index | Expected = %d, Got = %d", expected_start, got_start)
  }
  if got_end != expected_end {
    t.Errorf("Expected output mismatched for end index | Expected = %d, Got = %d", expected_end, got_end)
  }
}

func TestSlicingIndexLogic_Behavior_Corner_Case_End_Date_Lesser_Than_Prefs_Start(t *testing.T) {
  data  := computedDummyWithOriginalDataOnly
  prefs := UserPrefs{}
  os.Setenv(ENV_START_DATE_KEY, "2014-01-22")
  os.Setenv(ENV_END_DATE_KEY, "2014-01-26")

  prefs.fillDates("2014-01-23", "2014-01-22")

  expected_start, expected_end := 1, 4
  got_start, got_end := data.getSlicingIndexes(prefs)

  if got_start != expected_start {
    t.Errorf("Expected output mismatched for start index | Expected = %d, Got = %d", expected_start, got_start)
  }
  if got_end != expected_end {
    t.Errorf("Expected output mismatched for end index | Expected = %d, Got = %d", expected_end, got_end)
  }
}

func TestSlicingIndexLogic_Behavior_Corner_Case_Prefs_Dates_Equal(t *testing.T) {
  data  := computedDummyWithOriginalDataOnly
  prefs := UserPrefs{}
  os.Setenv(ENV_START_DATE_KEY, "2014-01-22")
  os.Setenv(ENV_END_DATE_KEY, "2014-01-26")

  prefs.fillDates("2014-01-24", "2014-01-24")

  expected_start, expected_end := 2, 4
  got_start, got_end := data.getSlicingIndexes(prefs)

  if got_start != expected_start {
    t.Errorf("Expected output mismatched for start index | Expected = %d, Got = %d", expected_start, got_start)
  }
  if got_end != expected_end {
    t.Errorf("Expected output mismatched for end index | Expected = %d, Got = %d", expected_end, got_end)
  }
}
