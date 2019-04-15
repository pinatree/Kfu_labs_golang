package main

import (
  "fmt"
  "math"
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

  max_index := -1
  max_value := math.MinInt32

  min_index := -1
  min_value := math.MaxInt32

  for x := 0; x < len(readed_data.Arr); x++ {
    if readed_data.Arr[x] < min_value {
      min_value = readed_data.Arr[x]
      min_index = x
    }

    if readed_data.Arr[x] > max_value {
      max_index = x
      max_value = readed_data.Arr[x]
    }
  }
  from := -1
  to := -1
  if max_index > min_index {
    from = min_index
    to = max_index
  } else {
    from = max_index
    to = min_index
  }

  new_data := make([]int, 0)
  sum := 0
  for x := from + 1; x < to; x++ {
    sum += readed_data.Arr[x]
    new_data = append(new_data, readed_data.Arr[x])
  }

  fmt.Println("Marshalling new data...")

  var new_user_data user_data
  new_user_data.Arr = new_data

  parsed_new, err := json.Marshal(new_user_data)
  if err != nil {
    panic(err.Error())
  }

  fmt.Println("Creating new file")

  file_path_new, _ := os.Getwd()
  file_path_new = file_path_new + "/range.json"

  file_new, _ := os.Create(file_path_new)
  file_new.Write(parsed_new)


}
