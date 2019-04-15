package main

import (
  "fmt"
  "math"
)

func main() {
  var sharping int
  var pi float64

  fmt.Println("Sharp row length")
  fmt.Scanf("%d", &sharping)

  insideHooks := 1.0

  for i := 1; i <= (sharping + 1); i++ {
    if i % 2 == 0 {
      insideHooks += currentValue(float64(i))
    } else {
      insideHooks -= currentValue(float64(i))
    }
  }

  pi = 2 * math.Sqrt(3.0) * insideHooks

  fmt.Print("pi with this accuracy = ")
  fmt.Println(pi)
}

//BEGINING FROM 1'st POSITION!!!!! NOT ZERO!!!
func currentValue(position float64) float64 {
  var result float64

  result = 1 / (math.Pow(3, position) * (3 + 2 * (position - 1)))

  return result
}
