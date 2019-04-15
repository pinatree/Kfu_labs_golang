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

  for x := 0; x < count; x++ {
    for y := 0; y < (count - 1); y++ {
      if slice[y + 1] > slice[y] {
        third := slice[y]
        slice[y] = slice[y + 1]
        slice[y + 1] = third        
      }
    }
  }

  fmt.Print("Result array = ")
  fmt.Println(slice)
}
