package parser

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

}
