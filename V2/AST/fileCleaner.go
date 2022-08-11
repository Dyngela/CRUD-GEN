package AST

import (
	"regexp"
	"strings"
	"unicode"
)

func cleanFile(sqlSplit []string) []string {
	if shouldPopLastIndex(sqlSplit[len(sqlSplit)-1]) {
		sqlSplit = sqlSplit[:len(sqlSplit)-1]
	}

	for i := 0; i < len(sqlSplit); i++ {
		sqlSplit[i] = cleanComment(sqlSplit[i])
		sqlSplit[i] = cleanString(sqlSplit[i])
	}
	return sqlSplit
}

func cleanComment(str string) string {
	return str
}

// ShouldPopLastIndex If the last index of an array is empty, then we trim it.
// If there's only non printable character then we send back true so we trim the string.
// Else we send false and do nothing.
func shouldPopLastIndex(sqlSplit string) bool {
	sqlSplit = strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, sqlSplit)

	if len(sqlSplit) == 0 {
		return true
	}
	return false
}

// Clean any non-printable character
func cleanString(str string) string {
	str = strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, str)

	return str
}

func cleanLastParenthesis(str string) string {
	return strings.Replace(str, ")", "", 1)
}

func cleanDoubleWhiteSpace(str string) string {
	space := regexp.MustCompile(`\s+`)
	cleanedString := space.ReplaceAllString(str, " ")
	return cleanedString
}

func cleanInParenthesisWhiteSpace(str string) string {
	//parenthesis := regexp.MustCompile(`\s*\((.*?)\)`)
	parenthesis := regexp.MustCompile(`\s*\(`)
	cleanedString := parenthesis.ReplaceAllString(str, "(")
	//log.Println(cleanedString)
	return cleanedString
}

// numeric\s*\((.*?)\) -> get numeric and value
// Je veux effacer tout espace dans les parenthése et tout les espace a gauche de la
// parenthèse
