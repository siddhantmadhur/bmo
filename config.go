package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
)



type Config struct {
    Name string `json:"name"`
    BuildCommands []string `json:"commands"`
    BinaryCommand string `json:"binary_command"`
    ExcludeFiles []string `json:"exclude_files"`
}

func (c *Config) GetBuildCommand(raw string) (string, string) {
    spl := strings.Split(raw, ":") 
    ext := spl[0]
    com := spl[1]
    return ext, com
}

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

func (c *Config) Read() error {
    file, err := os.ReadFile("./bmo.toml")
    if err != nil {
        return err
    }
    err = toml.Unmarshal(file, &c)
    return err
}

func (c *Config) Write() error {
    data, err := toml.Marshal(c)
    if err != nil {
        return err
    }
    err = os.WriteFile("bmo.toml", data, 0777)
    return err
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
        newConfig.ExcludeFiles = []string{"tmp", ".git"}
        var data []byte
        switch((*args)[2]) {
            case "go":
                newConfig.Name = "go"
                newConfig.BinaryCommand = "go build -o tmp/main ."
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
