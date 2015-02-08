package main

import (
  "fmt"
  "net/http"
  "github.com/zenazn/goji"
  "github.com/zenazn/goji/web"
  "github.com/zenazn/goji/web/middleware"
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
  fmt.Println(report_names)
  fmt.Println(prefs)

  fmt.Fprintf(w, "TODO: NOT YET IMPLEMENTED")
}

func main() {
  reports := web.New()
  goji.Handle("/reports/*", reports)

  reports.Use(middleware.SubRouter)
  reports.Use(ReportNamesParser)
  reports.Use(UserPrefsParser)
  reports.Get("/", ReportsController)

  goji.Serve()
}
