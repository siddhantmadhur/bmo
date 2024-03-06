package runner

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)



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
