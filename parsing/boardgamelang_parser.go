// Code generated from ./BoardGameLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // BoardGameLang
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type BoardGameLangParser struct {
	*antlr.BaseParser
}

var BoardGameLangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func boardgamelangParserInit() {
	staticData := &BoardGameLangParserStaticData
	staticData.LiteralNames = []string{
		"", "'('", "')'", "'{'", "'}'", "'['", "']'", "'^'", "'&&'", "'||'",
		"'.'", "','", "'='", "'apply'", "'while'", "'if'", "'else'", "'fn'",
		"'return'", "';'", "", "", "", "", "", "", "'*'", "'/'", "'%'", "'+'",
		"'-'", "'..'", "'..='", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "LINE_COMMENT", "MULTILINE_COMMENT", "BOOL", "NAME", "NUMBER",
		"STRING", "MUL", "DIV", "MOD", "ADD", "SUB", "RANGE", "RANGE_INCLUSIVE",
		"EQUAL", "NOT_EQUAL", "LESS_THAN", "GREATER_THAN", "LESS_EQUAL", "GREATER_EQUAL",
		"WS",
	}
	staticData.RuleNames = []string{
		"expression", "member", "literal", "argList", "expressionList", "namedArgs",
		"namedArg", "statementList", "statement", "functionDeclArgs", "nameList",
		"namedArgsDeclList", "namedArgDecl",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 39, 253, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 3, 0, 34, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 63, 8, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 5, 0, 70, 8, 0, 10, 0, 12, 0, 73, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 5, 1, 81, 8, 1, 10, 1, 12, 1, 84, 9, 1, 1, 2, 1, 2, 1, 2, 1,
		2, 5, 2, 90, 8, 2, 10, 2, 12, 2, 93, 9, 2, 1, 2, 3, 2, 96, 8, 2, 3, 2,
		98, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 104, 8, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 3, 3, 112, 8, 3, 1, 3, 3, 3, 115, 8, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 3, 3, 122, 8, 3, 1, 3, 1, 3, 3, 3, 126, 8, 3, 1, 4, 1, 4, 1,
		4, 5, 4, 131, 8, 4, 10, 4, 12, 4, 134, 9, 4, 1, 5, 1, 5, 1, 5, 5, 5, 139,
		8, 5, 10, 5, 12, 5, 142, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 5, 7, 149,
		8, 7, 10, 7, 12, 7, 152, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 178, 8, 8, 10, 8, 12, 8, 181, 9, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 188, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 3, 8, 209, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		3, 9, 217, 8, 9, 1, 9, 3, 9, 220, 8, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3,
		9, 227, 8, 9, 1, 9, 1, 9, 3, 9, 231, 8, 9, 1, 10, 1, 10, 1, 10, 5, 10,
		236, 8, 10, 10, 10, 12, 10, 239, 9, 10, 1, 11, 1, 11, 1, 11, 5, 11, 244,
		8, 11, 10, 11, 12, 11, 247, 9, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 0,
		2, 0, 2, 13, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 0, 4, 1, 0,
		26, 28, 1, 0, 29, 30, 1, 0, 31, 32, 1, 0, 33, 38, 281, 0, 33, 1, 0, 0,
		0, 2, 74, 1, 0, 0, 0, 4, 103, 1, 0, 0, 0, 6, 125, 1, 0, 0, 0, 8, 127, 1,
		0, 0, 0, 10, 135, 1, 0, 0, 0, 12, 143, 1, 0, 0, 0, 14, 150, 1, 0, 0, 0,
		16, 208, 1, 0, 0, 0, 18, 230, 1, 0, 0, 0, 20, 232, 1, 0, 0, 0, 22, 240,
		1, 0, 0, 0, 24, 248, 1, 0, 0, 0, 26, 27, 6, 0, -1, 0, 27, 28, 5, 1, 0,
		0, 28, 29, 3, 0, 0, 0, 29, 30, 5, 2, 0, 0, 30, 34, 1, 0, 0, 0, 31, 34,
		3, 2, 1, 0, 32, 34, 3, 4, 2, 0, 33, 26, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0,
		33, 32, 1, 0, 0, 0, 34, 71, 1, 0, 0, 0, 35, 36, 10, 9, 0, 0, 36, 37, 5,
		7, 0, 0, 37, 70, 3, 0, 0, 10, 38, 39, 10, 8, 0, 0, 39, 40, 7, 0, 0, 0,
		40, 70, 3, 0, 0, 9, 41, 42, 10, 7, 0, 0, 42, 43, 7, 1, 0, 0, 43, 70, 3,
		0, 0, 8, 44, 45, 10, 6, 0, 0, 45, 46, 7, 2, 0, 0, 46, 70, 3, 0, 0, 7, 47,
		48, 10, 5, 0, 0, 48, 49, 7, 3, 0, 0, 49, 70, 3, 0, 0, 6, 50, 51, 10, 4,
		0, 0, 51, 52, 5, 8, 0, 0, 52, 70, 3, 0, 0, 5, 53, 54, 10, 3, 0, 0, 54,
		55, 5, 9, 0, 0, 55, 70, 3, 0, 0, 4, 56, 57, 10, 11, 0, 0, 57, 62, 3, 6,
		3, 0, 58, 59, 5, 3, 0, 0, 59, 60, 3, 14, 7, 0, 60, 61, 5, 4, 0, 0, 61,
		63, 1, 0, 0, 0, 62, 58, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 70, 1, 0, 0,
		0, 64, 65, 10, 10, 0, 0, 65, 66, 5, 5, 0, 0, 66, 67, 3, 8, 4, 0, 67, 68,
		5, 6, 0, 0, 68, 70, 1, 0, 0, 0, 69, 35, 1, 0, 0, 0, 69, 38, 1, 0, 0, 0,
		69, 41, 1, 0, 0, 0, 69, 44, 1, 0, 0, 0, 69, 47, 1, 0, 0, 0, 69, 50, 1,
		0, 0, 0, 69, 53, 1, 0, 0, 0, 69, 56, 1, 0, 0, 0, 69, 64, 1, 0, 0, 0, 70,
		73, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 1, 1, 0, 0,
		0, 73, 71, 1, 0, 0, 0, 74, 75, 6, 1, -1, 0, 75, 76, 5, 23, 0, 0, 76, 82,
		1, 0, 0, 0, 77, 78, 10, 2, 0, 0, 78, 79, 5, 10, 0, 0, 79, 81, 5, 23, 0,
		0, 80, 77, 1, 0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82, 83,
		1, 0, 0, 0, 83, 3, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 85, 97, 5, 5, 0, 0,
		86, 91, 3, 0, 0, 0, 87, 88, 5, 11, 0, 0, 88, 90, 3, 0, 0, 0, 89, 87, 1,
		0, 0, 0, 90, 93, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92,
		95, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 94, 96, 5, 11, 0, 0, 95, 94, 1, 0,
		0, 0, 95, 96, 1, 0, 0, 0, 96, 98, 1, 0, 0, 0, 97, 86, 1, 0, 0, 0, 97, 98,
		1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 104, 5, 6, 0, 0, 100, 104, 5, 22, 0,
		0, 101, 104, 5, 24, 0, 0, 102, 104, 5, 25, 0, 0, 103, 85, 1, 0, 0, 0, 103,
		100, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 102, 1, 0, 0, 0, 104, 5, 1,
		0, 0, 0, 105, 106, 5, 1, 0, 0, 106, 126, 5, 2, 0, 0, 107, 108, 5, 1, 0,
		0, 108, 111, 3, 8, 4, 0, 109, 110, 5, 11, 0, 0, 110, 112, 3, 10, 5, 0,
		111, 109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 114, 1, 0, 0, 0, 113,
		115, 5, 11, 0, 0, 114, 113, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 116,
		1, 0, 0, 0, 116, 117, 5, 2, 0, 0, 117, 126, 1, 0, 0, 0, 118, 119, 5, 1,
		0, 0, 119, 121, 3, 10, 5, 0, 120, 122, 5, 11, 0, 0, 121, 120, 1, 0, 0,
		0, 121, 122, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 124, 5, 2, 0, 0, 124,
		126, 1, 0, 0, 0, 125, 105, 1, 0, 0, 0, 125, 107, 1, 0, 0, 0, 125, 118,
		1, 0, 0, 0, 126, 7, 1, 0, 0, 0, 127, 132, 3, 0, 0, 0, 128, 129, 5, 11,
		0, 0, 129, 131, 3, 0, 0, 0, 130, 128, 1, 0, 0, 0, 131, 134, 1, 0, 0, 0,
		132, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 9, 1, 0, 0, 0, 134, 132,
		1, 0, 0, 0, 135, 140, 3, 12, 6, 0, 136, 137, 5, 11, 0, 0, 137, 139, 3,
		12, 6, 0, 138, 136, 1, 0, 0, 0, 139, 142, 1, 0, 0, 0, 140, 138, 1, 0, 0,
		0, 140, 141, 1, 0, 0, 0, 141, 11, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 143,
		144, 5, 23, 0, 0, 144, 145, 5, 12, 0, 0, 145, 146, 3, 0, 0, 0, 146, 13,
		1, 0, 0, 0, 147, 149, 3, 16, 8, 0, 148, 147, 1, 0, 0, 0, 149, 152, 1, 0,
		0, 0, 150, 148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 15, 1, 0, 0, 0,
		152, 150, 1, 0, 0, 0, 153, 154, 5, 13, 0, 0, 154, 155, 3, 0, 0, 0, 155,
		156, 5, 3, 0, 0, 156, 157, 3, 14, 7, 0, 157, 158, 5, 4, 0, 0, 158, 209,
		1, 0, 0, 0, 159, 160, 5, 14, 0, 0, 160, 161, 3, 0, 0, 0, 161, 162, 5, 3,
		0, 0, 162, 163, 3, 14, 7, 0, 163, 164, 5, 4, 0, 0, 164, 209, 1, 0, 0, 0,
		165, 166, 5, 15, 0, 0, 166, 167, 3, 0, 0, 0, 167, 168, 5, 3, 0, 0, 168,
		169, 3, 14, 7, 0, 169, 179, 5, 4, 0, 0, 170, 171, 5, 16, 0, 0, 171, 172,
		5, 15, 0, 0, 172, 173, 3, 0, 0, 0, 173, 174, 5, 3, 0, 0, 174, 175, 3, 14,
		7, 0, 175, 176, 5, 4, 0, 0, 176, 178, 1, 0, 0, 0, 177, 170, 1, 0, 0, 0,
		178, 181, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180,
		187, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 182, 183, 5, 16, 0, 0, 183, 184,
		5, 3, 0, 0, 184, 185, 3, 14, 7, 0, 185, 186, 5, 4, 0, 0, 186, 188, 1, 0,
		0, 0, 187, 182, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 209, 1, 0, 0, 0,
		189, 190, 5, 17, 0, 0, 190, 191, 5, 23, 0, 0, 191, 192, 3, 18, 9, 0, 192,
		193, 5, 3, 0, 0, 193, 194, 3, 14, 7, 0, 194, 195, 5, 4, 0, 0, 195, 209,
		1, 0, 0, 0, 196, 197, 5, 18, 0, 0, 197, 198, 3, 0, 0, 0, 198, 199, 5, 19,
		0, 0, 199, 209, 1, 0, 0, 0, 200, 201, 3, 0, 0, 0, 201, 202, 5, 19, 0, 0,
		202, 209, 1, 0, 0, 0, 203, 204, 5, 23, 0, 0, 204, 205, 5, 12, 0, 0, 205,
		206, 3, 0, 0, 0, 206, 207, 5, 19, 0, 0, 207, 209, 1, 0, 0, 0, 208, 153,
		1, 0, 0, 0, 208, 159, 1, 0, 0, 0, 208, 165, 1, 0, 0, 0, 208, 189, 1, 0,
		0, 0, 208, 196, 1, 0, 0, 0, 208, 200, 1, 0, 0, 0, 208, 203, 1, 0, 0, 0,
		209, 17, 1, 0, 0, 0, 210, 211, 5, 1, 0, 0, 211, 231, 5, 2, 0, 0, 212, 213,
		5, 1, 0, 0, 213, 216, 3, 20, 10, 0, 214, 215, 5, 11, 0, 0, 215, 217, 3,
		22, 11, 0, 216, 214, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 219, 1, 0,
		0, 0, 218, 220, 5, 11, 0, 0, 219, 218, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0,
		220, 221, 1, 0, 0, 0, 221, 222, 5, 2, 0, 0, 222, 231, 1, 0, 0, 0, 223,
		224, 5, 1, 0, 0, 224, 226, 3, 22, 11, 0, 225, 227, 5, 11, 0, 0, 226, 225,
		1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 229, 5, 2,
		0, 0, 229, 231, 1, 0, 0, 0, 230, 210, 1, 0, 0, 0, 230, 212, 1, 0, 0, 0,
		230, 223, 1, 0, 0, 0, 231, 19, 1, 0, 0, 0, 232, 237, 5, 23, 0, 0, 233,
		234, 5, 11, 0, 0, 234, 236, 5, 23, 0, 0, 235, 233, 1, 0, 0, 0, 236, 239,
		1, 0, 0, 0, 237, 235, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 21, 1, 0,
		0, 0, 239, 237, 1, 0, 0, 0, 240, 245, 3, 24, 12, 0, 241, 242, 5, 11, 0,
		0, 242, 244, 3, 24, 12, 0, 243, 241, 1, 0, 0, 0, 244, 247, 1, 0, 0, 0,
		245, 243, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 23, 1, 0, 0, 0, 247, 245,
		1, 0, 0, 0, 248, 249, 5, 23, 0, 0, 249, 250, 5, 12, 0, 0, 250, 251, 3,
		0, 0, 0, 251, 25, 1, 0, 0, 0, 25, 33, 62, 69, 71, 82, 91, 95, 97, 103,
		111, 114, 121, 125, 132, 140, 150, 179, 187, 208, 216, 219, 226, 230, 237,
		245,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// BoardGameLangParserInit initializes any static state used to implement BoardGameLangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBoardGameLangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BoardGameLangParserInit() {
	staticData := &BoardGameLangParserStaticData
	staticData.once.Do(boardgamelangParserInit)
}

// NewBoardGameLangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBoardGameLangParser(input antlr.TokenStream) *BoardGameLangParser {
	BoardGameLangParserInit()
	this := new(BoardGameLangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &BoardGameLangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "BoardGameLang.g4"

	return this
}

// BoardGameLangParser tokens.
const (
	BoardGameLangParserEOF               = antlr.TokenEOF
	BoardGameLangParserT__0              = 1
	BoardGameLangParserT__1              = 2
	BoardGameLangParserT__2              = 3
	BoardGameLangParserT__3              = 4
	BoardGameLangParserT__4              = 5
	BoardGameLangParserT__5              = 6
	BoardGameLangParserT__6              = 7
	BoardGameLangParserT__7              = 8
	BoardGameLangParserT__8              = 9
	BoardGameLangParserT__9              = 10
	BoardGameLangParserT__10             = 11
	BoardGameLangParserT__11             = 12
	BoardGameLangParserT__12             = 13
	BoardGameLangParserT__13             = 14
	BoardGameLangParserT__14             = 15
	BoardGameLangParserT__15             = 16
	BoardGameLangParserT__16             = 17
	BoardGameLangParserT__17             = 18
	BoardGameLangParserT__18             = 19
	BoardGameLangParserLINE_COMMENT      = 20
	BoardGameLangParserMULTILINE_COMMENT = 21
	BoardGameLangParserBOOL              = 22
	BoardGameLangParserNAME              = 23
	BoardGameLangParserNUMBER            = 24
	BoardGameLangParserSTRING            = 25
	BoardGameLangParserMUL               = 26
	BoardGameLangParserDIV               = 27
	BoardGameLangParserMOD               = 28
	BoardGameLangParserADD               = 29
	BoardGameLangParserSUB               = 30
	BoardGameLangParserRANGE             = 31
	BoardGameLangParserRANGE_INCLUSIVE   = 32
	BoardGameLangParserEQUAL             = 33
	BoardGameLangParserNOT_EQUAL         = 34
	BoardGameLangParserLESS_THAN         = 35
	BoardGameLangParserGREATER_THAN      = 36
	BoardGameLangParserLESS_EQUAL        = 37
	BoardGameLangParserGREATER_EQUAL     = 38
	BoardGameLangParserWS                = 39
)

// BoardGameLangParser rules.
const (
	BoardGameLangParserRULE_expression        = 0
	BoardGameLangParserRULE_member            = 1
	BoardGameLangParserRULE_literal           = 2
	BoardGameLangParserRULE_argList           = 3
	BoardGameLangParserRULE_expressionList    = 4
	BoardGameLangParserRULE_namedArgs         = 5
	BoardGameLangParserRULE_namedArg          = 6
	BoardGameLangParserRULE_statementList     = 7
	BoardGameLangParserRULE_statement         = 8
	BoardGameLangParserRULE_functionDeclArgs  = 9
	BoardGameLangParserRULE_nameList          = 10
	BoardGameLangParserRULE_namedArgsDeclList = 11
	BoardGameLangParserRULE_namedArgDecl      = 12
)

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BoardGameLangParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CallContext struct {
	ExpressionContext
}

func NewCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallContext {
	var p = new(CallContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CallContext) ArgList() IArgListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *CallContext) StatementList() IStatementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *CallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterCall(s)
	}
}

