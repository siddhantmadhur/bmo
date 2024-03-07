package runner

import (
	"io"
	"os"

	"bmo.siddhantsoftware.com/config"
	"github.com/fsnotify/fsnotify"
)

func NewRunner(c *config.Config) *Runner {
    var runner Runner
    runner.config = c
    runner.stop = make(chan bool)
    return &runner
}

type Runner struct {
    watcher *fsnotify.Watcher
    config *config.Config
    buildProcess *os.Process
    ioRead io.ReadCloser
    stop chan bool
}


