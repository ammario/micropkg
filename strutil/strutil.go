//Package strutil provides string utilities
package strutil

import (
	"bytes"
	"regexp"
	"strings"
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

var camel = regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)")

// CamelToSnake converts a camel case string to it's snake representation.
// Taken from https://gist.github.com/regeda/969a067ff4ed6ffa8ed6.
func CamelToSnake(s string) string {
	var a []string
	for _, sub := range camel.FindAllStringSubmatch(s, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}
	return strings.ToLower(strings.Join(a, "_"))
}
