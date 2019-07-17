package socket

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"os"
	"sort"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"
)

// ID  socket唯一标识
type ID [rawLen]byte

const (
	encodedLen = 20                                 //24
	rawLen     = 12                                 //15
	encoding   = "0123456789abcdefghijklmnopqrstuv" //"0123456789abcdefghijklmnopqrstuvwxyz"
)

var (
	//ErrInvalidObjectID is returned when try to parse an invlid ID
	ErrInvalidObjectID = errors.New("Invalid ID")
	// IDCounter is rand int64
	IDCounter = randInt()
	//机器id
	machineID = getMachineID()
	/// 进程id
	pid = os.Getpid()
	/// 非空标识
	nilObjectID ID
	/// decoding map for base36 encoding
	dec [256]byte
)

func init() {
	for idx := 0; idx < len(dec); idx++ {
		dec[idx] = 0xFF
	}
	for idx := 0; idx < len(encoding); idx++ {
		dec[encoding[idx]] = byte(idx)
	}
	b, err := ioutil.ReadFile("/proc/self/cpuset")
	if err == nil && len(b) > 1 {
		pid ^= int(crc32.ChecksumIEEE(b))
	}
}

/// linux系统机器码
func linuxMachineID() (string, error) {
	b, err := ioutil.ReadFile("/sys/class/dmi/id/product_uuid")
	return string(b), err
}

/// freebsd系统机器码
func freeBSDMachineID() (string, error) {
	return syscall.Sysctl("kern.hostuuid")
}

/// 无法系统机器码
func fallbackMachineID() (string, error) {
	return "", errors.New("no implemented machine id")
}

/// darwin系统机器码
func darwinPlatformMachineID() (string, error) {
	return syscall.Sysctl("kern.uuid")
}

// randInt generates a random uint64
func randInt64() uint64 {
	b := make([]byte, 6)
	if _, err := rand.Reader.Read(b); err != nil {
		panic(fmt.Errorf("cannot generate random number: %v;", err))
	}
	return uint64(b[0])<<48 | uint64(b[1])<<32 | uint64(b[2])<<24 | uint64(b[3])<<16 | uint64(b[4])<<8 | uint64(b[5])
}

func randInt() uint32 {
	b := make([]byte, 3)
	if _, err := rand.Reader.Read(b); err != nil {
		panic(fmt.Errorf("cannot generate random number: %v;", err))
	}
	return uint32(b[0])<<16 | uint32(b[1])<<8 | uint32(b[2])
}

/**
 * 机器ID
 */
func getMachineID() []byte {
	var mID string
	var mErr error
	var tempErr error
	id := make([]byte, 3)
	if mID, tempErr = darwinPlatformMachineID(); tempErr != nil {
		mErr = tempErr
	}

	if mID, tempErr = freeBSDMachineID(); tempErr != nil {
		mErr = tempErr
	}

	if mID, tempErr = linuxMachineID(); tempErr != nil {
		mErr = tempErr
	}
	if mErr != nil || len(mID) == 0 {
		mID, tempErr = os.Hostname()
		if tempErr != nil {
			panic(fmt.Errorf("cannot get hostname, err: %v mErr: %v", tempErr, mErr))
		} else {
			if _, randErr := rand.Reader.Read(id); randErr != nil {
				panic(fmt.Errorf("cannot generate a random number err: %v", randErr))
			}
		}
	}
	hw := md5.New()
	hw.Write([]byte(mID))
	copy(id, hw.Sum(nil))
	return id
}

// NewSocketID 以当前时间初始化
func NewSocketID() ID {
	return NewSocketIDWithTime(time.Now())
}

// NewSocketIDWithTime  以指定时间初始化
func NewSocketIDWithTime(t time.Time) ID {
	var id ID
	// timestamp 4bytes
	binary.BigEndian.PutUint32(id[:], uint32(t.Unix()))
	//machine id 3 bytes
	id[4] = machineID[0]
	id[5] = machineID[1]
	id[6] = machineID[2]
	//pid 2bytes
	id[7] = byte(pid >> 8)
	id[8] = byte(pid)
	// increment counter, 4bytes
	i := atomic.AddUint32(&IDCounter, 1)
	id[9] = byte(i >> 16)
	id[10] = byte(i >> 8)
	id[11] = byte(i)
	return id
}

// FromString 以指定字符串生成ObjectID
func FromString(id string) (ID, error) {
	i := &ID{}
	err := i.UnmarshalText([]byte(id))
	return *i, err
}

// UnmarshalText 解析指定字节
func (id *ID) UnmarshalText(bytes []byte) error {
	if len(bytes) != encodedLen {
		return ErrInvalidObjectID
	}
	for _, b := range bytes {
		if dec[b] == 0xFF {
			return ErrInvalidObjectID
		}
	}
	decode(id, bytes)
	return nil
}

// Time 从ObjectID解析时间
func (id ID) Time() time.Time {
	secs := int64(binary.BigEndian.Uint32(id[0:4]))
	return time.Unix(secs, 0)
}

// Machine 从ObjectID解析机器ID
func (id ID) Machine() []byte {
	return id[4:7]
}

