package token

func GetNumber(index *int, fileContents *[]byte, line *int) string {
	var str string
	for i := *index; i < len(*fileContents); i++ {
		if (isDigit((*fileContents)[i]) || (*fileContents)[i] == '.') {
			str += string((*fileContents)[i])
		} else {
			*index = i - 1
			break;
		}

		if i + 1 == len(*fileContents) {
			*index = i
			break
		}
	}
	return str
}