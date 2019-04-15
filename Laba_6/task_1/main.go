package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func main(){
  var main_str string
  var sub_str string

  fmt.Println("Get me main string")
  myscanner := bufio.NewScanner(os.Stdin)
  myscanner.Scan()
  main_str = myscanner.Text()

  fmt.Println("Get me substring")
  myscanner.Scan()
  sub_str = myscanner.Text()

  if strings.Count(main_str, sub_str) > 0 {
    fmt.Println("First string contains second string")
  } else {
      fmt.Println("First string doesn't contains second string")
  }
}
