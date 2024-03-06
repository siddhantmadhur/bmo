package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)



type Runner struct {
    watcher *fsnotify.Watcher
    config *Config
}

func (r *Runner) Init(config *Config) {
   r.config = config; 
}

func (r *Runner) addFilePaths() {
    filepath.WalkDir("./", func(path string, d fs.DirEntry, err error) error {
        for _, file := range r.config.ExcludeFiles {
            if file == path {
                return nil;
            }
        }
        err = r.watcher.Add(path)
        return err
    }) 
}



func (r *Runner) Start() {
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
                } else if event.Has(fsnotify.Create) {
                    flg := true
                    for _, file := range r.config.ExcludeFiles {
                        if strings.Index(event.Name, file) != -1 {
                            flg = false
                        }
                    }
                    if flg {
                        r.watcher.Add(event.Name)
                    }
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


