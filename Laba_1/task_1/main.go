package main

import(
  "fmt"
)

func main(){
  var r1,r2,r3,result float64
  fmt.Println("Give me r1, r2 and r3")
  fmt.Scanf("%f", &r1)
  fmt.Scanf("%f", &r2)
  fmt.Scanf("%f", &r3)
  result = 1/((1/r1)+(1/r2)+(1/r3))
  fmt.Print("R = ")
  fmt.Println(result)
}
