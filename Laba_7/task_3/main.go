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
  Arr []int32
}

func main(){
  file_path, _ := os.Getwd()
  file_path = file_path + "/f"

  var arr_size int
  fmt.Println("Get me (10 positive, 10 negative numbers) period count")
  fmt.Scanf("%d", &arr_size)

  var data user_data

  data.Arr = make([]int32, 0)

  fmt.Println("New array elements: ")
  for i := 0; i < arr_size; i++ {
    for x := 0; x < 10; x++ {
      data.Arr = append(data.Arr, rand.Int31n(8) + 1)
    }
    for x := 0; x < 10; x++ {
      data.Arr = append(data.Arr, ((-1) * (rand.Int31n(8) + 1)))
    }
  }

  fmt.Println(data)

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

  positive := make([]int32, 0, 0)
  negative := make([]int32, 0, 0)

  cycle_counter := 0
  for true {
    //positive
    for n := 0; n < 10; n++ {
      positive = append(positive, readed_data.Arr[cycle_counter])
      cycle_counter++
    }
    //negative
    for n := 0; n < 10; n++ {
      negative = append(negative, readed_data.Arr[cycle_counter])
      cycle_counter++
    }
    if cycle_counter >= len(data.Arr) - 1 {
      break
    }
  }

  var new_data []int32
  for x := 0; x < (len(positive) / 5); x++ {
    for n := 0; n < 5; n++ {
      new_data = append(new_data, positive[x * 5 + n])
    }
    for n := 0; n < 5; n++ {
      new_data = append(new_data, negative[x * 5 + n])
    }
  }

  fmt.Println("New data:")
  fmt.Println(new_data)

  var new_user_data user_data
  new_user_data.Arr = new_data

  fmt.Println("Marshalling new data...")

  parsed_new, err := json.Marshal(new_user_data)
  if err != nil {
    panic(err.Error())
  }

  fmt.Println("Creating new file")
  new_direct, _ := os.Getwd()
  new_direct = new_direct + "/g"
  file_new, _ := os.Create(new_direct)
  file_new.Write(parsed_new)
  fmt.Println("All done")
}
