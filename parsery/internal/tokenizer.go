package parsery

// var nodes []Node

type Tokenizer struct {
	Content []rune
	Index   int
	Max     int
	Nodes   []Node
}

type Node struct {
	TokType Token
	Value   string
}

func NewTokenizer(content []rune) *Tokenizer {
	return &Tokenizer{
		Content: content,
	}
}

func (t *Tokenizer) GetTokens() []Node {
	return t.process()
}

func (t *Tokenizer) process() []Node {

	for t.Index < len(t.Content) {

		if t.isDigit(t.Content[t.Index]) {
			_val := ""
			for t.notEnd() && t.isDigit(t.Content[t.Index]) {
				_val += string(t.Content[t.Index])
				logc(t.Content[t.Index])
				t.Index += 1
			}

			t.Nodes = append(t.Nodes, Node{TokType: NUMBER, Value: _val})

			continue
		}

		if string(t.Content[t.Index]) == tokens[ADD] {
			t.Nodes = append(t.Nodes, Node{TokType: ADD, Value: tokens[ADD]})
			logc(t.Content[t.Index])
			t.Index += 1
			continue
		}

		t.Index += 1
	}
	log(t.Nodes)
	t.fPrint()
	return []Node{}
}

// func (t *Tokenizer) getTokType(target Token) Token{
// 	switch target{
// 	case ADD:
// 		return ADD
// 	}
// }

func (t *Tokenizer) isDigit(target rune) bool {
	return target >= '0' && target <= '9'
}

func (t *Tokenizer) notEnd() bool {
	return t.Index < len(t.Content)
}

func (t *Tokenizer) fPrint() {
	for _, item := range t.Nodes {
		// log(item.TokType)
		// log(tokens[ADD])
		// return
		logf("%s: %s\n", tokens[item.TokType], item.Value)
	}
}

// func isLetter(ch rune) bool {
// 	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch >= utf8.RuneSelf && unicode.IsLetter(ch)
// }

// func isDigit(ch rune) bool {
// 	return '0' <= ch && ch <= '9' || ch >= utf8.RuneSelf && unicode.IsDigit(ch)
// }
