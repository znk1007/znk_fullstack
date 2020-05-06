package skframe

//FrameType type of frames
type FrameType byte

const (
	//FrameString identifies a string frame
	FrameString FrameType = iota
	//FrameBinary identifier a binary frame
	FrameBinary
)

//ByteToFrameType converts a byte to FrameType
func ByteToFrameType(b byte) FrameType {
	return FrameType(b)
}

//Byte returns type in byte
func (ft FrameType) Byte() byte {
	return byte(ft)
}

type FrameReader interface {
	NextReader() (FrameType, Package)
}
