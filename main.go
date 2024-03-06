package main

import (
	"fmt"
	"time"
)


func main() {
    fmt.Printf("Hello, world! \n")
    var i = 0
    for i <= 10 {
        fmt.Printf("\rSecond: %d", i)
        time.Sleep(time.Second)
        i = i + 1
    }
    time.Sleep(time.Second)
    fmt.Printf("\nBye, world! \n")
}
