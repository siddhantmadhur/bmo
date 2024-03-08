package runner

import (
	"io/fs"
	"path/filepath"
)



func (r *Runner) addFilePath(dir string) {
    r.watcher.Add(dir)
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
}
