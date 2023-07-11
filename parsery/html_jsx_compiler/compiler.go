package main

import "os"

// type Compiler interface {
// 	Read(file string) (string, error)
// 	Tokenize(content string) (*Token, error)
// 	Parse()
// 	Executer()
// }

type HtmlJsxCompiler struct {
	Content            []byte
	Index              int
	Length             int
	timeToGetAttribute bool
	timeToGetInnerText bool
	closeTag           bool
	Tokens             []Token
}

func NewHtmlJsxCompiler(path string) (*HtmlJsxCompiler, error) {

	c := new(HtmlJsxCompiler)
	content, err := c.Read(path)
	if err != nil {
		return nil, err
	}

	return &HtmlJsxCompiler{
		Content:            content,
		Index:              0,
		Length:             len(content),
		timeToGetAttribute: false,
		timeToGetInnerText: false,
	}, nil
}

func (h *HtmlJsxCompiler) Read(file string) ([]byte, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return []byte{}, err
	}
	return content, nil
}
func (h *HtmlJsxCompiler) Tokenize(content []byte) (*Token, error) {
	for h.Index < h.Length {

		if h.Index >= h.Length {
			return nil, nil
		}

		if h.timeToGetAttribute || h.timeToGetInnerText {

			if h.timeToGetAttribute {
				h.Index += 1
				if isChar(h.Content[h.Index]) {
					str := h.tokAttribute()
					h.Tokens = append(h.Tokens, Token{
						Symbol: "ATTR",
						Value:  str,
					})
				}
				h.timeToGetAttribute = false
				lg(h.Tokens)
				continue
			}

			if h.timeToGetInnerText {
				str := ""
				for isChar(h.Content[h.Index]) {
					str += string(h.Content[h.Index])
					h.Index += 1
				}
				h.Tokens = append(h.Tokens, Token{
					Symbol: "INNERTEXT",
					Value:  str,
				})
			}
			h.timeToGetInnerText = false
			lg(h.Tokens)
			continue
		}

		if isLeft(h.Content[h.Index]) {
			tok := Token{
				Symbol: "LEFT",
				Value:  "<",
			}
			h.Index += 1
			h.Tokens = append(h.Tokens, tok)
			continue
		}

		if isSlash(h.Content[h.Index]) {
			str := "/"
			h.Index += 1
			for isChar(h.Content[h.Index]) {
				str += string(h.Content[h.Index])
				h.Index += 1
			}
			h.Tokens = append(h.Tokens, Token{
				Symbol: "CLOSETAG",
				Value:  str,
			})
			lg(h.Tokens)
		}

		if isChar(h.Content[h.Index]) {
			str := ""
			for isChar(h.Content[h.Index]) {
				str += string(h.Content[h.Index])
				h.Index += 1
			}

			if str == "div" {
				tok := Token{
					Symbol: "TAG",
					Value:  "div",
				}
				h.Tokens = append(h.Tokens, tok)
				lg(h.Tokens)
				h.timeToGetAttribute = true
				continue
			}
		}

		if isRight(h.Content[h.Index]) {
			tok := Token{
				Symbol: "RIGHT",
				Value:  ">",
			}
			h.Index += 1
			h.Tokens = append(h.Tokens, tok)
			lg(h.Tokens)
			h.timeToGetInnerText = true
			continue
		}

		h.Index += 1
	}
	return nil, nil
}
func (h *HtmlJsxCompiler) Parse()    {}
func (h *HtmlJsxCompiler) Executer() {}

func isLeft(c byte) bool {
	return c == byte('<')
}

func isRight(c byte) bool {
	return c == byte('>')
}

func isSlash(c byte) bool {
	return c == byte('/')
}

func isChar(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c >= 'Z'
}

func (h *HtmlJsxCompiler) tokAttribute() string {
	str := ""
	for h.Content[h.Index] != '>' {
		str += string(h.Content[h.Index])
		h.Index += 1
	}
	return str
}
