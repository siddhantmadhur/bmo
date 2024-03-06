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
    ExcludeFiles []string `json:"exclude_files"`
}


func (c *Config) Read() {

}

func (c *Config) Init(args *[]string) {
    _, err := os.ReadFile("./bmo.toml")

    if len(*args) < 3 {
        fmt.Println("Default configuration not provided \nExample: bmo init go")
        os.Exit(1)
    }

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
        switch((*args)[2]) {
            case "go":
                newConfig.Name = "go"
                newConfig.BinaryCommand = "go build -o tmp/main ."
                newConfig.ExcludeFiles = []string{"./tmp"}
                fmt.Println("Don't forget to add ./tmp to .gitignore!")
        }
        data, err = toml.Marshal(newConfig)
        newFile.Write(data)
        if err == nil {
            fmt.Println("Created default configration file")
        }
    }
}


func (c *Config) Build() {
}
