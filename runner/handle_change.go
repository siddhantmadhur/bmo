package runner

import (
	"strings"
	"sync"
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
                var wg sync.WaitGroup
                go r.BuildDeps(&wg)
                if r.proc.Process != nil {
                    r.proc.Process.Kill()
                }
                r.stop <- true
                r.proc.Process.Wait()
                wg.Wait()
                go r.Run()
            }
        }
    }
}
