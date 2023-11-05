package main

import (
 "fmt"
 "math"
)

func main() {
 a := 2.25
 xValues := []float64{1.31, 1.39, 1.44, 1.56, 1.92}

 for _, x := range xValues {
  result := math.Pow(a, math.Pow(x*x-1, 1)) - math.Log10(x*x-1) + math.Pow(x*x-1, 1.0/3.0)
  fmt.Printf("x = %.2f, y = %.2f\n", x, result)
 }
}