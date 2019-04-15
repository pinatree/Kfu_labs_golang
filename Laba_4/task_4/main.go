package main

import (
  "fmt"
)

func main(){
  const count int = 16
  slice := make([]int, count, count)
  var positive_slice []int
  var negative_slice []int

  fmt.Println("Get me 16 items")

  for x := 0; x < count; x++ {
    var nextint int
    fmt.Print("Get me item №")
    fmt.Println(x + 1)
    fmt.Scanf("%d", &nextint)
    slice[x] = nextint
  }

  for _, x := range slice {
    if x >= 0 {
      positive_slice = append(positive_slice, x)
    } else if x < 0 {
      negative_slice = append(negative_slice, x)
    }
  }

  slice = append(negative_slice, positive_slice...)

  fmt.Println("Result array:")
  fmt.Println(slice)
}
