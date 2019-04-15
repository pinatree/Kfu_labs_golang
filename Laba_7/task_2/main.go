package main

import (
  "os/exec"
)

func main(){
  exec.Command("bash", "-c", "cat f1 > h").Run()
  exec.Command("bash", "-c", "cat f2 > f1").Run()
  exec.Command("bash", "-c", "cat h > f2").Run()
  exec.Command("bash", "-c", "rm h").Run()
}
