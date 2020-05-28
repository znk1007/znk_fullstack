package ngxio

//SessionIDGenerator generates new session id.
type SessionIDGenerator interface {
	NewID() string
}
