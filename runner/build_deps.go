package runner

import (
	"os/exec"
	"strings"
	"sync"
)


func (r *Runner) BuildDeps(wg *sync.WaitGroup) {
    defer wg.Done()
    var wgB sync.WaitGroup
    for _, dep := range r.config.BuildCommands {
            wgB.Add(1)
            go func(){
                defer wgB.Done()
                _, com := r.config.GetBuildCommand(dep)
                args := strings.Fields(com)
                cmd := exec.Command(args[0], args[1:]...) 
                err := cmd.Run()
                if err != nil {
                    panic(err)
                }
            }()
    }

    args := strings.Fields(r.config.FinalBuildCommand)
    cmd := exec.Command(args[0], args[1:]...) 
    err := cmd.Run()
    if err != nil {
        panic(err)
    }
}


