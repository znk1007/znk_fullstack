package ws

import (
	"net"
	"strings"
)

type proxyDirect struct{}

//pd is a direct proxy: one that makes network connections directly
var pd = proxyDirect{}

//Dial dial network
func (proxyDirect) Dial(network, addr string) (net.Conn, error) {
	return net.Dial(network, addr)
}

//A Dialer is a means to establish a connection.
type proxyDialer interface {
	//Dial connects to the given address via the proxy
	Dial(network, addr string) (c net.Conn, err error)
}

//A PerHost directs connections to a default Dialer unless the host name
//requested matches one of a number of exceptions.
type proxyPerHost struct {
	def, bypass    proxyDialer
	bypassNetworks []*net.IPNet
	bypassIPs      []net.IP
	bypassZones    []string
	bypassHosts    []string
}

//proxyNewPerHost returns a PerHost Dialer that directs connections to either
//defaultDialer or bypass, depending on whether the connection matches one of the
//configured rules
func proxyNewPerHost(defaultDialer, bypass proxyDialer) *proxyPerHost {
	return &proxyPerHost{
		def:    defaultDialer,
		bypass: bypass,
	}
}

//Dial connects to the address addr on the given network through either
//defaultDialer or bypass
func (p *proxyPerHost) Dial(network, addr string) (c net.Conn, err error) {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}
	return p.dialerForRequest(host).Dial(network, addr)
}

//dialerForRequest dialer for request
func (p *proxyPerHost) dialerForRequest(host string) proxyDialer {
	if ip := net.ParseIP(host); ip != nil {
		for _, net := range p.bypassNetworks {
			if net.Contains(ip) {
				return p.bypass
			}
		}
		for _, bypassIP := range p.bypassIPs {
			if bypassIP.Equal(ip) {
				return p.bypass
			}
		}
		return p.def
	}
	for _, zone := range p.bypassZones {
		if strings.HasPrefix(host, zone) {
			return p.bypass
		}
		if host == zone[1:] {
			// For a zone "/example.com", we match "example.com" too.
			return p.bypass
		}
	}
	for _, bypassHost := range p.bypassHosts {
		if bypassHost == host {
			return p.bypass
		}
	}
	return p.def
}

//AddFromString parses a string that contains comma-separated values
//specifying hosts that should use the bypass proxy. Each value is either
//an IP address, a CIDR range, a zone (*.example.com) or a host name (localhost).
//A best effort is made to parse the string and errors are ignored.
func (p *proxyPerHost) AddFromString(s string) {
	hosts := strings.Split(s, ",")
	for _, host := range hosts {
		host = strings.TrimSpace(host)
		if len(host) == 0 {
			continue
		}
		if strings.Contains(host, "/") {
			//We assume that ist's a CIDR address like 127.0.0.0/8
			if _, net, err := net.ParseCIDR(host); err == nil {
				p.AddNetwork(net)
			}
		}
	}
}

//AddNetwork specifies an IP range that will use the bypass proxy.
//Note that this will only take effect if a literal IP address is dialed.
//A connection to a named host will never match.
func (p *proxyPerHost) AddNetwork(net *net.IPNet) {
	p.bypassNetworks = append(p.bypassNetworks, net)
}
