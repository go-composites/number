package Number

import (
	"fmt"
	"math"
	"strconv"

	Boolean "github.com/go-composites/boolean/src"
	Error "github.com/go-composites/error/src"
	Result "github.com/go-composites/result/src"
)

/*
Number is a numeric composite over a Go int64 or float64 value.

Its fallible operations (notably Div) return a Result.Interface so that
failures — such as a division by zero — are values rather than panics.
*/
type Interface interface {
	ToGoInt() int64
	ToGoFloat() float64
	ToGoString() string
	IsInt() bool
	IsFloat() bool
	IsNull() bool
	Add(Interface) Result.Interface
	Sub(Interface) Result.Interface
	Mul(Interface) Result.Interface
	Div(Interface) Result.Interface
	Mod(Interface) Result.Interface
	Abs() Result.Interface
	Neg() Result.Interface
	Equal(Interface) Boolean.Interface
	LessThan(Interface) Boolean.Interface
	GreaterThan(Interface) Boolean.Interface
	Inspect() String
}

// String is the lightweight inspection representation of a Number.
type String = string

type data struct {
	value float64
	isInt bool
}

type Option func(*data)

/*
WithInt is a functional parameter setting the Number value from a Go int64.

The resulting Number reports IsInt() == true.
*/
func WithInt(value int64) Option {
	return func(d *data) {
		d.value = float64(value)
		d.isInt = true
	}
}

/*
WithFloat is a functional parameter setting the Number value from a Go float64.

The resulting Number reports IsFloat() == true.
*/
func WithFloat(value float64) Option {
	return func(d *data) {
		d.value = value
		d.isInt = false
	}
}

/*
New is the Number constructor.

Called without a functional parameter it yields the integer zero.

	n1 := Number.New()                       // 0   (int)
	n2 := Number.New(Number.WithInt(42))     // 42  (int)
	n3 := Number.New(Number.WithFloat(3.5))  // 3.5 (float)
*/
func New(options ...Option) Interface {
	d := &data{
		value: 0,
		isInt: true,
	}
	for _, opt := range options {
		opt(d)
	}
	return d
}

/*
ToGoInt returns the value truncated to a Go int64.
*/
func (d data) ToGoInt() int64 {
	return int64(d.value)
}

/*
ToGoFloat returns the value as a Go float64.
*/
func (d data) ToGoFloat() float64 {
	return d.value
}

/*
ToGoString returns the textual representation of the value.

Integers are rendered without a decimal point; floats use the shortest
representation that round-trips.
*/
func (d data) ToGoString() string {
	if d.isInt {
		return strconv.FormatInt(int64(d.value), 10)
	}
	return strconv.FormatFloat(d.value, 'g', -1, 64)
}

/*
IsInt reports whether the Number carries an integer value.
*/
func (d data) IsInt() bool {
	return d.isInt
}

/*
IsFloat reports whether the Number carries a floating-point value.
*/
func (d data) IsFloat() bool {
	return !d.isInt
}

/*
IsNull reports whether the Number is the Null-Object variant.

A concrete Number is never null.
*/
func (d data) IsNull() bool {
	return false
}

// build constructs a Number preserving integer-ness only when both operands
// are integers (mirroring Go's numeric tower for these operations).
func build(value float64, lhsInt bool, rhs Interface) Interface {
	if lhsInt && rhs.IsInt() {
		return &data{value: value, isInt: true}
	}
	return &data{value: value, isInt: false}
}

/*
Add returns a Result whose payload is the sum of the receiver and other.
*/
func (d data) Add(other Interface) Result.Interface {
	return Result.New(
		Result.WithPayload(
			build(d.value+other.ToGoFloat(), d.isInt, other),
		),
	)
}

/*
Sub returns a Result whose payload is the difference of the receiver and other.
*/
func (d data) Sub(other Interface) Result.Interface {
	return Result.New(
		Result.WithPayload(
			build(d.value-other.ToGoFloat(), d.isInt, other),
		),
	)
}

/*
Mul returns a Result whose payload is the product of the receiver and other.
*/
func (d data) Mul(other Interface) Result.Interface {
	return Result.New(
		Result.WithPayload(
			build(d.value*other.ToGoFloat(), d.isInt, other),
		),
	)
}

/*
Div returns a Result whose payload is the quotient of the receiver and other.

When other is zero the Result carries an Error ("division by zero") instead of
a payload — the division never panics and never returns nil.
*/
func (d data) Div(other Interface) Result.Interface {
	if other.ToGoFloat() == 0 {
		return Result.New(
			Result.WithError(
				Error.New("division by zero"),
			),
		)
	}
	return Result.New(
		Result.WithPayload(
			build(d.value/other.ToGoFloat(), d.isInt, other),
		),
	)
}

/*
Mod returns a Result whose payload is the remainder of the receiver divided by
other.

Integer operands use Go's % operator; if either operand is a float the
remainder is computed with math.Mod. When other is zero the Result carries an
Error ("modulo by zero") instead of a payload — the operation never panics and
never returns nil.
*/
func (d data) Mod(other Interface) Result.Interface {
	if other.ToGoFloat() == 0 {
		return Result.New(
			Result.WithError(
				Error.New("modulo by zero"),
			),
		)
	}
	if d.isInt && other.IsInt() {
		return Result.New(
			Result.WithPayload(
				build(float64(d.ToGoInt()%other.ToGoInt()), d.isInt, other),
			),
		)
	}
	return Result.New(
		Result.WithPayload(
			build(math.Mod(d.value, other.ToGoFloat()), d.isInt, other),
		),
	)
}

/*
Abs returns a Result whose payload is the absolute value of the receiver,
preserving its integer or floating-point kind.
*/
func (d data) Abs() Result.Interface {
	return Result.New(
		Result.WithPayload(
			&data{value: math.Abs(d.value), isInt: d.isInt},
		),
	)
}

/*
Neg returns a Result whose payload is the negation of the receiver, preserving
its integer or floating-point kind.
*/
func (d data) Neg() Result.Interface {
	return Result.New(
		Result.WithPayload(
			&data{value: -d.value, isInt: d.isInt},
		),
	)
}

/*
Equal reports, as a Boolean.Interface, whether the receiver and other hold the
same numeric value.
*/
func (d data) Equal(other Interface) Boolean.Interface {
	return Boolean.New(d.value == other.ToGoFloat())
}

/*
LessThan reports, as a Boolean.Interface, whether the receiver is strictly less
than other.
*/
func (d data) LessThan(other Interface) Boolean.Interface {
	return Boolean.New(d.value < other.ToGoFloat())
}

/*
GreaterThan reports, as a Boolean.Interface, whether the receiver is strictly
greater than other.
*/
func (d data) GreaterThan(other Interface) Boolean.Interface {
	return Boolean.New(d.value > other.ToGoFloat())
}

/*
Inspect returns a one-line representation of the Number with its address, kind
and value — mirroring the style of the Boolean composite.
*/
func (d data) Inspect() String {
	kind := "float"
	if d.isInt {
		kind = "int"
	}
	return fmt.Sprintf(
		"<Number:%p kind=%s value=%s>",
		&d, kind, d.ToGoString(),
	)
}
