package runner

import (
	"fmt"
	"os/exec"
	"strings"
)


func (r *Runner) Run () {
    args := strings.Fields(r.config.BinaryCommand)
    cmd := exec.Command(args[0], args[1:]...) 
    r.buildProcess = cmd.Process
    fmt.Println("Starting now")
    err := cmd.Start()
    if err != nil {
        fmt.Println("Error: ",err)
    } 
    for {
        select {
            case <-r.stop: 
                fmt.Println("Stop signal given")
                cmd.Process.Kill()
                return
            default:
        }
    }
}
