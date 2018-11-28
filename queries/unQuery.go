package queries

import (
	"strings"

	"github.com/HedvigInsurance/meerkat/constants"
	"github.com/HedvigInsurance/meerkat/mappers"
	"github.com/HedvigInsurance/meerkat/utils"
)

func QueryUNsanctionList(query []string, unList mappers.IndividualRoot) (result constants.SanctionResult) {
	for i := 0; i < len(unList.IndividualRoot.IndividualChilds); i++ {
		var individual mappers.Individual = unList.IndividualRoot.IndividualChilds[i]
		name := individual.FirstName + individual.SecondName + individual.ThirdName + individual.FourthName
		if strings.ToLower(utils.TrimCharacters(strings.Join(query, ""))) == strings.ToLower(utils.TrimCharacters(name)) {
			return constants.FullHit
		}

		for j := 0; j < len(individual.IndividualAlias); j++ {
			if strings.ToLower(utils.TrimCharacters(strings.Join(query, " "))) == strings.ToLower(utils.TrimCharacters(individual.IndividualAlias[j].AliasName)) {
				return constants.FullHit
			}
		}

	}
	return constants.NoHit
}
