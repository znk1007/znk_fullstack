package skpkg

//PacketType type of packet
type PacketType int

const (
	//Open is sent from the server when a new transport is opened (recheck)
	Open PacketType = iota
	//Close is request the close of this transport but does not shutdown the connection itself
	Close
	//Ping is sent by the client. Server should answer with a pong packet containing the same data.
	Ping
	//Pong is sent by the server to respond to ping packet
	Pong
	//Message is actual message, client and server should call their callbacks with the data
	Message
	// Upgrade is sent before engine.io switches a transport to test if server and
	// client can communicate over this transport. If this test succeed, the client sends
	// an upgrade packets which requests the server to flush its cache on the old transport
	// and switch to the new transport.
	Upgrade
	//Noop is a noop packet. Used primarily to force a poll cycle when an incoming
	//websocket connection is received.
	Noop
)

func (pt PacketType) String() string {
	switch pt {
	case Open:
		return "open"
	case Close:
		return "close"
	case Ping:
		return "ping"
	case Pong:
		return "pong"
	case Message:
		return "message"
	case Upgrade:
		return "upgrade"
	case Noop:
		return "noop"
	}
	return "unknown"
}

//StringByte converts a PacketType to byte in string
func (pt PacketType) StringByte() byte {
	return byte(pt) + '0'
}

//BinaryByte converts a PacketType to byte in binary
func (pt PacketType) BinaryByte() byte {
	return byte(pt)
}

//ByteToPacketType converts a byte to PacketType.
func ByteToPacketType(b byte, ft skframe.FrameType) PacketType {

}
