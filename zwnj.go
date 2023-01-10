package zwnj

import (
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
	replacer = strings.NewReplacer(string(ZWSP), empty, string(ZWNBSP), empty, string(ZWJ), empty, string(ZWNJ), empty)

	runes = map[rune]struct{}{
		ZWSP:   struct{}{},
		ZWNBSP: struct{}{},
		ZWJ:    struct{}{},
		ZWNJ:   struct{}{},
	}
)

// Trim removes the specified runes from the string. By default, removes
// all zero width characters.
func Trim(str string, rs ...rune) string {
	if len(rs) == 0 {
		for r := range runes {
			rs = append(rs, r)
		}
	}

	return replacer.Replace(str)
}

// Present returns true if the string contains any zero width character.
func Present(str string) bool {
	for _, r := range str {
		if _, ok := runes[r]; ok {
			return true
		}
	}

	return false
}
