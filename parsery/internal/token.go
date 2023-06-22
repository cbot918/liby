package parsery

type Token int

const (
	ADD Token = iota
	NUMBER
)

var tokens = [...]string{
	ADD:    "+",
	NUMBER: "NUMBER",
}
