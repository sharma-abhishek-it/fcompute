package main

import (
  "fmt"
  "net/http"
  "github.com/zenazn/goji/web"
  "fcompute/lib"
  "encoding/json"
)

func ReportsController(c web.C, w http.ResponseWriter, r *http.Request) {
  report_names, prefs := []string{}, fcompute.UserPrefs {}
  if c.Env["prefs"] == nil || c.Env["reports"] == nil {
    return
  } else {
    prefs = c.Env["prefs"].(fcompute.UserPrefs)
    report_names = c.Env["reports"].([]string)
  }

  reports := []interface{} {}
  fData := fcompute.GetBookKeepingData()
  fData.PreCompute(prefs)

  for _, name := range report_names {
    var report interface{}
    switch name {
      case "pnl_report":
        r := fcompute.ProfitNLossReport{}
        r.Generate(fData)
        report = interface{}(r)
      case "net_returns_report":
        r := fcompute.NetReturnsReport{}
        r.Generate(fData)
        report = interface{}(r)
      case "annualized_returns_report":
        r := fcompute.AnnualizedReturnsReport{}
        r.Generate(fData)
        report = interface{}(r)
      case "max_drawdown_report":
        r := fcompute.MaximumDrawdownReport{}
        r.Generate(fData)
        report = interface{}(r)
    }

    if report != nil {
      reports = append(reports, report)
    }
  }
  json_response, err := json.Marshal(reports)

  if err != nil {
    http.Error(w, "{'error': 'Something bad happened'}", http.StatusInternalServerError)
  } else {
    fmt.Fprintf(w, string(json_response))
  }
}
