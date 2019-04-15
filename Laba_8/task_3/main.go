package main

import (
  "fmt"
  "os"
  "encoding/json"
  "bufio"
  "io/ioutil"
  "strings"
)

type book struct {
  Name string
  Author_FIO string
  Publisher string
  Price int
}

type book_database struct {
  Books []book
}

var database_path string

func main(){

  file_path, _ := os.Getwd()
  database_path = file_path + "/employees.json"

  initDatabase(database_path)

  var author, publisher string
  fmt.Println("Get me author, whom you like")
  myscanner := bufio.NewScanner(os.Stdin)
  myscanner.Scan()
  author = myscanner.Text()

  fmt.Println("Get me publisher, whom you don't like")
  myscanner.Scan()
  publisher = myscanner.Text()

  database_handle := getDataBase(database_path)

  for _, item := range database_handle.Books {
    if ((item.Publisher != publisher) && (strings.Contains(item.Author_FIO, author))) {
      fmt.Print("Book name: ")
      fmt.Println(item.Name)

      fmt.Print("Book author: ")
      fmt.Println(item.Author_FIO)

      fmt.Print("Book publisher: ")
      fmt.Println(item.Publisher)

      fmt.Print("Book price: ")
      fmt.Println(item.Price)

      fmt.Println("")
    }
  }
}

/*понимаю, что маленький json-файлик не достоин
называться целой базой данных, ну да думаю и так сойдет*/
func getDataBase(db_path string) book_database{
  file_bytes, err := ioutil.ReadFile(db_path)

  if err != nil{
    panic(err.Error())
  }

  var readed_data book_database

  json.Unmarshal(file_bytes, &readed_data)

  return readed_data
}


func initDatabase(init_path string){
  var test_db = book_database {
    Books: []book {
      book{
        Name: "Zov ktulhu",
        Author_FIO: "Govard Lovecraft",
        Publisher: "Eksmo",
        Price: 200,
      },
      book{
        Name: "Hyperion quadrology",
        Author_FIO: "Den Symons",
        Publisher: "Mir",
        Price: 1200,
      },
      book{
        Name: "Metro trilogy",
        Author_FIO: "Dmitry Glukhovsky",
        Publisher: "Eksmo",
        Price: 1100,
      },
      book{
        Name: "Ridges of madness",
        Author_FIO: "Govard Lovecraft",
        Publisher: "Eksmo",
        Price: 300,
      },
      book{
        Name: "Dunwich horor",
        Author_FIO: "Govard Lovecraft",
        Publisher: "Mir",
        Price: 300,
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
