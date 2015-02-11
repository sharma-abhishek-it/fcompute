package main

import (
  "github.com/zenazn/goji"
)

func main() {
  goji.Handle("/reports/*", ReportsRoutesHandler())
  goji.Serve()
}