func (s *CallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitCall(s)
	}
}

func (s *CallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulDivModContext struct {
	ExpressionContext
	op antlr.Token
}

func NewMulDivModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivModContext {
	var p = new(MulDivModContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MulDivModContext) GetOp() antlr.Token { return s.op }

func (s *MulDivModContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivModContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MulDivModContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MulDivModContext) MUL() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserMUL, 0)
}

func (s *MulDivModContext) DIV() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserDIV, 0)
}

func (s *MulDivModContext) MOD() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserMOD, 0)
}

func (s *MulDivModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterMulDivMod(s)
	}
}

func (s *MulDivModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitMulDivMod(s)
	}
}

func (s *MulDivModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitMulDivMod(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrContext struct {
	ExpressionContext
	op antlr.Token
}

func NewOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrContext {
	var p = new(OrContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *OrContext) GetOp() antlr.Token { return s.op }

func (s *OrContext) SetOp(v antlr.Token) { s.op = v }

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *OrContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterOr(s)
	}
}

func (s *OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitOr(s)
	}
}

func (s *OrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MemberExprContext struct {
	ExpressionContext
}

func NewMemberExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberExprContext {
	var p = new(MemberExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MemberExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberExprContext) Member() IMemberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMemberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMemberContext)
}

