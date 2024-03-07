package runner

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)



func (r *Runner) Start() {
    r.BuildDeps()
    r.Run()    
    var err error
    r.watcher, err = fsnotify.NewWatcher()

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
                    return 
                }
                if event.Has(fsnotify.Write)   {
                    fmt.Println("modified file: ", event.Name)
                    r.HandleChange(event.Name)
                } else if event.Has(fsnotify.Create) {
                    r.addFilePath(event.Name)
                }
            case err, ok := <-r.watcher.Errors:
                if !ok {
                    return
                }
                fmt.Println("error: ", err)
            }
        }
    }()    

    r.addFilePaths()
    

    <-make(chan struct {})

}
