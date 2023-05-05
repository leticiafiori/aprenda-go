package kata_test

var _ = Describe("Test Example", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(BoolToWord(true)).To(Equal("Yes"))
		Expect(BoolToWord(false)).To(Equal("No"))
	})
})
