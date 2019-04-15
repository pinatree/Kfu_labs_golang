package main

import (
  "fmt"
  "os"
  "bufio"
  "encoding/json"
  "io/ioutil"
  "strings"
  "time"
  "math"
)

type employee struct {
  FIO string
  WorkPlace string
  YearWorkingFrom int
  Salary int
}

type employees_database struct {
  Employees []employee
}

var database_path string

func main(){

  file_path, _ := os.Getwd()
  database_path = file_path + "/employees.json"

  initDatabase(database_path)

  fmt.Println("Welcome to my program! Write 'help' to get commands")

  for true {
    var command string
    myscanner := bufio.NewScanner(os.Stdin)
    myscanner.Scan()
    command = myscanner.Text()

    switch command {
      case "help": {
        fmt.Println("'help' to get commands")
        fmt.Println("'addnew' to add new employee")
        fmt.Println("'search' to find employee")
        fmt.Println("'get_senior_managers' to get manager whom works more 4 years")
        fmt.Println("'get_highest_salary_emp' to get employee with highest salary")
        fmt.Println("'exit' to... exit program")
        break
      }
      case "addnew": {
        addNewEmployee()
        break
      }
      case "search": {
        searchEmployee()
        break
      }
      case "get_senior_managers": {
        getSeniorManagers()
        break
      }
      case "get_highest_salary_emp": {
        get_highest_salary_emp()
        break
      }
      case "exit": {
        fmt.Println("Good bye")
        return
      }
    }
    fmt.Println()
  }
}

func get_highest_salary_emp() {
  db := getDataBase(database_path)

  if len(db.Employees) == 0 {
    fmt.Println("There is no employees")
    return
  }

  salary_max := math.MinInt32
  salary_max_index := -1

  for x := 0; x < len(db.Employees); x++ {
    if db.Employees[x].Salary > salary_max {
      salary_max = db.Employees[x].Salary
      salary_max_index = x
    }
  }

  fmt.Print("Subname and initials: ")
  fmt.Println(db.Employees[salary_max_index].FIO)

  fmt.Print("Workplace: ")
  fmt.Println(db.Employees[salary_max_index].WorkPlace)

  fmt.Print("Year of begining work: ")
  fmt.Println(db.Employees[salary_max_index].YearWorkingFrom)

  fmt.Print("Salary: ")
  fmt.Println(db.Employees[salary_max_index].Salary)

  fmt.Println()
}

func getSeniorManagers() {
  db := getDataBase(database_path)
  year, _, _ := time.Now().Date()
  count := 0
  for _, item := range db.Employees {
    if (year - item.YearWorkingFrom > 4) && (item.WorkPlace == "manager") {
      count++

      fmt.Print("Subname and initials: ")
      fmt.Println(item.FIO)

      fmt.Print("Workplace: ")
      fmt.Println(item.WorkPlace)

      fmt.Print("Year of begining work: ")
      fmt.Println(item.YearWorkingFrom)

      fmt.Print("Salary: ")
      fmt.Println(item.Salary)

      fmt.Println()
    }
  }
  if count == 0 {
    fmt.Println("There is no senior managers")
  }
}

func searchEmployee() {
  var subname string
  myscanner := bufio.NewScanner(os.Stdin)

  fmt.Println("Get me subname of employee, and i will try search him")

  myscanner.Scan()
  subname = myscanner.Text()

  db := getDataBase(database_path)
  count := 0
  for _, item := range db.Employees {
    if strings.Contains(item.FIO, subname) {
      count++
      fmt.Print("Subname and initials: ")
      fmt.Println(item.FIO)

      fmt.Print("Workplace: ")
      fmt.Println(item.WorkPlace)

      fmt.Print("Year of begining work: ")
      fmt.Println(item.YearWorkingFrom)

      fmt.Print("Salary: ")
      fmt.Println(item.Salary)

      fmt.Println()
    }
  }

  if count == 0 {
    fmt.Println("Oops.. 0 employees found.")
  }
}

func addNewEmployee() {
  var fio, workplace string
  var year_from, salary int

  myscanner := bufio.NewScanner(os.Stdin)

  fmt.Println("New employee creating...")

  fmt.Println("Get subname and initials")
  myscanner.Scan()
  fio = myscanner.Text()

  fmt.Println("Get workplace")
  myscanner.Scan()
  workplace = myscanner.Text()

  fmt.Println("Get year of start working")
  fmt.Scanf("%d", &year_from)

  fmt.Println("Get salary")
  fmt.Scanf("%d", &salary)

  employee_new := employee {
    FIO: fio,
    WorkPlace: workplace,
    YearWorkingFrom: year_from,
    Salary: salary,
  }

  db := getDataBase(database_path)
  db.Employees = append(db.Employees, employee_new)

  reWriteDatabase(db)
}

func reWriteDatabase(db employees_database) {
  parsed, err := json.Marshal(db)
  if err != nil {
    panic(err.Error())
  }
  file, _ := os.Create(database_path)
  file.Write(parsed)
  fmt.Println("Database re-writed")
}

/*понимаю, что маленький json-файлик не достоин
называться целой базой данных, ну да думаю и так сойдет*/
func getDataBase(db_path string) employees_database{
  file_bytes, err := ioutil.ReadFile(db_path)

  if err != nil{
    panic(err.Error())
  }

  var readed_data employees_database

  json.Unmarshal(file_bytes, &readed_data)

  return readed_data
}


func initDatabase(init_path string){
  _, err := os.Stat(init_path)
  if err == nil {
    fmt.Println("Database file is exists.")
  } else if os.IsNotExist(err) {
    fmt.Println("Database file is not exists, creating...")

    var test_db = employees_database {

    }

    parsed, err := json.Marshal(test_db)
    if err != nil {
      panic(err.Error())
    }

    file, _ := os.Create(init_path)
    file.Write(parsed)

    fmt.Println("Database created!")
  }
}
