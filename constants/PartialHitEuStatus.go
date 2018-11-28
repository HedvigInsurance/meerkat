package constants

type PartialHitEuStatus uint16

const (
	PartialHitEuNone                  PartialHitEuStatus = 0
	PartialHitEuFirstName             PartialHitEuStatus = 1
	PartialHitEuMiddleName            PartialHitEuStatus = 2
	PartialHitEuLastName              PartialHitEuStatus = 3
	PartialHitEuWholeNameNoMiddleName PartialHitEuStatus = 4 // Only First and Last name matched
	PartialHitEuWholeName             PartialHitEuStatus = 5 // First, Last and Middle name matched
)

func (status PartialHitEuStatus) ToString() string {
	switch status {
	case PartialHitEuNone:
		return "EuNone"
	case PartialHitEuFirstName:
		return "EuFirstName"
	case PartialHitEuMiddleName:
		return "MiddleName"
	case PartialHitEuLastName:
		return "LastName"
	case PartialHitEuWholeNameNoMiddleName:
		return "WholeNameNoMiddleName"
	case PartialHitEuWholeName:
		return "WholeName"
	}
	return ""
}
