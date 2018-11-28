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
	var unList mappers.IndividualRoot

	BeforeSuite(func() {
		euList = mappers.MapEuSanctionList()
		unList = mappers.MapUnSanctionList()
	})

	Describe("Testing Querying function", func() {
		Context("query a non-terrosit", func() {
			It("should be a No-HIT", func() {

				var queryNoTerrorist []string = []string{"Meletis", "Mazarakiotis"}

				Expect(queries.QueryEUsanctionList(queryNoTerrorist, euList)).To(Equal(constants.NoHit))
			})
		})
		Context("query a terrosit", func() {
			It("should be a Full-HIT with wholeNameNoMiddleName", func() {

				var queryTerroristNoWholeName []string = []string{"Ricardo", "Ayeras"}

				Expect(queries.QueryEUsanctionList(queryTerroristNoWholeName, euList)).To(Equal(constants.FullHit))
			})
		})
		Context("query a non - terrosit with suspicious first name", func() {
			It("should be a Partial-HIT", func() {

				var queryPartialHit []string = []string{"Ricardo", "Mazarakiotis"}

				Expect(queries.QueryEUsanctionList(queryPartialHit, euList)).To(Equal(constants.PartialHit))
			})
		})
		Context("query a confirmed terrosit", func() {
			It("should be a Full-HIT with wholename match", func() {

				var queryWholeName []string = []string{"Robert", "Gabriel", "Mugabe"}

				Expect(queries.QueryEUsanctionList(queryWholeName, euList)).To(Equal(constants.FullHit))
			})
		})
		Context("query a confirmed terrosit from un", func() {
			It("should be a Full-HIT from UN", func() {

				var queryWholeName []string = []string{"Abdul Aziz", "Abbasin"}

				Expect(queries.QueryUNsanctionList(queryWholeName, unList)).To(Equal(constants.FullHit))
			})
		})
	})

})
