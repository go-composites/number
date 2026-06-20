package NullNumber_test

import (
	Number "github.com/go-composites/number/src"
	NullNumber "github.com/go-composites/number/src/null"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("NullNumber", func() {
	var n NullNumber.Interface
	ginkgo.BeforeEach(func() {
		n = NullNumber.New()
	})

	ginkgo.It("satisfies the Number interface", func() {
		var _ Number.Interface = n
	})
	ginkgo.It("reports IsNull() true", func() {
		gomega.Expect(n.IsNull()).To(gomega.BeTrue())
	})
	ginkgo.It("is neither int nor float", func() {
		gomega.Expect(n.IsInt()).To(gomega.BeFalse())
		gomega.Expect(n.IsFloat()).To(gomega.BeFalse())
	})
	ginkgo.It("converts to zero values", func() {
		gomega.Expect(n.ToGoInt()).To(gomega.BeEquivalentTo(0))
		gomega.Expect(n.ToGoFloat()).To(gomega.BeEquivalentTo(0.0))
		gomega.Expect(n.ToGoString()).To(gomega.Equal(``))
	})

	ginkgo.It("Add returns an error result", func() {
		gomega.Expect(n.Add(Number.New()).HasError()).To(gomega.BeTrue())
	})
	ginkgo.It("Sub returns an error result", func() {
		gomega.Expect(n.Sub(Number.New()).HasError()).To(gomega.BeTrue())
	})
	ginkgo.It("Mul returns an error result", func() {
		gomega.Expect(n.Mul(Number.New()).HasError()).To(gomega.BeTrue())
	})
	ginkgo.It("Div returns an error result", func() {
		r := n.Div(Number.New())
		gomega.Expect(r.HasError()).To(gomega.BeTrue())
		gomega.Expect(r.Error().Message()).To(gomega.ContainSubstring("Div"))
	})
	ginkgo.It("Mod returns an error result", func() {
		r := n.Mod(Number.New())
		gomega.Expect(r.HasError()).To(gomega.BeTrue())
		gomega.Expect(r.Error().Message()).To(gomega.ContainSubstring("Mod"))
	})
	ginkgo.It("Abs returns an error result", func() {
		r := n.Abs()
		gomega.Expect(r.HasError()).To(gomega.BeTrue())
		gomega.Expect(r.Error().Message()).To(gomega.ContainSubstring("Abs"))
	})
	ginkgo.It("Neg returns an error result", func() {
		r := n.Neg()
		gomega.Expect(r.HasError()).To(gomega.BeTrue())
		gomega.Expect(r.Error().Message()).To(gomega.ContainSubstring("Neg"))
	})
	ginkgo.It("Equal is true only against another null", func() {
		gomega.Expect(n.Equal(NullNumber.New()).ToGoBool()).To(gomega.BeTrue())
		gomega.Expect(n.Equal(Number.New()).ToGoBool()).To(gomega.BeFalse())
	})
	ginkgo.It("LessThan is always false", func() {
		gomega.Expect(n.LessThan(Number.New()).ToGoBool()).To(gomega.BeFalse())
	})
	ginkgo.It("GreaterThan is always false", func() {
		gomega.Expect(n.GreaterThan(Number.New()).ToGoBool()).To(gomega.BeFalse())
	})
	ginkgo.It("Inspect renders the null marker", func() {
		gomega.Expect(n.Inspect()).To(gomega.Equal(`<NullNumber>`))
	})
})
