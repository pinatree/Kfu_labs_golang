package main

import (
  "fmt"
  "math/rand"
)

func main(){
  const count int = 10
  var min int
  var max int

  slice := make([][]int, count, count)

  fmt.Println("Matrix with size 10 X 10")
  fmt.Println("Get me min value for random")
  fmt.Scanf("%d", &min)
  fmt.Println("Get me max value for random")
  fmt.Scanf("%d", &max)

  fmt.Println("______________________________________")
  for x := 0; x < count; x++ {
    slice[x] = make([]int, count, count)
    for y := 0; y < count; y++ {
      slice[x][y] = rand.Intn(max - min + 1) + min

      fmt.Print(slice[x][y])
      fmt.Print(" | ")
    }
    fmt.Println("")
  }
  fmt.Println("______________________________________")

  local_minimum_count := 0

  for x := 0; x < count; x++ {
    for y := 0; y < count; y++ {
      is_minimum := true

      //check top
      if x != 0 {
        if slice[x - 1][y] < slice[x][y] {
          is_minimum = false
        }
      }

      //check left
      if y != 0 {
        if slice[x][y - 1] < slice[x][y] {
          is_minimum = false
        }
      }

      //check right
      if y != (count - 1) {
        if slice[x][y + 1] < slice[x][y] {
          is_minimum = false
        }
      }

      //check bottom
      if x != (count - 1) {
        if slice[x + 1][y] < slice[x][y] {
          is_minimum = false
        }
      }

      if is_minimum {
        local_minimum_count += 1
      }
    }
  }

  fmt.Print("Local mimimum's count = ")
  fmt.Println(local_minimum_count)
}
