package proxy

import (
	"embed"
	"fmt"
	"strings"

	"bmo.siddhantsoftware.com/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

type Proxy struct {
    config *config.Config
}


//go:embed proxy.js
var f embed.FS


func (p *Proxy) Start(c *config.Config) {
    p.config = c
    fmt.Println("Proxy starting...") 

    app := fiber.New()
    
    app.All("/*", func(c *fiber.Ctx) error {
        if err := proxy.Do(c, p.config.WebServerUrl + c.Path()); err != nil {
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

