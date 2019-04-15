package main

import (
  "fmt"
  "math/rand"
)

func main(){
  const count int = 6
  var min_1 int
  var max_1 int

  var min_2 int
  var max_2 int

  slice_one := make([][]int, count, count)
  slice_two := make([][]int, count, count)
  slice_result := make([]int, count, count)

  fmt.Println("Matrix with size 6 X 6")

  fmt.Println("Get me min value for random in FIRST matrix")
  fmt.Scanf("%d", &min_1)
  fmt.Println("Get me max value for random in FIRST matrix")
  fmt.Scanf("%d", &max_1)

  fmt.Println("Get me min value for random in SECOND matrix")
  fmt.Scanf("%d", &min_2)
  fmt.Println("Get me max value for random in SECOND matrix")
  fmt.Scanf("%d", &max_2)

  fmt.Println("______________________")
  for x := 0; x < count; x++ {
    slice_one[x] = make([]int, count, count)
    for y := 0; y < count; y++ {
      slice_one[x][y] = rand.Intn(max_1 - min_1 + 1) + min_1

      fmt.Print(slice_one[x][y])
      fmt.Print(" | ")
    }
    fmt.Println("")
  }
  fmt.Println("______________________")

  fmt.Println("______________________")
  for x := 0; x < count; x++ {
    slice_two[x] = make([]int, count, count)
    for y := 0; y < count; y++ {
      slice_two[x][y] = rand.Intn(max_2 - min_2 + 1) + min_2

      fmt.Print(slice_two[x][y])
      fmt.Print(" | ")
    }
    fmt.Println("")
  }
  fmt.Println("______________________")

  for x := 0; x < count; x++ {
    all_more := true
    for y := 0; y < count; y++ {
      if slice_one[x][y] <= slice_two[x][y] {
        all_more = false
        break
      }
    }
    if all_more {
      slice_result[x] = 1;
    } else {
      slice_result[x] = 0
    }
  }

  fmt.Println("Result array = ")
  fmt.Println(slice_result)

}
