package id

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShort(t *testing.T) {
	last, next := ShortTimeReset()
	fmt.Printf("Short Reset: Last: %v, Next: %v\n", last, next)

	t.Run("GenShort() + Split()", func(t *testing.T) {
		inc = 0
		for i := 0; i < 5; i++ {
			s := GenShort()
			sid, createdAt, inc := SplitShort(s)
			//eyeball testing
			fmt.Printf("%v: %s (%b) | sid: %x, createdAt: %v, inc: %x\n", i, s, s, sid, createdAt, inc)
			assert.EqualValues(t, sid, HostID())
			assert.EqualValues(t, i+1, inc)
		}
	})

	// t.Run("String()", func(t *testing.T) {
	// 	assert.Equal(t, "0000000000014288", Short(0x014288).String())
	// })

}

func BenchmarkShort(b *testing.B) {
	b.Run("GenShort()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			GenShort()
		}
	})

	// b.Run("String()", func(b *testing.B) {
	// 	short := GenShort()
	// 	b.ResetTimer()

	// 	for i := 0; i < b.N; i++ {
	// 		 short.String()
	// 	}
	// })
}
