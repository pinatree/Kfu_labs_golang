package main

import (
  "fmt"
  "math/rand"
)

func main(){
  const count int = 15
  var min int
  var max int

  slice := make([][]int, count, count)

  fmt.Println("Matrix with size 15 X 15")
  fmt.Println("Get me min value for random")
  fmt.Scanf("%d", &min)
  fmt.Println("Get me max value for random")
  fmt.Scanf("%d", &max)

  not_zero_element_x := -1
  not_zero_element_y := -1
  has_zero_elements := false

  fmt.Println("__________________________________________________________")
  for x := 0; x < count; x++ {
    slice[x] = make([]int, count, count)
    for y := 0; y < count; y++ {
      slice[x][y] = rand.Intn(max - min + 1) + min

      if slice[x][y] != 0  {
        not_zero_element_x = x
        not_zero_element_y = y
      } else {
        has_zero_elements = true
      }

      fmt.Print(slice[x][y])
      fmt.Print(" | ")
    }
    fmt.Println("")
  }
  fmt.Println("__________________________________________________________")

  if has_zero_elements {
    fmt.Println("Matrix has zero elements")
  } else {
    fmt.Println("Matrix has no zero elements")
  }

  if not_zero_element_x != -1 {
    fmt.Println("For example of not zero elements, i can give you element")
    fmt.Print("string = ")
    fmt.Print(not_zero_element_x)
    fmt.Print("; column = ")
    fmt.Println(not_zero_element_y)
  } else {
    fmt.Println("All martix elements == 0, or martix is empty")
  }
}
