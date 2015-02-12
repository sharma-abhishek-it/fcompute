package main

import (
  "testing"
  "net/http"
  "net/http/httptest"
  "github.com/zenazn/goji/web"
  "os"
  "github.com/garyburd/redigo/redis"
)

func testing_data()(*web.Mux, web.C) {
  handler := ReportsRoutesHandler()
  context := web.C{
    URLParams: make(map[string]string),
    Env:       make(map[interface{}]interface{}),
  }

  return handler, context
}

func setup_test_db() {
  conn,_ := redis.Dial("tcp", os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"))
  os.Setenv("REDIS_DB", "10")
  conn.Do("SELECT", "10")

  conn.Send("SADD",  "all_sectors",  "Automobiles","Medicals")

  conn.Send("SADD",  "Automobiles:Products", "Hyundai",  "Suzuki")
  conn.Send("HMSET", "Automobiles:Hyundai", "name", "Hyundai",  "data", "[10,20,30]")
  conn.Send("HMSET", "Automobiles:Suzuki",  "name", "Suzuki",   "data", "[18,19,20]")


  conn.Send("SADD",  "Medicals:Products", "Himalaya", "Glaxo")
  conn.Send("HMSET", "Medicals:Himalaya",   "name", "Himalaya", "data", "[10,9,8]")
  conn.Send("HMSET", "Medicals:Glaxo",      "name", "Glaxo",    "data", "[3,4,5]")

  conn.Flush()
}

func teardown_test_db() {
  conn,_ := redis.Dial("tcp", os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"))
  conn.Do("SELECT", "10")
  conn.Do("FLUSHDB")
}

func TestReportsRequest_Correct_Params_200_Ok(t *testing.T) {
  setup_test_db()
  request_handler, context := testing_data()

  params := "reports=[\"pnl_report\",\"net_returns_report\",\"max_drawdown_report\",\"annualized_returns_report\"]"
  params += "&prefs={\"investment\": 100000, \"frequency\": 2, \"weights\":[0.25, 0.25, 0.25, 0.25]}"

  request, _ := http.NewRequest("GET", "/?"+params, nil)
  recorder   := httptest.NewRecorder()
  request_handler.ServeHTTPC(context, recorder, request)
  if recorder.Code != http.StatusOK {
      t.Errorf("API didn't return %v", http.StatusOK)
  }
  teardown_test_db()
}

func TestReportsRequest_Incorrect_Params_400_BadRequest(t *testing.T) {
  setup_test_db()
  request_handler, context := testing_data()
  params := "reports=[\"a\",\"b\",\"c\"]"

  request, _ := http.NewRequest("GET", "/?"+params, nil)
  recorder   := httptest.NewRecorder()
  request_handler.ServeHTTPC(context, recorder, request)
  if recorder.Code != http.StatusBadRequest {
      t.Errorf("API didn't return %v", http.StatusBadRequest)
  }
  teardown_test_db()
}


/* This test is an integration test and should only be run with original 15 products data, uncomment
   for integration testing
*/
// func TestReportsRequest_Correct_Params_200_Ok_Original_Data(t *testing.T) {
//   os.Setenv("REDIS_DB", "1")
//   request_handler, context := testing_data()
//
//   params := "reports=[\"pnl_report\",\"net_returns_report\",\"max_drawdown_report\",\"annualized_returns_report\"]"
//   params += `&prefs={
//   "investment": 100000,
//   "frequency": 7,
//   "weights":[0.07, 0.07, 0.07, 0.07, 0.07, 0.07, 0.07, 0.07, 0.07, 0.07, 0.07, 0.07, 0.07, 0.07, 0.02],
//   "start_date": "2008-01-01",
//   "end_date": "2013-11-24"}`
//
//   request, _ := http.NewRequest("GET", "/?"+params, nil)
//   recorder   := httptest.NewRecorder()
//   request_handler.ServeHTTPC(context, recorder, request)
//   if recorder.Code != http.StatusOK {
//       t.Errorf("API didn't return %v", http.StatusOK)
//   }
//   teardown_test_db()
// }


func BenchmarkReportsRequest(b *testing.B) {
  os.Setenv("REDIS_DB", "1")
  request_handler, context := testing_data()
  for i := 0; i < b.N; i++ {
    params := "reports=[\"pnl_report\",\"net_returns_report\",\"max_drawdown_report\",\"annualized_returns_report\"]"
    params += "&prefs={\"investment\": 100000, \"frequency\": 7, \"weights\":[0.07, 0.07, 0.07, 0.07, 0.07, 0.07, 0.07, 0.07, 0.07, 0.07, 0.07, 0.07, 0.07, 0.07, 0.02]}"

    request, _ := http.NewRequest("GET", "/?"+params, nil)
    recorder   := httptest.NewRecorder()
    request_handler.ServeHTTPC(context, recorder, request)
  }
}
