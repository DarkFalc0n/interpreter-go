package token

func isAlpha (b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || b == '_'
}

func isAlphaNumeric (b byte) bool {
	return isAlpha(b) || isDigit(b)
}

func isReservedKeyword (str string) int {
	for KEYWORD, keyword := range reservedKeywords {
		if keyword == str {
			return KEYWORD
		}
	}
	return -1
}

func GetIdentifier (index *int, fileContents *[]byte, line *int) string {
	var str string
	for i := *index; i < len(*fileContents); i++ {
		*index = i
		if !isAlphaNumeric((*fileContents)[i]) {
			*index = i - 1
			break
		}
		str += string((*fileContents)[i])

	}
	return str
}