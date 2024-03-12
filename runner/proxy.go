package runner

import (
	"embed"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

//go:embed proxy.js
var f embed.FS

func (r *Runner) StartProxyServer () {
    app := fiber.New()

    go func() {
        openbrowser("http://localhost:9090")
    }()


    app.Get("/_bmo/proxy.js", func(c *fiber.Ctx) error {
        data, _ := f.ReadFile("proxy.js")
        return c.SendString(string(data)) 
    })

    app.Use("/ws", func(c *fiber.Ctx) error {
        // IsWebSocketUpgrade returns true if the client
        // requested upgrade to the WebSocket protocol.
        if websocket.IsWebSocketUpgrade(c) {
            c.Locals("allowed", true)
            return c.Next()
        }
        return fiber.ErrUpgradeRequired
    })

    app.Get("/ws/proxy", websocket.New(func(c *websocket.Conn) {
        // websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
        var (
            err error
        )
        for {
            select {
                case <-r.Queue: 
                if err = c.WriteMessage(websocket.TextMessage, []byte("BMO: Update detected!")); err != nil {
                    r.Queue <- false
                }
                break
            }

        }
    }))

    app.All("/*", func(c *fiber.Ctx) error {
        if err := proxy.Do(c, fmt.Sprintf("http://localhost:%d" + c.Path(), r.Cfg.Build.WebServerPort)); err != nil {
            return err
        }
        response := c.Response().Body()
        spl := strings.Split(string(response), "</body>")
        if len(spl) > 1 {
            response = []byte(strings.Join(spl, `<script type="text/javascript" src="/_bmo/proxy.js"></script></body>`)) 
        }

        c.Response().SetBody(response)
        return nil
    }) 

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
