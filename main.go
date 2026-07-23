package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"

	"pawn/ast"
	"pawn/interpreter"
	"pawn/parsing"
)

func main() {
	switch len(os.Args) {
	case 1:
		repl()
	case 2:
		runFile(os.Args[1])
	default:
		fmt.Printf("ERROR: expected 0 or 1 argument, but %d were provided\n", len(os.Args))
		fmt.Printf("Usage: %s [path]\n", os.Args[0])
	}
}

func repl() {
	reader := bufio.NewReader(os.Stdin)
	environment := interpreter.CreateGlobalScope()

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

func runFile(path string) {
	sourceBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	source := string(sourceBytes)
	input := antlr.NewInputStream(source)
	lexer := parsing.NewBoardGameLangLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parsing.NewBoardGameLangParser(stream)
	parseTree := parser.StatementList()
	tree := ast.GenerateStatements(parseTree.AllStatement())

	defer func() {
		if r := recover(); r != nil {
			if pawnError, ok := r.(interpreter.PawnError); ok {
				fmt.Printf("ERROR: %s\n", pawnError.Message)
			} else {
				panic(r)
			}
		}
	}()

	environment := interpreter.CreateGlobalScope()
	interpreter.EvaluateStatements(tree, environment)
}

func execute(tree []ast.Statement, environment interpreter.Environment) {
	defer func() {
		if r := recover(); r != nil {
			if pawnError, ok := r.(interpreter.PawnError); ok {
				fmt.Printf("ERROR: %s\n", pawnError.Message)
			} else {
				panic(r)
			}
		}
	}()

	for _, statement := range tree {
		if expressionStatement, ok := statement.(ast.ExpressionStatement); ok {
			result := interpreter.EvaluateExpressionStatement(expressionStatement, environment)
			if result != nil {
				fmt.Println(result)
			}
		} else {
			interpreter.EvaluateStatement(statement, environment)
		}
	}
}
