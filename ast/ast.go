package ast

import "github.com/Bo0km4n/9cc-go/token"

const (
	ND_NUM = iota
)

type NodeType int

type ExpressionNode struct {
	Type  NodeType
	lhs   *ExpressionNode
	rhs   *ExpressionNode
	Value interface{}
}

type Parser struct {
	tokens   []*token.Token
	position int
}

func NewParser(tokens []*token.Token) *Parser {
	return &Parser{
		tokens:   tokens,
		position: 0,
	}
}
func (p *Parser) Parse() *ExpressionNode {
	return nil
}
