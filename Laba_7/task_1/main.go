package main

import (
  "fmt"
  //"bufio"
  "os"
  "math/rand"
  "encoding/json"
  "io/ioutil"
)

type user_data struct {
  Arr []int
}

func main(){
  file_path, _ := os.Getwd()
  file_path = file_path + "/elements.json"

  fmt.Print("Direcrory of elements repository will be ")
  fmt.Println(file_path)

  //var file_path string

  fmt.Println("Generating a new file")

  var arr_size, min, max int
  var data user_data

  fmt.Println("Get me new array size")
  fmt.Scanf("%d", &arr_size)

  fmt.Println("Get me new array elements min")
  fmt.Scanf("%d", &min)

  fmt.Println("Get me new array elements max")
  fmt.Scanf("%d", &max)

  data.Arr = make([]int, arr_size, arr_size)

  fmt.Println("New array elements: ")

  for x := 0; x < arr_size; x++ {
    data.Arr[x] = rand.Intn(max - min) + min
    fmt.Print(data.Arr[x])
    fmt.Print(" ")
  }
  fmt.Println("")

  fmt.Println("Marshalling...")

  parsed, err := json.Marshal(data)
  if err != nil {
    panic(err.Error())
  }

  fmt.Println("Creating new file")

  file, _ := os.Create(file_path)
  file.Write(parsed)

  fmt.Println("Reading file bytes")

  file_bytes, err := ioutil.ReadFile(file_path)

  if err != nil{
    panic(err.Error())
  }

  var readed_data user_data

  fmt.Println("Unmarshalling")

  json.Unmarshal(file_bytes, &readed_data)

  fmt.Println("Chetn elements")

  for x := 0; x < len(readed_data.Arr); x++ {
    if x % 2 == 0 {
      fmt.Print(x)
      fmt.Print(" ")
    }
  }
  fmt.Println("")

  fmt.Println("Nechetn elements")
  for x := 0; x < len(readed_data.Arr); x++ {
    if x % 2 == 1 {
      fmt.Print(x)
      fmt.Print(" ")
    }
  }
  fmt.Println("")
}