// Pid 从ObjectID解析进程ID
func (id ID) Pid() uint16 {
	return binary.BigEndian.Uint16(id[7:9])
}

// Counter 从ObjectID解析随机数
func (id ID) Counter() int32 {
	b := id[9:12]
	// Counter is stored as big-endian 3-byte value
	return int32(uint32(b[0])<<16 | uint32(b[1])<<8 | uint32(b[2]))
}

// Scan 扫码值类型，并返回结果
func (id *ID) Scan(value interface{}) (err error) {
	switch val := value.(type) {
	case string:
		return id.UnmarshalText([]byte(val))
	case []byte:
		return id.UnmarshalText(val)
	case nil:
		*id = nilObjectID
		return nil
	default:
		return fmt.Errorf("scanning unsupported type: %T", value)
	}
}

// IsNil ObjectID是否为空
func (id ID) IsNil() bool {
	return id == nilObjectID
}

// String ObjectID字符串显示
func (id ID) String() string {
	text := make([]byte, encodedLen)
	encode(text, id[:])
	return *(*string)(unsafe.Pointer(&text))
}

// MarshalJSON ObjectIDJSON字符串显示
func (id ID) MarshalJSON() ([]byte, error) {
	if id.IsNil() {
		return []byte("null"), nil
	}
	text := make([]byte, encodedLen+2)
	encode(text[1:encodedLen+1], id[:])
	text[0], text[encodedLen+1] = '"', '"'
	return text, nil
}

// Bytes 返回所有字节
func (id ID) Bytes() []byte {
	return id[:]
}

// FromBytes 根据bytes转ObjectID
func FromBytes(b []byte) (ID, error) {
	var id ID
	if len(b) != rawLen {
		return id, ErrInvalidObjectID
	}
	copy(id[:], b)
	return id, nil
}

// Compare 对比两个ObjectID
func (id ID) Compare(other ID) int {
	return bytes.Compare(id[:], other[:])
}

type sorter []ID

// ObjectID长度
func (s sorter) Len() int {
	return len(s)
}

// ObjectID字节比较大小
func (s sorter) Less(i, j int) bool {
	return s[i].Compare(s[j]) < 0
}

func (s sorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Sort ObjectID排序
func Sort(ids []ID) {
	sort.Sort(sorter(ids))
}

/**
 * 编码
 */
func encode(dst, id []byte) {
	_ = dst[19]
	_ = id[11]

	dst[19] = encoding[(id[11]<<4)&0x1F]
	dst[18] = encoding[(id[11]>>1)&0x1F]
	dst[17] = encoding[(id[11]>>6)&0x1F|(id[10]<<2)&0x1F]
	dst[16] = encoding[id[10]>>3]
	dst[15] = encoding[id[9]&0x1F]
	dst[14] = encoding[(id[9]>>5)|(id[8]<<3)&0x1F]
	dst[13] = encoding[(id[8]>>2)&0x1F]
	dst[12] = encoding[id[8]>>7|(id[7]<<1)&0x1F]
	dst[11] = encoding[(id[7]>>4)&0x1F|(id[6]<<4)&0x1F]
	dst[10] = encoding[(id[6]>>1)&0x1F]
	dst[9] = encoding[(id[6]>>6)&0x1F|(id[5]<<2)&0x1F]
	dst[8] = encoding[id[5]>>3]
	dst[7] = encoding[id[4]&0x1F]
	dst[6] = encoding[id[4]>>5|(id[3]<<3)&0x1F]
	dst[5] = encoding[(id[3]>>2)&0x1F]
	dst[4] = encoding[id[3]>>7|(id[2]<<1)&0x1F]
	dst[3] = encoding[(id[2]>>4)&0x1F|(id[1]<<4)&0x1F]
	dst[2] = encoding[(id[1]>>1)&0x1F]
	dst[1] = encoding[(id[1]>>6)&0x1F|(id[0]<<2)&0x1F]
	dst[0] = encoding[id[0]>>3]
}

/**
 * 解码
 */
func decode(id *ID, src []byte) {
	_ = src[19]
	_ = id[11]
	id[11] = dec[src[17]]<<6 | dec[src[18]]<<1 | dec[src[19]]>>4
	id[10] = dec[src[16]]<<3 | dec[src[17]]>>2
	id[9] = dec[src[14]]<<5 | dec[src[15]]
	id[8] = dec[src[12]]<<7 | dec[src[13]]<<2 | dec[src[14]]>>3
	id[7] = dec[src[11]]<<4 | dec[src[12]]>>1
	id[6] = dec[src[9]]<<6 | dec[src[10]]<<1 | dec[src[11]]>>4
	id[5] = dec[src[8]]<<3 | dec[src[9]]>>2
	id[4] = dec[src[6]]<<5 | dec[src[7]]
	id[3] = dec[src[4]]<<7 | dec[src[5]]<<2 | dec[src[6]]>>3
	id[2] = dec[src[3]]<<4 | dec[src[4]]>>1
	id[1] = dec[src[1]]<<6 | dec[src[2]]<<1 | dec[src[3]]>>4
	id[0] = dec[src[0]]<<3 | dec[src[1]]>>2
}
