package main

import (
  "fmt"
  "math"
)

func main(){
  var x,y float64
  fmt.Println("Get me x and y")
  fmt.Scanf("%f", &x)
  fmt.Scanf("%f", &y)

  if x < 0 {
    if (math.Abs(y) < 0.5 * math.Abs(x)) {
      fmt.Println("Point inside")
      return
    } else {
      fmt.Println("Point outside")
      return
    }
  } else {
    if math.Sqrt(y * y + x * x) < 1 {
      fmt.Println("Point inside")
      return
    } else {
      fmt.Println("Point outside")
      return
    }
  }
}
