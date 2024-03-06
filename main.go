package main

import "os"



func main() {
    var config Config
    config.Read()
    args := os.Args

    for _, arg := range args {
        switch(arg) {
            case "run":
                var runner Runner
                runner.Start(&config)
        }
    }
}
