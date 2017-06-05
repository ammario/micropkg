package id

import (
	"hash/fnv"
	"os"
	"sync/atomic"
	"time"
)

// bit counts.
const (
	HostIDBits      = 8
	TimeBits        = 48
	IncrementorBits = 8
)

var inc int64
var hostID int64

func init() {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	hash := fnv.New64a()
	hash.Write([]byte(hostname))
	hostID = (int64(hash.Sum64()) << (64 - HostIDBits))

	// make sure first bit is 0
	hostID &= int64(0x7FFFFFFFFFFFFFFF)
}

//HostID returns the host ID for the current machine.
//The first bit is always 0
func HostID() int64 {
	return hostID >> (64 - HostIDBits)
}

// GenShort returns a 64 bit id.
// The first 8 bits are from the server id
// The next 48 bits are from the timestamp with microsecond precision
// The last 8 bits are from an incrementor
func GenShort() int64 {
	micros := time.Now().UnixNano() / (1000)
	//clear server ID bits with enough room for inc bits then create room for server id bits
	time := (int64(micros) << (HostIDBits + IncrementorBits)) >> HostIDBits

	//clear last 50 bits, then move down 50
	i := (atomic.AddInt64(&inc, 1) << (HostIDBits + TimeBits)) >> (HostIDBits + TimeBits)

	return hostID + time + i
}

// SplitShort extracts information from the short ID
// The time parameter may only be useful if the ID was created in
// the latest 9 year interval (see ShortTimeReset).
func SplitShort(s int64) (hostID int64, createdAt time.Time, incrementor int64) {

	//shift by extra one to account for the forcefully left 0 last bit
	hostID = int64(s) >> (64 - (HostIDBits))

	micros := int64(time.Now().UnixNano()) / (1000)

	//clear time bits and make room for them
	micros = (micros >> TimeBits) << TimeBits

	//add microseconds from short
	micros += ((int64(s) >> IncrementorBits) << (IncrementorBits + HostIDBits)) >> (IncrementorBits + HostIDBits)

	incrementor = (int64(s) << (64 - IncrementorBits)) >> (64 - IncrementorBits)

	return hostID, time.Unix(0, int64(micros*1000)), incrementor
}

// ShortTimeReset the current bounds of
// validity for timestamps extracted from shorts
func ShortTimeReset() (last time.Time, next time.Time) {
	_, last, _ = SplitShort(0)
	_, next, _ = SplitShort(^(int64(0) & 0))
	return last, next
}
