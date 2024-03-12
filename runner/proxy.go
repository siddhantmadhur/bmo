package runner

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/gofiber/fiber/v2"
)


func (r *Runner) StartProxyServer () {
    app := fiber.New()

    go func() {
        openbrowser("http://localhost:8080")
    }()

    app.Listen(":9090")
}


func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
