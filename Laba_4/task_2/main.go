package main

import (
  "fmt"
)

func main(){
  var count int

  fmt.Println("Get me items count")
  fmt.Scanf("%d", &count)

  slice := make([]int, count, count)

  for x := 0; x < count; x++ {
    var nextint int
    fmt.Print("Get me item №")
    fmt.Println(x + 1)
    fmt.Scanf("%d", &nextint)
    slice[x] = nextint
  }

  more_seven_count := 0

  for x := 0; x < count; x++ {
    if(slice[x] > 7) {
      slice[x] = 7
      more_seven_count += 1
    }
  }

  fmt.Print("Sum of (>7) elements = ")
  fmt.Println(more_seven_count)

  fmt.Print("Result array = ")
  fmt.Println(slice)
}
