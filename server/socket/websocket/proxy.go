package ws

import (
	"errors"
	"io"
	"net"
	"net/url"
	"os"
	"strconv"
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

type proxySocks5 struct {
	user, passowrd string
	network, addr  string
	forward        proxyDialer
}

const proxySocks5Version = 5

const (
	proxySocks5AuthNone     = 0
	proxySocks5AuthPassword = 2
)

const proxySocks5Connect = 1

const (
	proxySocks5IP4    = 1
	proxySocks5Domain = 3
	proxySocks5IP6    = 4
)

var proxySocks5Errors = []string{
	"",
	"general failure",
	"connection forbidden",
	"network unreachable",
	"host unreachable",
	"connection refused",
	"TTL expired",
	"command not supppored",
	"address type not supported",
}

//proxySOCK5 returns a Dialer that makes SOCKSv5 connections to the given address
//with an optional username and password. See RFC 1928 and RFC 1929.
func proxySOCK5(network, addr string, auth *proxyAuth, forward proxyDialer) (proxyDialer, error) {
	s := &proxySocks5{
		network: network,
		addr:    addr,
		forward: forward,
	}
	if auth != nil {
		s.user = auth.User
		s.passowrd = auth.Password
	}
	return s, nil
}

func (s *proxySocks5) Dial(network, addr string) (net.Conn, error) {
	switch network {
	case "tcp", "tcp6", "tcp4":
	default:
		return nil, errors.New("proxy: no support for SOCKS5 proxy connections of type " + network)
	}
	conn, err := s.forward.Dial(s.network, s.addr)
	if err != nil {
		return nil, err
	}
	if err := s.connect(conn, addr); err != nil {
		conn.Close()
		return nil, err
	}
	return conn, nil
}

//connect
func (s *proxySocks5) connect(conn net.Conn, target string) error {
	host, portStr, err := net.SplitHostPort(target)
	if err != nil {
		return err
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return errors.New("proxy: failed to parse port number: " + portStr)
	}
	if port < 1 || port > 0xffff {
		return errors.New("proxy: port number out of range: " + portStr)
	}
	buf := make([]byte, 0, 6+len(host))
	buf = append(buf, proxySocks5Version)
	if len(s.user) > 0 && len(s.user) < 256 && len(s.passowrd) < 256 {
		buf = append(buf, 2 /*num auth methods*/, proxySocks5AuthNone, proxySocks5AuthPassword)
	} else {
		buf = append(buf, 1 /*num auth methods*/, proxySocks5AuthNone)
	}
	if _, err := conn.Write(buf); err != nil {
		return errors.New("proxy: failed to write greeting to SOCKS5 proxy at " + s.addr + ": " + err.Error())
	}
	if _, err := io.ReadFull(conn, buf[:2]); err != nil {
		return errors.New("proxy: failed to read greeting from SOCKS5 proxy at " + s.addr + ": " + err.Error())
	}
	if buf[0] != 5 {
		return errors.New("proxy: SOCKS5 proxy at " + s.addr + " has unexpected version " + strconv.Itoa(int(buf[0])))
	}
	if buf[1] == 0xff {
		return errors.New("proxy: SOCKS5 proxy at " + s.addr + " requires authentication")
	}
	//See RFC 1929
	if buf[1] == proxySocks5AuthPassword {
		buf = buf[:0]
		buf = append(buf, 1 /*password protocol version*/)
		buf = append(buf, uint8(len(s.user)))
		buf = append(buf, s.user...)
		buf = append(buf, uint8(len(s.passowrd)))
		buf = append(buf, s.passowrd...)
		if _, err := conn.Write(buf); err != nil {
			return errors.New("proxy: failed to write authentication request to SOCKS5 proxy at " + s.addr + ": " + err.Error())
		}
		if _, err := io.ReadFull(conn, buf[:2]); err != nil {
			return errors.New("proxy: failed to read authentication reply from SOCKS5 proxy at " + s.addr + ": " + err.Error())
		}
		if buf[1] != 0 {
			return errors.New("proxy: SOCKS5 proxy at " + s.addr + " rejected username/password")
		}
	}
	buf = buf[:0]
	buf = append(buf, proxySocks5Version, proxySocks5Connect, 0 /*reserved*/)
	if ip := net.ParseIP(host); ip != nil {
		if ip4 := ip.To4(); ip4 != nil {
			buf = append(buf, proxySocks5IP4)
			ip = ip4
		} else {
			buf = append(buf, proxySocks5IP6)
		}
		buf = append(buf, ip...)
	} else {
		if len(host) > 255 {
			return errors.New("proxy: destination host name too long: " + host)
		}
		buf = append(buf, proxySocks5Domain)
		buf = append(buf, byte(len(host)))
		buf = append(buf, host...)
	}
	buf = append(buf, byte(port>>8), byte(port))
	if _, err := conn.Write(buf); err != nil {
		return errors.New("proxy: failed to write connect request to SOCKS5 proxy at " + s.addr + ": " + err.Error())
	}
	if _, err := io.ReadFull(conn, buf[:4]); err != nil {
		return errors.New("proxy: failed to read connect reply from SOCKS5 proxy at " + s.addr + ": " + err.Error())
	}
	failure := "unknown error"
	if int(buf[1]) < len(proxySocks5Errors) {
		failure = proxySocks5Errors[buf[1]]
	}
	if len(failure) > 0 {
		return errors.New("proxy: SOCKS5 proxy at " + s.addr + " failed to connect: " + failure)
	}
	bytesToDiscard := 0
	switch buf[3] {
	case proxySocks5IP4:
		bytesToDiscard = net.IPv4len
	case proxySocks5IP6:
		bytesToDiscard = net.IPv6len
	case proxySocks5Domain:
		_, err := io.ReadFull(conn, buf[:1])
		if err != nil {
			return errors.New("proxy: failed to read domain length from SOCKS5 proxy at " + s.addr + ": " + err.Error())
		}
		bytesToDiscard = int(buf[0])
	default:
		return errors.New("proxy: got unkown address type " + strconv.Itoa(int(buf[3])) + " from SOCKS5 proxy at " + s.addr)
	}
	if cap(buf) < bytesToDiscard {
		buf = make([]byte, bytesToDiscard)
	} else {
		buf = buf[:bytesToDiscard]
	}
	if _, err := io.ReadFull(conn, buf); err != nil {
		return errors.New("proxy: failed to read address from SOCKS5 proxy at " + s.addr + ": " + err.Error())
	}
	//Also need to discard the port number
	if _, err := io.ReadFull(conn, buf[:2]); err != nil {
		return errors.New("proxy: failed to read port from SOCKS5 proxy at " + s.addr + ": " + err.Error())
	}
	return nil
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
	var auth *proxyAuth
	if u.User != nil {
		auth = new(proxyAuth)
		auth.User = u.User.Username()
		if p, ok := u.User.Password(); ok {
			auth.Password = p
		}
	}
	switch u.Scheme {
	case "socks5":
		return prox_so
	}
}

//proxyFromEnvironment returns the dialer specified by the proxy related variables
//in the environment
func proxyFromEnvironment() proxyDialer {
	allProxy := proxyAllProxyEnv.Get()
	if len(allProxy) == 0 {
		return pd
	}
	proxyURL, err := url.Parse(allProxy)
}
