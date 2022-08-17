package parser

import (
	"regexp"
	"strings"
	"unicode"
)

/* cleanFile
Clean the fine from every comment and handle if necessary a bad EOF. For example empty line or no final semicolon
Also clean every invisible character induced from some IDE
*/
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

/* cleanComment
Clean every comment of the sql file
*/
func cleanComment(str string) string {
	commentWithDoubleDash := regexp.MustCompile(`--.*`)
	commentWithoutMultilineComment := regexp.MustCompile("/\\*[\\s\\S]*?\\*/")
	commentWithDiese := regexp.MustCompile(`#.*`)

	cleanedString := commentWithDoubleDash.ReplaceAllString(str, "")
	cleanedString = commentWithoutMultilineComment.ReplaceAllString(cleanedString, "")
	cleanedString = commentWithDiese.ReplaceAllString(cleanedString, "")
	return cleanedString
}

/* ShouldPopLastIndex
If the last index of an array is empty, then we trim it.
If there's only non-printable character then we send back true, so we trim the string.
Else we send false and do nothing.
*/
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

/* cleanString
Clean any non-printable character
*/
func cleanString(str string) string {
	str = strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, str)

	return str
}

/* cleanDoubleWhiteSpace
Clean every double whitespace
*/
func cleanDoubleWhiteSpace(str string) string {
	space := regexp.MustCompile(`\s+`)
	cleanedString := space.ReplaceAllString(str, " ")
	return cleanedString
}

/* cleanInParenthesisWhiteSpace
clean space before parenthesis in a given string
*/
func cleanInParenthesisWhiteSpace(str string) string {
	parenthesis := regexp.MustCompile(`\s*\(`)
	cleanedString := parenthesis.ReplaceAllString(str, "(")
	return cleanedString
}
