package main

import (
    "fmt"
    "log"
    "os"
)

var LOGFILE = "mGo.log"

func main() {
    f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()

    iLog := log.New(f, "customLog ", log.LstdFlags)
    iLog.Println("Hello there!")
    iLog.Println("Another log entry!")
}
