# ZW: Zero Width

A golang package for detecting and handling zero width runes.

## Package Usage

```golang
package main

import (
  "github.com/prophittcorey/zw"
)

const (
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
