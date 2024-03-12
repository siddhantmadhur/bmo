package main

import (
	"fmt"
	"os"

	"bmo.siddhantsoftware.com/v2/runner"
	"github.com/TwiN/go-color"
)


func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        fmt.Println(color.Ize(color.Red, "\t[BMO] You have not provided any arguments"))
        os.Exit(0)
    }
    
    switch args[0] {
        case "run": 
            var bmo = runner.New();
            bmo.Start()

    }


}
