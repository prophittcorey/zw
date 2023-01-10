package zwnj

import "testing"

func TestPresence(t *testing.T) {
	given := []string{
		"\u200Bfrom the start",
		"at the end\u200D",
		"in the\u200C middle",
		"in the\uFEFFmiddle",
	}

	for _, str := range given {
		if !Present(str) {
			t.Fatalf("failed to detect zero width presence in %s", str)
		}
	}

	given = []string{
		"",
		"just a string",
		"\n",
	}

	for _, str := range given {
		if Present(str) {
			t.Fatalf("falsely detected a zero width presence in %v", str)
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
		trimmed := Trim(str)

		if trimmed != expect[i] {
			t.Fatalf("failed to trim %s; got %s", str, trimmed)
		}
	}
}