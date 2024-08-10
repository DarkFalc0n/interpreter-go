package token

var TokenList = []Token{}

var HasError bool = false

func Tokenize(fileContents []byte) int{
	var line int = 1
	if len(fileContents) > 0 {
		for _, b := range fileContents {
			switch b {
			case '(':
				TokenList = append(TokenList, *NewToken(LEFT_PAREN, string(b), "null", line))
			case ')':
				TokenList = append(TokenList, *NewToken(RIGHT_PAREN, string(b), "null", line))
			case '{':
				TokenList = append(TokenList, *NewToken(LEFT_BRACE, string(b), "null", line))
			case '}':
				TokenList = append(TokenList, *NewToken(RIGHT_BRACE, string(b), "null", line))
				case ',':
				TokenList = append(TokenList, *NewToken(COMMA, string(b), "null", line))
			case '.':
				TokenList = append(TokenList, *NewToken(DOT, string(b), "null", line))
			case '-':
				TokenList = append(TokenList, *NewToken(MINUS, string(b), "null", line))
			case '+':
				TokenList = append(TokenList, *NewToken(PLUS, string(b), "null", line))
			case ';':
				TokenList = append(TokenList, *NewToken(SEMICOLON, string(b), "null", line))
			case '*':
				TokenList = append(TokenList, *NewToken(STAR, string(b), "null", line))
			case '\n':
				line++
			default:
				//eliminate whitespaces
				if (b != ' ' && b != '\t' && b != '\r') {
					TokenList = append(TokenList, *NewToken(INVALID, string(b), "null", line))
				}
			}
		}
	}
	return line
}