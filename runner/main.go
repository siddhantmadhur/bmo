package runner

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"sync"

	"github.com/TwiN/go-color"
	"github.com/fsnotify/fsnotify"
)


func New() Runner {
    queue := make(chan []string) 
    var waitGroup sync.WaitGroup
    return Runner{
        Queue: queue,
        WaitGroup: waitGroup,
    }
}


type Runner struct {
     Queue chan []string        
     WaitGroup sync.WaitGroup
}


func (r *Runner) Start() {
    fmt.Println(color.Ize(color.Blue, "\t[BMO] Starting event listener..."));
    

    go r.Listener()
    r.DetectFileChanges()


    for {}

}

func (r *Runner) DetectFileChanges() {
    watcher, err := fsnotify.NewWatcher()

    if err != nil {
        fmt.Println(color.Ize(color.Red, "\t[BMO] There was an issue with the file listener."))
    }

    go func () {
        for {
            select {
                case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                if event.Has(fsnotify.Write)  {
                    
                    r.Queue <- []string{event.Name} 
                }
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                log.Println("error:", err)
            }
        }
    }()

   
    add_all_paths(watcher)

    if err != nil {
        fmt.Println(color.Ize(color.Red, "\t[BMO] There was an issue with the file listener.\n"))
    }

    <-make(chan struct{})

}

func (r *Runner) Listener() {

    for {
        if len(<-r.Queue) > 0 {
            fmt.Println("Detected change")
        } 
    }

}


func add_all_paths(notify *fsnotify.Watcher) {
    filepath.Walk("./", func(path string, info fs.FileInfo, err error) error {
        notify.Add(path)
        return err
    })
}



