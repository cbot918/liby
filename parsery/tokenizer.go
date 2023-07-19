package main

type Token struct {
	Symbol string
	Value  string
}

type Tokenizer struct {
	content []byte
	index   int
	length  int
	tokens  []Token
}

func NewTokenizer(file []byte) *Tokenizer {
	return &Tokenizer{
		content: file,
		index:   0,
		length:  len(file),
	}
}

func (t *Tokenizer) GetTokens() ([]Token, error) {

	for t.index < t.length {
		if isChar(t.content[t.index]) {
			bytes := []byte{}
			for isChar(t.content[t.index]) {
				bytes = append(bytes, t.content[t.index])
				t.index += 1
			}
			str := string(bytes)

			if str == "package" {
				t.tokens = t.addToken("PACKAGE", str)
				t.index += 1
				continue
			}
			if str == "func" {
				t.tokens = t.addToken("FUNC", str)
				t.index += 1
				continue
			}

			t.tokens = t.addToken("LITERAL", str)
			continue
		}

		if isDoubleQoute(t.content[t.index]) {
			t.index += 1
			bytes := []byte{}
			for isChar(t.content[t.index]) {
				bytes = append(bytes, t.content[t.index])
				t.index += 1
			}
			t.tokens = t.addToken("STRING", string(bytes))

			// skip right doubleQoute
			if t.content[t.index] == '"' {
				t.index += 1
			}
			continue
		}

		if isLeftParen(t.content[t.index]) {
			t.tokens = t.addToken("LEFTPAREN", "(")
			t.index += 1
			continue
		}
		if isRightParen(t.content[t.index]) {
			t.tokens = t.addToken("RIGHTPAREN", ")")
			t.index += 1
			continue
		}
		if isLeftBracket(t.content[t.index]) {
			t.tokens = t.addToken("LEFTBRACKET", "{")
			t.index += 1
			continue
		}
		if isRightBracket(t.content[t.index]) {
			t.tokens = t.addToken("RIGHTBRACKET", "}")
			t.index += 1
			continue
		}

		t.index += 1
	}
	return t.tokens, nil
}

func (t *Tokenizer) addToken(symbol string, value string) []Token {
	return append(t.tokens, Token{
		Symbol: symbol,
		Value:  value,
	})
}

func isChar(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}
func isLeftParen(c byte) bool {
	return c == '('
}
func isRightParen(c byte) bool {
	return c == ')'
}
func isLeftBracket(c byte) bool {
	return c == '{'
}
func isRightBracket(c byte) bool {
	return c == '}'
}
func isDoubleQoute(c byte) bool {
	return c == '"'
}
