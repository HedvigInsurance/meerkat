package constants

type SanctionResult uint16

const (
	NoHit      SanctionResult = 0
	PartialHit SanctionResult = 1
	FullHit    SanctionResult = 2
)

func (s SanctionResult) ToString() string {
	switch s {
	case NoHit:
		return "NoHit"
	case PartialHit:
		return "PartialHit"
	case FullHit:
		return "FullHit"
	}
	return ""
}
