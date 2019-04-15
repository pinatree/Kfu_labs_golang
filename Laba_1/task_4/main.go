package main

import(
  "fmt"
  "math"
)

func main(){
  var x1,x2,x3,y1,y2,y3,P,S float64

  fmt.Println("Give me x1, y1")
  fmt.Scanf("%f", &x1)
  fmt.Scanf("%f", &y1)

  fmt.Println("Give me x2, y2")
  fmt.Scanf("%f", &x2)
  fmt.Scanf("%f", &y2)

  fmt.Println("Give me x3, y3")
  fmt.Scanf("%f", &x3)
  fmt.Scanf("%f", &y3)

  a := math.Sqrt((x2 - x1)*(x2 - x1) + (y2 - y1) * (y2 - y1))
  b := math.Sqrt((x3 - x2)*(x3 - x2) + (y3 - y2) * (y3 - y2))
  c := math.Sqrt((x3 - x1)*(x3 - x1) + (y3 - y1) * (y3 - y1))

  P = a + b + c
  pinadva := P / 2
  S = math.Sqrt(pinadva * (pinadva - a) * (pinadva - b) * (pinadva - c))

  fmt.Print("a = ")
  fmt.Println(a)

  fmt.Print("b = ")
  fmt.Println(b)

  fmt.Print("c = ")
  fmt.Println(c)

  fmt.Print("P = ")
  fmt.Println(P)

  fmt.Print("S = ")
  fmt.Println(S)
}
