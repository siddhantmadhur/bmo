package main

import (
	"fmt"
	"os"

	"bmo.siddhantsoftware.com/v2/config"
	"bmo.siddhantsoftware.com/v2/runner"
	"github.com/TwiN/go-color"
)

var version = "dev" 

func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        fmt.Println(color.Ize(color.Red, "\t[BMO] You have not provided any arguments"))
        os.Exit(0)
    }
    switch args[0] {
        case "run": 
        cfg := config.New()    
        runner.New(&cfg);
        <- make(chan struct{})
    case "v":
        fmt.Printf("BMO \nby Siddhant Madhur \nVersion: %s\n", version)
    case "init":
        config.Init()

    }


}
