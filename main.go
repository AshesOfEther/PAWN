package main

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"

	"pawn/parsing"
)

func main() {
	filePath := os.Args[1]

	input, _ := antlr.NewFileStream(filePath)
	lexer := parsing.NewBoardGameLangLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parsing.NewBoardGameLangParser(stream)
	tree := parser.StatementList()

	fmt.Printf("%#v\n", generateStatements(tree.AllStatement()))
}
