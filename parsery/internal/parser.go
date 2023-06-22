package parsery

type Parser struct {
	Nodes []Node
}

func NewParser(nodes []Node) *Parser {
	return &Parser{
		Nodes: nodes,
	}
}
