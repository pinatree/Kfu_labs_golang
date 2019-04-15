package main

import (
  "fmt"
  "math"
)

var x float64

func main() {
  var sharping int

  fmt.Println("Get row length")
  fmt.Scanf("%d", &sharping)

  fmt.Println("Get me x")
  fmt.Scanf("%f", &x)

  result := 1.0

  for i := 1; i <= (sharping + 1); i++ {
    if i % 2 == 0 {
      result += currentValue(i)
    } else {
      result -= currentValue(i)
    }
  }

  fmt.Print("result = ")
  fmt.Println(result)
}

//BEGINING FROM 1'st POSITION!!!!! NOT ZERO!!!
func currentValue(position int) float64 {
  var result float64

  result = math.Pow(x, float64(position)) / float64(getFact(position))

  return result
}

func getFact(n int) int {
  result := 1

  for x := 1; x <= n; x++ {
    result *= x
  }

  return result
}
