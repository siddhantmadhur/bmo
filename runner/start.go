package runner

import (
	"fmt"
	"regexp"
	"sync"

	"github.com/fsnotify/fsnotify"
)



func (r *Runner) Start() {
    var wg sync.WaitGroup
    wg.Add(1)
    go r.BuildDeps(&wg)
    var err error
    r.watcher, err = fsnotify.NewWatcher()

    r.addFilePaths()
    if err != nil {
        fmt.Println("There was an error in watching")
        panic(err)
    }

    defer r.watcher.Close()
    

    go func() {
        for {
            select {
            case event, ok := <-r.watcher.Events:
                if !ok {
                    panic("Watcher error")
                    return 
                }
                flg := true
                for _, reg := range r.config.ExcludeRegex {
                    match, err := regexp.Match(reg, []byte(event.Name))
                    if err != nil {
                        panic(err)
                    }
                    if match {
                        flg = false
                    }
                }
                if flg {
                    if event.Has(fsnotify.Write)   {
                        go r.HandleChange(event.Name)
                    } else if event.Has(fsnotify.Create) {
                        r.addFilePath(event.Name)
                    }
                }
            case err, ok := <-r.watcher.Errors:
                if !ok {
                    panic("There was a watcher error")
                    return
                }
                fmt.Println("error: ", err)
            }
        }
    }()    

    
    wg.Wait()
    go r.Run()      

    <-make(chan struct {})

}
