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
		"", "'('", "')'", "'{'", "'}'", "'['", "']'", "'^'", "'.'", "','", "'='",
		"';'", "'apply'", "'while'", "'if'", "'else'", "'fn'", "'return'", "",
		"", "", "", "", "", "'*'", "'/'", "'%'", "'+'", "'-'", "'..'", "'..='",
		"'=='", "'!='", "'<'", "'>'", "'<='", "'>='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "LINE_COMMENT", "MULTILINE_COMMENT", "BOOL", "NAME", "NUMBER", "STRING",
		"MUL", "DIV", "MOD", "ADD", "SUB", "RANGE", "RANGE_INCLUSIVE", "EQUAL",
		"NOT_EQUAL", "LESS_THAN", "GREATER_THAN", "LESS_EQUAL", "GREATER_EQUAL",
		"WS",
	}
	staticData.RuleNames = []string{
		"expression", "member", "literal", "list", "argList", "expressionList",
		"namedArgs", "namedArg", "statementList", "statement", "apply", "while",
		"if", "functionDecl", "functionDeclArgs", "nameList", "namedArgsDeclList",
		"namedArgDecl", "return", "assignment",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 37, 271, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 48, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 71, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5,
		0, 78, 8, 0, 10, 0, 12, 0, 81, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		5, 1, 89, 8, 1, 10, 1, 12, 1, 92, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 98,
		8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 104, 8, 3, 10, 3, 12, 3, 107, 9, 3,
		1, 3, 3, 3, 110, 8, 3, 3, 3, 112, 8, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 3, 4, 122, 8, 4, 1, 4, 3, 4, 125, 8, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 3, 4, 132, 8, 4, 1, 4, 1, 4, 3, 4, 136, 8, 4, 1, 5, 1, 5, 1,
		5, 5, 5, 141, 8, 5, 10, 5, 12, 5, 144, 9, 5, 1, 6, 1, 6, 1, 6, 5, 6, 149,
		8, 6, 10, 6, 12, 6, 152, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 5, 8, 159,
		8, 8, 10, 8, 12, 8, 162, 9, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 177, 8, 9, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		5, 12, 203, 8, 12, 10, 12, 12, 12, 206, 9, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 3, 12, 213, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 228, 8, 14, 1,
		14, 3, 14, 231, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 238, 8,
		14, 1, 14, 1, 14, 3, 14, 242, 8, 14, 1, 15, 1, 15, 1, 15, 5, 15, 247, 8,
		15, 10, 15, 12, 15, 250, 9, 15, 1, 16, 1, 16, 1, 16, 5, 16, 255, 8, 16,
		10, 16, 12, 16, 258, 9, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 0, 2, 0, 2, 20, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 0, 4, 1, 0, 24,
		26, 1, 0, 27, 28, 1, 0, 29, 30, 1, 0, 31, 36, 290, 0, 47, 1, 0, 0, 0, 2,
		82, 1, 0, 0, 0, 4, 97, 1, 0, 0, 0, 6, 99, 1, 0, 0, 0, 8, 135, 1, 0, 0,
		0, 10, 137, 1, 0, 0, 0, 12, 145, 1, 0, 0, 0, 14, 153, 1, 0, 0, 0, 16, 160,
		1, 0, 0, 0, 18, 176, 1, 0, 0, 0, 20, 178, 1, 0, 0, 0, 22, 184, 1, 0, 0,
		0, 24, 190, 1, 0, 0, 0, 26, 214, 1, 0, 0, 0, 28, 241, 1, 0, 0, 0, 30, 243,
		1, 0, 0, 0, 32, 251, 1, 0, 0, 0, 34, 259, 1, 0, 0, 0, 36, 263, 1, 0, 0,
		0, 38, 266, 1, 0, 0, 0, 40, 41, 6, 0, -1, 0, 41, 42, 5, 1, 0, 0, 42, 43,
		3, 0, 0, 0, 43, 44, 5, 2, 0, 0, 44, 48, 1, 0, 0, 0, 45, 48, 3, 2, 1, 0,
		46, 48, 3, 4, 2, 0, 47, 40, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 47, 46, 1,
		0, 0, 0, 48, 79, 1, 0, 0, 0, 49, 50, 10, 7, 0, 0, 50, 51, 5, 7, 0, 0, 51,
		78, 3, 0, 0, 8, 52, 53, 10, 6, 0, 0, 53, 54, 7, 0, 0, 0, 54, 78, 3, 0,
		0, 7, 55, 56, 10, 5, 0, 0, 56, 57, 7, 1, 0, 0, 57, 78, 3, 0, 0, 6, 58,
		59, 10, 4, 0, 0, 59, 60, 7, 2, 0, 0, 60, 78, 3, 0, 0, 5, 61, 62, 10, 3,
		0, 0, 62, 63, 7, 3, 0, 0, 63, 78, 3, 0, 0, 4, 64, 65, 10, 9, 0, 0, 65,
		70, 3, 8, 4, 0, 66, 67, 5, 3, 0, 0, 67, 68, 3, 16, 8, 0, 68, 69, 5, 4,
		0, 0, 69, 71, 1, 0, 0, 0, 70, 66, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 78,
		1, 0, 0, 0, 72, 73, 10, 8, 0, 0, 73, 74, 5, 5, 0, 0, 74, 75, 3, 10, 5,
		0, 75, 76, 5, 6, 0, 0, 76, 78, 1, 0, 0, 0, 77, 49, 1, 0, 0, 0, 77, 52,
		1, 0, 0, 0, 77, 55, 1, 0, 0, 0, 77, 58, 1, 0, 0, 0, 77, 61, 1, 0, 0, 0,
		77, 64, 1, 0, 0, 0, 77, 72, 1, 0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 77, 1,
		0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 1, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 82,
		83, 6, 1, -1, 0, 83, 84, 5, 21, 0, 0, 84, 90, 1, 0, 0, 0, 85, 86, 10, 2,
		0, 0, 86, 87, 5, 8, 0, 0, 87, 89, 5, 21, 0, 0, 88, 85, 1, 0, 0, 0, 89,
		92, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 3, 1, 0, 0,
		0, 92, 90, 1, 0, 0, 0, 93, 98, 3, 6, 3, 0, 94, 98, 5, 20, 0, 0, 95, 98,
		5, 22, 0, 0, 96, 98, 5, 23, 0, 0, 97, 93, 1, 0, 0, 0, 97, 94, 1, 0, 0,
		0, 97, 95, 1, 0, 0, 0, 97, 96, 1, 0, 0, 0, 98, 5, 1, 0, 0, 0, 99, 111,
		5, 5, 0, 0, 100, 105, 3, 0, 0, 0, 101, 102, 5, 9, 0, 0, 102, 104, 3, 0,
		0, 0, 103, 101, 1, 0, 0, 0, 104, 107, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0,
		105, 106, 1, 0, 0, 0, 106, 109, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 108,
		110, 5, 9, 0, 0, 109, 108, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 112,
		1, 0, 0, 0, 111, 100, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 113, 1, 0,
		0, 0, 113, 114, 5, 6, 0, 0, 114, 7, 1, 0, 0, 0, 115, 116, 5, 1, 0, 0, 116,
		136, 5, 2, 0, 0, 117, 118, 5, 1, 0, 0, 118, 121, 3, 10, 5, 0, 119, 120,
		5, 9, 0, 0, 120, 122, 3, 12, 6, 0, 121, 119, 1, 0, 0, 0, 121, 122, 1, 0,
		0, 0, 122, 124, 1, 0, 0, 0, 123, 125, 5, 9, 0, 0, 124, 123, 1, 0, 0, 0,
		124, 125, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 127, 5, 2, 0, 0, 127,
		136, 1, 0, 0, 0, 128, 129, 5, 1, 0, 0, 129, 131, 3, 12, 6, 0, 130, 132,
		5, 9, 0, 0, 131, 130, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 133, 1, 0,
		0, 0, 133, 134, 5, 2, 0, 0, 134, 136, 1, 0, 0, 0, 135, 115, 1, 0, 0, 0,
		135, 117, 1, 0, 0, 0, 135, 128, 1, 0, 0, 0, 136, 9, 1, 0, 0, 0, 137, 142,
		3, 0, 0, 0, 138, 139, 5, 9, 0, 0, 139, 141, 3, 0, 0, 0, 140, 138, 1, 0,
		0, 0, 141, 144, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0,
		143, 11, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 145, 150, 3, 14, 7, 0, 146,
		147, 5, 9, 0, 0, 147, 149, 3, 14, 7, 0, 148, 146, 1, 0, 0, 0, 149, 152,
		1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 13, 1, 0,
		0, 0, 152, 150, 1, 0, 0, 0, 153, 154, 5, 21, 0, 0, 154, 155, 5, 10, 0,
		0, 155, 156, 3, 0, 0, 0, 156, 15, 1, 0, 0, 0, 157, 159, 3, 18, 9, 0, 158,
		157, 1, 0, 0, 0, 159, 162, 1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 160, 161,
		1, 0, 0, 0, 161, 17, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 163, 177, 3, 20,
		10, 0, 164, 177, 3, 22, 11, 0, 165, 177, 3, 24, 12, 0, 166, 177, 3, 26,
		13, 0, 167, 168, 3, 36, 18, 0, 168, 169, 5, 11, 0, 0, 169, 177, 1, 0, 0,
		0, 170, 171, 3, 0, 0, 0, 171, 172, 5, 11, 0, 0, 172, 177, 1, 0, 0, 0, 173,
		174, 3, 38, 19, 0, 174, 175, 5, 11, 0, 0, 175, 177, 1, 0, 0, 0, 176, 163,
		1, 0, 0, 0, 176, 164, 1, 0, 0, 0, 176, 165, 1, 0, 0, 0, 176, 166, 1, 0,
		0, 0, 176, 167, 1, 0, 0, 0, 176, 170, 1, 0, 0, 0, 176, 173, 1, 0, 0, 0,
		177, 19, 1, 0, 0, 0, 178, 179, 5, 12, 0, 0, 179, 180, 3, 0, 0, 0, 180,
		181, 5, 3, 0, 0, 181, 182, 3, 16, 8, 0, 182, 183, 5, 4, 0, 0, 183, 21,
		1, 0, 0, 0, 184, 185, 5, 13, 0, 0, 185, 186, 3, 0, 0, 0, 186, 187, 5, 3,
		0, 0, 187, 188, 3, 16, 8, 0, 188, 189, 5, 4, 0, 0, 189, 23, 1, 0, 0, 0,
		190, 191, 5, 14, 0, 0, 191, 192, 3, 0, 0, 0, 192, 193, 5, 3, 0, 0, 193,
		194, 3, 16, 8, 0, 194, 204, 5, 4, 0, 0, 195, 196, 5, 15, 0, 0, 196, 197,
		5, 14, 0, 0, 197, 198, 3, 0, 0, 0, 198, 199, 5, 3, 0, 0, 199, 200, 3, 16,
		8, 0, 200, 201, 5, 4, 0, 0, 201, 203, 1, 0, 0, 0, 202, 195, 1, 0, 0, 0,
		203, 206, 1, 0, 0, 0, 204, 202, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205,
		212, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 207, 208, 5, 15, 0, 0, 208, 209,
		5, 3, 0, 0, 209, 210, 3, 16, 8, 0, 210, 211, 5, 4, 0, 0, 211, 213, 1, 0,
		0, 0, 212, 207, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 25, 1, 0, 0, 0,
		214, 215, 5, 16, 0, 0, 215, 216, 5, 21, 0, 0, 216, 217, 3, 28, 14, 0, 217,
		218, 5, 3, 0, 0, 218, 219, 3, 16, 8, 0, 219, 220, 5, 4, 0, 0, 220, 27,
		1, 0, 0, 0, 221, 222, 5, 1, 0, 0, 222, 242, 5, 2, 0, 0, 223, 224, 5, 1,
		0, 0, 224, 227, 3, 30, 15, 0, 225, 226, 5, 9, 0, 0, 226, 228, 3, 32, 16,
		0, 227, 225, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 230, 1, 0, 0, 0, 229,
		231, 5, 9, 0, 0, 230, 229, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 232,
		1, 0, 0, 0, 232, 233, 5, 2, 0, 0, 233, 242, 1, 0, 0, 0, 234, 235, 5, 1,
		0, 0, 235, 237, 3, 32, 16, 0, 236, 238, 5, 9, 0, 0, 237, 236, 1, 0, 0,
		0, 237, 238, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 240, 5, 2, 0, 0, 240,
		242, 1, 0, 0, 0, 241, 221, 1, 0, 0, 0, 241, 223, 1, 0, 0, 0, 241, 234,
		1, 0, 0, 0, 242, 29, 1, 0, 0, 0, 243, 248, 5, 21, 0, 0, 244, 245, 5, 9,
		0, 0, 245, 247, 5, 21, 0, 0, 246, 244, 1, 0, 0, 0, 247, 250, 1, 0, 0, 0,
		248, 246, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 31, 1, 0, 0, 0, 250, 248,
		1, 0, 0, 0, 251, 256, 3, 34, 17, 0, 252, 253, 5, 9, 0, 0, 253, 255, 3,
		34, 17, 0, 254, 252, 1, 0, 0, 0, 255, 258, 1, 0, 0, 0, 256, 254, 1, 0,
		0, 0, 256, 257, 1, 0, 0, 0, 257, 33, 1, 0, 0, 0, 258, 256, 1, 0, 0, 0,
		259, 260, 5, 21, 0, 0, 260, 261, 5, 10, 0, 0, 261, 262, 3, 0, 0, 0, 262,
		35, 1, 0, 0, 0, 263, 264, 5, 17, 0, 0, 264, 265, 3, 0, 0, 0, 265, 37, 1,
		0, 0, 0, 266, 267, 5, 21, 0, 0, 267, 268, 5, 10, 0, 0, 268, 269, 3, 0,
		0, 0, 269, 39, 1, 0, 0, 0, 25, 47, 70, 77, 79, 90, 97, 105, 109, 111, 121,
		124, 131, 135, 142, 150, 160, 176, 204, 212, 227, 230, 237, 241, 248, 256,
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
	BoardGameLangParserLINE_COMMENT      = 18
	BoardGameLangParserMULTILINE_COMMENT = 19
	BoardGameLangParserBOOL              = 20
	BoardGameLangParserNAME              = 21
	BoardGameLangParserNUMBER            = 22
	BoardGameLangParserSTRING            = 23
	BoardGameLangParserMUL               = 24
	BoardGameLangParserDIV               = 25
	BoardGameLangParserMOD               = 26
	BoardGameLangParserADD               = 27
	BoardGameLangParserSUB               = 28
	BoardGameLangParserRANGE             = 29
	BoardGameLangParserRANGE_INCLUSIVE   = 30
	BoardGameLangParserEQUAL             = 31
	BoardGameLangParserNOT_EQUAL         = 32
	BoardGameLangParserLESS_THAN         = 33
	BoardGameLangParserGREATER_THAN      = 34
	BoardGameLangParserLESS_EQUAL        = 35
	BoardGameLangParserGREATER_EQUAL     = 36
	BoardGameLangParserWS                = 37
)

