package config

import (
	"fmt"
	"os"
	"strings"
)


func (c *Config) Add(args *[]string) error {
    if len(*args) < 3 {
        fmt.Println("No argument provided.")
        os.Exit(0)
    }
    command := (*args)[2]
    hand := strings.Split(command, ":")
    if len(hand) == 2 {
        c.BuildCommands = append(c.BuildCommands, command)
        err := c.Write()
        if err != nil {
            panic(err)
        }
    }else{
        fmt.Println("Function wasn't able to be parsed probably, make sure you're using quotes correctly.")
        os.Exit(0)
    }
    return nil
}
