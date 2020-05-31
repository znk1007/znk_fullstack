package parser

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type noBufferStruct struct {
	Str   string            `json:"str"`
	I     int               `json:"i"`
	Array []int             `json:"array"`
	Map   map[string]string `json:"map"`
}

type bufferStruct struct {
	I      int     `json:"i"`
	Buffer *Buffer `json:"buf"`
}

type bufferInnerStruct struct {
	I      int                `json:"i"`
	Buffer *Buffer            `json:"buf"`
	Inner  *bufferInnerStruct `json:"inner,omitempty"`
}

var tests = []struct {
	Name   string
	Header Header
	Event  string
	Var    []interface{}
	Datas  [][]byte
}{
	{"Empty", Header{Connect, "", 0, false}, "", nil, [][]byte{
		[]byte("0"),
	}},
	{"Data", Header{Error, "", 0, false}, "", []interface{}{"error"}, [][]byte{
		[]byte("4[\"error\"]\n"),
	}},
	{"BData", Header{Event, "", 0, false}, "msg", []interface{}{
		&Buffer{Data: []byte{1, 2, 3}},
	}, [][]byte{
		[]byte("51-[\"msg\",{\"_placeholder\":true,\"num\":0}]\n"),
		{1, 2, 3},
	}},
	{"ID", Header{Connect, "", 0, true}, "", nil, [][]byte{
		[]byte("00"),
	}},
	{"IDData", Header{Ack, "", 13, true}, "", []interface{}{"error"}, [][]byte{
		[]byte("313[\"error\"]\n"),
	}},
	{"IDBData", Header{Ack, "", 13, true}, "", []interface{}{
		&Buffer{
			Data: []byte{1, 2, 3},
		},
	}, [][]byte{
		[]byte("61-13[{\"_placeholder\":true,\"num\":0}]\n"),
		{1, 2, 3},
	}},
	{"Namespace", Header{Disconnect, "/woot", 0, false}, "", nil, [][]byte{
		[]byte("1/woot"),
	}},
	{"NamespaceData", Header{Event, "/woot", 0, false}, "msg", []interface{}{
		1,
	}, [][]byte{
		[]byte("2/woot,[\"msg\",1]\n"),
	}},
	{"NamespaceBData", Header{Event, "/woot", 0, false}, "msg", []interface{}{
		&Buffer{Data: []byte{2, 3, 4}},
	}, [][]byte{
		[]byte("51-/woot,[\"msg\",{\"_placeholder\":true,\"num\":0}]\n"),
		{2, 3, 4},
	}},
	{"NamespaceID", Header{Disconnect, "/woot", 1, true}, "", nil, [][]byte{
		[]byte("1/woot,1"),
	}},
	{"NamespaceIDData", Header{Event, "/woot", 1, true}, "msg", []interface{}{
		1,
	}, [][]byte{
		[]byte("2/woot,1[\"msg\",1]\n"),
	}},
	{"NamespaceIDBData", Header{Event, "/woot", 1, true}, "msg", []interface{}{
		&Buffer{Data: []byte{2, 3, 4}},
	}, [][]byte{
		[]byte("51-/woot,1[\"msg\",{\"_placeholder\":true,\"num\":0}]\n"),
		{2, 3, 4},
	}},
}

var attachmentTests = []struct {
	buffer         Buffer
	textEncoding   string
	binaryEncoding string
}{
	{
		Buffer{[]byte{1, 255}, false, 0},
		`{"type":"Buffer","data":[1,255]}`,
		`{"_placeholder":true,"num":0}`,
	},
	{
		Buffer{[]byte{}, false, 1},
		`{"type":"Buffer","data":[]}`,
		`{"_placeholder":true,"num":1}`,
	},
	{
		Buffer{nil, false, 2},
		`{"type":"Buffer","data":[]}`,
		`{"_placeholder":true,"num":2}`,
	},
}

func TestAttachmentEncodeText(t *testing.T) {
	should := assert.New(t)
	must := require.New(t)

	for _, test := range attachmentTests {
		b := test.buffer
		b.isBinary = false
		j, e := json.Marshal(b)
		must.Nil(e)
		should.Equal(test.textEncoding, string(j))
	}
}

func TestAttachmentEncodeBinary(t *testing.T) {
	should := assert.New(t)
	must := require.New(t)

	for _, test := range attachmentTests {
		b := test.buffer
		b.isBinary = true
		j, e := json.Marshal(b)
		must.Nil(e)
		should.Equal(test.binaryEncoding, string(j))
	}
}

func TestAttachmentDecodeText(t *testing.T) {
	should := assert.New(t)
	must := require.New(t)

	for _, test := range attachmentTests {
		var b Buffer
		err := json.Unmarshal([]byte(test.textEncoding), &b)
		must.Nil(err)
		should.False(b.isBinary)
		if len(test.buffer.Data) == 0 {
			continue
		}
		should.Equal(test.buffer.Data, b.Data)
	}
}

func TestAttachmentDecodeBinary(t *testing.T) {
	should := assert.New(t)
	must := require.New(t)

	for _, test := range attachmentTests {
		var b Buffer
		err := json.Unmarshal([]byte(test.binaryEncoding), &b)
		must.Nil(err)
		should.True(b.isBinary)
		should.Equal(test.buffer.num, b.num)
	}
}
