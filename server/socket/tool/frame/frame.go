package skframe

//FrameType type of frames
type FrameType byte

const (
	//FrameString identifies a string frame
	FrameString FrameType = iota
	//FrameBinary identifier a binary frame
	FrameBinary
)
