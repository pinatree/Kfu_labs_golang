package main

import(
  "fmt"
  "math"
)

func main() {
  var a,b float64

  fmt.Println("Get me two variables")
  fmt.Scanf("%f", &a)
  fmt.Scanf("%f", &b)

  if (a < 0 && b < 0) {
    a = math.Abs(a)
    b = math.Abs(b)
  } else if ((a < 0 && b >= 0) || (a >= 0 && b < 0)) {
    a += 0.5
    b += 0.5
  } else if !between(a) && !between(b) {
    a = a/10
    b = b/10
  }
  fmt.Println("Alghorithmt result:")

  fmt.Print("a = ")
  fmt.Println(a)

  fmt.Print("b = ")
  fmt.Println(b)
}

func between (x float64) bool {
  return x <= 2 && x >= 0.5
}
