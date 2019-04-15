package main

import (
  "fmt"
  "os"
  "encoding/json"
  "io/ioutil"
  "sort"
)

type human struct {
  FIO string
  BirthYear int
}

type humans_database struct {
  Humans []human
}

var database_path string

func main(){

  file_path, _ := os.Getwd()
  database_path = file_path + "/employees.json"

  initDatabase(database_path)

  database_handle := getDataBase(database_path)

  fmt.Println("Unsorted array:")
  fmt.Println()

  for _, item := range database_handle.Humans {
    showHumanData(item)
  }

  sort.Slice(database_handle.Humans, func(i int, j int) bool {
    return database_handle.Humans[i].FIO < database_handle.Humans[j].FIO
  })

  fmt.Println("Sorted array:")
  fmt.Println()

  for _, item := range database_handle.Humans {
    showHumanData(item)
  }
}

func showHumanData(person human) {
  fmt.Print("Subname and initials: ")
  fmt.Println(person.FIO)

  fmt.Print("Birty year: ")
  fmt.Println(person.BirthYear)

  fmt.Println()
}

/*понимаю, что маленький json-файлик не достоин
называться целой базой данных, ну да думаю и так сойдет*/
func getDataBase(db_path string) humans_database{
  file_bytes, err := ioutil.ReadFile(db_path)

  if err != nil{
    panic(err.Error())
  }

  var readed_data humans_database

  json.Unmarshal(file_bytes, &readed_data)

  return readed_data
}


func initDatabase(init_path string){
  var test_db = humans_database {
    Humans: []human {
      human{
        FIO: "Ivanov I.I.",
        BirthYear: 1977,
      },
      human{
        FIO: "Mykhailov A.A.",
        BirthYear: 1990,
      },
      human{
        FIO: "Andropov J.K.",
        BirthYear: 1980,
      },
      human{
        FIO: "Sadykov N.I.",
        BirthYear: 1977,
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
