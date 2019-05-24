package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argv := os.Args[1]

	in, _ := strconv.Atoi(argv)
	fmt.Printf(".intel_syntax noprefix\n")
	fmt.Printf(".global main\n")
	fmt.Printf("main:\n")
	fmt.Printf("	mov rax, %d\n", in)
	fmt.Printf("	ret\n")
}