// BoardGameLangParser rules.
const (
	BoardGameLangParserRULE_expression        = 0
	BoardGameLangParserRULE_member            = 1
	BoardGameLangParserRULE_literal           = 2
	BoardGameLangParserRULE_list              = 3
	BoardGameLangParserRULE_argList           = 4
	BoardGameLangParserRULE_expressionList    = 5
	BoardGameLangParserRULE_namedArgs         = 6
	BoardGameLangParserRULE_namedArg          = 7
	BoardGameLangParserRULE_statementList     = 8
	BoardGameLangParserRULE_statement         = 9
	BoardGameLangParserRULE_apply             = 10
	BoardGameLangParserRULE_while             = 11
	BoardGameLangParserRULE_if                = 12
	BoardGameLangParserRULE_functionDecl      = 13
	BoardGameLangParserRULE_functionDeclArgs  = 14
	BoardGameLangParserRULE_nameList          = 15
	BoardGameLangParserRULE_namedArgsDeclList = 16
	BoardGameLangParserRULE_namedArgDecl      = 17
	BoardGameLangParserRULE_return            = 18
	BoardGameLangParserRULE_assignment        = 19
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
	p.SetState(47)
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
			p.SetState(41)
			p.Match(BoardGameLangParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(42)
			p.expression(0)
		}
		{
			p.SetState(43)
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
			p.SetState(45)
			p.member(0)
		}

	case BoardGameLangParserT__4, BoardGameLangParserBOOL, BoardGameLangParserNUMBER, BoardGameLangParserSTRING:
		localctx = NewLiteralExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(46)
			p.Literal()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(79)
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
			p.SetState(77)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowerContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BoardGameLangParserRULE_expression)
				p.SetState(49)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(50)
					p.Match(BoardGameLangParserT__6)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(51)
					p.expression(8)
				}

			case 2:
				localctx = NewMulDivModContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BoardGameLangParserRULE_expression)
				p.SetState(52)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(53)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivModContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&117440512) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivModContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(54)
					p.expression(7)
				}

			case 3:
				localctx = NewAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BoardGameLangParserRULE_expression)
				p.SetState(55)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(56)

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
					p.SetState(57)
					p.expression(6)
				}

			case 4:
				localctx = NewRangeContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BoardGameLangParserRULE_expression)
				p.SetState(58)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(59)

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
					p.SetState(60)
					p.expression(5)
				}

			case 5:
				localctx = NewComparisonContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BoardGameLangParserRULE_expression)
				p.SetState(61)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(62)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ComparisonContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135291469824) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ComparisonContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(63)
					p.expression(4)
				}

			case 6:
				localctx = NewCallContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BoardGameLangParserRULE_expression)
				p.SetState(64)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(65)
					p.ArgList()
				}
				p.SetState(70)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(66)
						p.Match(BoardGameLangParserT__2)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					{
						p.SetState(67)
						p.StatementList()
					}
					{
						p.SetState(68)
						p.Match(BoardGameLangParserT__3)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				} else if p.HasError() { // JIM
					goto errorExit
				}

			case 7:
				localctx = NewIndexContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BoardGameLangParserRULE_expression)
				p.SetState(72)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(73)
					p.Match(BoardGameLangParserT__4)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(74)
					p.ExpressionList()
				}
				{
					p.SetState(75)
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
		p.SetState(81)
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

	// Getter signatures
	NAME() antlr.TerminalNode
	Member() IMemberContext

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

func (s *MemberContext) NAME() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserNAME, 0)
}

