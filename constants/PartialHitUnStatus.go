package constants

type PartialHitUnStatus uint16

const (
	PartialHitUnNone           PartialHitUnStatus = 0
	PartialHitUnFirstName      PartialHitUnStatus = 1
	PartialHitUnSecondName     PartialHitUnStatus = 2
	PartialHitUnThirdName      PartialHitUnStatus = 3
	PartialHitUnFourthName     PartialHitUnStatus = 4
	PartialHitUnWholeName      PartialHitUnStatus = 5 // First + Second + Third + Fourth
	PartialHitUnWholeNameAlias PartialHitUnStatus = 6 // only alias
)

func (status PartialHitUnStatus) ToString() string {
	switch status {
	case PartialHitUnNone:
		return "None"
	case PartialHitUnFirstName:
		return "FirstName"
	case PartialHitUnSecondName:
		return "SecondName"
	case PartialHitUnThirdName:
		return "ThirdName"
	case PartialHitUnFourthName:
		return "FourthName"
	case PartialHitUnWholeName:
		return "WholeName"
	case PartialHitUnWholeNameAlias:
		return "WholeNameAlias"
	}
	return ""
}
