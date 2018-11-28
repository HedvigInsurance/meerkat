package queries

import (
	"strings"

	"github.com/HedvigInsurance/meerkat/constants"
	"github.com/HedvigInsurance/meerkat/mappers"
)

func QueryEUsanctionList(query []string, euList mappers.SanctionEntites) (result constants.SanctionResult) {
	var partialHiited bool = false
	for i := 0; i < len(euList.Entites); i++ {
		for index := 0; index < len(euList.Entites[i].NameAlias); index++ {
			var hit constants.PartialHitEuStatus = constants.PartialHitEuNone
			for j := 0; j < len(query); j++ {
				if strings.ToLower(strings.TrimSpace(strings.Join(query, " "))) == strings.ToLower(strings.TrimSpace(euList.Entites[i].NameAlias[index].WholeName)) {
					return constants.FullHit
				} else {
					if strings.ToLower(query[j]) == strings.ToLower(strings.TrimSpace(euList.Entites[i].NameAlias[index].FirstName)) {
						hit = constants.PartialHitEuFirstName
					}
					if strings.ToLower(query[j]) == strings.ToLower(strings.TrimSpace(euList.Entites[i].NameAlias[index].LastName)) {
						if hit == constants.PartialHitEuFirstName {
							hit = constants.PartialHitEuWholeNameNoMiddleName
						} else {
							hit = constants.PartialHitEuLastName
						}
					}
					if strings.ToLower(query[j]) == strings.ToLower(strings.TrimSpace(euList.Entites[i].NameAlias[index].MiddleName)) {
						if hit == constants.PartialHitEuWholeNameNoMiddleName {
							hit = constants.PartialHitEuWholeName
						} else {
							hit = constants.PartialHitEuMiddleName
						}
					}
				}
			}
			if hit == constants.PartialHitEuWholeName || hit == constants.PartialHitEuWholeNameNoMiddleName {
				return constants.FullHit
			} else if hit > 0 { //PartialHitEuStatus is not EuNone
				partialHiited = true
			}
		}
	}
	if partialHiited {
		return constants.PartialHit
	} else {
		return constants.NoHit
	}
}
