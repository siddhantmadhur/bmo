package runner

import (
	"os/exec"
	"strings"
)


func (r *Runner) BuildDeps() {
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
