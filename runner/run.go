package runner

import (
	"os/exec"
	"strings"
)


func (r *Runner) Run () {
    args := strings.Fields(r.config.BinaryCommand)
    cmd := exec.Command(args[0], args[1:]...) 
    r.buildProcess = cmd.Process
    err := cmd.Run()
    if err != nil {
        panic(err)
    }
}
