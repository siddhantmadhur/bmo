package runner

import (
	"os/exec"
	"strings"
)



func (r *Runner) HandleChange (filePath string) {
    if r.buildProcess != nil {
        r.buildProcess.Kill()
    } 
    subPaths := strings.Split(filePath, "/")
    ext := strings.Split(subPaths[len(subPaths)-1], ".")
    if len(ext) < 2 {
        panic("Error")
    }else{
        for _, com := range r.config.BuildCommands {
            a, b := r.config.GetBuildCommand(com)
            if a == ext[1] {
                args := strings.Fields(b)
                cmd := exec.Command(args[0], args[1:]...) 
                r.buildProcess = cmd.Process
                var err error
                r.ioRead, err = cmd.StdoutPipe()
                if err != nil {
                    panic(err)
                }
                cmd.Run()
            }
        }
    }
}
