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

})
