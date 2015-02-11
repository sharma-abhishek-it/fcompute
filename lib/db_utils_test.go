package fcompute

import (
  "testing"
  "reflect"
  "os"
  "github.com/garyburd/redigo/redis"
)

func setup_test_db() {
  conn,_ := redis.Dial("tcp", os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"))

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
  conn.Do("FLUSHDB")
}

func TestDBFetchingSectorToProductMapping(t *testing.T) {
  setup_test_db()

  expected := map[string][]string {
    "Automobiles": []string{"Hyundai", "Suzuki"},
    "Medicals":    []string{"Glaxo", "Himalaya"},
  }

  fData := GetBookKeepingData()

  if !reflect.DeepEqual(fData.SectorToProducts, expected) {
    t.Error("Map from db is different -> ", fData.SectorToProducts, expected)
  }

  teardown_test_db()
}

func TestDBProductNamesOrder(t *testing.T) {
  setup_test_db()

  expected := []string {"Hyundai", "Suzuki", "Glaxo", "Himalaya"}

  fData := GetBookKeepingData()

  if !reflect.DeepEqual(fData.ProductNames, expected) {
    t.Error("Product names order from db is different -> ", fData.ProductNames, expected)
  }

  teardown_test_db()
}

func TestDBOriginalProductsData(t *testing.T) {
  setup_test_db()

  expected := [][]float64 {
    []float64{10,20,30}, // Hyundai
    []float64{18,19,20}, // Suzuki
    []float64{3,4,5},    // Glaxo
    []float64{10,9,8},   // Himalaya
  }

  fData := GetBookKeepingData()

  if !reflect.DeepEqual(fData.OriginalData, expected) {
    t.Error("Product data from db is different -> ", fData.OriginalData, expected)
  }

  teardown_test_db()
}
