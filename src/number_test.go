package Number_test

import (
	Number "github.com/go-composites/number/src"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Number", func() {

	ginkgo.Describe("constructors", func() {
		ginkgo.It("defaults to integer zero", func() {
			n := Number.New()
			gomega.Expect(n.ToGoInt()).To(gomega.BeEquivalentTo(0))
			gomega.Expect(n.IsInt()).To(gomega.BeTrue())
			gomega.Expect(n.IsFloat()).To(gomega.BeFalse())
		})
		ginkgo.It("can be built from a Go int64", func() {
			n := Number.New(Number.WithInt(42))
			gomega.Expect(n.ToGoInt()).To(gomega.BeEquivalentTo(42))
			gomega.Expect(n.IsInt()).To(gomega.BeTrue())
		})
		ginkgo.It("can be built from a Go float64", func() {
			n := Number.New(Number.WithFloat(3.5))
			gomega.Expect(n.ToGoFloat()).To(gomega.BeEquivalentTo(3.5))
			gomega.Expect(n.IsFloat()).To(gomega.BeTrue())
			gomega.Expect(n.IsInt()).To(gomega.BeFalse())
		})
		ginkgo.It("is never null", func() {
			gomega.Expect(Number.New().IsNull()).To(gomega.BeFalse())
		})
	})

	ginkgo.Describe("conversions", func() {
		ginkgo.It("renders an integer without a decimal point", func() {
			gomega.Expect(Number.New(Number.WithInt(7)).ToGoString()).To(gomega.Equal("7"))
		})
		ginkgo.It("renders a float with its shortest representation", func() {
			gomega.Expect(Number.New(Number.WithFloat(3.5)).ToGoString()).To(gomega.Equal("3.5"))
		})
		ginkgo.It("truncates a float when converting to a Go int", func() {
			gomega.Expect(Number.New(Number.WithFloat(3.9)).ToGoInt()).To(gomega.BeEquivalentTo(3))
		})
		ginkgo.It("widens an int when converting to a Go float", func() {
			gomega.Expect(Number.New(Number.WithInt(4)).ToGoFloat()).To(gomega.BeEquivalentTo(4.0))
		})
	})

	ginkgo.Describe("arithmetic", func() {
		var six = Number.New(Number.WithInt(6))
		var two = Number.New(Number.WithInt(2))
		var half = Number.New(Number.WithFloat(0.5))

		ginkgo.It("adds two numbers", func() {
			r := six.Add(two)
			gomega.Expect(r.HasError()).To(gomega.BeFalse())
			gomega.Expect(r.Payload().(Number.Interface).ToGoInt()).To(gomega.BeEquivalentTo(8))
		})
		ginkgo.It("subtracts two numbers", func() {
			r := six.Sub(two)
			gomega.Expect(r.Payload().(Number.Interface).ToGoInt()).To(gomega.BeEquivalentTo(4))
		})
		ginkgo.It("multiplies two numbers", func() {
			r := six.Mul(two)
			gomega.Expect(r.Payload().(Number.Interface).ToGoInt()).To(gomega.BeEquivalentTo(12))
		})
		ginkgo.It("divides two numbers", func() {
			r := six.Div(two)
			gomega.Expect(r.HasError()).To(gomega.BeFalse())
			gomega.Expect(r.Payload().(Number.Interface).ToGoInt()).To(gomega.BeEquivalentTo(3))
		})
		ginkgo.It("keeps integer-ness only when both operands are integers", func() {
			r := six.Add(half)
			gomega.Expect(r.Payload().(Number.Interface).IsFloat()).To(gomega.BeTrue())
			gomega.Expect(r.Payload().(Number.Interface).ToGoFloat()).To(gomega.BeEquivalentTo(6.5))
		})

		ginkgo.Describe("division by zero", func() {
			ginkgo.It("returns a Result carrying an error instead of panicking", func() {
				r := six.Div(Number.New())
				gomega.Expect(r.HasError()).To(gomega.BeTrue())
				gomega.Expect(r.Error().Message()).To(gomega.Equal("division by zero"))
			})
			ginkgo.It("also guards a float zero divisor", func() {
				r := six.Div(Number.New(Number.WithFloat(0.0)))
				gomega.Expect(r.HasError()).To(gomega.BeTrue())
			})
		})
	})

	ginkgo.Describe("comparisons", func() {
		var six = Number.New(Number.WithInt(6))
		var two = Number.New(Number.WithInt(2))

		ginkgo.It("reports equality", func() {
			gomega.Expect(six.Equal(six).ToGoBool()).To(gomega.BeTrue())
			gomega.Expect(six.Equal(two).ToGoBool()).To(gomega.BeFalse())
		})
		ginkgo.It("reports less-than", func() {
			gomega.Expect(two.LessThan(six).ToGoBool()).To(gomega.BeTrue())
			gomega.Expect(six.LessThan(two).ToGoBool()).To(gomega.BeFalse())
		})
		ginkgo.It("reports greater-than", func() {
			gomega.Expect(six.GreaterThan(two).ToGoBool()).To(gomega.BeTrue())
			gomega.Expect(two.GreaterThan(six).ToGoBool()).To(gomega.BeFalse())
		})
	})

	ginkgo.Describe("inspection", func() {
		ginkgo.It("renders an int Number", func() {
			gomega.Expect(Number.New(Number.WithInt(6)).Inspect()).To(gomega.ContainSubstring("kind=int value=6"))
		})
		ginkgo.It("renders a float Number", func() {
			gomega.Expect(Number.New(Number.WithFloat(1.5)).Inspect()).To(gomega.ContainSubstring("kind=float value=1.5"))
		})
	})
})
