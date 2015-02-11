package fcompute

func even(x int) bool { return x % 2 == 0 }
func odd(x int) bool  { return x % 2 != 0 }

func mean(arr []float64) float64 {
  sum := 0.0
  for _,v := range arr { sum += v }
  return sum / float64(len(arr))
}
