package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func main(){
  var working_str string
  var result_str string

  fmt.Println("Get me main string")
  myscanner := bufio.NewScanner(os.Stdin)
  myscanner.Scan()
  working_str = myscanner.Text()

  str_splited := strings.Split(working_str, " ")
  str_parsed := []string{}

  for x := 0; x < len(str_splited); x++ {
    if str_splited[x] != "" {
      str_parsed = append(str_parsed, str_splited[x])
    }
  }

  for _, x := range str_parsed {
    result_str += (x + " ")
  }

  fmt.Println("Result string: ")
  fmt.Println(result_str)
}