func (s *MemberExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterMemberExpr(s)
	}
}

func (s *MemberExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitMemberExpr(s)
	}
}

func (s *MemberExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitMemberExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubContext struct {
	ExpressionContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AddSubContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddSubContext) ADD() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserADD, 0)
}

func (s *AddSubContext) SUB() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserSUB, 0)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitAddSub(s)
	}
}

func (s *AddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type ComparisonContext struct {
	ExpressionContext
	op antlr.Token
}

func NewComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparisonContext {
	var p = new(ComparisonContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ComparisonContext) GetOp() antlr.Token { return s.op }

func (s *ComparisonContext) SetOp(v antlr.Token) { s.op = v }

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ComparisonContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ComparisonContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserEQUAL, 0)
}

func (s *ComparisonContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserNOT_EQUAL, 0)
}

func (s *ComparisonContext) LESS_EQUAL() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserLESS_EQUAL, 0)
}

func (s *ComparisonContext) GREATER_EQUAL() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserGREATER_EQUAL, 0)
}

func (s *ComparisonContext) LESS_THAN() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserLESS_THAN, 0)
}

func (s *ComparisonContext) GREATER_THAN() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserGREATER_THAN, 0)
}

func (s *ComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterComparison(s)
	}
}

