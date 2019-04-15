package main

import (
  "fmt"
)

func main(){
  var count int
  var slice []int

  fmt.Println("Get me items count")
  fmt.Scanf("%d", &count)

  for x := 0; x < count; x++ {
    var nextint int
    fmt.Print("Get me item №")
    fmt.Println(x + 1)
    fmt.Scanf("%d", &nextint)
    slice = append(slice, nextint)
  }

  sum := 0
  otritz_count := 0

  for _, x := range slice {
    if x > 0 {
      sum += x
    } else if x < 0 {
      otritz_count = otritz_count + 1
    }
  }

  fmt.Print("Sum of positive elements = ")
  fmt.Println(sum)

  fmt.Print("Count of negative elements = ")
  fmt.Println(otritz_count)
}
