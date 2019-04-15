package main

import (
  "fmt"
  "math"
)

func main(){
  var a,b,c, max float64

  fmt.Println("Get me three lines length")
  fmt.Scanf("%f", &a)
  fmt.Scanf("%f", &b)
  fmt.Scanf("%f", &c)

  if (a > b + c) || (b > a + c) || (c > a + b) {
    fmt.Println("Triangle is not correct")
    return
  } else {
    fmt.Println("Triangle is correct")
  }

  max = math.Max(c, math.Max(a,b))

  result := 0.0

  if a == max {
    result = (b*b + c*c - a*a) / (2 * b * c)
  } else if b == max {
    result = (a*a + c*c - b*b) / (2 * a * c)
  } else if c == max {
    result = (a*a + b*b - c*c) / (2 * a * b)
  }

  if result > 0 {
    fmt.Println("Sharp trangle")
  } else if result == 0 {
    fmt.Println("Straight trangle")
  } else {
      fmt.Println("Dull trangle")
  }

}
