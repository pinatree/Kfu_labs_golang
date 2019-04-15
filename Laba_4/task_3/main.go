package main

import (
  "fmt"
  "math/rand"
)

func main(){
  const count int = 95
  var min, max int

  fmt.Println("For your convenience, i will generate 95 numbers. Get me max and" +
     "min value, it will determine the probability of double and triple zero.")

  fmt.Println("Get me min value")
  fmt.Scanf("%d", &min)
  fmt.Println("Get me max value")
  fmt.Scanf("%d", &max)

  slice := make([]int, count, count)

  for x := 0; x < count; x++ {
    slice[x] = rand.Intn(max - min) + min
  }
  fmt.Println("Result array")
  fmt.Println(slice)

  double_checked := false
  triple_checked := false

  current_zeros := 0

  for x := 0; x < count; x++ {
    if slice[x] == 0 {
      current_zeros += 1
      if current_zeros >= 2 {
        double_checked = true
        if current_zeros >= 3 {
          triple_checked = true
          break
        }
      }
    } else {
      current_zeros = 0
    }
  }

  if triple_checked {
    fmt.Println("Array has triple and double zeros")
  } else if double_checked {
    fmt.Println("Array has only double zeros")
  } else {
    fmt.Println("Array hasn't double or more- zeros.")
  }
}
