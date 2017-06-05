package id

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLong(t *testing.T) {
	t.Run("GenLong()", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			fmt.Printf("Long: %v\n", GenLong())
		}
	})

	t.Run("ParseLong()", func(t *testing.T) {
		t.Run("Good", func(t *testing.T) {
			want := GenLong()
			got, err := ParseLong(want.String())
			require.Nil(t, err)
			require.Equal(t, want, got)
		})

		t.Run("Bad Size", func(t *testing.T) {
			_, err := ParseLong(GenLong().String() + "ab")
			require.NotNil(t, err)
		})

		t.Run("Bad Hex", func(t *testing.T) {
			str := GenLong().String()
			str = "O" + str[1:]
			_, err := ParseLong(str)
			require.NotNil(t, err)
		})
	})
}

func BenchmarkLong(b *testing.B) {
	b.Run("GenLong()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			GenLong()
		}
	})
}
