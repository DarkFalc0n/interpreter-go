package token

var TokenList = []Token{}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func Tokenize(fileContents []byte) int{
	var line int = 1
	if len(fileContents) > 0 {
		for index := 0; index < len(fileContents); index++ {
			b := fileContents[index]
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
			case '=':
				if index + 1 < len(fileContents) && fileContents[index + 1] == '=' {
					TokenList = append(TokenList, *NewToken(EQUAL_EQUAL, "==", "null", line))
					index++
				} else {
					TokenList = append(TokenList, *NewToken(EQUAL, string(b), "null", line))
				}
			case '!':
				if index + 1 < len(fileContents) && fileContents[index + 1] == '=' {
					TokenList = append(TokenList, *NewToken(BANG_EQUAL, "!=", "null", line))
					index++
				} else {
					TokenList = append(TokenList, *NewToken(BANG, string(b), "null", line))
				}
			case '<':
				if index + 1 < len(fileContents) && fileContents[index + 1] == '=' {
					TokenList = append(TokenList, *NewToken(LESS_EQUAL, "<=", "null", line))
					index++
				} else {
					TokenList = append(TokenList, *NewToken(LESS, string(b), "null", line))
				}
			case '>':
				if index + 1 < len(fileContents) && fileContents[index + 1] == '=' {
					TokenList = append(TokenList, *NewToken(GREATER_EQUAL, ">=", "null", line))
					index++
				} else {
					TokenList = append(TokenList, *NewToken(GREATER, string(b), "null", line))
				}
			case '/':
				if index + 1 < len(fileContents) && fileContents[index + 1] == '/' {
					for fileContents[index] != '\n' && index < len(fileContents) - 1 {
						index++
					}
					if fileContents[index] == '\n' {
						line++
					}
				} else {
					TokenList = append(TokenList, *NewToken(SLASH, string(b), "null", line))
				}
			case '"':
				index++
				var str string = GetString(&index, &fileContents, &line)
				if TokenError == UNTERMINATED_STRING {					
					ThrowUnterminatedStringError(line)
				} else {
					TokenList = append(TokenList, *NewToken(STRING, "\""+str+"\"", str, line))
				}
			case '\n':
				line++
			default:
				//handle numbers
				if isDigit(b) {
					var num string = GetNumber(&index, &fileContents, &line)
					TokenList = append(TokenList, *NewToken(NUMBER, num, parseNumber(num), line))
				} else if isAlpha(b) {
					var str string = GetIdentifier(&index, &fileContents, &line)
					var KEYWORD int = isReservedKeyword(str)
					if KEYWORD == -1 {
					TokenList = append(TokenList, *NewToken(IDENTIFIER, str, "null", line))
					} else {
						TokenList = append(TokenList, *NewToken(KEYWORD, str, "null", line))
					}
				} else if (b != ' ' && b != '\t' && b != '\r') {
					TokenError = UNEXPECTED_TOKEN
					TokenList = append(TokenList, *NewToken(INVALID, string(b), "null", line))
				}
			}
		}
	}
	return line
}