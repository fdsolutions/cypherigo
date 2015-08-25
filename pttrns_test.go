package cypherigo_test

import (
	. "github.com/fdsolutions/cypherigo"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pttrns", func() {
	Describe("Node", func() {
		Context("Node(a)", func() {
			It("returns (a)", func() {
				n := Node(`a`)
				Expect(n.ToCypher()).To(Equal(`(a)`))
			})
		})
	})
})
