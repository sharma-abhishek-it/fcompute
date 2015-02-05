package fcompute

import (
  "testing"
  "reflect"
  "github.com/garyburd/redigo/redis"
)

func setup_test_db() {
  conn,_ := redis.Dial("tcp", "127.0.0.1:6379")

  conn.Send("RPUSH",  "all_sectors",  "Automobiles","Medicals")

  conn.Send("RPUSH",  "Automobiles:Products", "Hyundai",  "Suzuki")
  conn.Send("HMSET", "Automobiles:Hyundai", "name", "Hyundai",  "data", "[10,20,30]")
  conn.Send("HMSET", "Automobiles:Suzuki",  "name", "Suzuki",   "data", "[18,19,20]")


  conn.Send("RPUSH",  "Medicals:Products", "Himalaya", "Glaxo")
  conn.Send("HMSET", "Medicals:Himalaya",   "name", "Himalaya", "data", "[10,9,8]")
  conn.Send("HMSET", "Medicals:Glaxo",      "name", "Glaxo",    "data", "[3,4,5]")

  conn.Flush()
}

func teardown_test_db() {
  conn,_ := redis.Dial("tcp", "127.0.0.1:6379")
  conn.Do("FLUSHDB")
}

func TestDBFetchingSectorToProductMapping(t *testing.T) {
  setup_test_db()

  expected := map[string][]string {
    "Automobiles": []string{"Hyundai", "Suzuki"},
    "Medicals":    []string{"Glaxo", "Himalaya"},
  }

  _,_,sectorToProducts := GetBookKeepingData(0)

  if !reflect.DeepEqual(sectorToProducts, expected) {
    t.Error("Map from db is different -> ", sectorToProducts, expected)
  }

  teardown_test_db()
}

func TestDBProductNamesOrder(t *testing.T) {
  setup_test_db()

  expected := []string {"Hyundai", "Suzuki", "Glaxo", "Himalaya"}

  _,productNames,_ := GetBookKeepingData(0)

  if !reflect.DeepEqual(productNames, expected) {
    t.Error("Product names order from db is different -> ", productNames, expected)
  }

  teardown_test_db()
}

func TestDBProductsFData(t *testing.T) {
  setup_test_db()

  expected := [][]float64 {
    []float64{10,20,30}, // Hyundai
    []float64{18,19,20}, // Suzuki
    []float64{3,4,5},    // Glaxo
    []float64{10,9,8},   // Himalaya
  }

  productsFData,_,_ := GetBookKeepingData(0)

  if !reflect.DeepEqual(productsFData, expected) {
    t.Error("Product data from db is different -> ", productsFData, expected)
  }

  teardown_test_db()
}
