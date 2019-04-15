package main

import (
  "fmt"
  "math"
)

func main(){
  var a, z1, z2 float64
  fmt.Println("Get me a")
  fmt.Scanf("%f", &a)

  z1 = (math.Sin(2 * a) + math.Sin(5 * a) - math.Sin(3 * a)) / (math.Cos(a) - math.Cos(3 * a) + math.Cos(5 * a))
  z2 = math.Tan(3 * a)

  fmt.Print("z1 = ")
  fmt.Println(z1)

  fmt.Print("z2 = ")
  fmt.Println(z2)
}