func (s *ComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitComparison(s)
	}
}

func (s *ComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParensContext struct {
	ExpressionContext
}

func NewParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensContext {
	var p = new(ParensContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterParens(s)
	}
}

func (s *ParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitParens(s)
	}
}

func (s *ParensContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitParens(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndContext struct {
	ExpressionContext
	op antlr.Token
}

func NewAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndContext {
	var p = new(AndContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AndContext) GetOp() antlr.Token { return s.op }

func (s *AndContext) SetOp(v antlr.Token) { s.op = v }

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AndContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitAnd(s)
	}
}

func (s *AndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralExprContext struct {
	ExpressionContext
}

func NewLiteralExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExprContext {
	var p = new(LiteralExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LiteralExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExprContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterLiteralExpr(s)
	}
}

func (s *LiteralExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitLiteralExpr(s)
	}
}

func (s *LiteralExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitLiteralExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IndexContext struct {
	ExpressionContext
}

func NewIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexContext {
	var p = new(IndexContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *IndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterIndex(s)
	}
}

func (s *IndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitIndex(s)
	}
}

func (s *IndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

type RangeContext struct {
	ExpressionContext
	op antlr.Token
}

func NewRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RangeContext {
	var p = new(RangeContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *RangeContext) GetOp() antlr.Token { return s.op }

func (s *RangeContext) SetOp(v antlr.Token) { s.op = v }

func (s *RangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *RangeContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RangeContext) RANGE_INCLUSIVE() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserRANGE_INCLUSIVE, 0)
}

func (s *RangeContext) RANGE() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserRANGE, 0)
}

func (s *RangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterRange(s)
	}
}

func (s *RangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitRange(s)
	}
}

func (s *RangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitRange(s)

	default:
		return t.VisitChildren(s)
	}
}

type PowerContext struct {
	ExpressionContext
}

func NewPowerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowerContext {
	var p = new(PowerContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *PowerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *PowerContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PowerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterPower(s)
	}
}

func (s *PowerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitPower(s)
	}
}

