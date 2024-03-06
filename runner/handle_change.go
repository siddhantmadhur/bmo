package runner

import (
	"fmt"
	"strings"
)



func (r *Runner) HandleChange (filePath string) {
    subPaths := strings.Split(filePath, "/")
    ext := strings.Split(subPaths[len(subPaths)-1], ".")
    if len(ext) < 2 {
        panic("Error")
    }else{
        for _, com := range r.config.BuildCommands {
            a, b := r.config.GetBuildCommand(com)
            if a == ext[1] {
                fmt.Println(b)
            }
        }
    }
}
