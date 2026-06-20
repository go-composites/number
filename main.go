package main

import (
	"fmt"

	Error "github.com/go-composites/error/src"
	Number "github.com/go-composites/number/src"
	Result "github.com/go-composites/result/src"
)

func report(label string, result Result.Interface) {
	if result.HasError() {
		fmt.Printf("%s -> error: %s\n", label, result.Error().Message())
		return
	}
	fmt.Printf("%s -> %s\n", label, result.Payload().(Number.Interface).ToGoString())
}

func main() {
	six := Number.New(Number.WithInt(6))
	two := Number.New(Number.WithInt(2))
	zero := Number.New()

	report("6 + 2", six.Add(two))
	report("6 - 2", six.Sub(two))
	report("6 * 2", six.Mul(two))
	report("6 / 2", six.Div(two))

	// The canonical Result use-case: division by zero is a value, not a panic.
	divByZero := six.Div(zero)
	fmt.Println("6 / 0 has error:", divByZero.HasError())
	report("6 / 0", divByZero)

	// Errors are first-class values.
	var _ Error.Interface = divByZero.Error()

	fmt.Println("6 == 2 :", six.Equal(two).ToGoString())
	fmt.Println("6 < 2  :", six.LessThan(two).ToGoString())
	fmt.Println("6 > 2  :", six.GreaterThan(two).ToGoString())
	fmt.Println(six.Inspect())
}
