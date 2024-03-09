package main

import (
	"fmt"
	"os"

	"bmo.siddhantsoftware.com/config"
	"bmo.siddhantsoftware.com/proxy"
	"bmo.siddhantsoftware.com/runner"
)



func main() {
    var proxy proxy.Proxy
    config, err := config.NewConfig()
    runner := runner.NewRunner(config)
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
            go proxy.Start(config)
            runner.Start()
        case "init":
            config.Init(&args)
        case "add":
            config.Add(&args)
        }
    }
}
