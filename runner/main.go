package runner

import (
	"os/exec"

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
    stop chan bool
    proc *exec.Cmd
}


