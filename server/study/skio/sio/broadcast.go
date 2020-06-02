package sio

import "sync"

type EachFunc func(Conn)

//Broadcast is the adaptor to handle broadcasts & rooms for socket.io server API
type Broadcast interface {
	Join(room string, connection Conn)            //causes the connection to join a room
	Leave(room string, connection Conn)           //causes the connection to leave room
	LeaveAll()                                    // causes given connection to leave all rooms
	Clear(room string)                            //causes removal of all connections from the room
	Send(room, event string, args ...interface{}) // will send an event with args to the room
	SendAll(event string, args ...interface{})    //will send an event with args to all the rooms
	ForEach(room string, f EachFunc)              //refresh room for the Conn
	Len(room string) int                          //gives number of connections in the room
	Rooms(connection Conn) []string               //gives list of all the rooms if no connection given, else list of all the rooms that the connection joined
}

type broadcast struct {
	rooms map[string]map[string]Conn //map of rooms where each room contains a map of connection id to connections in that room
	lock  sync.RWMutex
}

//NewBroadcast creates a new broadcast adapter
func NewBroadcast() *broadcast {
	return &broadcast{
		rooms: make(map[string]map[string]Conn),
	}
}

//Join joins the given connection to the broadcast room
func (bc *broadcast) Join(room string, connection Conn) {
	//get write lock
	bc.lock.Lock()
	defer bc.lock.Unlock()
	//check if room already has connection mappings, create one if not
	if _, ok := bc.rooms[room]; !ok {
		bc.rooms[room] = make(map[string]Conn)
	}
	//add the connection to the rooms connectin map
	bc.rooms[room][connection.ID()] = connection
}

//Leave leaves the given connection from given room if exist
func (bc *broadcast) Leave(room string, connection Conn) {
	//get write lock
	bc.lock.Lock()
	defer bc.lock.Unlock()

	//check if rooms connection
	if connections, ok := bc.rooms[room]; ok {
		delete(connections, connection.ID())
		//check if no more connection is left to the room, then delete the room
		if len(connections) == 0 {
			delete(bc.rooms, room)
		}
	}
}

//LeaveAll leaves the given connection from all rooms
func (bc *broadcast) LeaveAll(connection Conn) {
	bc.lock.Lock()
	defer bc.lock.Unlock()
	// iterate through each room
	for room, connections := range bc.rooms {
		//remove the connection from the rooms connections
		delete(connections, connection.ID())
		//check if no more connection is left to the room, then delete the room
		if len(connections) == 0 {
			delete(bc.rooms, room)
		}
	}
}

//Clear clears the room
func (bc *broadcast) Clear(room string) {
	//get write lock
	bc.lock.Lock()
	defer bc.lock.Unlock()
	//delete the room
	delete(bc.rooms, room)
}

//Send sends given event & args to all the connections in the specified room
func (bc *broadcast) Send(room, event string, args ...interface{}) {
	//get a read lock
	bc.lock.RLock()
	defer bc.lock.RUnlock()

	//iterate through each connection in the room
	for _, connection := range bc.rooms[room] {
		//emit the event to the connection
		connection.Emit(event, args...)
	}
}

//SendAll sends given event & args to all the connections to all the rooms
func (bc *broadcast) SendAll(event string, args ...interface{}) {
	//get a read lock
	bc.lock.RLock()
	defer bc.lock.RUnlock()
	//iterate through each room
	for _, connections := range bc.rooms {
		//iterate through each connection in the room
		for _, connection := range connections {
			//emit the event to the connection
			connection.Emit(event, args...)
		}
	}
}

//ForEach sends data returned by DataFunc, if the return is 'ok' (second return)
func (bc *broadcast) ForEach(room string, f EachFunc) {
	//get a read lock
	bc.lock.RLock()
	defer bc.lock.RUnlock()

	occupants, ok := bc.rooms[room]
	if !ok {
		return
	}
	for _, connection := range occupants {
		f(connection)
	}
}

//Len gives number of connections in the room
func (bc *broadcast) Len(room string) int {
	//get a read lock
	bc.lock.RLock()
	defer bc.lock.RUnlock()
	return len(bc.rooms[room])
}

//Rooms gives the list of all the rooms available for broadcast in case of
//no connection is given, in case of a connection is given, it gives list
//of all the rooms the connection is joined to
func (bc *broadcast) Rooms(connection Conn) []string {
	//get a read lock
	bc.lock.RLock()
	defer bc.lock.RUnlock()

	rooms := make([]string, 0)
}
