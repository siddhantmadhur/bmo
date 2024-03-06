package main

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)



type Config struct {
    Name string `json:"name"`
    BuildCommands []string `json:"commands"`
    BinaryCommand string `json:"binary_command"`
    Files []string `json:"files"`
}


func (c *Config) Read() {

}

func (c *Config) Init() {
    _, err := os.ReadFile("./bmo.toml")
    if err == nil {
        //file exists
        fmt.Println("Configuration already exists")        
    } else {
        //file doesn't exist
        newFile, err := os.Create("bmo.toml")
        if err != nil {
            panic(err)
        } 
        var newConfig Config
        var data []byte

        data, err = toml.Marshal(newConfig)
        newFile.Write(data)
        if err == nil {
            fmt.Println("Created default configration file")
        }
    }
}



