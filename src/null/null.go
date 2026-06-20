package NullNumber

import (
	Boolean "github.com/go-composites/boolean/src"
	MethodNotImplementedError "github.com/go-composites/error/src/method_not_implemented"
	Number "github.com/go-composites/number/src"
	Result "github.com/go-composites/result/src"
)

/*
NullNumber is the Null-Object variant of Number.

It satisfies Number.Interface so callers never have to test for a bare nil:
its value is zero, its arithmetic yields a Result carrying a
"method not implemented" Error, and IsNull() returns true.
*/
type Interface interface {
	Number.Interface
}

type data struct{}

/*
New returns a NullNumber.
*/
func New() Interface {
	return &data{}
}

func (d data) ToGoInt() int64 {
	return 0
}

func (d data) ToGoFloat() float64 {
	return 0
}

func (d data) ToGoString() string {
	return ``
}

func (d data) IsInt() bool {
	return false
}

func (d data) IsFloat() bool {
	return false
}

func (d data) IsNull() bool {
	return true
}

func notImplemented(methodName string) Result.Interface {
	return Result.New(
		Result.WithError(
			MethodNotImplementedError.New(methodName),
		),
	)
}

func (d data) Add(Number.Interface) Result.Interface {
	return notImplemented(`Add`)
}

func (d data) Sub(Number.Interface) Result.Interface {
	return notImplemented(`Sub`)
}

func (d data) Mul(Number.Interface) Result.Interface {
	return notImplemented(`Mul`)
}

func (d data) Div(Number.Interface) Result.Interface {
	return notImplemented(`Div`)
}

func (d data) Mod(Number.Interface) Result.Interface {
	return notImplemented(`Mod`)
}

func (d data) Abs() Result.Interface {
	return notImplemented(`Abs`)
}

func (d data) Neg() Result.Interface {
	return notImplemented(`Neg`)
}

func (d data) Equal(other Number.Interface) Boolean.Interface {
	return Boolean.New(other.IsNull())
}

func (d data) LessThan(Number.Interface) Boolean.Interface {
	return Boolean.False()
}

func (d data) GreaterThan(Number.Interface) Boolean.Interface {
	return Boolean.False()
}

func (d data) Inspect() Number.String {
	return `<NullNumber>`
}
