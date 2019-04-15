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

  b_count := 0

  result_str += working_str[0 : 1]
    
  for x := 1; x < len(working_str); x++ {
    if (strings.ToUpper(working_str[x : x + 1]) == "B") && (strings.ToUpper(working_str[x - 1 : x]) == "A") {
      b_count += 1
    } else {
      result_str += working_str[x : x + 1]
    }
  }

  fmt.Println("'ab' combinations count: ")
  fmt.Println(b_count)

  fmt.Println("Result string: ")
  fmt.Println(result_str)
}
