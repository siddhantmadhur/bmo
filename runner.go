package main

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
)



type Runner struct {
    watcher *fsnotify.Watcher

}


func (r *Runner) addRecursively() {}

func (r *Runner) Start(config *Config) {
    var err error
    r.watcher, err = fsnotify.NewWatcher()
    directories, err := os.ReadDir("./")

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
                } else if event.Has(fsnotify.Create) {
                    r.watcher.Add(event.Name)
                }
            case err, ok := <-r.watcher.Errors:
                if !ok {
                    return
                }
                fmt.Println("error: ", err)
            }
        }
    }()      
    r.watcher.Add("")   
    for _, dir := range directories {
        err = r.watcher.Add(dir.Name())

    }
    for _, file := range config.ExcludeFiles {
        err = r.watcher.Remove(file)
        if err != nil {
            panic(err)
        }
    }

    <-make(chan struct {})

}
