# shortscale

[![CI](https://github.com/jldec/shortscale-go/workflows/CI/badge.svg)](https://github.com/jldec/shortscale-go/actions)  
[Rust docs](https://docs.rs/shortscale) | [crates.io](https://crates.io/crates/shortscale)

Go library to convert numbers into English words.

This was originally written as an exploration of JavaScript and Rust [documented here](https://jldec.me/forays-from-node-to-rust).

The [short scale](https://en.wikipedia.org/wiki/Long_and_short_scales#Comparison),
has different words for each power of 1000.

This library expresses numbers from zero to thousands,
millions, billions, trillions, and quadrillions, up to 999_999_999_999_999_999.

### Function
```go
func Shortscale(n uint64) string
```

### Example
```go
import (
    "fmt"
    shortscale "github.com/jldec/shortscale-go"
)

// four hundred and twenty billion nine hundred and ninety nine thousand and fifteen
fmt.println(shortscale.Shortscale(420_000_999_015))
```