func (s *PowerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitPower(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BoardGameLangParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *BoardGameLangParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, BoardGameLangParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BoardGameLangParserT__0:
		localctx = NewParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(27)
			p.Match(BoardGameLangParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(28)
			p.expression(0)
		}
		{
			p.SetState(29)
			p.Match(BoardGameLangParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BoardGameLangParserNAME:
		localctx = NewMemberExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(31)
			p.member(0)
		}

	case BoardGameLangParserT__4, BoardGameLangParserBOOL, BoardGameLangParserNUMBER, BoardGameLangParserSTRING:
		localctx = NewLiteralExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(32)
			p.Literal()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(69)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowerContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BoardGameLangParserRULE_expression)
				p.SetState(35)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(36)
					p.Match(BoardGameLangParserT__6)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(37)
					p.expression(10)
				}

			case 2:
				localctx = NewMulDivModContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BoardGameLangParserRULE_expression)
				p.SetState(38)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(39)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivModContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&469762048) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivModContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(40)
					p.expression(9)
				}

			case 3:
				localctx = NewAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BoardGameLangParserRULE_expression)
				p.SetState(41)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(42)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == BoardGameLangParserADD || _la == BoardGameLangParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(43)
					p.expression(8)
				}

			case 4:
				localctx = NewRangeContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BoardGameLangParserRULE_expression)
				p.SetState(44)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(45)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RangeContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == BoardGameLangParserRANGE || _la == BoardGameLangParserRANGE_INCLUSIVE) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RangeContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(46)
					p.expression(7)
				}

			case 5:
				localctx = NewComparisonContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BoardGameLangParserRULE_expression)
				p.SetState(47)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(48)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ComparisonContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&541165879296) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ComparisonContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(49)
					p.expression(6)
				}

			case 6:
				localctx = NewAndContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BoardGameLangParserRULE_expression)
				p.SetState(50)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(51)

					var _m = p.Match(BoardGameLangParserT__7)

					localctx.(*AndContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(52)
					p.expression(5)
				}

			case 7:
				localctx = NewOrContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BoardGameLangParserRULE_expression)
				p.SetState(53)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(54)

					var _m = p.Match(BoardGameLangParserT__8)

					localctx.(*OrContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(55)
					p.expression(4)
				}

			case 8:
				localctx = NewCallContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BoardGameLangParserRULE_expression)
				p.SetState(56)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(57)
					p.ArgList()
				}
				p.SetState(62)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(58)
						p.Match(BoardGameLangParserT__2)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					{
						p.SetState(59)
						p.StatementList()
					}
					{
						p.SetState(60)
						p.Match(BoardGameLangParserT__3)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				} else if p.HasError() { // JIM
					goto errorExit
				}

			case 9:
				localctx = NewIndexContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BoardGameLangParserRULE_expression)
				p.SetState(64)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(65)
					p.Match(BoardGameLangParserT__4)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(66)
					p.ExpressionList()
				}
				{
					p.SetState(67)
					p.Match(BoardGameLangParserT__5)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMemberContext is an interface to support dynamic dispatch.
type IMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMemberContext differentiates from other interfaces.
	IsMemberContext()
}

type MemberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberContext() *MemberContext {
	var p = new(MemberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_member
	return p
}

func InitEmptyMemberContext(p *MemberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_member
}

func (*MemberContext) IsMemberContext() {}

func NewMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberContext {
	var p = new(MemberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BoardGameLangParserRULE_member

	return p
}

func (s *MemberContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberContext) CopyAll(ctx *MemberContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *MemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PropertyContext struct {
	MemberContext
}

func NewPropertyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropertyContext {
	var p = new(PropertyContext)

	InitEmptyMemberContext(&p.MemberContext)
	p.parser = parser
	p.CopyAll(ctx.(*MemberContext))

	return p
}

func (s *PropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyContext) Member() IMemberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMemberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMemberContext)
}

func (s *PropertyContext) NAME() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserNAME, 0)
}

func (s *PropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterProperty(s)
	}
}

func (s *PropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitProperty(s)
	}
}

func (s *PropertyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitProperty(s)

	default:
		return t.VisitChildren(s)
	}
}

type NameContext struct {
	MemberContext
}

func NewNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NameContext {
	var p = new(NameContext)

	InitEmptyMemberContext(&p.MemberContext)
	p.parser = parser
	p.CopyAll(ctx.(*MemberContext))

	return p
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) NAME() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserNAME, 0)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitName(s)
	}
}

func (s *NameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BoardGameLangParser) Member() (localctx IMemberContext) {
	return p.member(0)
}

func (p *BoardGameLangParser) member(_p int) (localctx IMemberContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewMemberContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMemberContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, BoardGameLangParserRULE_member, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewNameContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(75)
		p.Match(BoardGameLangParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPropertyContext(p, NewMemberContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, BoardGameLangParserRULE_member)
			p.SetState(77)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(78)
				p.Match(BoardGameLangParserT__9)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(79)
				p.Match(BoardGameLangParserNAME)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BoardGameLangParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyAll(ctx *LiteralContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NumberContext struct {
	LiteralContext
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserNUMBER, 0)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolContext struct {
	LiteralContext
}

func NewBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolContext {
	var p = new(BoolContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *BoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolContext) BOOL() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserBOOL, 0)
}

func (s *BoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterBool(s)
	}
}

func (s *BoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitBool(s)
	}
}

