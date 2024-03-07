package runner

import (
	"strings"
)



func (r *Runner) HandleChange (filePath string) {
    subPaths := strings.Split(filePath, "/")
    ext := strings.Split(subPaths[len(subPaths)-1], ".")
    if len(ext) < 2 {
        panic("Error")
    }else{
        for _, com := range r.config.BuildCommands {
            a, _ := r.config.GetBuildCommand(com)
            if a == ext[1] {
                if r.proc.Process != nil {
                    r.proc.Process.Kill()
                    r.proc.Process.Wait()
                }
                r.stop <- true
                r.BuildDeps()
                go r.Run()
            }
        }
    }
}
