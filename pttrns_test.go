package cypherigo_test

import (
	. "github.com/fdsolutions/cypherigo"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pttrns", func() {
	Describe("Node", func() {
		Context("Node", func() {
			It("returns: ()", func() {
				n := Node
				Expect(n.ToCypher()).To(Equal(`()`))
			})
		})
		Context("Node.Ident(`a`)", func() {
			It("returns: (a)", func() {
				n := Node.Ident("a")
				Expect(n.ToCypher()).To(Equal(`(a)`))
			})
		})
		Context("Node.Ident(`a`).WithLabels(`Person`, `Swedish`)", func() {
			It("returns: (a:Person:Swedish)", func() {
				n := Node.Ident("a").WithLabels("Person", "Swedish")
				Expect(n.ToCypher()).To(Equal(`(a:Person:Swedish)`))
			})
		})
	})
})
