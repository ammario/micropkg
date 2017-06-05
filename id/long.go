package id

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/ammario/frandom"
	"github.com/pkg/errors"
)

// Long defines a long id
type Long [16]byte

// parse errors
var (
	ErrWrongSize = errors.Errorf("id in string form should be exactly 33 bytes")
)

// GenLong generates an ID of which the first 8 bytes of the long ID are from a Short()
// The last 8 bytes are from a CSPRNG.
// The generated long is a suitable for use as a secret.
func GenLong() (id Long) {
	binary.BigEndian.PutUint64(id[:8], uint64(GenShort()))
	if _, err := frandom.Read(id[8:]); err != nil {
		panic(err)
	}
	return
}

// String returns a hex representation of l.
// It splits the Short and random component with a `-`
func (l Long) String() string {
	return fmt.Sprintf("%016x-%016x", l[:8], l[8:])
}

// ParseLong parses the String() representation of a long ID
func ParseLong(l string) (long Long, err error) {
	if len(l) != 33 {
		return long, ErrWrongSize
	}

	short, err := hex.DecodeString(l[:16])
	if err != nil {
		return long, errors.Wrap(err, "failed to decode short portion")
	}

	rand, err := hex.DecodeString(l[17:])
	if err != nil {
		return long, errors.Wrap(err, "failed to decode rand portion")
	}

	copy(long[:8], short)
	copy(long[8:], rand)

	return
}
