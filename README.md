<p align="center"><img src="https://raw.githubusercontent.com/go-composites/brand/main/social/golang-oop.png" alt="go-composites/number" width="720"></p>

# Number

A small numeric composite for **Composition-Oriented Programming**. A `Number`
wraps a Go `int64` or `float64` value and exposes its arithmetic as **fallible
operations that return a `Result`** — so failures (the canonical example being a
division by zero) are *values*, never panics and never `nil`.

This is the canonical `Result` use-case in the `go-composites` family:

```golang
quotient := numerator.Div(denominator)
if quotient.HasError() {
    fmt.Println(quotient.Error().Message()) // "division by zero"
} else {
    fmt.Println(quotient.Payload().(Number.Interface).ToGoString())
}
```

`Number` follows the org's Null-Object / never-nil invariant (enforced by the
`nonnil` CI analyzer): the `NullNumber` variant in `src/null` satisfies the same
`Interface` and reports `IsNull() == true`.

## Install

```bash
export GOPRIVATE=github.com/go-composites GOPROXY=direct GOSUMDB=off
go get github.com/go-composites/number@main
```

## Usage

> [!NOTE] main.go

```golang
package main

import (
    "fmt"

    Number "github.com/go-composites/number/src"
)

func main() {
    six := Number.New(Number.WithInt(6))
    two := Number.New(Number.WithInt(2))
    zero := Number.New() // default zero (int)

    // Arithmetic returns a Result.
    sum := six.Add(two)
    fmt.Println(sum.Payload().(Number.Interface).ToGoString()) // 8

    // Division by zero is a value, not a panic.
    div := six.Div(zero)
    fmt.Println("has error:", div.HasError())      // true
    fmt.Println(div.Error().Message())             // division by zero

    // Comparisons return a go-composites/boolean.
    fmt.Println(six.GreaterThan(two).ToGoString()) // "true"
    fmt.Println(six.Inspect())                     // <Number:0x... kind=int value=6>
}
```

```bash
$ task build
```

## API

Constructors

- `New(opts ...Option) Interface` — default integer zero.
- `WithInt(int64) Option` / `WithFloat(float64) Option`.
- `null.New() Interface` — the `NullNumber` Null-Object (`IsNull() == true`).

Conversions

- `ToGoInt() int64`, `ToGoFloat() float64`, `ToGoString() string`.
- `IsInt() bool`, `IsFloat() bool`, `IsNull() bool`.

Arithmetic (each returns `Result.Interface`)

- `Add(other) Result` / `Sub(other) Result` / `Mul(other) Result`.
- `Div(other) Result` — a `Result` carrying `Error.New("division by zero")`
  when `other` is zero.

Comparisons (each returns `Boolean.Interface`)

- `Equal(other)` / `LessThan(other)` / `GreaterThan(other)`.

Inspection

- `Inspect() string` — `<Number:0x... kind=int value=...>`.

## License

BSD-3-Clause — see [LICENSE](./LICENSE).
