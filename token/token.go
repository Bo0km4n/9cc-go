package token

import (
	"fmt"
	"strconv"
)

const (
	PLUS  = "+"
	MINUS = "-"

	TK_PLUS = iota
	TK_MINUS
	TK_NUM
	TK_EOF
)

var tokenMap = map[string]int64{
	PLUS:  TK_PLUS,
	MINUS: TK_NUM,
}

type Token struct {
	Ty    int
	Val   int
	Input string
}

func GetTokenType(char string) (int64, error) {
	if isDigit(char) {
		return TK_NUM, nil
	}
	val, ok := tokenMap[char]
	if !ok {
		return 0, fmt.Errorf("Not found token. got=%s", char)
	}
	return val, nil

}

func isDigit(char string) bool {
	if _, err := strconv.ParseInt(char, 10, 64); err != nil {
		return false
	}
	return true
}
