// @Author: 艾珩
// @Date 2023/4/10 13:44
// report bug to <chenjiamin.cjm@alibaba-inc.com>

package customplugin

import (
	"context"
	"net"

	"github.com/coredns/coredns/plugin"
	clog "github.com/coredns/coredns/plugin/pkg/log"
	"github.com/coredns/coredns/request"
	"github.com/miekg/dns"
)

var log = clog.NewWithPlugin("auto")

type CustomPlugin struct {
	Next plugin.Handler
}

func (c CustomPlugin) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	state := request.Request{W: w, Req: r}
	qname := state.QName()

	log.Infof("qname is %s", qname)
	if qname == "example.com." {
		m := new(dns.Msg)
		m.SetReply(r)
		m.Authoritative = true
		rr := new(dns.A)
		rr.Hdr = dns.RR_Header{Name: qname, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60}
		rr.A = net.ParseIP("1.2.3.4").To4()
		m.Answer = append(m.Answer, rr)

		w.WriteMsg(m)
		return dns.RcodeSuccess, nil
	}

	return plugin.NextOrFailure(c.Name(), c.Next, ctx, w, r)
}

func (c CustomPlugin) Name() string {
	return "customplugin"
}
