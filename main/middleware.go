package main

import (
  "net/http"
  "encoding/json"
  "github.com/zenazn/goji/web"
  "fcompute/lib"
)

func UserPrefsParser(c *web.C, h http.Handler) http.Handler {
  handler := func(w http.ResponseWriter, r *http.Request) {
    prefs := fcompute.UserPrefs {}
    err := json.Unmarshal([]byte(r.URL.Query().Get("prefs")), &prefs)

    if err == nil {
      c.Env["prefs"] = prefs
    } else {
      http.Error(w, "{'error': 'User preferences not passed correctly!'}", http.StatusBadRequest)
    }
    h.ServeHTTP(w, r)
  }
  return http.HandlerFunc(handler)
}

func ReportNamesParser(c *web.C, h http.Handler) http.Handler {
  handler := func(w http.ResponseWriter, r *http.Request) {
    report_names := []string{}
    err := json.Unmarshal([]byte(r.URL.Query().Get("reports")), &report_names)

    if err == nil {
      c.Env["reports"] = report_names
    } else {
      http.Error(w, "{'error': 'Report names not passed correctly!'}", http.StatusBadRequest)
    }
    h.ServeHTTP(w, r)
  }
  return http.HandlerFunc(handler)
}
