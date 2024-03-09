package runner

import (
	"embed"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)



//go:embed proxy.js
var f embed.FS


func (r *Runner) StartProxy() {
    fmt.Println("Proxy starting...") 

    app := fiber.New()
   
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
        log.Println(c.Locals("allowed"))  // true
        log.Println(c.Params("id"))       // 123
        log.Println(c.Query("v"))         // 1.0
        log.Println(c.Cookies("session")) // ""

        // websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
        var (
            err error
        )
        
        for {
            select {
                case <- r.stop: 
                if err = c.WriteMessage(websocket.TextMessage, []byte("BMO: Update detected!")); err != nil {
                    log.Println("write:", err)
                    break
                }

            }
        }
    }))

    app.All("/*", func(c *fiber.Ctx) error {
        if err := proxy.Do(c, r.config.WebServerUrl + c.Path()); err != nil {
            return err
        }
        response := c.Response().Body()
        spl := strings.Split(string(response), "</body>")
        if len(spl) > 1 {
            data, _ := f.ReadFile("proxy.js")
            response = []byte(strings.Join(spl, fmt.Sprintf(`<script type="text/javascript">%s</script></body>`, data))) 
        }

        c.Response().SetBody(response)
        return nil
    }) 

    app.Listen(":9090") 
}

