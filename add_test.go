package add_test

import (
	. "add"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Add", func() {

	Context("When passed an empty string", func() {

		It("should return 0", func() {
			Expect(Add("")).To(Equal(0))
		})

	})

	Context("When passed 69 as a string", func() {

		It("should return 69 as a number", func() {
			Expect(Add("69")).To(Equal(69))
		})

	})

	Context("When passed 70 as a string", func() {

		It("should return 70 as a number", func() {
			Expect(Add("70")).To(Equal(70))
		})

	})

})
