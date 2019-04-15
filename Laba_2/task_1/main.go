package main

import (
  "fmt"
  "math"
)

func main(){
  var x,y float64

  fmt.Println("Get me x")
  fmt.Scanf("%f", &x)

  if x >= -3 && x < -2 {
    y = -2 * x - 5
  } else if x >= -2 && x < 0 {
    y = -1 * math.Sqrt(1 - (x + 1)*(x + 1)) - 1
  } else if x >=0 && x < 1 {
    y = x - 1
  } else if x >= 1 && x < 3 {
    y = math.Sqrt(1 - (x - 2)*(x - 2))
  } else {
    fmt.Println("X in out of range")
    return
  }

  fmt.Print("y = ")
  fmt.Println(y)
}
