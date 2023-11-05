package main

import (
 "fmt"
 "math"
)

func main() {
 a := 2.25
 deltaX := 0.3
 xStart := 1.2
 xEnd := 2.7

 for x := xStart; x <= xEnd; x += deltaX {
  result := math.Pow(a, math.Pow(x*x-1, 1)) - math.Log10(x*x-1) + math.Pow(x*x-1, 1.0/3.0)
  fmt.Printf("x = %.2f, y = %.2f\n", x, result)
 }
}