func (s *MemberContext) Member() IMemberContext {
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

func (s *MemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterMember(s)
	}
}

func (s *MemberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitMember(s)
	}
}

func (s *MemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitMember(s)

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
	{
		p.SetState(83)
		p.Match(BoardGameLangParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(90)
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
			localctx = NewMemberContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, BoardGameLangParserRULE_member)
			p.SetState(85)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(86)
				p.Match(BoardGameLangParserT__7)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(87)
				p.Match(BoardGameLangParserNAME)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(92)
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

	// Getter signatures
	List() IListContext
	BOOL() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	STRING() antlr.TerminalNode

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

func (s *LiteralContext) List() IListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *LiteralContext) BOOL() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserBOOL, 0)
}

func (s *LiteralContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserNUMBER, 0)
}

func (s *LiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(BoardGameLangParserSTRING, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BoardGameLangParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BoardGameLangParserRULE_literal)
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BoardGameLangParserT__4:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)
			p.List()
		}

	case BoardGameLangParserBOOL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(94)
			p.Match(BoardGameLangParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BoardGameLangParserNUMBER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(95)
			p.Match(BoardGameLangParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BoardGameLangParserSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(96)
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

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_list
	return p
}

func InitEmptyListContext(p *ListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_list
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BoardGameLangParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

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

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *BoardGameLangParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BoardGameLangParserRULE_list)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(BoardGameLangParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15728674) != 0 {
		{
			p.SetState(100)
			p.expression(0)
		}
		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(101)
					p.Match(BoardGameLangParserT__8)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(102)
					p.expression(0)
				}

			}
			p.SetState(107)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == BoardGameLangParserT__8 {
			{
				p.SetState(108)
				p.Match(BoardGameLangParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(113)
		p.Match(BoardGameLangParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
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
	p.EnterRule(localctx, 8, BoardGameLangParserRULE_argList)
	var _la int

	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(115)
			p.Match(BoardGameLangParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
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

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.Match(BoardGameLangParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)
			p.ExpressionList()
		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(119)
				p.Match(BoardGameLangParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(120)
				p.NamedArgs()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == BoardGameLangParserT__8 {
			{
				p.SetState(123)
				p.Match(BoardGameLangParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(126)
			p.Match(BoardGameLangParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(128)
			p.Match(BoardGameLangParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)
			p.NamedArgs()
		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == BoardGameLangParserT__8 {
			{
				p.SetState(130)
				p.Match(BoardGameLangParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(133)
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
	p.EnterRule(localctx, 10, BoardGameLangParserRULE_expressionList)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
		p.expression(0)
	}
	p.SetState(142)
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
				p.SetState(138)
				p.Match(BoardGameLangParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(139)
				p.expression(0)
			}

		}
		p.SetState(144)
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
	p.EnterRule(localctx, 12, BoardGameLangParserRULE_namedArgs)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.NamedArg()
	}
	p.SetState(150)
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
				p.SetState(146)
				p.Match(BoardGameLangParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(147)
				p.NamedArg()
			}

		}
		p.SetState(152)
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
	p.EnterRule(localctx, 14, BoardGameLangParserRULE_namedArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Match(BoardGameLangParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
		p.Match(BoardGameLangParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
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
	p.EnterRule(localctx, 16, BoardGameLangParserRULE_statementList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15953954) != 0 {
		{
			p.SetState(157)
			p.Statement()
		}

		p.SetState(162)
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

	// Getter signatures
	Apply() IApplyContext
	While() IWhileContext
	If_() IIfContext
	FunctionDecl() IFunctionDeclContext
	Return_() IReturnContext
	Expression() IExpressionContext
	Assignment() IAssignmentContext

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

func (s *StatementContext) Apply() IApplyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IApplyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IApplyContext)
}

func (s *StatementContext) While() IWhileContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileContext)
}

func (s *StatementContext) If_() IIfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfContext)
}

func (s *StatementContext) FunctionDecl() IFunctionDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclContext)
}

func (s *StatementContext) Return_() IReturnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnContext)
}

func (s *StatementContext) Expression() IExpressionContext {
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

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BoardGameLangListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BoardGameLangVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BoardGameLangParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BoardGameLangParserRULE_statement)
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(163)
			p.Apply()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(164)
			p.While()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(165)
			p.If_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(166)
			p.FunctionDecl()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(167)
			p.Return_()
		}
		{
			p.SetState(168)
			p.Match(BoardGameLangParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(170)
			p.expression(0)
		}
		{
			p.SetState(171)
			p.Match(BoardGameLangParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(173)
			p.Assignment()
		}
		{
			p.SetState(174)
			p.Match(BoardGameLangParserT__10)
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

// IApplyContext is an interface to support dynamic dispatch.
type IApplyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	StatementList() IStatementListContext

	// IsApplyContext differentiates from other interfaces.
	IsApplyContext()
}

type ApplyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyApplyContext() *ApplyContext {
	var p = new(ApplyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_apply
	return p
}

func InitEmptyApplyContext(p *ApplyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_apply
}

func (*ApplyContext) IsApplyContext() {}

func NewApplyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ApplyContext {
	var p = new(ApplyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BoardGameLangParserRULE_apply

	return p
}

func (s *ApplyContext) GetParser() antlr.Parser { return s.parser }

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

func (s *ApplyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApplyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *BoardGameLangParser) Apply() (localctx IApplyContext) {
	localctx = NewApplyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BoardGameLangParserRULE_apply)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Match(BoardGameLangParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(179)
		p.expression(0)
	}
	{
		p.SetState(180)
		p.Match(BoardGameLangParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
		p.StatementList()
	}
	{
		p.SetState(182)
		p.Match(BoardGameLangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
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

// IWhileContext is an interface to support dynamic dispatch.
type IWhileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	StatementList() IStatementListContext

	// IsWhileContext differentiates from other interfaces.
	IsWhileContext()
}

type WhileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileContext() *WhileContext {
	var p = new(WhileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_while
	return p
}

func InitEmptyWhileContext(p *WhileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_while
}

func (*WhileContext) IsWhileContext() {}

func NewWhileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileContext {
	var p = new(WhileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BoardGameLangParserRULE_while

	return p
}

func (s *WhileContext) GetParser() antlr.Parser { return s.parser }

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

func (s *WhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *BoardGameLangParser) While() (localctx IWhileContext) {
	localctx = NewWhileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BoardGameLangParserRULE_while)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Match(BoardGameLangParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(185)
		p.expression(0)
	}
	{
		p.SetState(186)
		p.Match(BoardGameLangParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(187)
		p.StatementList()
	}
	{
		p.SetState(188)
		p.Match(BoardGameLangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
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

// IIfContext is an interface to support dynamic dispatch.
type IIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllStatementList() []IStatementListContext
	StatementList(i int) IStatementListContext

	// IsIfContext differentiates from other interfaces.
	IsIfContext()
}

type IfContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfContext() *IfContext {
	var p = new(IfContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_if
	return p
}

func InitEmptyIfContext(p *IfContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_if
}

func (*IfContext) IsIfContext() {}

func NewIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfContext {
	var p = new(IfContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BoardGameLangParserRULE_if

	return p
}

func (s *IfContext) GetParser() antlr.Parser { return s.parser }

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

func (s *IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *BoardGameLangParser) If_() (localctx IIfContext) {
	localctx = NewIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BoardGameLangParserRULE_if)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Match(BoardGameLangParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(191)
		p.expression(0)
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
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(195)
				p.Match(BoardGameLangParserT__14)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(196)
				p.Match(BoardGameLangParserT__13)
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
				p.Match(BoardGameLangParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(199)
				p.StatementList()
			}
			{
				p.SetState(200)
				p.Match(BoardGameLangParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BoardGameLangParserT__14 {
		{
			p.SetState(207)
			p.Match(BoardGameLangParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(208)
			p.Match(BoardGameLangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(209)
			p.StatementList()
		}
		{
			p.SetState(210)
			p.Match(BoardGameLangParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IFunctionDeclContext is an interface to support dynamic dispatch.
type IFunctionDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode
	FunctionDeclArgs() IFunctionDeclArgsContext
	StatementList() IStatementListContext

	// IsFunctionDeclContext differentiates from other interfaces.
	IsFunctionDeclContext()
}

type FunctionDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclContext() *FunctionDeclContext {
	var p = new(FunctionDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_functionDecl
	return p
}

func InitEmptyFunctionDeclContext(p *FunctionDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_functionDecl
}

func (*FunctionDeclContext) IsFunctionDeclContext() {}

func NewFunctionDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclContext {
	var p = new(FunctionDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BoardGameLangParserRULE_functionDecl

	return p
}

func (s *FunctionDeclContext) GetParser() antlr.Parser { return s.parser }

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

func (s *FunctionDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *BoardGameLangParser) FunctionDecl() (localctx IFunctionDeclContext) {
	localctx = NewFunctionDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BoardGameLangParserRULE_functionDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.Match(BoardGameLangParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.Match(BoardGameLangParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(216)
		p.FunctionDeclArgs()
	}
	{
		p.SetState(217)
		p.Match(BoardGameLangParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(218)
		p.StatementList()
	}
	{
		p.SetState(219)
		p.Match(BoardGameLangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
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
	p.EnterRule(localctx, 28, BoardGameLangParserRULE_functionDeclArgs)
	var _la int

	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(221)
			p.Match(BoardGameLangParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(222)
			p.Match(BoardGameLangParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
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
			p.NameList()
		}
		p.SetState(227)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(225)
				p.Match(BoardGameLangParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(226)
				p.NamedArgsDeclList()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == BoardGameLangParserT__8 {
			{
				p.SetState(229)
				p.Match(BoardGameLangParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(232)
			p.Match(BoardGameLangParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(234)
			p.Match(BoardGameLangParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(235)
			p.NamedArgsDeclList()
		}
		p.SetState(237)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == BoardGameLangParserT__8 {
			{
				p.SetState(236)
				p.Match(BoardGameLangParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(239)
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
	p.EnterRule(localctx, 30, BoardGameLangParserRULE_nameList)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Match(BoardGameLangParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(248)
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
				p.SetState(244)
				p.Match(BoardGameLangParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(245)
				p.Match(BoardGameLangParserNAME)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(250)
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
	p.EnterRule(localctx, 32, BoardGameLangParserRULE_namedArgsDeclList)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.NamedArgDecl()
	}
	p.SetState(256)
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
				p.SetState(252)
				p.Match(BoardGameLangParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(253)
				p.NamedArgDecl()
			}

		}
		p.SetState(258)
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
	p.EnterRule(localctx, 34, BoardGameLangParserRULE_namedArgDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.Match(BoardGameLangParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(260)
		p.Match(BoardGameLangParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
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

// IReturnContext is an interface to support dynamic dispatch.
type IReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsReturnContext differentiates from other interfaces.
	IsReturnContext()
}

type ReturnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnContext() *ReturnContext {
	var p = new(ReturnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_return
	return p
}

func InitEmptyReturnContext(p *ReturnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_return
}

func (*ReturnContext) IsReturnContext() {}

func NewReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnContext {
	var p = new(ReturnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BoardGameLangParserRULE_return

	return p
}

func (s *ReturnContext) GetParser() antlr.Parser { return s.parser }

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

func (s *ReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *BoardGameLangParser) Return_() (localctx IReturnContext) {
	localctx = NewReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, BoardGameLangParserRULE_return)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)
		p.Match(BoardGameLangParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(264)
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

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode
	Expression() IExpressionContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BoardGameLangParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BoardGameLangParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

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

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *BoardGameLangParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, BoardGameLangParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.Match(BoardGameLangParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(267)
		p.Match(BoardGameLangParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(268)
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
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *BoardGameLangParser) Member_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
