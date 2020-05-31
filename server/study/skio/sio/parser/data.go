package parser

import (
	"bytes"
	"encoding/json"
	"strconv"
)

//Type of packet
type Type byte

const (
	//Connect type, present connect to server from client
	Connect Type = iota
	//Disconnect type, present disconnect to server from client
	Disconnect
	//Event type, present event from client
	Event
	//Ack type, present ack info from client
	Ack
	//Error type, an error occur when transport
	Error
	//binaryEvent type, present bianry event type when transport
	binaryEvent
	//binaryAck type, present binary ack when transport
	binaryAck
	//typeMax placeholder type
	typeMax
)

//Header of packet
type Header struct {
	Type      Type
	Namespace string
	ID        uint64
	NeedAck   bool
}

//Buffer is an binary buffer handler used in emit args.
//All buffers will be sent as binary in the transport layer.
type Buffer struct {
	Data     []byte
	isBinary bool
	num      uint64
}

//MarshalJSON marshals to JSON data from buffer
func (b Buffer) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	if err := b.marshalJSONBuf(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (b *Buffer) marshalJSONBuf(buf *bytes.Buffer) error {
	encode := b.encodeText
	if b.isBinary {
		encode = b.encodeBinary
	}
	return encode(buf)
}

func (b *Buffer) encodeText(buf *bytes.Buffer) error {
	buf.WriteString("{\"type\":\"Buffer\", \"data\":[")
	for i, d := range b.Data {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(strconv.Itoa(int(d)))
	}
	buf.WriteString("]}")
	return nil
}

func (b *Buffer) encodeBinary(buf *bytes.Buffer) error {
	buf.WriteString("\"_placeholder\":true,\"num\":")
	buf.WriteString(strconv.FormatUint(b.num, 10))
	return nil
}

//UnmarshalJSON unmarshals data from JSON
func (b *Buffer) UnmarshalJSON(bt []byte) error {
	var data struct {
		Data        []byte
		Placeholder bool `json:"_placeholder"`
		Num         uint64
	}
	if err := json.Unmarshal(bt, &data); err != nil {
		return err
	}
	b.isBinary = data.Placeholder
	b.Data = data.Data
	b.num = data.Num
	return nil
}
