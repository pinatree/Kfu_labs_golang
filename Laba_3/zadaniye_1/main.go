package main

import (
  "fmt"
  "math"
)

func main() {
  var start, end, step float64
  var sharping int

  fmt.Println("Get me start (-1<)")
  fmt.Scanf("%f", &start)

  fmt.Println("Get me end (>=1)")
  fmt.Scanf("%f", &end)

  fmt.Println("Get me step")
  fmt.Scanf("%f", &step)

  fmt.Println("Get me sharping")
  fmt.Scanf("%d", &sharping)

  fmt.Println("X value    |    func result    |    accuracy")

  for x := start; x <= end; x = x + step {
    fmt.Print(fmt.Sprintf("%.3f", float64(x)))
    fmt.Print("         ")
    fmt.Print(fmt.Sprintf("%.3f", float64(calculateElement(x, sharping))))
    fmt.Print("              ")
    fmt.Println(fmt.Sprintf("%.3f", float64(sharping)))
  }
}

func calculateElement(val float64, sharping int) float64 {
  result := val

  for x := 2; x <= (sharping + 1); x++ {
    if x % 2 == 0 {
      result -= (math.Pow(val, float64(x)) / float64(x))
    } else {
      result += (math.Pow(val, float64(x)) / float64(x))
    }
  }
  return result
}
