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
)
