package strutil

import (
	"strings"
	"testing"
)

func BenchmarkEllipsis(b *testing.B) {
	b.SetBytes(100)
	tstString := strings.Repeat("h", 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Ellipsis(tstString, 100)
	}
}
