package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func main(){
  var working_str string

  fmt.Println("Get me main string")
  myscanner := bufio.NewScanner(os.Stdin)
  myscanner.Scan()
  working_str = myscanner.Text()

  for true {
    first_skob := strings.Index(working_str, "(")
    second_skob := strings.Index(working_str, ")")

    if first_skob == -1 || second_skob == -1 {
      break
    }

    working_str = working_str[:first_skob] + " " + working_str[(second_skob + 1):]
  }

  fmt.Println("Result string: ")
  fmt.Println(working_str)
}