func (s *BoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitBool(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListContext struct {
	LiteralContext
}

func NewListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListContext {
	var p = new(ListContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitList(s)
	}
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringContext struct {
	LiteralContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BoardGameLangParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BoardGameLangParserRULE_literal)
	var _la int

	var _alt int

	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BoardGameLangParserT__4:
		localctx = NewListContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(85)
			p.Match(BoardGameLangParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62914594) != 0 {
			{
				p.SetState(86)
				p.expression(0)
			}
			p.SetState(91)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					{
						p.SetState(87)
						p.Match(BoardGameLangParserT__10)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					{
						p.SetState(88)
						p.expression(0)
					}

				}
				p.SetState(93)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}
			p.SetState(95)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == BoardGameLangParserT__10 {
				{
					p.SetState(94)
					p.Match(BoardGameLangParserT__10)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}

		}
		{
			p.SetState(99)
			p.Match(BoardGameLangParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BoardGameLangParserBOOL:
		localctx = NewBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(100)
			p.Match(BoardGameLangParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BoardGameLangParserNUMBER:
		localctx = NewNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(101)
			p.Match(BoardGameLangParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BoardGameLangParserSTRING:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(102)
			p.Match(BoardGameLangParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgListContext is an interface to support dynamic dispatch.
type IArgListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExpressionList() IExpressionListContext
	NamedArgs() INamedArgsContext

	// IsArgListContext differentiates from other interfaces.
	IsArgListContext()
}

type ArgListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgListContext() *ArgListContext {
	var p = new(ArgListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_argList
	return p
}

func InitEmptyArgListContext(p *ArgListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_argList
}

func (*ArgListContext) IsArgListContext() {}

func NewArgListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgListContext {
	var p = new(ArgListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BoardGameLangParserRULE_argList

	return p
}

func (s *ArgListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgListContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ArgListContext) NamedArgs() INamedArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamedArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamedArgsContext)
}

func (s *ArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterArgList(s)
	}
}

func (s *ArgListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitArgList(s)
	}
}

func (s *ArgListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitArgList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BoardGameLangParser) ArgList() (localctx IArgListContext) {
	localctx = NewArgListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BoardGameLangParserRULE_argList)
	var _la int

	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(105)
			p.Match(BoardGameLangParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(BoardGameLangParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(107)
			p.Match(BoardGameLangParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)
			p.ExpressionList()
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(109)
				p.Match(BoardGameLangParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(110)
				p.NamedArgs()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == BoardGameLangParserT__10 {
			{
				p.SetState(113)
				p.Match(BoardGameLangParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(116)
			p.Match(BoardGameLangParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(118)
			p.Match(BoardGameLangParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(119)
			p.NamedArgs()
		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == BoardGameLangParserT__10 {
			{
				p.SetState(120)
				p.Match(BoardGameLangParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(123)
			p.Match(BoardGameLangParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_expressionList
	return p
}

func InitEmptyExpressionListContext(p *ExpressionListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_expressionList
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BoardGameLangParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BoardGameLangParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BoardGameLangParserRULE_expressionList)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.expression(0)
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(128)
				p.Match(BoardGameLangParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(129)
				p.expression(0)
			}

		}
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INamedArgsContext is an interface to support dynamic dispatch.
type INamedArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNamedArg() []INamedArgContext
	NamedArg(i int) INamedArgContext

	// IsNamedArgsContext differentiates from other interfaces.
	IsNamedArgsContext()
}

type NamedArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedArgsContext() *NamedArgsContext {
	var p = new(NamedArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_namedArgs
	return p
}

func InitEmptyNamedArgsContext(p *NamedArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_namedArgs
}

func (*NamedArgsContext) IsNamedArgsContext() {}

func NewNamedArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedArgsContext {
	var p = new(NamedArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BoardGameLangParserRULE_namedArgs

	return p
}

func (s *NamedArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedArgsContext) AllNamedArg() []INamedArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INamedArgContext); ok {
			len++
		}
	}

	tst := make([]INamedArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INamedArgContext); ok {
			tst[i] = t.(INamedArgContext)
			i++
		}
	}

	return tst
}

func (s *NamedArgsContext) NamedArg(i int) INamedArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamedArgContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamedArgContext)
}

func (s *NamedArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterNamedArgs(s)
	}
}

func (s *NamedArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitNamedArgs(s)
	}
}

func (s *NamedArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitNamedArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BoardGameLangParser) NamedArgs() (localctx INamedArgsContext) {
	localctx = NewNamedArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BoardGameLangParserRULE_namedArgs)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.NamedArg()
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(136)
				p.Match(BoardGameLangParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(137)
				p.NamedArg()
			}

		}
		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INamedArgContext is an interface to support dynamic dispatch.
type INamedArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode
	Expression() IExpressionContext

	// IsNamedArgContext differentiates from other interfaces.
	IsNamedArgContext()
}

type NamedArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedArgContext() *NamedArgContext {
	var p = new(NamedArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_namedArg
	return p
}

func InitEmptyNamedArgContext(p *NamedArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_namedArg
}

func (*NamedArgContext) IsNamedArgContext() {}

func NewNamedArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedArgContext {
	var p = new(NamedArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BoardGameLangParserRULE_namedArg

	return p
}

func (s *NamedArgContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedArgContext) NAME() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserNAME, 0)
}

func (s *NamedArgContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NamedArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterNamedArg(s)
	}
}

func (s *NamedArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitNamedArg(s)
	}
}

func (s *NamedArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitNamedArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BoardGameLangParser) NamedArg() (localctx INamedArgContext) {
	localctx = NewNamedArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BoardGameLangParserRULE_namedArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(BoardGameLangParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Match(BoardGameLangParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_statementList
	return p
}

func InitEmptyStatementListContext(p *StatementListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_statementList
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BoardGameLangParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementListContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterStatementList(s)
	}
}

func (s *StatementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitStatementList(s)
	}
}

func (s *StatementListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitStatementList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BoardGameLangParser) StatementList() (localctx IStatementListContext) {
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BoardGameLangParserRULE_statementList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&63365154) != 0 {
		{
			p.SetState(147)
			p.Statement()
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BoardGameLangParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyAll(ctx *StatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssignmentContext struct {
	StatementContext
}

func NewAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentContext {
	var p = new(AssignmentContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) NAME() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserNAME, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnContext struct {
	StatementContext
}

func NewReturnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnContext {
	var p = new(ReturnContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterReturn(s)
	}
}

func (s *ReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitReturn(s)
	}
}

func (s *ReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

type ApplyContext struct {
	StatementContext
}

func NewApplyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ApplyContext {
	var p = new(ApplyContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ApplyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApplyContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ApplyContext) StatementList() IStatementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *ApplyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterApply(s)
	}
}

func (s *ApplyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitApply(s)
	}
}

func (s *ApplyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitApply(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionStatementContext struct {
	StatementContext
}

func NewExpressionStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ExpressionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitExpressionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type WhileContext struct {
	StatementContext
}

func NewWhileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileContext {
	var p = new(WhileContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *WhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileContext) StatementList() IStatementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *WhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterWhile(s)
	}
}

func (s *WhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitWhile(s)
	}
}

func (s *WhileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitWhile(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfContext struct {
	StatementContext
}

func NewIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfContext {
	var p = new(IfContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *IfContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfContext) AllStatementList() []IStatementListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementListContext); ok {
			len++
		}
	}

	tst := make([]IStatementListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementListContext); ok {
			tst[i] = t.(IStatementListContext)
			i++
		}
	}

	return tst
}

func (s *IfContext) StatementList(i int) IStatementListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *IfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterIf(s)
	}
}

func (s *IfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitIf(s)
	}
}

func (s *IfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitIf(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionDeclContext struct {
	StatementContext
}

func NewFunctionDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionDeclContext {
	var p = new(FunctionDeclContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *FunctionDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclContext) NAME() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserNAME, 0)
}

func (s *FunctionDeclContext) FunctionDeclArgs() IFunctionDeclArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDeclArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclArgsContext)
}

func (s *FunctionDeclContext) StatementList() IStatementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *FunctionDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterFunctionDecl(s)
	}
}

