package runner

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)


func (r *Runner) Run () {
    args := strings.Fields(r.config.BinaryCommand)
    r.proc = exec.Command(args[0], args[1:]...) 
    stdout, err := r.proc.StdoutPipe()
    if err != nil {
        fmt.Println("Error: ",err)
    } 
    fmt.Println("Starting now")
    err = r.proc.Start()
    if err != nil {
        fmt.Println("Error: ",err)
        panic(err)
    } 
    buf := bufio.NewReader(stdout)
    defer stdout.Close()
    for {
        select {
            case <-r.stop: 
                fmt.Println("Stop signal given")
                return
            default:
                line, _, _ := buf.ReadLine()
                //TODO: make the program not pause waiting for a new line
                fmt.Println(string(line))
        }
    }
}
