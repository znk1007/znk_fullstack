package base

//PacketType is the type of packet
type PacketType int

const (
	//OPEN is sent from the server when a new transport is opened (recheck).
	OPEN PacketType = iota
	//CLOSE is request the close of this transport but does not shutdown the
	//connection itself.
	CLOSE
	//PING is sent by client. Server should answer with a pong packet
	//containing the same data.
	PING
	//PONG is sent by the server to respond to ping packets.
	PONG
	//MESSAGE is actual message, client and server should call their callbacks
	//with the data.
	MESSAGE
	//UPGRADE is sent before ngxio switches a transport to test if server
	//and client can communicate over this transport. If this test succeed,
	//when client sends an upgrade packets which requests the server to flush
	//tis cache on the old transport and switch to the new transport.
	UPGRADE
)
