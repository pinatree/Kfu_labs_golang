package main

import(
  "fmt"
)

func main() {
  var x,y float64

  fmt.Println("x and y of point")
  fmt.Scanf("%f", &x)
  fmt.Scanf("%f", &y)

  if x == 0 && y == 0 {
    fmt.Println("Point in begining of coordinates")
  } else if x != 0 && y == 0 {
      fmt.Println("Point in y line")
  } else if x == 0 && y != 0 {
      fmt.Println("Point in x line")
  } else if x > 0 && y > 0 {
      fmt.Println("Point in first quater")
  } else if x > 0 && y < 0 {
      fmt.Println("Point in second quater")
  } else if x < 0 && y < 0 {
      fmt.Println("Point in third quater")
  } else if x > 0 && y < 0 {
      fmt.Println("Point in fourth quater")
  }
}
