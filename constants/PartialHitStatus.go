package constants

type PartialHitStatus uint16

const (
	None                  PartialHitStatus = 0
	FirstName             PartialHitStatus = 1
	MiddleName            PartialHitStatus = 2
	LastName              PartialHitStatus = 3
	WholeNameNoMiddleName PartialHitStatus = 4 // Only First and Last name matched
	WholeName             PartialHitStatus = 5 // First, Last and Middle name matched
)

func (status PartialHitStatus) ToString() string {
	switch status {
	case None:
		return "None"
	case FirstName:
		return "FirstName"
	case MiddleName:
		return "MiddleName"
	case LastName:
		return "LastName"
	case WholeNameNoMiddleName:
		return "WholeNameNoMiddleName"
	case WholeName:
		return "WholeName"
	}
	return ""
}
