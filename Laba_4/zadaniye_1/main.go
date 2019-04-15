package main

import (
  "fmt"
  "math"
)

func main(){
  var count int

  fmt.Println("Get me items count")
  fmt.Scanf("%d", &count)

  slice := make([]int, count, count)

  for x := 0; x < count; x++ {
    var nextint int
    fmt.Print("Get me item №")
    fmt.Println(x)
    fmt.Scanf("%d", &nextint)
    slice[x] = nextint
  }

  nechet_sum := 0
  between_sum := 0

  first_negative_pos := -1
  last_negative_pos := -1




  for x := 0; x < count; x++ {
    if x % 2 != 0 {
      nechet_sum += slice[x]
    }

    if slice[x] < 0 {
      if first_negative_pos == -1 {
        first_negative_pos = x
      }

      last_negative_pos = x
    }
  }

  result_slice := []int{}

  for x := first_negative_pos + 1; x < last_negative_pos; x++ {
    between_sum += slice[x]
  }

  for x := 0; x < count; x++ {
    if math.Abs(float64(slice[x])) > 1 {
      result_slice = append(result_slice, slice[x])
    }
  }

  fmt.Print("Sum of nechet elements = ")
  fmt.Println(nechet_sum)

  fmt.Print("Sum betveen first and last negative position = ")
  fmt.Println(between_sum)
  
  fmt.Print("Compressed slice = ")
  fmt.Println(result_slice)
}
