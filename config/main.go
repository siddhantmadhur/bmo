package config

import (
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

func NewConfig() (*Config, error) {
   var config Config
   err := config.Read()
   return &config, err
}

type Config struct {
    Name string `json:"name"`
    BuildCommands []string `json:"commands"`
    FinalBuildCommand string `json:"final_command"`
    BinaryCommand string `json:"binary_command"`
    ExcludeFiles []string `json:"exclude_files"`
    ExcludeRegex []string `json:"exclude_regex"`
    ProxyPort int `json:"proxy_port"`
    WebServerUrl string `json:"web_server_url"`
}

func (c *Config) GetBuildCommand(raw string) (string, string) {
    spl := strings.Split(raw, ":") 
    if len(spl) != 2 {
        panic("Not appropriate: " + strings.Join(spl, " "))
    }
    ext := spl[0]
    com := spl[1]
    
    return ext, com
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


