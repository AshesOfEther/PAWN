package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"

	"pawn/ast"
	"pawn/parsing"
)

func main() {
	repl()
}

func repl() {
	reader := bufio.NewReader(os.Stdin)
	environment := newEnvironment(nil)

	for true {
		fmt.Print("> ")
		source, _ := reader.ReadString('\n')

		input := antlr.NewInputStream(source)
		lexer := parsing.NewBoardGameLangLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		parser := parsing.NewBoardGameLangParser(stream)
		parseTree := parser.StatementList()
		tree := ast.GenerateStatements(parseTree.AllStatement())

		execute(tree, environment)
	}
}

func execute(tree []ast.Statement, environment Environment) {
	defer func() {
		if r := recover(); r != nil {
			if pawnError, ok := r.(PawnError); ok {
				fmt.Printf("ERROR: %s\n", pawnError.message)
			} else {
				panic(r)
			}
		}
	}()

	for _, statement := range tree {
		if expressionStatement, ok := statement.(ast.ExpressionStatement); ok {
			result := evaluateExpression(expressionStatement.Expression, environment)
			fmt.Println(result)
		} else {
			evaluateStatement(statement, environment)
		}
	}
}
