// @Author: 艾珩
// @Date 2023/4/10 13:43
// report bug to <chenjiamin.cjm@alibaba-inc.com>

package customplugin

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

func init() {
	plugin.Register("customplugin", setup)
}

func setup(c *caddy.Controller) error {
	log.Infof("customplug start")
	c.Next() // Ignore the 'customplugin' directive
	if c.NextArg() {
		return plugin.Error("customplugin", c.ArgErr())
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return CustomPlugin{Next: next}
	})

	return nil
}
