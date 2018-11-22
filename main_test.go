package main_test

import (
	"github.com/HedvigInsurance/meerkat/constants"
	"github.com/HedvigInsurance/meerkat/mappers"
	"github.com/HedvigInsurance/meerkat/queries"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {

	var euList mappers.SanctionEntites

	BeforeSuite(func() {
		euList = mappers.MapEuSanctionList()
	})

	Describe("Testing Querying function", func() {
		Context("query a non-terrosit", func() {
			It("should be a No-HIT", func() {

				query := []string{"Meletis", "Mazarakiotis"}

				Expect(queries.QueryEUsanctionList(query, euList)).To(Equal(constants.NoHit))
			})
		})
		Context("query a terrosit", func() {
			It("should be a Parcial-HIT", func() {

				query := []string{"Ricardo", "Ayeras"}

				Expect(queries.QueryEUsanctionList(query, euList)).To(Equal(constants.FullHit))
			})
		})
		Context("query a terrosit", func() {
			It("should be a Full-HIT", func() {

				query := []string{"Ricardo", "Mazarakiotis"}

				Expect(queries.QueryEUsanctionList(query, euList)).To(Equal(constants.PartialHit))
			})
		})
	})

})
