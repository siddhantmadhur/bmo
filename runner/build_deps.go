package runner

import (
	"os/exec"
	"strings"
	"sync"
)


func (r *Runner) BuildDeps(wg *sync.WaitGroup) {
    wg.Add(1)
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
