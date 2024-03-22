package runner

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"bmo.siddhantsoftware.com/v2/config"
	"github.com/TwiN/go-color"
	"github.com/fsnotify/fsnotify"
)


func New(cfg *config.Config) Runner {
    queue := make(chan bool) 
    var waitGroup sync.WaitGroup

    

    var runner = Runner{
        Queue: queue,
        WaitGroup: &waitGroup,
        Cfg: cfg,
    }

    waitGroup.Add(1); 
    go runner.Start()
    go runner.RunWebServer()
    return runner;
}


func (r *Runner) RunWebServer() {
    defer r.WaitGroup.Done();

    var side_grp sync.WaitGroup
    for _, command := range r.Cfg.Build.BuildAssetsCommand {
        side_grp.Add(1) 
        go func() {
            defer side_grp.Done()
            cur := strings.Split(command, " ")
            fmt.Println(cur)
            exec.Command(cur[0], cur[1:]...).Run()
        }()
    }
    side_grp.Wait()
    cur := strings.Split(r.Cfg.Build.BuildBinaryCommand, " ")
    exec.Command(cur[0], cur[1:]...).Run()
    
    cur = strings.Split(r.Cfg.Build.RunBinaryCommand, " ")

    r.Process = exec.Command(cur[0], cur[1:]...)
    stdout, err := r.Process.StdoutPipe()
    if err != nil {
        fmt.Println(color.Ize(color.Red, "\t[BMO] There was an error: "), err.Error())
        os.Exit(1)
    }
    err = r.Process.Start()
    if err != nil {
        fmt.Println(color.Ize(color.Red, "\t[BMO] There was an error: "), err.Error())
        os.Exit(1)
    }
    
    fmt.Println(color.Ize(color.White, fmt.Sprintf("\t[BMO] Process is running at %d ", r.Process.Process.Pid)))

    if err != nil {
        fmt.Println(color.Ize(color.Red, "\t[BMO] There was an error: "), err.Error())
        os.Exit(1)
    }
    r.Queue <- true
    scanner := bufio.NewScanner(stdout)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        m := scanner.Text()
        fmt.Println(m)
    }

}


type Runner struct {
     Queue chan bool        
     WaitGroup *sync.WaitGroup
     Process *exec.Cmd
     Cfg *config.Config
}


func (r *Runner) Start() {
    fmt.Println(color.Ize(color.Blue, "\t[BMO] Starting event listener..."));
    

    go r.DetectFileChanges()

    r.StartProxyServer()
}

func (r *Runner) DetectFileChanges() {
    watcher, err := fsnotify.NewWatcher()
    r.add_all_paths(watcher)
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
                flg := true
                for _, reg := range r.Cfg.Build.ExcludedRegex {
                    match, err := regexp.Match(reg, []byte(event.Name))
                    if err != nil {
                        fmt.Println(color.Ize(color.Red, "\t[BMO] There was an error in checking excluded regex"))
                        os.Exit(1)
                    }
                    if match {
                        flg = false
                    }
                }
                if event.Has(fsnotify.Write) && flg {
                    fmt.Println(color.Ize(color.Blue, "\t[BMO] Detected change..."))
                    if r.Process != nil {
                        r.Process.Process.Kill()
                        r.WaitGroup.Add(1)
                        go r.RunWebServer()
                    }
                }
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                log.Println("error:", err)
            }
        }
    }()

   

    if err != nil {
        fmt.Println(color.Ize(color.Red, "\t[BMO] There was an issue with the file listener.\n"))
    }

    <-make(chan struct{})

}



func (r *Runner) add_all_paths(notify *fsnotify.Watcher) {
    filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
        flg := true
        for _, exc_path := range r.Cfg.Build.ExcludedDirs {
            if strings.Contains(path, exc_path) {
                flg = false
                break;
            }
        }
        if flg {
            notify.Add(path)
        }
        return err
    })
}



