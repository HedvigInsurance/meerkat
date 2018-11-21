package mappers

import (
	"encoding/xml"
	"fmt"
	"log"
	"meerkat/utils"
	"time"
)

type SanctionEntites struct {
	Entites []SanctionEntity `xml:"sanctionEntity"`
}

type SanctionEntity struct {
	NameAlias []NameAlias `xml:"nameAlias"`
}

type NameAlias struct {
	FirstName  string `xml:"firstName,attr"`
	LastName   string `xml:"lastName,attr"`
	MiddleName string `xml:"middleName,attr"`
	WholeName  string `xml:"wholeName,attr"`
}

func MapEuSanctionList() (euSanctionList SanctionEntites) {
	startEU := time.Now()
	if xmlStr, err := utils.FetchXmlFromUrl("https://webgate.ec.europa.eu/europeaid/fsd/fsf/public/files/xmlFullSanctionsList/content?token=dG9rZW4tMjAxNw"); err != nil {
		fmt.Printf("Failed to get XML: %v", err)
	} else {
		// fmt.Println("Received XML")
		xml.Unmarshal(xmlStr, &euSanctionList)
	}
	log.Println("EU -- Downloading took ", time.Since(startEU))
	return euSanctionList
}
