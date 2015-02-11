package main

import (
  "testing"
  "net/http"
  "net/http/httptest"
  "github.com/zenazn/goji/web"
)

func testing_data()(*web.Mux, web.C) {
  handler := ReportsRoutesHandler()
  context := web.C{
    URLParams: make(map[string]string),
    Env:       make(map[interface{}]interface{}),
  }

  return handler, context
}

func TestReportsRequest_Correct_Params_200_Ok(t *testing.T) {
  request_handler, context := testing_data()

  params := "reports=[\"a\",\"b\",\"c\"]"
  params += "&prefs={\"investment\": 100000}"

  request, _ := http.NewRequest("GET", "/?"+params, nil)
  recorder   := httptest.NewRecorder()
  request_handler.ServeHTTPC(context, recorder, request)
  if recorder.Code != http.StatusOK {
      t.Errorf("API didn't return %v", http.StatusOK)
  }
}

func TestReportsRequest_Incorrect_Params_400_BadRequest(t *testing.T) {
  request_handler, context := testing_data()
  params := "reports=[\"a\",\"b\",\"c\"]"

  request, _ := http.NewRequest("GET", "/?"+params, nil)
  recorder   := httptest.NewRecorder()
  request_handler.ServeHTTPC(context, recorder, request)
  if recorder.Code != http.StatusBadRequest {
      t.Errorf("API didn't return %v", http.StatusOK)
  }
}

func BenchmarkReportsRequest(b *testing.B) {
  request_handler, context := testing_data()
  for i := 0; i < b.N; i++ {
    params := "reports=[\"a\",\"b\",\"c\"]"
    params += "&prefs={\"investment\": 100000}"

    request, _ := http.NewRequest("GET", "/?"+params, nil)
    recorder   := httptest.NewRecorder()
    request_handler.ServeHTTPC(context, recorder, request)
  }
}
