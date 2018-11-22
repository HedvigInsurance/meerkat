package queries

import (
	"log"
	"strings"
	"time"

	"github.com/HedvigInsurance/meerkat/constants"
	"github.com/HedvigInsurance/meerkat/mappers"
)

func QueryEUsanctionList(query []string, euList mappers.SanctionEntites) (result constants.SanctionResult) {

	start_eu_sanct := time.Now()

	var hitted bool = false
	for i := 0; i < len(euList.Entites); i++ {
		for index := 0; index < len(euList.Entites[i].NameAlias); index++ {
			var hit uint = 0
			for j := 0; j < len(query); j++ {
				if strings.ToLower(query[j]) == strings.ToLower(strings.TrimSpace(euList.Entites[i].NameAlias[index].FirstName)) {
					hit++
				}
				if strings.ToLower(query[j]) == strings.ToLower(strings.TrimSpace(euList.Entites[i].NameAlias[index].LastName)) {
					hit++
				}
				if strings.ToLower(query[j]) == strings.ToLower(strings.TrimSpace(euList.Entites[i].NameAlias[index].MiddleName)) {
					hit++
				}
				// if strings.ToLower(strings.TrimSpace(strings.Join(query))) == strings.ToLower(strings.TrimSpace(euList.Entites[i].NameAlias[index].WholeName)) {
				// 	log.Println("\nComparison Whole : ", strings.ToLower(query[j]), strings.ToLower(strings.TrimSpace(euList.Entites[i].NameAlias[index].WholeName)))
				// 	hit++
				// }
			}
			if hit > 1 {
				log.Println("Sanctionlist took ", time.Since(start_eu_sanct))
				return constants.FullHit
			} else if hit == 1 {
				hitted = true
			}
		}
	}
	if hitted {
		return constants.PartialHit
		// w.WriteHeader(200)
		// w.Write([]byte("PARTIAL hit"))
	} else {
		// 	w.WriteHeader(201)
		// 	w.Write([]byte("NO HIT"))
		return constants.NoHit
	}
}
