package ws

import (
	"net/http"
	"reflect"
	"testing"
)

var equalASCIIFoldTests = []struct {
	t, s string
	eq   bool
}{
	{"WebSocket", "websocket", true},
	{"websocket", "WebSocket", true},
	{"Öyster", "öyster", false},
	{"WebSocket", "WetSocket", false},
}

func TestEqualASCIIFold(t *testing.T) {
	for _, tt := range equalASCIIFoldTests {
		eq := equalASCIIFold(tt.s, tt.t)
		if eq != tt.eq {
			t.Errorf("equalASCIIFold(%q, %q) = %v, want %v", tt.s, tt.t, eq, tt.eq)
		}
	}
}

var tokenListContainsValueTests = []struct {
	value string
	ok    bool
}{
	{"WebSocket", true},
	{"WEBSOCKET", true},
	{"websocket", true},
	{"websockets", false},
	{"x websocket", false},
	{"websocket x", false},
	{"other,websocket,more", true},
	{"other, websocket, more", true},
}

func TestTokenListContainsValue(t *testing.T) {
	for _, tt := range tokenListContainsValueTests {
		h := http.Header{"Upgrade": {tt.value}}
		ok := tokenListContainsValue(h, "Upgrade", "websocket")
		if ok != tt.ok {
			t.Errorf("TokenListContainsValue(h, n, %q) = %v, want %v", tt.value, ok, tt.ok)
		}
	}
}

var parseExtensionTests = []struct {
	value      string
	extensions []map[string]string
}{
	{`foo`, []map[string]string{{"": "foo"}}},
	{`foo, bar; baz=2`, []map[string]string{
		{"": "foo"},
		{"": "bar", "baz": "2"}}},
	{`foo; bar="b,a;z"`, []map[string]string{
		{"": "foo", "bar": "b,a;z"}}},
	{`foo , bar; baz = 2`, []map[string]string{
		{"": "foo"},
		{"": "bar", "baz": "2"}}},
	{`foo, bar; baz=2 junk`, []map[string]string{
		{"": "foo"}}},
	{`foo junk, bar; baz=2 junk`, nil},
	{`mux; max-channels=4; flow-control, deflate-stream`, []map[string]string{
		{"": "mux", "max-channels": "4", "flow-control": ""},
		{"": "deflate-stream"}}},
	{`permessage-foo; x="10"`, []map[string]string{
		{"": "permessage-foo", "x": "10"}}},
	{`permessage-foo; use_y, permessage-foo`, []map[string]string{
		{"": "permessage-foo", "use_y": ""},
		{"": "permessage-foo"}}},
	{`permessage-deflate; client_max_window_bits; server_max_window_bits=10 , permessage-deflate; client_max_window_bits`, []map[string]string{
		{"": "permessage-deflate", "client_max_window_bits": "", "server_max_window_bits": "10"},
		{"": "permessage-deflate", "client_max_window_bits": ""}}},
	{"permessage-deflate; server_no_context_takeover; client_max_window_bits=15", []map[string]string{
		{"": "permessage-deflate", "server_no_context_takeover": "", "client_max_window_bits": "15"},
	}},
}

func TestParseExtensions(t *testing.T) {
	for _, tt := range parseExtensionTests {
		h := http.Header{http.CanonicalHeaderKey("Sec-WebSocket-Extensions"): {tt.value}}
		extensions := parseExtensions(h)
		if !reflect.DeepEqual(extensions, tt.extensions) {
			t.Errorf("ParseExtensions(%q)\n = %v,\nwant %v", tt.value, extensions, tt.extensions)
		}
	}
}

/*
BenchmarkEmptyStr1-4   	1000000000	         0.330 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/znk_fullstack/server/socket/websocket	0.862s
Success: Benchmarks passed.
*/
func BenchmarkEmptyStr1(t *testing.B) {
	var str = "test"
	for i := 0; i < t.N; i++ {
		if str == "" {

		}
	}
}

/*
BenchmarkEmptyStr2-4   	1000000000	         0.328 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/znk_fullstack/server/socket/websocket	0.790s
Success: Benchmarks passed.
*/
func BenchmarkEmptyStr2(t *testing.B) {
	var str = "test"
	for i := 0; i < t.N; i++ {
		if len(str) == 0 {

		}
	}
}
