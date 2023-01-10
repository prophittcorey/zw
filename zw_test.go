package zwnj

import (
	"bytes"
	"strings"
	"testing"
)

func TestPresent(t *testing.T) {
	given := []string{
		"\u200Bfrom the start",
		"at the end\u200D",
		"in the\u200C middle",
		"in the\uFEFFmiddle",
	}

	for _, str := range given {
		if !Present(strings.NewReader(str)) {
			t.Fatalf("failed to detect zero width presence in %s", str)
		}
	}

	given = []string{
		"",
		"just a string",
		"\n",
	}

	for _, str := range given {
		if Present(strings.NewReader(str)) {
			t.Fatalf("falsely detected a zero width presence in %v", str)
		}
	}
}

func TestPresentBytes(t *testing.T) {
	given := [][]byte{
		[]byte("\u200Bfrom the start"),
		[]byte("at the end\u200D"),
		[]byte("in the\u200C middle"),
		[]byte("in the\uFEFFmiddle"),
	}

	for _, bs := range given {
		if !Present(bytes.NewReader(bs)) {
			t.Fatalf("failed to detect zero width presence in %s", string(bs))
		}
	}

	given = [][]byte{
		[]byte(""),
		[]byte("just a string"),
		[]byte("\n"),
	}

	for _, bs := range given {
		if Present(bytes.NewReader(bs)) {
			t.Fatalf("falsely detected a zero width presence in %s", string(bs))
		}
	}
}

func TestTrim(t *testing.T) {
	given := []string{
		"",
		"\u200Bfrom the start",
		"at the end\u200D",
		"in the\u200C middle",
		"in the\uFEFFmiddle",
		"\t\n",
	}

	expect := []string{
		"",
		"from the start",
		"at the end",
		"in the middle",
		"in themiddle",
		"\t\n",
	}

	for i, str := range given {
		trimmed := Trim(strings.NewReader(str))

		if string(trimmed) != expect[i] {
			t.Fatalf("failed to trim %s; got %s", str, trimmed)
		}
	}
}

func TestTrimBytes(t *testing.T) {
	given := [][]byte{
		[]byte(""),
		[]byte("\u200Bfrom the start"),
		[]byte("at the end\u200D"),
		[]byte("in the\u200C middle"),
		[]byte("in the\uFEFFmiddle"),
	}

	expect := [][]byte{
		[]byte{},
		[]byte("from the start"),
		[]byte("at the end"),
		[]byte("in the middle"),
		[]byte("in themiddle"),
		[]byte("\t\n"),
	}

	for i, bs := range given {
		trimmed := Trim(bytes.NewReader(bs))

		if !bytes.Equal(trimmed, expect[i]) {
			t.Fatalf("failed to trim %s; got %v", string(bs), trimmed)
		}
	}
}
