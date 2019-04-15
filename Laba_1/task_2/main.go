package main

import(
  "fmt"
  "math"
  "time"
)

func main(){
  fmt.Println("Give me first hour, minute and second")
  first_time := gettime()

  fmt.Println("Give me second hour, minute and sec")
  second_time := gettime()

  duration := first_time.Sub(second_time)

  fmt.Println("Different between first and second time:")

  hours := math.RoundToEven(math.Abs(duration.Hours()))
  minutes := math.RoundToEven(math.RoundToEven(math.Abs(duration.Minutes())) - (hours * 60))
  seconds := math.RoundToEven(math.RoundToEven(math.Abs(duration.Seconds())) - (minutes * 60) - (hours * 60 * 60))

  fmt.Print("Hours: ")
  fmt.Println(hours)

  fmt.Print("Minutes: ")
  fmt.Println(minutes)

  fmt.Print("Seconds: ")
  fmt.Println(seconds)
}


func gettime() (time.Time) {
  var sec,min,hour int

  fmt.Println("Get me hour")
  fmt.Scanf("%d", &hour)

  fmt.Println("Get me minute")
  fmt.Scanf("%d", &min)

  fmt.Println("Get me second")
  fmt.Scanf("%d", &sec)

  return time.Date(2000, 01, 01, hour, min, sec, 0, time.UTC)
}
