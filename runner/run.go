package runner

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)


func (r *Runner) Run () {
    args := strings.Fields(r.config.BinaryCommand)
    cmd := exec.Command(args[0], args[1:]...) 
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        fmt.Println("Error: ",err)
    } 
    fmt.Println("Starting now")
    err = cmd.Start()
    if err != nil {
        fmt.Println("Error: ",err)
    } 
    buf := bufio.NewReader(stdout)
    for {
        select {
            case <-r.stop: 
                fmt.Println("Stop signal given")
                cmd.Process.Kill()
                return
            default:
                line, _, _ := buf.ReadLine()
                fmt.Println(string(line))
        }
    }
}
