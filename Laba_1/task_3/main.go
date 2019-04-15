package main

import(
  "fmt"
  "math"
)

func main(){
  var r_in, r_out, S float64
  r_in = 20

  for true {
    fmt.Println("Give me outside radius (>20)")
    fmt.Scanf("%f", &r_out)
    if r_out > 20 {
      break
    }
    fmt.Printf("Outside radius must me more 20!")
  }

  S = math.Pi * (r_out - r_in) * (r_out + r_in)

  fmt.Print("S = ")
  fmt.Println(S)
}
