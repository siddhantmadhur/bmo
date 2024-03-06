package main

import (
	"fmt"
	"os"
)



func main() {
    var config Config
    config.Read()
    args := os.Args
    if len(args) < 2 {
        fmt.Println("No command provided")
    } else {
        switch(args[1]) {
        case "run":
            var runner Runner
            runner.Start(&config)
        case "init":
            config.Init(&args)
        }
    }
}
