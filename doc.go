/*
Package exit 接收系统退出信号，并发送给注册的业务端，然后等待业务端处理完成后，执行os.Exit(0)

Example:

  package main

  import (
    "fmt"

    "github.com/sunreaver/exit"
  )

  func main() {
    c, over := exit.RegistExiter()

    select {
    case <-c:
      fmt.Println("exit chan")
    }

    // do something
    // example: flush buffer

    close(over)

    select {}

    fmt.Println("ok")
  }

*/
package exit
