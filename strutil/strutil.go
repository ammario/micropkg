//Package strutil provides string utilities
package strutil

import (
	"bytes"
	"unicode/utf8"
)

//Ellipsis modifies str to end with an ... if it is too long.
//
//Returned string will have a rune length of maxlen at most
func Ellipsis(str string, maxLen int) string {
	if len(str) < maxLen {
		//fast path if there's no way string can be  longer than maxlen
		return str
	}

	lookbehindPos := [...]int{0, 0, 0}

	var (
		count int
		pos   int
	)

	for {
		_, size := utf8.DecodeRuneInString(str[pos:])
		pos += size
		count++
		if count == maxLen || pos == len(str) {
			break
		}

		//shift from the right
		copy(lookbehindPos[:], lookbehindPos[1:])
		lookbehindPos[2] = pos
	}

	if count <= maxLen && pos == len(str) {
		return str
	}

	ret := &bytes.Buffer{}
	ret.WriteString(str[:lookbehindPos[0]])
	ret.WriteString("...")

	return ret.String()
}
