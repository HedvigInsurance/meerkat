package mappers

import (
	"encoding/xml"
	"fmt"
	"log"
	"time"

	"github.com/HedvigInsurance/meerkat/utils"
)

type IndividualRoot struct {
	IndividualRoot Individuals `xml:"INDIVIDUALS"`
}

type Individuals struct {
	IndividualChilds []Individual `xml:"INDIVIDUAL"`
}

type Individual struct {
	FirstName       string            `xml:"FIRST_NAME"`
	SecondName      string            `xml:"SECOND_NAME"`
	ThirdName       string            `xml:"THIRD_NAME"`
	FourthName      string            `xml:"FOURTH_NAME"`
	IndividualAlias []IndividualAlias `xml:"INDIVIDUAL_ALIAS"`
}

type IndividualAlias struct {
	AliasName string `xml:"ALIAS_NAME"`
}

func MapUnSanctionList() (unSanctionList IndividualRoot) {
	startUN := time.Now()
	if xmlStr, err := utils.FetchXmlFromUrl("https://scsanctions.un.org/resources/xml/en/consolidated.xml"); err != nil {
		fmt.Printf("Failed to get XML: %v", err)
	} else {
		xml.Unmarshal(xmlStr, &unSanctionList)
	}
	log.Println("UN -- Downloading  took ", time.Since(startUN))
	return unSanctionList
}
