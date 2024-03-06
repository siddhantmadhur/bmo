package main

import (
	"fmt"
	"os"
)



func main() {
    var config Config
    err := config.Read()
    var runner Runner
    runner.Init(&config)
    args := os.Args
    if len(args) < 2 {
        fmt.Println("No command provided")
    } else {
        if args[1] != "init" && err != nil {
            fmt.Println("Configuration file does not exist")
            os.Exit(0)
        }
        switch(args[1]) {
        case "start":
            runner.Start()
        case "init":
            config.Init(&args)
        case "add":
            config.Add(&args)
        }
    }
}
