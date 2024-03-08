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
        var wg sync.WaitGroup
        wg.Add(1)
        go r.BuildDeps(&wg)
        if r.proc.Process != nil {
            r.proc.Process.Kill()
        }
        wg.Wait()
        r.proc.Process.Wait()
        r.stop <- true
        go r.Run()
    }
}
