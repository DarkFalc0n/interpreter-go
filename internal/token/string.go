package token

func GetString(index *int, fileContents *[]byte, line *int) string {
	var str string
	for i := *index; i < len(*fileContents); i++ {
		if (*fileContents)[i] == '\n' {
			*line++
		}
		if (*fileContents)[i] == '"' {
			*index = i
			break
		}
		if (i + 1) == len(*fileContents) {
			TokenError = UNTERMINATED_STRING
			*index = i
			return ""
		}
		str += string((*fileContents)[i])
	}
	return str
}
	