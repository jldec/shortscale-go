# shortscale

[![CI](https://github.com/jldec/shortscale-go/workflows/CI/badge.svg)](https://github.com/jldec/shortscale-go/actions)  
Go.dev [github.com/jldec/shortscale-go](https://pkg.go.dev/github.com/jldec/shortscale-go)

Go package to convert numbers into English words.

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
fmt.Println(shortscale.Shortscale(420_000_999_015))
```

### Benchmarks

```
goos: darwin
goarch: amd64
pkg: github.com/jldec/shortscale-go
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
4053015	       291.8 ns/op	     120 B/op	       4 allocs/op

goos: linux
goarch: amd64
pkg: github.com/jldec/shortscale-go
2569018	       472 ns/op	     120 B/op	       4 allocs/op
```
