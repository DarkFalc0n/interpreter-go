package token

func GetNumber(index *int, fileContents *[]byte, line *int) string {
	var str string
	var didScanDecimal bool = false
	for i := *index; i < len(*fileContents); i++ {
		if (isDigit((*fileContents)[i]) ) {
			str += string((*fileContents)[i])
		} else if (*fileContents)[i] == '.'{
			if didScanDecimal {
				*index = i - 1
				break
			} else {
				str += string((*fileContents)[i])
				didScanDecimal = true
			}
		}else {
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