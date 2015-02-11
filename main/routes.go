package main

import (
  "github.com/zenazn/goji/web"
  "github.com/zenazn/goji/web/middleware"
)

func ReportsRoutesHandler() *web.Mux {
  reports := web.New()
  reports.Use(middleware.SubRouter)
  reports.Use(ReportNamesParser)
  reports.Use(UserPrefsParser)
  reports.Get("/", ReportsController)

  return reports
}
