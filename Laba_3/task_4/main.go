package main

import (
  "fmt"
)

var maximum int32

func main() {
  var n int32

  fmt.Println("Get me n (>1)")
  fmt.Scanf("%d", &n)

  fmt.Print(n)

  counter := 0

  for true {
    if (n % 2) == 0 {
      n = chetn(n)
      counter += 1
    } else {
      n = nechet(n)
      counter += 2
    }
    if n == 1 {
      break
    }
  }

  fmt.Println("")
  fmt.Print("Steps: ")
  fmt.Println(counter)
  fmt.Print("Maximum: ")
  fmt.Println(maximum)
}

//2 steps
func nechet(n int32) int32 {
  n *= 3

  if n > maximum {
    maximum = n
  }

  fmt.Print(" => ")
  fmt.Print(n)
  n += 1

  if n > maximum {
    maximum = n
  }

  fmt.Print(" => ")
  fmt.Print(n)

  return n
}

//1 step
func chetn(n int32) int32 {
  n = n / 2

  if n > maximum {
    maximum = n
  }

  fmt.Print(" => ")
  fmt.Print(n)

  return n
}