func (s *FunctionDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitFunctionDecl(s)
	}
}

func (s *FunctionDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitFunctionDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BoardGameLangParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BoardGameLangParserRULE_statement)
	var _la int

	var _alt int

	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		localctx = NewApplyContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.Match(BoardGameLangParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.expression(0)
		}
		{
			p.SetState(155)
			p.Match(BoardGameLangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)
			p.StatementList()
		}
		{
			p.SetState(157)
			p.Match(BoardGameLangParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewWhileContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(159)
			p.Match(BoardGameLangParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.expression(0)
		}
		{
			p.SetState(161)
			p.Match(BoardGameLangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)
			p.StatementList()
		}
		{
			p.SetState(163)
			p.Match(BoardGameLangParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(165)
			p.Match(BoardGameLangParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)
			p.expression(0)
		}
		{
			p.SetState(167)
			p.Match(BoardGameLangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.StatementList()
		}
		{
			p.SetState(169)
			p.Match(BoardGameLangParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(170)
					p.Match(BoardGameLangParserT__15)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(171)
					p.Match(BoardGameLangParserT__14)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(172)
					p.expression(0)
				}
				{
					p.SetState(173)
					p.Match(BoardGameLangParserT__2)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(174)
					p.StatementList()
				}
				{
					p.SetState(175)
					p.Match(BoardGameLangParserT__3)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(181)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == BoardGameLangParserT__15 {
			{
				p.SetState(182)
				p.Match(BoardGameLangParserT__15)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(183)
				p.Match(BoardGameLangParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(184)
				p.StatementList()
			}
			{
				p.SetState(185)
				p.Match(BoardGameLangParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 4:
		localctx = NewFunctionDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(189)
			p.Match(BoardGameLangParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
			p.Match(BoardGameLangParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)
			p.FunctionDeclArgs()
		}
		{
			p.SetState(192)
			p.Match(BoardGameLangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.StatementList()
		}
		{
			p.SetState(194)
			p.Match(BoardGameLangParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(196)
			p.Match(BoardGameLangParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.expression(0)
		}
		{
			p.SetState(198)
			p.Match(BoardGameLangParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewExpressionStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(200)
			p.expression(0)
		}
		{
			p.SetState(201)
			p.Match(BoardGameLangParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(203)
			p.Match(BoardGameLangParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)
			p.Match(BoardGameLangParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(205)
			p.expression(0)
		}
		{
			p.SetState(206)
			p.Match(BoardGameLangParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionDeclArgsContext is an interface to support dynamic dispatch.
type IFunctionDeclArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NameList() INameListContext
	NamedArgsDeclList() INamedArgsDeclListContext

	// IsFunctionDeclArgsContext differentiates from other interfaces.
	IsFunctionDeclArgsContext()
}

type FunctionDeclArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclArgsContext() *FunctionDeclArgsContext {
	var p = new(FunctionDeclArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_functionDeclArgs
	return p
}

func InitEmptyFunctionDeclArgsContext(p *FunctionDeclArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_functionDeclArgs
}

func (*FunctionDeclArgsContext) IsFunctionDeclArgsContext() {}

func NewFunctionDeclArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclArgsContext {
	var p = new(FunctionDeclArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BoardGameLangParserRULE_functionDeclArgs

	return p
}

func (s *FunctionDeclArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclArgsContext) NameList() INameListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameListContext)
}

func (s *FunctionDeclArgsContext) NamedArgsDeclList() INamedArgsDeclListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamedArgsDeclListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamedArgsDeclListContext)
}

func (s *FunctionDeclArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterFunctionDeclArgs(s)
	}
}

func (s *FunctionDeclArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitFunctionDeclArgs(s)
	}
}

func (s *FunctionDeclArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitFunctionDeclArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BoardGameLangParser) FunctionDeclArgs() (localctx IFunctionDeclArgsContext) {
	localctx = NewFunctionDeclArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BoardGameLangParserRULE_functionDeclArgs)
	var _la int

	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(210)
			p.Match(BoardGameLangParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(211)
			p.Match(BoardGameLangParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(212)
			p.Match(BoardGameLangParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.NameList()
		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(214)
				p.Match(BoardGameLangParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(215)
				p.NamedArgsDeclList()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(219)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == BoardGameLangParserT__10 {
			{
				p.SetState(218)
				p.Match(BoardGameLangParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(221)
			p.Match(BoardGameLangParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(223)
			p.Match(BoardGameLangParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)
			p.NamedArgsDeclList()
		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == BoardGameLangParserT__10 {
			{
				p.SetState(225)
				p.Match(BoardGameLangParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(228)
			p.Match(BoardGameLangParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INameListContext is an interface to support dynamic dispatch.
type INameListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNAME() []antlr.TerminalNode
	NAME(i int) antlr.TerminalNode

	// IsNameListContext differentiates from other interfaces.
	IsNameListContext()
}

type NameListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameListContext() *NameListContext {
	var p = new(NameListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_nameList
	return p
}

func InitEmptyNameListContext(p *NameListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_nameList
}

func (*NameListContext) IsNameListContext() {}

func NewNameListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameListContext {
	var p = new(NameListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BoardGameLangParserRULE_nameList

	return p
}

func (s *NameListContext) GetParser() antlr.Parser { return s.parser }

func (s *NameListContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(BoardGameLangParserNAME)
}

func (s *NameListContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserNAME, i)
}

func (s *NameListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterNameList(s)
	}
}

func (s *NameListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitNameList(s)
	}
}

func (s *NameListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitNameList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BoardGameLangParser) NameList() (localctx INameListContext) {
	localctx = NewNameListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BoardGameLangParserRULE_nameList)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		p.Match(BoardGameLangParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(233)
				p.Match(BoardGameLangParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(234)
				p.Match(BoardGameLangParserNAME)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INamedArgsDeclListContext is an interface to support dynamic dispatch.
type INamedArgsDeclListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNamedArgDecl() []INamedArgDeclContext
	NamedArgDecl(i int) INamedArgDeclContext

	// IsNamedArgsDeclListContext differentiates from other interfaces.
	IsNamedArgsDeclListContext()
}

type NamedArgsDeclListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedArgsDeclListContext() *NamedArgsDeclListContext {
	var p = new(NamedArgsDeclListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_namedArgsDeclList
	return p
}

func InitEmptyNamedArgsDeclListContext(p *NamedArgsDeclListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_namedArgsDeclList
}

func (*NamedArgsDeclListContext) IsNamedArgsDeclListContext() {}

func NewNamedArgsDeclListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedArgsDeclListContext {
	var p = new(NamedArgsDeclListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BoardGameLangParserRULE_namedArgsDeclList

	return p
}

func (s *NamedArgsDeclListContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedArgsDeclListContext) AllNamedArgDecl() []INamedArgDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INamedArgDeclContext); ok {
			len++
		}
	}

	tst := make([]INamedArgDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INamedArgDeclContext); ok {
			tst[i] = t.(INamedArgDeclContext)
			i++
		}
	}

	return tst
}

func (s *NamedArgsDeclListContext) NamedArgDecl(i int) INamedArgDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamedArgDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamedArgDeclContext)
}

func (s *NamedArgsDeclListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedArgsDeclListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedArgsDeclListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterNamedArgsDeclList(s)
	}
}

func (s *NamedArgsDeclListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitNamedArgsDeclList(s)
	}
}

func (s *NamedArgsDeclListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitNamedArgsDeclList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BoardGameLangParser) NamedArgsDeclList() (localctx INamedArgsDeclListContext) {
	localctx = NewNamedArgsDeclListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BoardGameLangParserRULE_namedArgsDeclList)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		p.NamedArgDecl()
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(241)
				p.Match(BoardGameLangParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(242)
				p.NamedArgDecl()
			}

		}
		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INamedArgDeclContext is an interface to support dynamic dispatch.
type INamedArgDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode
	Expression() IExpressionContext

	// IsNamedArgDeclContext differentiates from other interfaces.
	IsNamedArgDeclContext()
}

type NamedArgDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedArgDeclContext() *NamedArgDeclContext {
	var p = new(NamedArgDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_namedArgDecl
	return p
}

func InitEmptyNamedArgDeclContext(p *NamedArgDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_namedArgDecl
}

func (*NamedArgDeclContext) IsNamedArgDeclContext() {}

func NewNamedArgDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedArgDeclContext {
	var p = new(NamedArgDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BoardGameLangParserRULE_namedArgDecl

	return p
}

func (s *NamedArgDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedArgDeclContext) NAME() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserNAME, 0)
}

func (s *NamedArgDeclContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NamedArgDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedArgDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedArgDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterNamedArgDecl(s)
	}
}

func (s *NamedArgDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitNamedArgDecl(s)
	}
}

func (s *NamedArgDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitNamedArgDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BoardGameLangParser) NamedArgDecl() (localctx INamedArgDeclContext) {
	localctx = NewNamedArgDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BoardGameLangParserRULE_namedArgDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(248)
		p.Match(BoardGameLangParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.Match(BoardGameLangParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(250)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *BoardGameLangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 1:
		var t *MemberContext = nil
		if localctx != nil {
			t = localctx.(*MemberContext)
		}
		return p.Member_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *BoardGameLangParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 10)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *BoardGameLangParser) Member_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 9:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
