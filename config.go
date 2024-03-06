package main

import "os"



type Config struct {
    Name string `json:"name"`
    BuildCommands []string `json:"commands"`
    BinaryCommand string `json:"binary_command"`
    Files []string `json:"files"`
}


func (c *Config) Read() {
    os.ReadFile("./bmo.json")
    c.Name = "Hello"
    c.Files = []string{"./"}

}



