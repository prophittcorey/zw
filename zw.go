package zw

import (
	"bytes"
	"io"
)

const (
	empty = ""

	// ZWSP represents zero-width space.
	ZWSP = '\u200B'

	// ZWNBSP represents zero-width no-break space.
	ZWNBSP = '\uFEFF'

	// ZWJ represents zero-width joiner.
	ZWJ = '\u200D'

	// ZWNJ represents zero-width non-joiner.
	ZWNJ = '\u200C'
)

var (
	/* default runes for Trim */
	runes = []rune{ZWJ, ZWNJ, ZWSP, ZWNBSP}

	/* use for O(1) lookups */
	runemap = map[rune]struct{}{
		ZWSP:   struct{}{},
		ZWNBSP: struct{}{},
		ZWJ:    struct{}{},
		ZWNJ:   struct{}{},
	}
)

// Trim scans the reader and outputs a new slice of bytes containing only
// non-zero width runes.
func Trim(rr io.RuneReader, rs ...rune) []byte {
	if len(rs) == 0 {
		rs = runes
	}

	var buffer bytes.Buffer

	for {
		r, _, err := rr.ReadRune()

		if err != nil {
			break
		}

		if _, ok := runemap[r]; ok {
			continue
		}

		buffer.WriteRune(r)
	}

	return buffer.Bytes()
}

// Present returns true if any known zero width runes are present in the stream.
func Present(rr io.RuneReader) bool {
	for {
		r, _, err := rr.ReadRune()

		if err != nil {
			break
		}

		if _, ok := runemap[r]; ok {
			return true
		}
	}

	return false
}
