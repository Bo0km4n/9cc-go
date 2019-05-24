package main

import (
	"github.com/Bo0km4n/9cc-go/ast"
	"github.com/Bo0km4n/9cc-go/lexer"
)

func main() {
	// argv := os.Args[1]

	// in, _ := strconv.Atoi(argv)
	// fmt.Printf(".intel_syntax noprefix\n")
	// fmt.Printf(".global main\n")
	// fmt.Printf("main:\n")
	// fmt.Printf("	mov rax, %d\n", in)
	// fmt.Printf("	ret\n")
	l := lexer.NewLexer("1 + 2 * 3 / 10")
	tokens := l.Parse()
	p := ast.NewParser(tokens)
	p.Parse()
}
