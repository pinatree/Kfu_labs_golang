package main

import (
  "fmt"
  "math"
  "math/rand"
)

var slice_one [][]int
var count int

func main(){
  var min int
  var max int

  fmt.Println("Get me matrix size")
  fmt.Scanf("%d", &count)

  slice_one = make([][]int, count, count)
  has_negative_variables_arr := make([]bool, count, count)

  fmt.Println("Get me min value for random")
  fmt.Scanf("%d", &min)
  fmt.Println("Get me max value for random")
  fmt.Scanf("%d", &max)

  fmt.Println("______________________")
  for x := 0; x < count; x++ {
    slice_one[x] = make([]int, count, count)
    hasNegativeVariables := false
    for y := 0; y < count; y++ {
      slice_one[x][y] = rand.Intn(max - min + 1) + min

      if slice_one[x][y] < 0 {
        hasNegativeVariables = true
      }

      fmt.Print(slice_one[x][y])
      fmt.Print(" | ")
    }
    has_negative_variables_arr[x] = hasNegativeVariables
    fmt.Println("")
  }
  fmt.Println("______________________")

  for x := 0; x < count; x++ {
    if !has_negative_variables_arr[x] {
      multiplier := getStringMultiply(slice_one, x)
      fmt.Print("String ")
      fmt.Print(x)
      fmt.Print(" has no negative elements. Multiplier: ")
      fmt.Println(multiplier)
    } else {
      fmt.Print("String ")
      fmt.Print(x)
      fmt.Println(" have negative elements.")
    }
  }
  
  procedureStyle();
}

func procedureStyle(){
  max := math.MinInt32
  variable_for_local := 0
  //check main
  for x := 0; x < count; x++ {
    variable_for_local += slice_one[x][x]
  }

  if variable_for_local > max {
    max = variable_for_local
    variable_for_local = 0
  }

  //check to left
  for x := 1; x < count; x++ {
    local := 0
    for y := x; y < (count - x); y++ {
      local += slice_one[x][y]
    }
    if local > max {
      max = local
    }
  }

  //check to down
  for x := 1; x < count; x++ {
    local := 0
    for y := x; y < (count - x); y++ {
      local += slice_one[y][x]
    }
    if local > max {
      max = local
    }
  }

  fmt.Print("Maximum from parallel lines to main diagonal....")
  fmt.Println(max)
}

func getStringMultiply(matrix [][]int, stringIndex int) int {
  result := 1

  for x := 0; x < len(matrix); x++ {
    result *= matrix[stringIndex][x]
  }

  return result
}
