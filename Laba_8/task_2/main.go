package main

import (
  "fmt"
  "os"
  "encoding/json"
  "io/ioutil"
)

type telephone_data struct {
  Name string
  Weight int
  Price int
}

type telephones_database struct {
  Telephones []telephone_data
}

var database_path string

func main(){

  file_path, _ := os.Getwd()
  database_path = file_path + "/employees.json"

  initDatabase(database_path)

  var min_price, max_price int

  fmt.Println("Get me minimum price")
  fmt.Scanf("%d", &min_price)

  fmt.Println("Get me maximum price")
  fmt.Scanf("%d", &max_price)

  database_handle := getDataBase(database_path)

  for _, item := range database_handle.Telephones {
    if ((item.Weight < 500) && (item.Price >= min_price) && (item.Price <= max_price)) {
      fmt.Print("Mobile name: ")
      fmt.Println(item.Name)

      fmt.Print("Mobile price: ")
      fmt.Println(item.Price)

      fmt.Print("Mobole weight: ")
      fmt.Println(item.Weight)

      fmt.Println("")
    }
  }


}

/*понимаю, что маленький json-файлик не достоин
называться целой базой данных, ну да думаю и так сойдет*/
func getDataBase(db_path string) telephones_database{
  file_bytes, err := ioutil.ReadFile(db_path)

  if err != nil{
    panic(err.Error())
  }

  var readed_data telephones_database

  json.Unmarshal(file_bytes, &readed_data)

  return readed_data
}


func initDatabase(init_path string){
  var test_db = telephones_database {
    Telephones: []telephone_data {
      telephone_data{
        Name: "IPhone 10",
        Weight: 450,
        Price: 100000,
      },
      telephone_data{
        Name: "Galaxy S6",
        Weight: 450,
        Price: 60000,
      },
      telephone_data{
        Name: "ZTE cheto tam",
        Weight: 350,
        Price: 18000,
      },
      telephone_data{
        Name: "IPhone 6",
        Weight: 300,
        Price: 15000,
      },
      telephone_data{
        Name: "Samsung che to tam",
        Weight: 500,
        Price: 12000,
      },
    },
  }

  parsed, err := json.Marshal(test_db)
  if err != nil {
    panic(err.Error())
  }

  file, _ := os.Create(init_path)
  file.Write(parsed)
}
