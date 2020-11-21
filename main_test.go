package main_test

import (
	"github.com/HedvigInsurance/meerkat/constants"
	"github.com/HedvigInsurance/meerkat/mappers"
	"github.com/HedvigInsurance/meerkat/queries"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {

	var euList mappers.SanctionEntities
	var unList mappers.IndividualRoot

	BeforeSuite(func() {
		euList = mappers.MapEuSanctionList()
		unList = mappers.MapUnSanctionList()
	})

	Describe("Testing Querying function", func() {
		Context("query a non-sanctioned person", func() {
			It("should be a No-HIT", func() {

				var queryNoTerrorist []string = []string{"Meletis", "Mazarakiotis"}

				Expect(queries.QueryEuSanctionList(queryNoTerrorist, euList)).To(Equal(constants.NoHit))
			})
		})
		Context("query a sanctioned person", func() {
			It("should be a Full-HIT with wholeNameNoMiddleName", func() {

				var queryTerroristNoWholeName []string = []string{"Ricardo", "Ayeras"}

				Expect(queries.QueryEuSanctionList(queryTerroristNoWholeName, euList)).To(Equal(constants.FullHit))
			})
		})
		Context("query a non - sanctioned with suspicious first name", func() {
			It("should be a Partial-HIT", func() {

				var queryPartialHit []string = []string{"Ricardo", "Mazarakiotis"}

				Expect(queries.QueryEuSanctionList(queryPartialHit, euList)).To(Equal(constants.PartialHit))
			})
		})
		Context("query a confirmed sanctioned", func() {
			It("should be a Full-HIT with wholename match", func() {

				var queryWholeName []string = []string{"Abdul Aziz", "Abbasin"}

				Expect(queries.QueryEuSanctionList(queryWholeName, euList)).To(Equal(constants.FullHit))
			})
		})
		Context("query a confirmed sanctioned from un", func() {
			It("should be a Full-HIT from UN", func() {

				var queryWholeName []string = []string{"Abdul Aziz", "Abbasin"}

				Expect(queries.QueryUnSanctionList(queryWholeName, unList)).To(Equal(constants.FullHit))
			})
		})
	})

})
