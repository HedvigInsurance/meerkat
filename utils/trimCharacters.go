package utils

import (
	"strings"
)

const SPACE string = " "
const EMPTY string = ""

var speacialCharactes = []string{"-", ".", "'"}

func TrimCharacters(stringToTrim string) string {
	stringWithoutSpaces := trimSpaces(stringToTrim)

	var stringWithoutSpecialCHaractes string = stringWithoutSpaces
	for i := 0; i < len(speacialCharactes); i++ {
		stringWithoutSpecialCHaractes = trimSpecialCharacters(stringWithoutSpecialCHaractes, speacialCharactes[i])
	}

	return stringWithoutSpecialCHaractes
}

func trimSpaces(stringToTrim string) string {
	return strings.Replace(stringToTrim, SPACE, EMPTY, -1)
}

func trimSpecialCharacters(stringToTrim string, specialcharacter string) string {
	return strings.Replace(stringToTrim, specialcharacter, EMPTY, -1)
}
