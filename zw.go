package zwnj

import (
	"io"
	"strings"
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

	replacer = strings.NewReplacer(string(ZWSP), empty, string(ZWNBSP), empty, string(ZWJ), empty, string(ZWNJ), empty)
)

// Trim removes the specified runes from the string. By default, removes
// all zero width characters.
func Trim(str string, rs ...rune) string {
	if len(rs) == 0 {
		rs = runes
	}

	return replacer.Replace(str)
}

// Present returns true if any known zero width runes are present in the stream.
func Present(rs io.RuneReader) bool {
	for {
		r, _, err := rs.ReadRune()

		if _, ok := runemap[r]; ok {
			return true
		}

		if err != nil {
			break
		}
	}

	return false
}
