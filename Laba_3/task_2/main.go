package main

import (
  "fmt"
)

func main() {
  var x int

  fmt.Println("Get me value")
  fmt.Scanf("%d", &x)

  var count = 1
  var first_val = 0
  var sum = 0

  for true {
    current_symbol := x % 10
    sum += current_symbol
    x = x/10
    if x == 0 {
      first_val = current_symbol
      break
    }

    count++
  }

  fmt.Print("numbers count = ")
  fmt.Println(count)

  fmt.Print("sum of numbers = ")
  fmt.Println(sum)

  fmt.Print("first number = ")
  fmt.Println(first_val)
}
