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


func (r *Runner) HandleChange (filePath string) {
    subPaths := strings.Split(filePath, "/")
    ext := strings.Split(subPaths[len(subPaths)-1], ".")
    if len(ext) < 2 {
        panic("Error")
    }else{
        for _, com := range r.config.BuildCommands {
            a, b := r.config.GetBuildCommand(com)
            if a == ext[1] {
                fmt.Println(b)
            }
        }
    }
}


func (r *Runner) Init(config *Config) {
   r.config = config; 
}

func (r *Runner) addFilePath(dir string) {
    flg := true
    for _, file := range r.config.ExcludeFiles {
        if strings.Index(dir, file) != -1 {
            flg = false
        }
    }
    if flg {
        r.watcher.Add(dir)
    }
}



func (r *Runner) addFilePaths() {
    filepath.WalkDir("./", func(path string, d fs.DirEntry, err error) error {
        err = r.watcher.Add(path)
        return err
    }) 
    
    for _, file := range r.config.ExcludeFiles {
       
        filepath.WalkDir(file, func(path string, d fs.DirEntry, err error) error {
            err = r.watcher.Remove(path)
            return err
        }) 
    }
    fmt.Println(r.watcher.WatchList())
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


