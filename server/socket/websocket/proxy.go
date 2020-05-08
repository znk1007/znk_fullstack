package ws

import (
	"net"
	"net/url"
	"os"
	"strings"
	"sync"
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
			continue
		}
		if ip := net.ParseIP(host); ip != nil {
			p.AddIP(ip)
			continue
		}
		if strings.HasPrefix(host, "*.") {
			p.AddZone(host[1:])
			continue
		}
		p.AddHost(host)
	}
}

//AddNetwork specifies an IP range that will use the bypass proxy.
//Note that this will only take effect if a literal IP address is dialed.
//A connection to a named host will never match.
func (p *proxyPerHost) AddNetwork(net *net.IPNet) {
	p.bypassNetworks = append(p.bypassNetworks, net)
}

//AddZone specifies a DNS suffix that will use the bypass proxy.
//A zone of "example.com" matches "example.com" and all of its subdomains.
func (p *proxyPerHost) AddZone(zone string) {
	if strings.HasPrefix(zone, ".") {
		zone = zone[:len(zone)-1]
	}
	if !strings.HasPrefix(zone, ".") {
		zone = "." + zone
	}
	p.bypassZones = append(p.bypassZones, zone)
}

//AddIP specifies an IP address that will use the bypass proxy.
//Note that this will only take effect if a literal IP address is dialed.
//A connection to a named host will never match an IP.
func (p *proxyPerHost) AddIP(ip net.IP) {
	p.bypassIPs = append(p.bypassIPs, ip)
}

//AddHost specifies a host name that will use the bypass proxy.
func (p *proxyPerHost) AddHost(host string) {
	if strings.HasSuffix(host, ".") {
		host = host[:len(host)-1]
	}
	p.bypassHosts = append(p.bypassHosts, host)
}

//Auth contains authentication parameters that specific Dialers may require
type proxyAuth struct {
	User, Password string
}

//EnvOnce looks up an environment variable (optionally by multiple names) once.
//It mitigates expensive lookups on some platforms (e.g. Windows).
//(Borrowed from net/http/transport.go)
type proxyEnvOnce struct {
	names []string
	once  sync.Once
	val   string
}

func (e *proxyEnvOnce) Get() string {
	e.once.Do(e.init)
	return e.val
}

//init 初始化
func (e *proxyEnvOnce) init() {
	for _, n := range e.names {
		e.val = os.Getenv(n)
		if e.val != "" {
			return
		}
	}
}

var (
	proxyAllProxyEnv = &proxyEnvOnce{
		names: []string{"ALL_PROXY", "all_proxy"},
	}
	proxyNoProxyEnv = &proxyEnvOnce{
		names: []string{"NO_PROXY", "no_proxy"},
	}
)

//proxyProxySchemes is a map from URL schemes to a function that creates a Dialer
//from a URL with such a scheme.
var proxyProxySchemes map[string]func(*url.URL, proxyDialer) (proxyDialer, error)

//proxyRegisterDialerType takes a URL scheme and a function to generate Dialers from
//a URL with that scheme and a forwarding Dialer.
//Registered schemes are used by FromURL.
func proxyRegisterDialerType(scheme string, f func(*url.URL, proxyDialer) (proxyDialer, error)) {
	if proxyProxySchemes == nil {
		proxyProxySchemes = make(map[string]func(*url.URL, proxyDialer) (proxyDialer, error))
	}
	proxyProxySchemes[scheme] = f
}

func proxyFromURL(u *url.URL, forward proxyDialer) (proxyDialer, error) {

}

//proxyFromEnvironment returns the dialer specified by the proxy related variables
//in the environment
func proxyFromEnvironment() proxyDialer {
	allProxy := proxy
}
