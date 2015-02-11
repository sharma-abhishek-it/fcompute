package main

import (
  "fmt"
  "net/http"
  "github.com/zenazn/goji/web"
  "fcompute/lib"
)

func ReportsController(c web.C, w http.ResponseWriter, r *http.Request) {
  report_names, prefs := []string{}, fcompute.UserPrefs {}
  if c.Env["prefs"] == nil || c.Env["reports"] == nil {
    return
  } else {
    prefs = c.Env["prefs"].(fcompute.UserPrefs)
    report_names = c.Env["reports"].([]string)
  }
  if (len(report_names) > 0 && prefs.Investment != 0) {
    fmt.Fprintf(w, "TODO: NOT YET IMPLEMENTED")
  }

}
