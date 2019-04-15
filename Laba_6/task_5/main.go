package main

import (
  "fmt"
  "bufio"
  "os"
)

func main(){
  var copy_str string

  fmt.Println("Get me main string")
  myscanner := bufio.NewScanner(os.Stdin)
  myscanner.Scan()
  copy_str = myscanner.Text()

  for true {
    original_str := copy_str

    for x := 0; x < len(copy_str); x++ {
      current_symbol := copy_str[x : x+1]
      next_hook := checkForNextHook(current_symbol, copy_str, x)

      if next_hook.newOpenHookWasFound {
        x = next_hook.position - 1
        continue
        } else {
          if next_hook.closeHookWasFound && next_hook.closeHookIsCorrect {
            fmt.Println("Destroying positions")
            fmt.Print(x)
            fmt.Print(" and ")
            fmt.Print(next_hook.position)
            fmt.Println("...")
            copy_str = copy_str[:x] + copy_str[next_hook.position + 1 :]
            fmt.Println(copy_str)
            break
            } else {
              hasHooksInside := false
              for i := 0; i < len(copy_str); i++ {
                if isHook(copy_str[i : i + 1]) {
                  hasHooksInside = true
                  break
                } else {
                  continue
                }
              }

              if hasHooksInside{
                fmt.Print("Its bad hooks :(")
                fmt.Print(" Thiere is problem between positions ")
                fmt.Print(x)
                fmt.Print(" and ")
                fmt.Println(next_hook.position)
                return
              } else {
                fmt.Println("All done! String is correct.")
                return
              }
            }
          }
        }
        if original_str == copy_str {
          fmt.Println("All done! String is correct.")
          return
        }
      }
    }
type hookData struct {
  closeHookWasFound bool
  closeHookIsCorrect bool
  newOpenHookWasFound bool
  position int
}

func checkForNextHook(hook_str string, main_str string, start_pos int) hookData {
  var result hookData

  for x := start_pos + 1; x < len(main_str); x++ {
    local_symbol := main_str[x : x+1]

    if isHook(local_symbol) {

      if isOpenHook(local_symbol) {

        result.newOpenHookWasFound = true
        result.position = x
        return result

      } else {
        if areInverseHooks(hook_str, local_symbol) {

          result.closeHookWasFound = true
          result.closeHookIsCorrect = true
          result.position = x
          return result

        } else {
          result.closeHookWasFound = true
          result.closeHookIsCorrect = false
          result.position = x
          return result

        }

      }
    }
  }

  result.newOpenHookWasFound = false
  result.closeHookWasFound = false
  result.closeHookIsCorrect = false
  result.position = -1
  return result
}

func isHook(str string) bool {
  return isOpenHook(str) || isCloseHook(str)
}

func isOpenHook(str string) bool {
  return str == "{" || str == "[" || str == "("
}
func isCloseHook(str string) bool {
  return str == "}" || str == "]" || str == ")"
}

func areInverseHooks(str1 string, str2 string) bool {
  switch str1 {
    case "[": {
      return str2 == "]"
    }
    case "{": {
      return str2 == "}"
    }
    case "(": {
      return str2 == ")"
    }
    default: {
      return false
    }
  }
}
