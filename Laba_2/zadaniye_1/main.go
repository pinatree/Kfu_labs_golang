package main

import(
  "fmt"
  "math"
)

func main() {

  var x, y, result float64

    fmt.Println("Give me x and y")
    fmt.Scanf("%f", &x)
    fmt.Scanf("%f", &y)

  if x > y {
    result = math.Atan(x) + (x-y)*(x-y)*(x-y)
  } else if x < y {
    result = math.Tan(x) + (y-x)*(y-x)*(y-x)
  } else {
    result = (y+x)*(y+x)*(y+x) + 0.5
  }

  fmt.Print("Result = ")
  fmt.Println(result)
}
