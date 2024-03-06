package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)



type Runner struct {
    

}



func (r *Runner) Start(config *Config) {
    files := config.Files 
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        fmt.Println("There was an error in watching")
        panic(err)
    }
    defer watcher.Close()


    go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return 
                }
                if event.Has(fsnotify.Write) {
                    fmt.Println("modified file: ", event.Name)
                }
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                fmt.Println("error: ", err)
            }
        }
    }()      
    

    cur := 0
    for cur < len(files) {
       err = watcher.Add(files[cur])
       if err != nil {
            panic(err)
       }
       cur += 1
    }

    <-make(chan struct {})

}
