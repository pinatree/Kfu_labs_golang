package main

import (
  "fmt"
)

func main(){
  var apart_num, floor, paradnaya int

  apart_per_paradnaya := 4 * 9

  fmt.Println("Give me apartament number")
  fmt.Scanf("%d", &apart_num)

  paradnaya = (apart_num - 1) / apart_per_paradnaya + 1

  floor = (apart_num - 1 - (paradnaya - 1) * apart_per_paradnaya) / 4 + 1

  fmt.Print("Paradnaya ")
  fmt.Println(paradnaya)

  fmt.Print("Floor ")
  fmt.Println(floor)
}
