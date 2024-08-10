package token

var TokenList = []Token{}

func Tokenize(fileContents []byte) {
	if len(fileContents) > 0 {
		for _, b := range fileContents {
			switch b {
			case '(':
				TokenList = append(TokenList, *NewToken(LEFT_PAREN, "null"))
			case ')':
				TokenList = append(TokenList, *NewToken(RIGHT_PAREN, "null"))
			case '{':
				TokenList = append(TokenList, *NewToken(LEFT_BRACE, "null"))
			case '}':
				TokenList = append(TokenList, *NewToken(RIGHT_BRACE, "null"))
			}
		}
	}
}