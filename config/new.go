package config

import (
	"fmt"
	"os"

	"github.com/TwiN/go-color"
	"github.com/pelletier/go-toml/v2"
)


func New () Config {
    content, err := os.ReadFile(".bmo.toml")
    if err != nil {
        fmt.Println(color.Ize(color.Red, "\t[BMO] There was an error \n\n\tMake sure configuration exists by using\n\n\tbmo init\n"))
        os.Exit(1)
    }
    var config Config
    toml.Unmarshal(content, &config)
    return config
}

func Init() {
    var config Config
    
    bytes, err := toml.Marshal(config)
    if err != nil {
        fmt.Println(color.Ize(color.Red, "\t[BMO] There was an error \n\n\tMake sure configuration exists by using\n\n\tbmo init\n"))
        os.Exit(1)
    }
    
    os.WriteFile(".bmo.toml", bytes, 0777)
}

type Config struct {
    Build Build `toml:"BUILD"`
}

type Build struct {
    BuildAssetsCommand []string `toml:"build_assets_cmd"`
    BuildBinaryCommand string `toml:"build_binary_cmd"`
    RunBinaryCommand string `toml:"run_binary_cmd"`
    ExcludedDirs []string `toml:"excluded_files"`
    ExcludedRegex []string `toml:"excluded_regex"`
    WebServerPort int `toml:"web_server_port"`
}
