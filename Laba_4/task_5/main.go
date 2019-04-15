package main

import (
  "fmt"
  "math/rand"
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

  rand_elements_count := rand.Intn(count - 1) + 1

  fmt.Print("Random-generated count...")
  fmt.Println(rand_elements_count)

  for x := 0; x < rand_elements_count; x++ {
    rand_position_after_x := rand.Intn(count - x) + x
    third := slice[rand_position_after_x]
    slice[rand_position_after_x] = slice[x]
    slice[x] = third
  }

  fmt.Print("Result array = ")
  fmt.Println(slice)
}
