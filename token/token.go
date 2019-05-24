package token

type TokenType int

const (
	TK_NUM = iota
	TK_PLUS
	TK_MINUS
	TK_MUL
	TK_DIV
	TK_INT
)

type Token struct {
	Type    TokenType
	Literal string
}
