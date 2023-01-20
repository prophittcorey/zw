# ZW: Zero Width

[![Go Reference](https://pkg.go.dev/badge/github.com/prophittcorey/zw.svg)](https://pkg.go.dev/github.com/prophittcorey/zw)

A golang package for detecting and trimming zero width runes.

## Why this package?

Other packages exist for cleaning zero width characters, but most other
packages work either on strings or bytes. The end result is a large number of
unnecessary (and expensive) conversions from one type to another.

This package was written with minimal memory reads, writes and copies in mind.

## Package Usage

```golang
package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/prophittcorey/zw"
)

var (
	text = "Here's some text with \u200Brunes\uFEFF."

	bs = []byte(text)
)

func main() {
	/* works on strings */
	if zw.Present(strings.NewReader(text)) {
		fmt.Println("It's true, the string contains zero width runes.")
	}

	/* same code works on byte slices */
	if zw.Present(bytes.NewReader(bs)) {
		fmt.Println("It's true, the bytes contain zero width runes.")
	}

	/* cleaning works much the same way */
	cleaned := zw.Trim(strings.NewReader(text))

	if string(cleaned) == "Here's some text with runes." {
		fmt.Println("All zero width runes were removed.")
	}
}
```

## License

The source code for this repository is licensed under the MIT license, which you can
find in the [LICENSE](LICENSE.md) file.
