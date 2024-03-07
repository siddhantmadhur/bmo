package runner

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
)


func (r *Runner) BuildDeps(wg *sync.WaitGroup) {
    defer wg.Done()
    for _, dep := range r.config.BuildCommands {
        _, com := r.config.GetBuildCommand(dep)
        args := strings.Fields(com)
        cmd := exec.Command(args[0], args[1:]...) 
        err := cmd.Run()
        if err != nil {
            panic(err)
        }
    }
}

func (r *Runner) BuildDep(wg *sync.WaitGroup, extension string) {
    defer wg.Done()
    for _, dep := range r.config.BuildCommands {
        ext, com := r.config.GetBuildCommand(dep)
        fmt.Println(extension == ext)
        if extension == ext {
            args := strings.Fields(com)
            cmd := exec.Command(args[0], args[1:]...) 
            err := cmd.Run()
            if err != nil {
                panic(err)
            }
        }
    }
}
