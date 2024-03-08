package proxy
import "bmo.siddhantsoftware.com/config"

type Proxy struct {
    config *config.Config
}


func (p *Proxy) Start(c *config.Config) {
   p.config = c


}
