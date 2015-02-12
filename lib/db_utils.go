package fcompute

import(
  "github.com/garyburd/redigo/redis"
  "encoding/json"
  "os"
)

// Before calculating a minimum set of information is needed
// Products, Sectors and financial data is stored in Redis.
// With as minimal round trips as possible we try to get that data
// and make that available in a consumable format
// Also the order of product data and product names in productsFData & productNames is same
// meaning that productNames[i] maps to productsFData[i]
//
// NOTE: We use Sort to minimie round trips and hence the need for MultiBulk
func GetBookKeepingData() (fData ComputedFData) {

  fData = ComputedFData{
    OriginalData: [][]float64{},
    ProductNames: []string{},
    SectorToProducts: map[string] []string{},
  }

  // sectorToProducts = make(map[string] []string)
  conn,_ := redis.Dial("tcp", os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"))
  conn.Do("SELECT", os.Getenv("REDIS_DB"))

  // Get a list of all the sectors.
  sectors := make([]string, 0)
  reply, _ := redis.MultiBulk(conn.Do("SORT", "all_sectors", "alpha"))
  for _, x := range reply {
    var v, ok = x.([]byte)
    if ok {
        sectors = append(sectors, string(v))
    }
  }

  // Make SORT for all sectors individually which returns
  // names and values for each product of a sector
  for _, sector := range sectors {
    conn.Send("SORT", sector+":Products", "alpha", "GET", sector+":*->name", "GET", sector+":*->data")
  }
  conn.Flush()
  for _, sector := range sectors {
    reply, _ := redis.MultiBulk(conn.Receive())

    // In each receive first the name of product comes and then the data of that product
    // So even odd logic needed for the same
    name_turn, data_turn, loop_count := even, odd, 0

    for _, x := range reply {
      var v, ok = x.([]byte)

      // Save sector product mappings and product names
      if ok && name_turn(loop_count) {
        fData.SectorToProducts[sector] = append(fData.SectorToProducts[sector], string(v))
        fData.ProductNames             = append(fData.ProductNames, string(v))
      }

      // save products data.
      if ok && data_turn(loop_count){
        data := make([]float64, 0)
        json.Unmarshal(v, &data)
        fData.OriginalData = append(fData.OriginalData, data)
      }

      loop_count++
    }
  }

  return fData
}
