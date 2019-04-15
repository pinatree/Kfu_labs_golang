package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "math"
)

func main(){
  var working_str string

  fmt.Println("Get string")
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

  var min_word string
  current_min_length := math.MaxInt32

  for _, x := range str_parsed {
    if len(x) < current_min_length {
      current_min_length = len(x)
      min_word = x
    }
  }

  fmt.Println("The shortest word: ")
  fmt.Println(min_word)
}
