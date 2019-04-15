package main

import (
  "fmt"
  "math"
)

func main(){
  var func_min = -3.0
  var func_max = 3.0
  var calculating_period = 0.25

  fmt.Println("Get me start (minimal) value")
  fmt.Scanf("%f", &func_min)

  fmt.Println("Get me end (maximal) value")
  fmt.Scanf("%f", &func_max)

  fmt.Println("Get me calculating period")
  fmt.Scanf("%f", &calculating_period)


  for i := func_min; i <= func_max; i = i + calculating_period {
    fmt.Print("exp(")
    fmt.Print(i)
    fmt.Print(" - 5) = ")
    fmt.Println(math.Exp(i - 5))
  }

}
