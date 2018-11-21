package constants

type SanctionResult uint16

const (
	NoHit      SanctionResult = 0
	PartialHit SanctionResult = 1
	FullHit    SanctionResult = 2
)
