package main

import (
  "fmt"
  "os"
  "bufio"
  "encoding/json"
  "io/ioutil"
)

type user_data struct {
  FIO string
  Birth_year int
  Salary int
}

type users_database struct {
  Users []user_data
}

var database_path string

func main(){

  file_path, _ := os.Getwd()
  database_path = file_path + "/employees.json"

  initDatabase(database_path)

  var lastname_path string

  fmt.Println("Get me the begining of last name (in russian language its call 'familiya')")
  myscanner := bufio.NewScanner(os.Stdin)
  myscanner.Scan()
  lastname_path = myscanner.Text()

  database_handle := getDataBase(database_path)

  sum_salary := 0
  employees_count := 0

  for _, item := range database_handle.Users {
    if item.FIO[:len(lastname_path)] == lastname_path {
      fmt.Print("Family, name and patronimyc: ")
      fmt.Println(item.FIO)

      fmt.Print("Employee salary: ")
      fmt.Println(item.Salary)

      fmt.Print("Employee birth date: ")
      fmt.Println(item.Birth_year)

      sum_salary += item.Salary
      employees_count += 1

      fmt.Println("")
      fmt.Println("")
    }
  }

  fmt.Print("Employees found: ")
  fmt.Println(employees_count)

  fmt.Print("Aberage salary: ")
  fmt.Println(float32(sum_salary) / float32(employees_count))

}

/*понимаю, что маленький json-файлик не достоин
называться целой базой данных, ну да думаю и так сойдет*/
func getDataBase(db_path string) users_database{
  file_bytes, err := ioutil.ReadFile(db_path)

  if err != nil{
    panic(err.Error())
  }

  var readed_data users_database

  json.Unmarshal(file_bytes, &readed_data)

  return readed_data
}


func initDatabase(init_path string){
  var test_db = users_database {
    Users: []user_data {
      user_data{
        FIO: "Abramov Alexey Ivanovich",
        Birth_year: 1990,
        Salary: 40000,
      },
      user_data{
        FIO: "Abramenko Ivan Alexeevych",
        Birth_year: 1995,
        Salary: 37500,
      },
      user_data{
        FIO: "Alexandrov Ivan Petrovych",
        Birth_year: 1987,
        Salary: 50000,
      },
      user_data{
        FIO: "Belov Viktor Mikhailovych",
        Birth_year: 1970,
        Salary: 120000,
      },
      user_data{
        FIO: "Vasilyev Eugeniy Petrovych",
        Birth_year: 1985,
        Salary: 49000,
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
