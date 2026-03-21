// Code generated from C:/Users/Lenovo/Desktop/l Semestre 2024/Compiladores e interpretes/Semana 15/proyecto3/MiniGoParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package generated // MiniGoParser
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

type MiniGoParser struct {
	*antlr.BaseParser
}

var MiniGoParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func minigoparserParserInit() {
	staticData := &MiniGoParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'package'", "';'", "':'", "'var'", "'('", "')'", "'='", "'type'",
		"'func'", "','", "'.'", "'['", "']'", "'struct'", "'{'", "'}'", "'*'",
		"'/'", "'%'", "'<<'", "'>>'", "'&'", "'&^'", "'+'", "'-'", "'|'", "'^'",
		"'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'&&'", "'||'", "'!'",
		"'append'", "'len'", "'cap'", "'print'", "'println'", "'return'", "'break'",
		"'continue'", "'++'", "'--'", "':='", "'+='", "'&='", "'-='", "'|='",
		"'*='", "'^='", "'<<='", "'>>='", "'&^='", "'%='", "'/='", "'if'", "'else'",
		"'for'", "'switch'", "'case'", "'default'",
	}
	staticData.SymbolicNames = []string{
		"", "PACKAGE", "PyCOMA", "DOSP", "VAR", "PARIZQ", "PARDER", "IGUAL",
		"TYPE", "FUNC", "COMA", "PUNTO", "CORCHIZQ", "CORCHDER", "STRUCT", "LLAVEIZQ",
		"LLAVEDER", "MULTI", "DIV", "PORCENT", "MAYORM", "MEMORM", "AND", "BITCLE",
		"SUMA", "RESTA", "OR", "XOR", "IDENTICO", "DIFERENTE", "MAYOR", "MAYORIG",
		"MENOR", "MENORIG", "ANDD", "ORD", "EXCL", "APPEND", "LEN", "CAP", "PRINT",
		"PRINTLN", "RETURN", "BREAK", "CONTINUE", "SUMAD", "RESTAD", "ASSIGN",
		"INCREM", "ANDCOMP", "DECREME", "ASSIGNOR", "ASSIGNMULT", "ASSIGNXOR",
		"ASSIGMA", "ASSIGNME", "ASSIGNBIT", "ASSIGNPOR", "ASSINGDIV", "IF",
		"ELSE", "FOR", "SWITCH", "CASE", "DEFAULT", "ID", "INT", "FLOAT", "RUNE",
		"RAWSTRING", "INTERPRETEDSTRING", "COMENT", "COMENTM", "WS",
	}
	staticData.RuleNames = []string{
		"root", "topDeclarationList", "variableDecl", "innerVarDecls", "singleVarDecl",
		"singleVarDeclNoExps", "typeDecl", "innerTypeDecls", "singleTypeDecl",
		"funcDecl", "funcFrontDecl", "funcArgDecls", "declType", "sliceDeclType",
		"arrayDeclType", "structDeclType", "structMemDecls", "identifierList",
		"expression", "expressionList", "primaryExpression", "operand", "literal",
		"index", "arguments", "selector", "appendExpression", "lengthExpression",
		"capExpression", "statementList", "block", "statement", "simpleStatement",
		"assignmentStatement", "ifStatement", "loop", "switch", "expressionCaseClauseList",
		"expressionCaseClause", "expressionSwitchCase",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 73, 596, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 5, 1, 90, 8, 1, 10, 1, 12, 1, 93, 9, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 3, 2, 109, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 116, 8, 3, 10,
		3, 12, 3, 119, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 3, 4, 131, 8, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 150, 8,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 157, 8, 7, 10, 7, 12, 7, 160, 9,
		7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10,
		3, 10, 173, 8, 10, 1, 10, 1, 10, 3, 10, 177, 8, 10, 1, 11, 1, 11, 1, 11,
		5, 11, 182, 8, 11, 10, 11, 12, 11, 185, 9, 11, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 195, 8, 12, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 3, 15, 209,
		8, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 218, 8,
		16, 10, 16, 12, 16, 221, 9, 16, 1, 17, 1, 17, 1, 17, 5, 17, 226, 8, 17,
		10, 17, 12, 17, 229, 9, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 3, 18, 241, 8, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 5, 18, 300, 8, 18, 10, 18, 12, 18, 303, 9, 18, 1, 19, 1, 19, 1, 19,
		5, 19, 308, 8, 19, 10, 19, 12, 19, 311, 9, 19, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 3, 20, 318, 8, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		5, 20, 326, 8, 20, 10, 20, 12, 20, 329, 9, 20, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 3, 21, 337, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		3, 22, 344, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 3, 24, 352,
		8, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 29, 5, 29, 377, 8, 29, 10, 29, 12, 29, 380, 9, 29, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 3, 31, 389, 8, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 396, 8, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 3, 31, 402, 8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 426, 8, 31, 1, 32, 1, 32, 1, 32,
		3, 32, 431, 8, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 438, 8, 32,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 488, 8, 33, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 528, 8, 34, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 3, 35, 551, 8,
		35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 579, 8, 36, 1, 37, 1, 37,
		1, 37, 1, 37, 3, 37, 585, 8, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1,
		39, 1, 39, 3, 39, 594, 8, 39, 1, 39, 0, 2, 36, 40, 40, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
		48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 0, 1, 1,
		0, 45, 46, 656, 0, 80, 1, 0, 0, 0, 2, 91, 1, 0, 0, 0, 4, 108, 1, 0, 0,
		0, 6, 110, 1, 0, 0, 0, 8, 130, 1, 0, 0, 0, 10, 132, 1, 0, 0, 0, 12, 149,
		1, 0, 0, 0, 14, 151, 1, 0, 0, 0, 16, 161, 1, 0, 0, 0, 18, 164, 1, 0, 0,
		0, 20, 168, 1, 0, 0, 0, 22, 178, 1, 0, 0, 0, 24, 194, 1, 0, 0, 0, 26, 196,
		1, 0, 0, 0, 28, 200, 1, 0, 0, 0, 30, 205, 1, 0, 0, 0, 32, 212, 1, 0, 0,
		0, 34, 222, 1, 0, 0, 0, 36, 240, 1, 0, 0, 0, 38, 304, 1, 0, 0, 0, 40, 317,
		1, 0, 0, 0, 42, 336, 1, 0, 0, 0, 44, 343, 1, 0, 0, 0, 46, 345, 1, 0, 0,
		0, 48, 349, 1, 0, 0, 0, 50, 355, 1, 0, 0, 0, 52, 358, 1, 0, 0, 0, 54, 365,
		1, 0, 0, 0, 56, 370, 1, 0, 0, 0, 58, 378, 1, 0, 0, 0, 60, 381, 1, 0, 0,
		0, 62, 425, 1, 0, 0, 0, 64, 437, 1, 0, 0, 0, 66, 487, 1, 0, 0, 0, 68, 527,
		1, 0, 0, 0, 70, 550, 1, 0, 0, 0, 72, 578, 1, 0, 0, 0, 74, 584, 1, 0, 0,
		0, 76, 586, 1, 0, 0, 0, 78, 593, 1, 0, 0, 0, 80, 81, 5, 1, 0, 0, 81, 82,
		5, 65, 0, 0, 82, 83, 5, 2, 0, 0, 83, 84, 3, 2, 1, 0, 84, 85, 5, 0, 0, 1,
		85, 1, 1, 0, 0, 0, 86, 90, 3, 4, 2, 0, 87, 90, 3, 12, 6, 0, 88, 90, 3,
		18, 9, 0, 89, 86, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 88, 1, 0, 0, 0, 90,
		93, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 3, 1, 0, 0,
		0, 93, 91, 1, 0, 0, 0, 94, 95, 5, 4, 0, 0, 95, 96, 3, 8, 4, 0, 96, 97,
		5, 2, 0, 0, 97, 109, 1, 0, 0, 0, 98, 99, 5, 4, 0, 0, 99, 100, 5, 5, 0,
		0, 100, 101, 3, 6, 3, 0, 101, 102, 5, 6, 0, 0, 102, 103, 5, 2, 0, 0, 103,
		109, 1, 0, 0, 0, 104, 105, 5, 4, 0, 0, 105, 106, 5, 5, 0, 0, 106, 107,
		5, 6, 0, 0, 107, 109, 5, 2, 0, 0, 108, 94, 1, 0, 0, 0, 108, 98, 1, 0, 0,
		0, 108, 104, 1, 0, 0, 0, 109, 5, 1, 0, 0, 0, 110, 111, 3, 8, 4, 0, 111,
		117, 5, 2, 0, 0, 112, 113, 3, 8, 4, 0, 113, 114, 5, 2, 0, 0, 114, 116,
		1, 0, 0, 0, 115, 112, 1, 0, 0, 0, 116, 119, 1, 0, 0, 0, 117, 115, 1, 0,
		0, 0, 117, 118, 1, 0, 0, 0, 118, 7, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 120,
		121, 3, 34, 17, 0, 121, 122, 3, 24, 12, 0, 122, 123, 5, 7, 0, 0, 123, 124,
		3, 38, 19, 0, 124, 131, 1, 0, 0, 0, 125, 126, 3, 34, 17, 0, 126, 127, 5,
		7, 0, 0, 127, 128, 3, 38, 19, 0, 128, 131, 1, 0, 0, 0, 129, 131, 3, 10,
		5, 0, 130, 120, 1, 0, 0, 0, 130, 125, 1, 0, 0, 0, 130, 129, 1, 0, 0, 0,
		131, 9, 1, 0, 0, 0, 132, 133, 3, 34, 17, 0, 133, 134, 3, 24, 12, 0, 134,
		11, 1, 0, 0, 0, 135, 136, 5, 8, 0, 0, 136, 137, 3, 16, 8, 0, 137, 138,
		5, 2, 0, 0, 138, 150, 1, 0, 0, 0, 139, 140, 5, 8, 0, 0, 140, 141, 5, 5,
		0, 0, 141, 142, 3, 14, 7, 0, 142, 143, 5, 6, 0, 0, 143, 144, 5, 2, 0, 0,
		144, 150, 1, 0, 0, 0, 145, 146, 5, 8, 0, 0, 146, 147, 5, 5, 0, 0, 147,
		148, 5, 6, 0, 0, 148, 150, 5, 2, 0, 0, 149, 135, 1, 0, 0, 0, 149, 139,
		1, 0, 0, 0, 149, 145, 1, 0, 0, 0, 150, 13, 1, 0, 0, 0, 151, 152, 3, 16,
		8, 0, 152, 158, 5, 2, 0, 0, 153, 154, 3, 16, 8, 0, 154, 155, 5, 2, 0, 0,
		155, 157, 1, 0, 0, 0, 156, 153, 1, 0, 0, 0, 157, 160, 1, 0, 0, 0, 158,
		156, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 15, 1, 0, 0, 0, 160, 158, 1,
		0, 0, 0, 161, 162, 5, 65, 0, 0, 162, 163, 3, 24, 12, 0, 163, 17, 1, 0,
		0, 0, 164, 165, 3, 20, 10, 0, 165, 166, 3, 60, 30, 0, 166, 167, 5, 2, 0,
		0, 167, 19, 1, 0, 0, 0, 168, 169, 5, 9, 0, 0, 169, 170, 5, 65, 0, 0, 170,
		172, 5, 5, 0, 0, 171, 173, 3, 22, 11, 0, 172, 171, 1, 0, 0, 0, 172, 173,
		1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 176, 5, 6, 0, 0, 175, 177, 3, 24,
		12, 0, 176, 175, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 21, 1, 0, 0, 0,
		178, 183, 3, 10, 5, 0, 179, 180, 5, 10, 0, 0, 180, 182, 3, 10, 5, 0, 181,
		179, 1, 0, 0, 0, 182, 185, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 184,
		1, 0, 0, 0, 184, 23, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 186, 187, 5, 5,
		0, 0, 187, 188, 3, 24, 12, 0, 188, 189, 5, 6, 0, 0, 189, 195, 1, 0, 0,
		0, 190, 195, 5, 65, 0, 0, 191, 195, 3, 26, 13, 0, 192, 195, 3, 28, 14,
		0, 193, 195, 3, 30, 15, 0, 194, 186, 1, 0, 0, 0, 194, 190, 1, 0, 0, 0,
		194, 191, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 194, 193, 1, 0, 0, 0, 195,
		25, 1, 0, 0, 0, 196, 197, 5, 12, 0, 0, 197, 198, 5, 13, 0, 0, 198, 199,
		3, 24, 12, 0, 199, 27, 1, 0, 0, 0, 200, 201, 5, 12, 0, 0, 201, 202, 5,
		66, 0, 0, 202, 203, 5, 13, 0, 0, 203, 204, 3, 24, 12, 0, 204, 29, 1, 0,
		0, 0, 205, 206, 5, 14, 0, 0, 206, 208, 5, 15, 0, 0, 207, 209, 3, 32, 16,
		0, 208, 207, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210,
		211, 5, 16, 0, 0, 211, 31, 1, 0, 0, 0, 212, 213, 3, 10, 5, 0, 213, 219,
		5, 2, 0, 0, 214, 215, 3, 10, 5, 0, 215, 216, 5, 2, 0, 0, 216, 218, 1, 0,
		0, 0, 217, 214, 1, 0, 0, 0, 218, 221, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0,
		219, 220, 1, 0, 0, 0, 220, 33, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 222, 227,
		5, 65, 0, 0, 223, 224, 5, 10, 0, 0, 224, 226, 5, 65, 0, 0, 225, 223, 1,
		0, 0, 0, 226, 229, 1, 0, 0, 0, 227, 225, 1, 0, 0, 0, 227, 228, 1, 0, 0,
		0, 228, 35, 1, 0, 0, 0, 229, 227, 1, 0, 0, 0, 230, 231, 6, 18, -1, 0, 231,
		241, 3, 40, 20, 0, 232, 233, 5, 24, 0, 0, 233, 241, 3, 36, 18, 4, 234,
		235, 5, 25, 0, 0, 235, 241, 3, 36, 18, 3, 236, 237, 5, 36, 0, 0, 237, 241,
		3, 36, 18, 2, 238, 239, 5, 27, 0, 0, 239, 241, 3, 36, 18, 1, 240, 230,
		1, 0, 0, 0, 240, 232, 1, 0, 0, 0, 240, 234, 1, 0, 0, 0, 240, 236, 1, 0,
		0, 0, 240, 238, 1, 0, 0, 0, 241, 301, 1, 0, 0, 0, 242, 243, 10, 23, 0,
		0, 243, 244, 5, 17, 0, 0, 244, 300, 3, 36, 18, 24, 245, 246, 10, 22, 0,
		0, 246, 247, 5, 18, 0, 0, 247, 300, 3, 36, 18, 23, 248, 249, 10, 21, 0,
		0, 249, 250, 5, 19, 0, 0, 250, 300, 3, 36, 18, 22, 251, 252, 10, 20, 0,
		0, 252, 253, 5, 20, 0, 0, 253, 300, 3, 36, 18, 21, 254, 255, 10, 19, 0,
		0, 255, 256, 5, 21, 0, 0, 256, 300, 3, 36, 18, 20, 257, 258, 10, 18, 0,
		0, 258, 259, 5, 22, 0, 0, 259, 300, 3, 36, 18, 19, 260, 261, 10, 17, 0,
		0, 261, 262, 5, 23, 0, 0, 262, 300, 3, 36, 18, 18, 263, 264, 10, 16, 0,
		0, 264, 265, 5, 24, 0, 0, 265, 300, 3, 36, 18, 17, 266, 267, 10, 15, 0,
		0, 267, 268, 5, 25, 0, 0, 268, 300, 3, 36, 18, 16, 269, 270, 10, 14, 0,
		0, 270, 271, 5, 26, 0, 0, 271, 300, 3, 36, 18, 15, 272, 273, 10, 13, 0,
		0, 273, 274, 5, 27, 0, 0, 274, 300, 3, 36, 18, 14, 275, 276, 10, 12, 0,
		0, 276, 277, 5, 28, 0, 0, 277, 300, 3, 36, 18, 13, 278, 279, 10, 11, 0,
		0, 279, 280, 5, 29, 0, 0, 280, 300, 3, 36, 18, 12, 281, 282, 10, 10, 0,
		0, 282, 283, 5, 30, 0, 0, 283, 300, 3, 36, 18, 11, 284, 285, 10, 9, 0,
		0, 285, 286, 5, 31, 0, 0, 286, 300, 3, 36, 18, 10, 287, 288, 10, 8, 0,
		0, 288, 289, 5, 32, 0, 0, 289, 300, 3, 36, 18, 9, 290, 291, 10, 7, 0, 0,
		291, 292, 5, 33, 0, 0, 292, 300, 3, 36, 18, 8, 293, 294, 10, 6, 0, 0, 294,
		295, 5, 34, 0, 0, 295, 300, 3, 36, 18, 7, 296, 297, 10, 5, 0, 0, 297, 298,
		5, 35, 0, 0, 298, 300, 3, 36, 18, 6, 299, 242, 1, 0, 0, 0, 299, 245, 1,
		0, 0, 0, 299, 248, 1, 0, 0, 0, 299, 251, 1, 0, 0, 0, 299, 254, 1, 0, 0,
		0, 299, 257, 1, 0, 0, 0, 299, 260, 1, 0, 0, 0, 299, 263, 1, 0, 0, 0, 299,
		266, 1, 0, 0, 0, 299, 269, 1, 0, 0, 0, 299, 272, 1, 0, 0, 0, 299, 275,
		1, 0, 0, 0, 299, 278, 1, 0, 0, 0, 299, 281, 1, 0, 0, 0, 299, 284, 1, 0,
		0, 0, 299, 287, 1, 0, 0, 0, 299, 290, 1, 0, 0, 0, 299, 293, 1, 0, 0, 0,
		299, 296, 1, 0, 0, 0, 300, 303, 1, 0, 0, 0, 301, 299, 1, 0, 0, 0, 301,
		302, 1, 0, 0, 0, 302, 37, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 304, 309, 3,
		36, 18, 0, 305, 306, 5, 10, 0, 0, 306, 308, 3, 36, 18, 0, 307, 305, 1,
		0, 0, 0, 308, 311, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 309, 310, 1, 0, 0,
		0, 310, 39, 1, 0, 0, 0, 311, 309, 1, 0, 0, 0, 312, 313, 6, 20, -1, 0, 313,
		318, 3, 42, 21, 0, 314, 318, 3, 52, 26, 0, 315, 318, 3, 54, 27, 0, 316,
		318, 3, 56, 28, 0, 317, 312, 1, 0, 0, 0, 317, 314, 1, 0, 0, 0, 317, 315,
		1, 0, 0, 0, 317, 316, 1, 0, 0, 0, 318, 327, 1, 0, 0, 0, 319, 320, 10, 6,
		0, 0, 320, 326, 3, 50, 25, 0, 321, 322, 10, 5, 0, 0, 322, 326, 3, 46, 23,
		0, 323, 324, 10, 4, 0, 0, 324, 326, 3, 48, 24, 0, 325, 319, 1, 0, 0, 0,
		325, 321, 1, 0, 0, 0, 325, 323, 1, 0, 0, 0, 326, 329, 1, 0, 0, 0, 327,
		325, 1, 0, 0, 0, 327, 328, 1, 0, 0, 0, 328, 41, 1, 0, 0, 0, 329, 327, 1,
		0, 0, 0, 330, 337, 3, 44, 22, 0, 331, 337, 5, 65, 0, 0, 332, 333, 5, 5,
		0, 0, 333, 334, 3, 36, 18, 0, 334, 335, 5, 6, 0, 0, 335, 337, 1, 0, 0,
		0, 336, 330, 1, 0, 0, 0, 336, 331, 1, 0, 0, 0, 336, 332, 1, 0, 0, 0, 337,
		43, 1, 0, 0, 0, 338, 344, 5, 66, 0, 0, 339, 344, 5, 67, 0, 0, 340, 344,
		5, 68, 0, 0, 341, 344, 5, 69, 0, 0, 342, 344, 5, 70, 0, 0, 343, 338, 1,
		0, 0, 0, 343, 339, 1, 0, 0, 0, 343, 340, 1, 0, 0, 0, 343, 341, 1, 0, 0,
		0, 343, 342, 1, 0, 0, 0, 344, 45, 1, 0, 0, 0, 345, 346, 5, 12, 0, 0, 346,
		347, 3, 36, 18, 0, 347, 348, 5, 13, 0, 0, 348, 47, 1, 0, 0, 0, 349, 351,
		5, 5, 0, 0, 350, 352, 3, 38, 19, 0, 351, 350, 1, 0, 0, 0, 351, 352, 1,
		0, 0, 0, 352, 353, 1, 0, 0, 0, 353, 354, 5, 6, 0, 0, 354, 49, 1, 0, 0,
		0, 355, 356, 5, 11, 0, 0, 356, 357, 5, 65, 0, 0, 357, 51, 1, 0, 0, 0, 358,
		359, 5, 37, 0, 0, 359, 360, 5, 5, 0, 0, 360, 361, 3, 36, 18, 0, 361, 362,
		5, 10, 0, 0, 362, 363, 3, 36, 18, 0, 363, 364, 5, 6, 0, 0, 364, 53, 1,
		0, 0, 0, 365, 366, 5, 38, 0, 0, 366, 367, 5, 5, 0, 0, 367, 368, 3, 36,
		18, 0, 368, 369, 5, 6, 0, 0, 369, 55, 1, 0, 0, 0, 370, 371, 5, 39, 0, 0,
		371, 372, 5, 5, 0, 0, 372, 373, 3, 36, 18, 0, 373, 374, 5, 6, 0, 0, 374,
		57, 1, 0, 0, 0, 375, 377, 3, 62, 31, 0, 376, 375, 1, 0, 0, 0, 377, 380,
		1, 0, 0, 0, 378, 376, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 59, 1, 0,
		0, 0, 380, 378, 1, 0, 0, 0, 381, 382, 5, 15, 0, 0, 382, 383, 3, 58, 29,
		0, 383, 384, 5, 16, 0, 0, 384, 61, 1, 0, 0, 0, 385, 386, 5, 40, 0, 0, 386,
		388, 5, 5, 0, 0, 387, 389, 3, 38, 19, 0, 388, 387, 1, 0, 0, 0, 388, 389,
		1, 0, 0, 0, 389, 390, 1, 0, 0, 0, 390, 391, 5, 6, 0, 0, 391, 426, 5, 2,
		0, 0, 392, 393, 5, 41, 0, 0, 393, 395, 5, 5, 0, 0, 394, 396, 3, 38, 19,
		0, 395, 394, 1, 0, 0, 0, 395, 396, 1, 0, 0, 0, 396, 397, 1, 0, 0, 0, 397,
		398, 5, 6, 0, 0, 398, 426, 5, 2, 0, 0, 399, 401, 5, 42, 0, 0, 400, 402,
		3, 36, 18, 0, 401, 400, 1, 0, 0, 0, 401, 402, 1, 0, 0, 0, 402, 403, 1,
		0, 0, 0, 403, 426, 5, 2, 0, 0, 404, 405, 5, 43, 0, 0, 405, 426, 5, 2, 0,
		0, 406, 407, 5, 44, 0, 0, 407, 426, 5, 2, 0, 0, 408, 409, 3, 64, 32, 0,
		409, 410, 5, 2, 0, 0, 410, 426, 1, 0, 0, 0, 411, 412, 3, 60, 30, 0, 412,
		413, 5, 2, 0, 0, 413, 426, 1, 0, 0, 0, 414, 415, 3, 72, 36, 0, 415, 416,
		5, 2, 0, 0, 416, 426, 1, 0, 0, 0, 417, 418, 3, 68, 34, 0, 418, 419, 5,
		2, 0, 0, 419, 426, 1, 0, 0, 0, 420, 421, 3, 70, 35, 0, 421, 422, 5, 2,
		0, 0, 422, 426, 1, 0, 0, 0, 423, 426, 3, 12, 6, 0, 424, 426, 3, 4, 2, 0,
		425, 385, 1, 0, 0, 0, 425, 392, 1, 0, 0, 0, 425, 399, 1, 0, 0, 0, 425,
		404, 1, 0, 0, 0, 425, 406, 1, 0, 0, 0, 425, 408, 1, 0, 0, 0, 425, 411,
		1, 0, 0, 0, 425, 414, 1, 0, 0, 0, 425, 417, 1, 0, 0, 0, 425, 420, 1, 0,
		0, 0, 425, 423, 1, 0, 0, 0, 425, 424, 1, 0, 0, 0, 426, 63, 1, 0, 0, 0,
		427, 438, 1, 0, 0, 0, 428, 430, 3, 36, 18, 0, 429, 431, 7, 0, 0, 0, 430,
		429, 1, 0, 0, 0, 430, 431, 1, 0, 0, 0, 431, 438, 1, 0, 0, 0, 432, 438,
		3, 66, 33, 0, 433, 434, 3, 38, 19, 0, 434, 435, 5, 47, 0, 0, 435, 436,
		3, 38, 19, 0, 436, 438, 1, 0, 0, 0, 437, 427, 1, 0, 0, 0, 437, 428, 1,
		0, 0, 0, 437, 432, 1, 0, 0, 0, 437, 433, 1, 0, 0, 0, 438, 65, 1, 0, 0,
		0, 439, 440, 3, 38, 19, 0, 440, 441, 5, 7, 0, 0, 441, 442, 3, 38, 19, 0,
		442, 488, 1, 0, 0, 0, 443, 444, 3, 36, 18, 0, 444, 445, 5, 48, 0, 0, 445,
		446, 3, 36, 18, 0, 446, 488, 1, 0, 0, 0, 447, 448, 3, 36, 18, 0, 448, 449,
		5, 49, 0, 0, 449, 450, 3, 36, 18, 0, 450, 488, 1, 0, 0, 0, 451, 452, 3,
		36, 18, 0, 452, 453, 5, 50, 0, 0, 453, 454, 3, 36, 18, 0, 454, 488, 1,
		0, 0, 0, 455, 456, 3, 36, 18, 0, 456, 457, 5, 51, 0, 0, 457, 458, 3, 36,
		18, 0, 458, 488, 1, 0, 0, 0, 459, 460, 3, 36, 18, 0, 460, 461, 5, 52, 0,
		0, 461, 462, 3, 36, 18, 0, 462, 488, 1, 0, 0, 0, 463, 464, 3, 36, 18, 0,
		464, 465, 5, 53, 0, 0, 465, 466, 3, 36, 18, 0, 466, 488, 1, 0, 0, 0, 467,
		468, 3, 36, 18, 0, 468, 469, 5, 54, 0, 0, 469, 470, 3, 36, 18, 0, 470,
		488, 1, 0, 0, 0, 471, 472, 3, 36, 18, 0, 472, 473, 5, 55, 0, 0, 473, 474,
		3, 36, 18, 0, 474, 488, 1, 0, 0, 0, 475, 476, 3, 36, 18, 0, 476, 477, 5,
		56, 0, 0, 477, 478, 3, 36, 18, 0, 478, 488, 1, 0, 0, 0, 479, 480, 3, 36,
		18, 0, 480, 481, 5, 57, 0, 0, 481, 482, 3, 36, 18, 0, 482, 488, 1, 0, 0,
		0, 483, 484, 3, 36, 18, 0, 484, 485, 5, 58, 0, 0, 485, 486, 3, 36, 18,
		0, 486, 488, 1, 0, 0, 0, 487, 439, 1, 0, 0, 0, 487, 443, 1, 0, 0, 0, 487,
		447, 1, 0, 0, 0, 487, 451, 1, 0, 0, 0, 487, 455, 1, 0, 0, 0, 487, 459,
		1, 0, 0, 0, 487, 463, 1, 0, 0, 0, 487, 467, 1, 0, 0, 0, 487, 471, 1, 0,
		0, 0, 487, 475, 1, 0, 0, 0, 487, 479, 1, 0, 0, 0, 487, 483, 1, 0, 0, 0,
		488, 67, 1, 0, 0, 0, 489, 490, 5, 59, 0, 0, 490, 491, 3, 36, 18, 0, 491,
		492, 3, 60, 30, 0, 492, 528, 1, 0, 0, 0, 493, 494, 5, 59, 0, 0, 494, 495,
		3, 36, 18, 0, 495, 496, 3, 60, 30, 0, 496, 497, 5, 60, 0, 0, 497, 498,
		3, 68, 34, 0, 498, 528, 1, 0, 0, 0, 499, 500, 5, 59, 0, 0, 500, 501, 3,
		36, 18, 0, 501, 502, 3, 60, 30, 0, 502, 503, 5, 60, 0, 0, 503, 504, 3,
		60, 30, 0, 504, 528, 1, 0, 0, 0, 505, 506, 5, 59, 0, 0, 506, 507, 3, 64,
		32, 0, 507, 508, 5, 2, 0, 0, 508, 509, 3, 36, 18, 0, 509, 510, 3, 60, 30,
		0, 510, 528, 1, 0, 0, 0, 511, 512, 5, 59, 0, 0, 512, 513, 3, 64, 32, 0,
		513, 514, 5, 2, 0, 0, 514, 515, 3, 36, 18, 0, 515, 516, 3, 60, 30, 0, 516,
		517, 5, 60, 0, 0, 517, 518, 3, 68, 34, 0, 518, 528, 1, 0, 0, 0, 519, 520,
		5, 59, 0, 0, 520, 521, 3, 64, 32, 0, 521, 522, 5, 2, 0, 0, 522, 523, 3,
		36, 18, 0, 523, 524, 3, 60, 30, 0, 524, 525, 5, 60, 0, 0, 525, 526, 3,
		60, 30, 0, 526, 528, 1, 0, 0, 0, 527, 489, 1, 0, 0, 0, 527, 493, 1, 0,
		0, 0, 527, 499, 1, 0, 0, 0, 527, 505, 1, 0, 0, 0, 527, 511, 1, 0, 0, 0,
		527, 519, 1, 0, 0, 0, 528, 69, 1, 0, 0, 0, 529, 530, 5, 61, 0, 0, 530,
		551, 3, 60, 30, 0, 531, 532, 5, 61, 0, 0, 532, 533, 3, 36, 18, 0, 533,
		534, 3, 60, 30, 0, 534, 551, 1, 0, 0, 0, 535, 536, 5, 61, 0, 0, 536, 537,
		3, 64, 32, 0, 537, 538, 5, 2, 0, 0, 538, 539, 3, 36, 18, 0, 539, 540, 5,
		2, 0, 0, 540, 541, 3, 64, 32, 0, 541, 542, 3, 60, 30, 0, 542, 551, 1, 0,
		0, 0, 543, 544, 5, 61, 0, 0, 544, 545, 3, 64, 32, 0, 545, 546, 5, 2, 0,
		0, 546, 547, 5, 2, 0, 0, 547, 548, 3, 64, 32, 0, 548, 549, 3, 60, 30, 0,
		549, 551, 1, 0, 0, 0, 550, 529, 1, 0, 0, 0, 550, 531, 1, 0, 0, 0, 550,
		535, 1, 0, 0, 0, 550, 543, 1, 0, 0, 0, 551, 71, 1, 0, 0, 0, 552, 553, 5,
		62, 0, 0, 553, 554, 3, 64, 32, 0, 554, 555, 5, 2, 0, 0, 555, 556, 3, 36,
		18, 0, 556, 557, 5, 15, 0, 0, 557, 558, 3, 74, 37, 0, 558, 559, 5, 16,
		0, 0, 559, 579, 1, 0, 0, 0, 560, 561, 5, 62, 0, 0, 561, 562, 3, 36, 18,
		0, 562, 563, 5, 15, 0, 0, 563, 564, 3, 74, 37, 0, 564, 565, 5, 16, 0, 0,
		565, 579, 1, 0, 0, 0, 566, 567, 5, 62, 0, 0, 567, 568, 3, 64, 32, 0, 568,
		569, 5, 2, 0, 0, 569, 570, 5, 15, 0, 0, 570, 571, 3, 74, 37, 0, 571, 572,
		5, 16, 0, 0, 572, 579, 1, 0, 0, 0, 573, 574, 5, 62, 0, 0, 574, 575, 5,
		15, 0, 0, 575, 576, 3, 74, 37, 0, 576, 577, 5, 16, 0, 0, 577, 579, 1, 0,
		0, 0, 578, 552, 1, 0, 0, 0, 578, 560, 1, 0, 0, 0, 578, 566, 1, 0, 0, 0,
		578, 573, 1, 0, 0, 0, 579, 73, 1, 0, 0, 0, 580, 585, 1, 0, 0, 0, 581, 582,
		3, 76, 38, 0, 582, 583, 3, 74, 37, 0, 583, 585, 1, 0, 0, 0, 584, 580, 1,
		0, 0, 0, 584, 581, 1, 0, 0, 0, 585, 75, 1, 0, 0, 0, 586, 587, 3, 78, 39,
		0, 587, 588, 5, 3, 0, 0, 588, 589, 3, 58, 29, 0, 589, 77, 1, 0, 0, 0, 590,
		591, 5, 63, 0, 0, 591, 594, 3, 38, 19, 0, 592, 594, 5, 64, 0, 0, 593, 590,
		1, 0, 0, 0, 593, 592, 1, 0, 0, 0, 594, 79, 1, 0, 0, 0, 37, 89, 91, 108,
		117, 130, 149, 158, 172, 176, 183, 194, 208, 219, 227, 240, 299, 301, 309,
		317, 325, 327, 336, 343, 351, 378, 388, 395, 401, 425, 430, 437, 487, 527,
		550, 578, 584, 593,
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

// MiniGoParserInit initializes any static state used to implement MiniGoParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMiniGoParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MiniGoParserInit() {
	staticData := &MiniGoParserParserStaticData
	staticData.once.Do(minigoparserParserInit)
}

// NewMiniGoParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMiniGoParser(input antlr.TokenStream) *MiniGoParser {
	MiniGoParserInit()
	this := new(MiniGoParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MiniGoParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "MiniGoParser.g4"

	return this
}

// MiniGoParser tokens.
const (
	MiniGoParserEOF               = antlr.TokenEOF
	MiniGoParserPACKAGE           = 1
	MiniGoParserPyCOMA            = 2
	MiniGoParserDOSP              = 3
	MiniGoParserVAR               = 4
	MiniGoParserPARIZQ            = 5
	MiniGoParserPARDER            = 6
	MiniGoParserIGUAL             = 7
	MiniGoParserTYPE              = 8
	MiniGoParserFUNC              = 9
	MiniGoParserCOMA              = 10
	MiniGoParserPUNTO             = 11
	MiniGoParserCORCHIZQ          = 12
	MiniGoParserCORCHDER          = 13
	MiniGoParserSTRUCT            = 14
	MiniGoParserLLAVEIZQ          = 15
	MiniGoParserLLAVEDER          = 16
	MiniGoParserMULTI             = 17
	MiniGoParserDIV               = 18
	MiniGoParserPORCENT           = 19
	MiniGoParserMAYORM            = 20
	MiniGoParserMEMORM            = 21
	MiniGoParserAND               = 22
	MiniGoParserBITCLE            = 23
	MiniGoParserSUMA              = 24
	MiniGoParserRESTA             = 25
	MiniGoParserOR                = 26
	MiniGoParserXOR               = 27
	MiniGoParserIDENTICO          = 28
	MiniGoParserDIFERENTE         = 29
	MiniGoParserMAYOR             = 30
	MiniGoParserMAYORIG           = 31
	MiniGoParserMENOR             = 32
	MiniGoParserMENORIG           = 33
	MiniGoParserANDD              = 34
	MiniGoParserORD               = 35
	MiniGoParserEXCL              = 36
	MiniGoParserAPPEND            = 37
	MiniGoParserLEN               = 38
	MiniGoParserCAP               = 39
	MiniGoParserPRINT             = 40
	MiniGoParserPRINTLN           = 41
	MiniGoParserRETURN            = 42
	MiniGoParserBREAK             = 43
	MiniGoParserCONTINUE          = 44
	MiniGoParserSUMAD             = 45
	MiniGoParserRESTAD            = 46
	MiniGoParserASSIGN            = 47
	MiniGoParserINCREM            = 48
	MiniGoParserANDCOMP           = 49
	MiniGoParserDECREME           = 50
	MiniGoParserASSIGNOR          = 51
	MiniGoParserASSIGNMULT        = 52
	MiniGoParserASSIGNXOR         = 53
	MiniGoParserASSIGMA           = 54
	MiniGoParserASSIGNME          = 55
	MiniGoParserASSIGNBIT         = 56
	MiniGoParserASSIGNPOR         = 57
	MiniGoParserASSINGDIV         = 58
	MiniGoParserIF                = 59
	MiniGoParserELSE              = 60
	MiniGoParserFOR               = 61
	MiniGoParserSWITCH            = 62
	MiniGoParserCASE              = 63
	MiniGoParserDEFAULT           = 64
	MiniGoParserID                = 65
	MiniGoParserINT               = 66
	MiniGoParserFLOAT             = 67
	MiniGoParserRUNE              = 68
	MiniGoParserRAWSTRING         = 69
	MiniGoParserINTERPRETEDSTRING = 70
	MiniGoParserCOMENT            = 71
	MiniGoParserCOMENTM           = 72
	MiniGoParserWS                = 73
)

// MiniGoParser rules.
const (
	MiniGoParserRULE_root                     = 0
	MiniGoParserRULE_topDeclarationList       = 1
	MiniGoParserRULE_variableDecl             = 2
	MiniGoParserRULE_innerVarDecls            = 3
	MiniGoParserRULE_singleVarDecl            = 4
	MiniGoParserRULE_singleVarDeclNoExps      = 5
	MiniGoParserRULE_typeDecl                 = 6
	MiniGoParserRULE_innerTypeDecls           = 7
	MiniGoParserRULE_singleTypeDecl           = 8
	MiniGoParserRULE_funcDecl                 = 9
	MiniGoParserRULE_funcFrontDecl            = 10
	MiniGoParserRULE_funcArgDecls             = 11
	MiniGoParserRULE_declType                 = 12
	MiniGoParserRULE_sliceDeclType            = 13
	MiniGoParserRULE_arrayDeclType            = 14
	MiniGoParserRULE_structDeclType           = 15
	MiniGoParserRULE_structMemDecls           = 16
	MiniGoParserRULE_identifierList           = 17
	MiniGoParserRULE_expression               = 18
	MiniGoParserRULE_expressionList           = 19
	MiniGoParserRULE_primaryExpression        = 20
	MiniGoParserRULE_operand                  = 21
	MiniGoParserRULE_literal                  = 22
	MiniGoParserRULE_index                    = 23
	MiniGoParserRULE_arguments                = 24
	MiniGoParserRULE_selector                 = 25
	MiniGoParserRULE_appendExpression         = 26
	MiniGoParserRULE_lengthExpression         = 27
	MiniGoParserRULE_capExpression            = 28
	MiniGoParserRULE_statementList            = 29
	MiniGoParserRULE_block                    = 30
	MiniGoParserRULE_statement                = 31
	MiniGoParserRULE_simpleStatement          = 32
	MiniGoParserRULE_assignmentStatement      = 33
	MiniGoParserRULE_ifStatement              = 34
	MiniGoParserRULE_loop                     = 35
	MiniGoParserRULE_switch                   = 36
	MiniGoParserRULE_expressionCaseClauseList = 37
	MiniGoParserRULE_expressionCaseClause     = 38
	MiniGoParserRULE_expressionSwitchCase     = 39
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_root
	return p
}

func InitEmptyRootContext(p *RootContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_root
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) CopyAll(ctx *RootContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RootASTContext struct {
	RootContext
}

func NewRootASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RootASTContext {
	var p = new(RootASTContext)

	InitEmptyRootContext(&p.RootContext)
	p.parser = parser
	p.CopyAll(ctx.(*RootContext))

	return p
}

func (s *RootASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootASTContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPACKAGE, 0)
}

func (s *RootASTContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniGoParserID, 0)
}

func (s *RootASTContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, 0)
}

func (s *RootASTContext) TopDeclarationList() ITopDeclarationListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopDeclarationListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITopDeclarationListContext)
}

func (s *RootASTContext) EOF() antlr.TerminalNode {
	return s.GetToken(MiniGoParserEOF, 0)
}

func (s *RootASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitRootAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MiniGoParserRULE_root)
	localctx = NewRootASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.Match(MiniGoParserPACKAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
		p.Match(MiniGoParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.Match(MiniGoParserPyCOMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
		p.TopDeclarationList()
	}
	{
		p.SetState(84)
		p.Match(MiniGoParserEOF)
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

// ITopDeclarationListContext is an interface to support dynamic dispatch.
type ITopDeclarationListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTopDeclarationListContext differentiates from other interfaces.
	IsTopDeclarationListContext()
}

type TopDeclarationListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopDeclarationListContext() *TopDeclarationListContext {
	var p = new(TopDeclarationListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_topDeclarationList
	return p
}

func InitEmptyTopDeclarationListContext(p *TopDeclarationListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_topDeclarationList
}

func (*TopDeclarationListContext) IsTopDeclarationListContext() {}

func NewTopDeclarationListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopDeclarationListContext {
	var p = new(TopDeclarationListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_topDeclarationList

	return p
}

func (s *TopDeclarationListContext) GetParser() antlr.Parser { return s.parser }

func (s *TopDeclarationListContext) CopyAll(ctx *TopDeclarationListContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TopDeclarationListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopDeclarationListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TopDeclarationListASTContext struct {
	TopDeclarationListContext
}

func NewTopDeclarationListASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TopDeclarationListASTContext {
	var p = new(TopDeclarationListASTContext)

	InitEmptyTopDeclarationListContext(&p.TopDeclarationListContext)
	p.parser = parser
	p.CopyAll(ctx.(*TopDeclarationListContext))

	return p
}

func (s *TopDeclarationListASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopDeclarationListASTContext) AllVariableDecl() []IVariableDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableDeclContext); ok {
			len++
		}
	}

	tst := make([]IVariableDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableDeclContext); ok {
			tst[i] = t.(IVariableDeclContext)
			i++
		}
	}

	return tst
}

func (s *TopDeclarationListASTContext) VariableDecl(i int) IVariableDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclContext); ok {
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

	return t.(IVariableDeclContext)
}

func (s *TopDeclarationListASTContext) AllTypeDecl() []ITypeDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeDeclContext); ok {
			len++
		}
	}

	tst := make([]ITypeDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeDeclContext); ok {
			tst[i] = t.(ITypeDeclContext)
			i++
		}
	}

	return tst
}

func (s *TopDeclarationListASTContext) TypeDecl(i int) ITypeDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDeclContext); ok {
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

	return t.(ITypeDeclContext)
}

func (s *TopDeclarationListASTContext) AllFuncDecl() []IFuncDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncDeclContext); ok {
			len++
		}
	}

	tst := make([]IFuncDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncDeclContext); ok {
			tst[i] = t.(IFuncDeclContext)
			i++
		}
	}

	return tst
}

func (s *TopDeclarationListASTContext) FuncDecl(i int) IFuncDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDeclContext); ok {
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

	return t.(IFuncDeclContext)
}

func (s *TopDeclarationListASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitTopDeclarationListAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) TopDeclarationList() (localctx ITopDeclarationListContext) {
	localctx = NewTopDeclarationListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MiniGoParserRULE_topDeclarationList)
	var _la int

	localctx = NewTopDeclarationListASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&784) != 0 {
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case MiniGoParserVAR:
			{
				p.SetState(86)
				p.VariableDecl()
			}

		case MiniGoParserTYPE:
			{
				p.SetState(87)
				p.TypeDecl()
			}

		case MiniGoParserFUNC:
			{
				p.SetState(88)
				p.FuncDecl()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(93)
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

// IVariableDeclContext is an interface to support dynamic dispatch.
type IVariableDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVariableDeclContext differentiates from other interfaces.
	IsVariableDeclContext()
}

type VariableDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclContext() *VariableDeclContext {
	var p = new(VariableDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_variableDecl
	return p
}

func InitEmptyVariableDeclContext(p *VariableDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_variableDecl
}

func (*VariableDeclContext) IsVariableDeclContext() {}

func NewVariableDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclContext {
	var p = new(VariableDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_variableDecl

	return p
}

func (s *VariableDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclContext) CopyAll(ctx *VariableDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VariableDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VariableDeclSingASTContext struct {
	VariableDeclContext
}

func NewVariableDeclSingASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableDeclSingASTContext {
	var p = new(VariableDeclSingASTContext)

	InitEmptyVariableDeclContext(&p.VariableDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclContext))

	return p
}

func (s *VariableDeclSingASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclSingASTContext) VAR() antlr.TerminalNode {
	return s.GetToken(MiniGoParserVAR, 0)
}

func (s *VariableDeclSingASTContext) SingleVarDecl() ISingleVarDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleVarDeclContext)
}

func (s *VariableDeclSingASTContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, 0)
}

func (s *VariableDeclSingASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitVariableDeclSingAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type VariableDeclPaRASTContext struct {
	VariableDeclContext
}

func NewVariableDeclPaRASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableDeclPaRASTContext {
	var p = new(VariableDeclPaRASTContext)

	InitEmptyVariableDeclContext(&p.VariableDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclContext))

	return p
}

func (s *VariableDeclPaRASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclPaRASTContext) VAR() antlr.TerminalNode {
	return s.GetToken(MiniGoParserVAR, 0)
}

func (s *VariableDeclPaRASTContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARIZQ, 0)
}

func (s *VariableDeclPaRASTContext) InnerVarDecls() IInnerVarDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInnerVarDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInnerVarDeclsContext)
}

func (s *VariableDeclPaRASTContext) PARDER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARDER, 0)
}

func (s *VariableDeclPaRASTContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, 0)
}

func (s *VariableDeclPaRASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitVariableDeclPaRAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type VariableDeclNullASTContext struct {
	VariableDeclContext
}

func NewVariableDeclNullASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableDeclNullASTContext {
	var p = new(VariableDeclNullASTContext)

	InitEmptyVariableDeclContext(&p.VariableDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclContext))

	return p
}

func (s *VariableDeclNullASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclNullASTContext) VAR() antlr.TerminalNode {
	return s.GetToken(MiniGoParserVAR, 0)
}

func (s *VariableDeclNullASTContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARIZQ, 0)
}

func (s *VariableDeclNullASTContext) PARDER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARDER, 0)
}

func (s *VariableDeclNullASTContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, 0)
}

func (s *VariableDeclNullASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitVariableDeclNullAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) VariableDecl() (localctx IVariableDeclContext) {
	localctx = NewVariableDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MiniGoParserRULE_variableDecl)
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVariableDeclSingASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(94)
			p.Match(MiniGoParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.SingleVarDecl()
		}
		{
			p.SetState(96)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewVariableDeclPaRASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(98)
			p.Match(MiniGoParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.Match(MiniGoParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
			p.InnerVarDecls()
		}
		{
			p.SetState(101)
			p.Match(MiniGoParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(102)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewVariableDeclNullASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(104)
			p.Match(MiniGoParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.Match(MiniGoParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(MiniGoParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.Match(MiniGoParserPyCOMA)
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

// IInnerVarDeclsContext is an interface to support dynamic dispatch.
type IInnerVarDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsInnerVarDeclsContext differentiates from other interfaces.
	IsInnerVarDeclsContext()
}

type InnerVarDeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInnerVarDeclsContext() *InnerVarDeclsContext {
	var p = new(InnerVarDeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_innerVarDecls
	return p
}

func InitEmptyInnerVarDeclsContext(p *InnerVarDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_innerVarDecls
}

func (*InnerVarDeclsContext) IsInnerVarDeclsContext() {}

func NewInnerVarDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InnerVarDeclsContext {
	var p = new(InnerVarDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_innerVarDecls

	return p
}

func (s *InnerVarDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *InnerVarDeclsContext) CopyAll(ctx *InnerVarDeclsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *InnerVarDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerVarDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type InnerVarDeclsASTContext struct {
	InnerVarDeclsContext
}

func NewInnerVarDeclsASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InnerVarDeclsASTContext {
	var p = new(InnerVarDeclsASTContext)

	InitEmptyInnerVarDeclsContext(&p.InnerVarDeclsContext)
	p.parser = parser
	p.CopyAll(ctx.(*InnerVarDeclsContext))

	return p
}

func (s *InnerVarDeclsASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerVarDeclsASTContext) AllSingleVarDecl() []ISingleVarDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleVarDeclContext); ok {
			len++
		}
	}

	tst := make([]ISingleVarDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleVarDeclContext); ok {
			tst[i] = t.(ISingleVarDeclContext)
			i++
		}
	}

	return tst
}

func (s *InnerVarDeclsASTContext) SingleVarDecl(i int) ISingleVarDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclContext); ok {
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

	return t.(ISingleVarDeclContext)
}

func (s *InnerVarDeclsASTContext) AllPyCOMA() []antlr.TerminalNode {
	return s.GetTokens(MiniGoParserPyCOMA)
}

func (s *InnerVarDeclsASTContext) PyCOMA(i int) antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, i)
}

func (s *InnerVarDeclsASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitInnerVarDeclsAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) InnerVarDecls() (localctx IInnerVarDeclsContext) {
	localctx = NewInnerVarDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MiniGoParserRULE_innerVarDecls)
	var _la int

	localctx = NewInnerVarDeclsASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.SingleVarDecl()
	}
	{
		p.SetState(111)
		p.Match(MiniGoParserPyCOMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MiniGoParserID {
		{
			p.SetState(112)
			p.SingleVarDecl()
		}
		{
			p.SetState(113)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(119)
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

// ISingleVarDeclContext is an interface to support dynamic dispatch.
type ISingleVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSingleVarDeclContext differentiates from other interfaces.
	IsSingleVarDeclContext()
}

type SingleVarDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleVarDeclContext() *SingleVarDeclContext {
	var p = new(SingleVarDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_singleVarDecl
	return p
}

func InitEmptySingleVarDeclContext(p *SingleVarDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_singleVarDecl
}

func (*SingleVarDeclContext) IsSingleVarDeclContext() {}

func NewSingleVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleVarDeclContext {
	var p = new(SingleVarDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_singleVarDecl

	return p
}

func (s *SingleVarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleVarDeclContext) CopyAll(ctx *SingleVarDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SingleVarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SingleVarDeclNoExpsSVDASTContext struct {
	SingleVarDeclContext
}

func NewSingleVarDeclNoExpsSVDASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleVarDeclNoExpsSVDASTContext {
	var p = new(SingleVarDeclNoExpsSVDASTContext)

	InitEmptySingleVarDeclContext(&p.SingleVarDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleVarDeclContext))

	return p
}

func (s *SingleVarDeclNoExpsSVDASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarDeclNoExpsSVDASTContext) SingleVarDeclNoExps() ISingleVarDeclNoExpsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleVarDeclNoExpsContext)
}

func (s *SingleVarDeclNoExpsSVDASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSingleVarDeclNoExpsSVDAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SingleVarDeclIdenASTContext struct {
	SingleVarDeclContext
}

func NewSingleVarDeclIdenASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleVarDeclIdenASTContext {
	var p = new(SingleVarDeclIdenASTContext)

	InitEmptySingleVarDeclContext(&p.SingleVarDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleVarDeclContext))

	return p
}

func (s *SingleVarDeclIdenASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarDeclIdenASTContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *SingleVarDeclIdenASTContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *SingleVarDeclIdenASTContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(MiniGoParserIGUAL, 0)
}

func (s *SingleVarDeclIdenASTContext) ExpressionList() IExpressionListContext {
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

func (s *SingleVarDeclIdenASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSingleVarDeclIdenAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SingleVarDeclIASTContext struct {
	SingleVarDeclContext
}

func NewSingleVarDeclIASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleVarDeclIASTContext {
	var p = new(SingleVarDeclIASTContext)

	InitEmptySingleVarDeclContext(&p.SingleVarDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleVarDeclContext))

	return p
}

func (s *SingleVarDeclIASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarDeclIASTContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *SingleVarDeclIASTContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(MiniGoParserIGUAL, 0)
}

func (s *SingleVarDeclIASTContext) ExpressionList() IExpressionListContext {
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

func (s *SingleVarDeclIASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSingleVarDeclIAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) SingleVarDecl() (localctx ISingleVarDeclContext) {
	localctx = NewSingleVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MiniGoParserRULE_singleVarDecl)
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSingleVarDeclIdenASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.IdentifierList()
		}
		{
			p.SetState(121)
			p.DeclType()
		}
		{
			p.SetState(122)
			p.Match(MiniGoParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.ExpressionList()
		}

	case 2:
		localctx = NewSingleVarDeclIASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			p.IdentifierList()
		}
		{
			p.SetState(126)
			p.Match(MiniGoParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.ExpressionList()
		}

	case 3:
		localctx = NewSingleVarDeclNoExpsSVDASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(129)
			p.SingleVarDeclNoExps()
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

// ISingleVarDeclNoExpsContext is an interface to support dynamic dispatch.
type ISingleVarDeclNoExpsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSingleVarDeclNoExpsContext differentiates from other interfaces.
	IsSingleVarDeclNoExpsContext()
}

type SingleVarDeclNoExpsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleVarDeclNoExpsContext() *SingleVarDeclNoExpsContext {
	var p = new(SingleVarDeclNoExpsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_singleVarDeclNoExps
	return p
}

func InitEmptySingleVarDeclNoExpsContext(p *SingleVarDeclNoExpsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_singleVarDeclNoExps
}

func (*SingleVarDeclNoExpsContext) IsSingleVarDeclNoExpsContext() {}

func NewSingleVarDeclNoExpsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleVarDeclNoExpsContext {
	var p = new(SingleVarDeclNoExpsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_singleVarDeclNoExps

	return p
}

func (s *SingleVarDeclNoExpsContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleVarDeclNoExpsContext) CopyAll(ctx *SingleVarDeclNoExpsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SingleVarDeclNoExpsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarDeclNoExpsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SingleVarDeclNoExpsASTContext struct {
	SingleVarDeclNoExpsContext
}

func NewSingleVarDeclNoExpsASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleVarDeclNoExpsASTContext {
	var p = new(SingleVarDeclNoExpsASTContext)

	InitEmptySingleVarDeclNoExpsContext(&p.SingleVarDeclNoExpsContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleVarDeclNoExpsContext))

	return p
}

func (s *SingleVarDeclNoExpsASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarDeclNoExpsASTContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *SingleVarDeclNoExpsASTContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *SingleVarDeclNoExpsASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSingleVarDeclNoExpsAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) SingleVarDeclNoExps() (localctx ISingleVarDeclNoExpsContext) {
	localctx = NewSingleVarDeclNoExpsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MiniGoParserRULE_singleVarDeclNoExps)
	localctx = NewSingleVarDeclNoExpsASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.IdentifierList()
	}
	{
		p.SetState(133)
		p.DeclType()
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

// ITypeDeclContext is an interface to support dynamic dispatch.
type ITypeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTypeDeclContext differentiates from other interfaces.
	IsTypeDeclContext()
}

type TypeDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclContext() *TypeDeclContext {
	var p = new(TypeDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_typeDecl
	return p
}

func InitEmptyTypeDeclContext(p *TypeDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_typeDecl
}

func (*TypeDeclContext) IsTypeDeclContext() {}

func NewTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclContext {
	var p = new(TypeDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_typeDecl

	return p
}

func (s *TypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclContext) CopyAll(ctx *TypeDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeDeclPARASTContext struct {
	TypeDeclContext
}

func NewTypeDeclPARASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeDeclPARASTContext {
	var p = new(TypeDeclPARASTContext)

	InitEmptyTypeDeclContext(&p.TypeDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeDeclContext))

	return p
}

func (s *TypeDeclPARASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclPARASTContext) TYPE() antlr.TerminalNode {
	return s.GetToken(MiniGoParserTYPE, 0)
}

func (s *TypeDeclPARASTContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARIZQ, 0)
}

func (s *TypeDeclPARASTContext) InnerTypeDecls() IInnerTypeDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInnerTypeDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInnerTypeDeclsContext)
}

func (s *TypeDeclPARASTContext) PARDER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARDER, 0)
}

func (s *TypeDeclPARASTContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, 0)
}

func (s *TypeDeclPARASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitTypeDeclPARAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeDeclTypeASTContext struct {
	TypeDeclContext
}

func NewTypeDeclTypeASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeDeclTypeASTContext {
	var p = new(TypeDeclTypeASTContext)

	InitEmptyTypeDeclContext(&p.TypeDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeDeclContext))

	return p
}

func (s *TypeDeclTypeASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclTypeASTContext) TYPE() antlr.TerminalNode {
	return s.GetToken(MiniGoParserTYPE, 0)
}

func (s *TypeDeclTypeASTContext) SingleTypeDecl() ISingleTypeDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleTypeDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleTypeDeclContext)
}

func (s *TypeDeclTypeASTContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, 0)
}

func (s *TypeDeclTypeASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitTypeDeclTypeAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeDeclNULPARASTContext struct {
	TypeDeclContext
}

func NewTypeDeclNULPARASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeDeclNULPARASTContext {
	var p = new(TypeDeclNULPARASTContext)

	InitEmptyTypeDeclContext(&p.TypeDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeDeclContext))

	return p
}

func (s *TypeDeclNULPARASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclNULPARASTContext) TYPE() antlr.TerminalNode {
	return s.GetToken(MiniGoParserTYPE, 0)
}

func (s *TypeDeclNULPARASTContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARIZQ, 0)
}

func (s *TypeDeclNULPARASTContext) PARDER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARDER, 0)
}

func (s *TypeDeclNULPARASTContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, 0)
}

func (s *TypeDeclNULPARASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitTypeDeclNULPARAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) TypeDecl() (localctx ITypeDeclContext) {
	localctx = NewTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MiniGoParserRULE_typeDecl)
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeDeclTypeASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)
			p.Match(MiniGoParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.SingleTypeDecl()
		}
		{
			p.SetState(137)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewTypeDeclPARASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(139)
			p.Match(MiniGoParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.Match(MiniGoParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.InnerTypeDecls()
		}
		{
			p.SetState(142)
			p.Match(MiniGoParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewTypeDeclNULPARASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(145)
			p.Match(MiniGoParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.Match(MiniGoParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.Match(MiniGoParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(148)
			p.Match(MiniGoParserPyCOMA)
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

// IInnerTypeDeclsContext is an interface to support dynamic dispatch.
type IInnerTypeDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsInnerTypeDeclsContext differentiates from other interfaces.
	IsInnerTypeDeclsContext()
}

type InnerTypeDeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInnerTypeDeclsContext() *InnerTypeDeclsContext {
	var p = new(InnerTypeDeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_innerTypeDecls
	return p
}

func InitEmptyInnerTypeDeclsContext(p *InnerTypeDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_innerTypeDecls
}

func (*InnerTypeDeclsContext) IsInnerTypeDeclsContext() {}

func NewInnerTypeDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InnerTypeDeclsContext {
	var p = new(InnerTypeDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_innerTypeDecls

	return p
}

func (s *InnerTypeDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *InnerTypeDeclsContext) CopyAll(ctx *InnerTypeDeclsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *InnerTypeDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerTypeDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type InnerTypeDeclsASTContext struct {
	InnerTypeDeclsContext
}

func NewInnerTypeDeclsASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InnerTypeDeclsASTContext {
	var p = new(InnerTypeDeclsASTContext)

	InitEmptyInnerTypeDeclsContext(&p.InnerTypeDeclsContext)
	p.parser = parser
	p.CopyAll(ctx.(*InnerTypeDeclsContext))

	return p
}

func (s *InnerTypeDeclsASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerTypeDeclsASTContext) AllSingleTypeDecl() []ISingleTypeDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleTypeDeclContext); ok {
			len++
		}
	}

	tst := make([]ISingleTypeDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleTypeDeclContext); ok {
			tst[i] = t.(ISingleTypeDeclContext)
			i++
		}
	}

	return tst
}

func (s *InnerTypeDeclsASTContext) SingleTypeDecl(i int) ISingleTypeDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleTypeDeclContext); ok {
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

	return t.(ISingleTypeDeclContext)
}

func (s *InnerTypeDeclsASTContext) AllPyCOMA() []antlr.TerminalNode {
	return s.GetTokens(MiniGoParserPyCOMA)
}

func (s *InnerTypeDeclsASTContext) PyCOMA(i int) antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, i)
}

func (s *InnerTypeDeclsASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitInnerTypeDeclsAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) InnerTypeDecls() (localctx IInnerTypeDeclsContext) {
	localctx = NewInnerTypeDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MiniGoParserRULE_innerTypeDecls)
	var _la int

	localctx = NewInnerTypeDeclsASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.SingleTypeDecl()
	}
	{
		p.SetState(152)
		p.Match(MiniGoParserPyCOMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MiniGoParserID {
		{
			p.SetState(153)
			p.SingleTypeDecl()
		}
		{
			p.SetState(154)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(160)
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

// ISingleTypeDeclContext is an interface to support dynamic dispatch.
type ISingleTypeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSingleTypeDeclContext differentiates from other interfaces.
	IsSingleTypeDeclContext()
}

type SingleTypeDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleTypeDeclContext() *SingleTypeDeclContext {
	var p = new(SingleTypeDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_singleTypeDecl
	return p
}

func InitEmptySingleTypeDeclContext(p *SingleTypeDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_singleTypeDecl
}

func (*SingleTypeDeclContext) IsSingleTypeDeclContext() {}

func NewSingleTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleTypeDeclContext {
	var p = new(SingleTypeDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_singleTypeDecl

	return p
}

func (s *SingleTypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleTypeDeclContext) CopyAll(ctx *SingleTypeDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SingleTypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleTypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SingleTypeDeclASTContext struct {
	SingleTypeDeclContext
}

func NewSingleTypeDeclASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleTypeDeclASTContext {
	var p = new(SingleTypeDeclASTContext)

	InitEmptySingleTypeDeclContext(&p.SingleTypeDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleTypeDeclContext))

	return p
}

func (s *SingleTypeDeclASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleTypeDeclASTContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniGoParserID, 0)
}

func (s *SingleTypeDeclASTContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *SingleTypeDeclASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSingleTypeDeclAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) SingleTypeDecl() (localctx ISingleTypeDeclContext) {
	localctx = NewSingleTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MiniGoParserRULE_singleTypeDecl)
	localctx = NewSingleTypeDeclASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(MiniGoParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.DeclType()
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

// IFuncDeclContext is an interface to support dynamic dispatch.
type IFuncDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFuncDeclContext differentiates from other interfaces.
	IsFuncDeclContext()
}

type FuncDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDeclContext() *FuncDeclContext {
	var p = new(FuncDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_funcDecl
	return p
}

func InitEmptyFuncDeclContext(p *FuncDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_funcDecl
}

func (*FuncDeclContext) IsFuncDeclContext() {}

func NewFuncDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDeclContext {
	var p = new(FuncDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_funcDecl

	return p
}

func (s *FuncDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDeclContext) CopyAll(ctx *FuncDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FuncDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncDeclASTContext struct {
	FuncDeclContext
}

func NewFuncDeclASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncDeclASTContext {
	var p = new(FuncDeclASTContext)

	InitEmptyFuncDeclContext(&p.FuncDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*FuncDeclContext))

	return p
}

func (s *FuncDeclASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclASTContext) FuncFrontDecl() IFuncFrontDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncFrontDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncFrontDeclContext)
}

func (s *FuncDeclASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncDeclASTContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, 0)
}

func (s *FuncDeclASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitFuncDeclAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) FuncDecl() (localctx IFuncDeclContext) {
	localctx = NewFuncDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MiniGoParserRULE_funcDecl)
	localctx = NewFuncDeclASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.FuncFrontDecl()
	}
	{
		p.SetState(165)
		p.Block()
	}
	{
		p.SetState(166)
		p.Match(MiniGoParserPyCOMA)
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

// IFuncFrontDeclContext is an interface to support dynamic dispatch.
type IFuncFrontDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFuncFrontDeclContext differentiates from other interfaces.
	IsFuncFrontDeclContext()
}

type FuncFrontDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncFrontDeclContext() *FuncFrontDeclContext {
	var p = new(FuncFrontDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_funcFrontDecl
	return p
}

func InitEmptyFuncFrontDeclContext(p *FuncFrontDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_funcFrontDecl
}

func (*FuncFrontDeclContext) IsFuncFrontDeclContext() {}

func NewFuncFrontDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncFrontDeclContext {
	var p = new(FuncFrontDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_funcFrontDecl

	return p
}

func (s *FuncFrontDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncFrontDeclContext) CopyAll(ctx *FuncFrontDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FuncFrontDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncFrontDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncFrontDeclASTContext struct {
	FuncFrontDeclContext
}

func NewFuncFrontDeclASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncFrontDeclASTContext {
	var p = new(FuncFrontDeclASTContext)

	InitEmptyFuncFrontDeclContext(&p.FuncFrontDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*FuncFrontDeclContext))

	return p
}

func (s *FuncFrontDeclASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncFrontDeclASTContext) FUNC() antlr.TerminalNode {
	return s.GetToken(MiniGoParserFUNC, 0)
}

func (s *FuncFrontDeclASTContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniGoParserID, 0)
}

func (s *FuncFrontDeclASTContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARIZQ, 0)
}

func (s *FuncFrontDeclASTContext) PARDER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARDER, 0)
}

func (s *FuncFrontDeclASTContext) FuncArgDecls() IFuncArgDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncArgDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncArgDeclsContext)
}

func (s *FuncFrontDeclASTContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *FuncFrontDeclASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitFuncFrontDeclAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) FuncFrontDecl() (localctx IFuncFrontDeclContext) {
	localctx = NewFuncFrontDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MiniGoParserRULE_funcFrontDecl)
	var _la int

	localctx = NewFuncFrontDeclASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(MiniGoParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Match(MiniGoParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Match(MiniGoParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MiniGoParserID {
		{
			p.SetState(171)
			p.FuncArgDecls()
		}

	}
	{
		p.SetState(174)
		p.Match(MiniGoParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-5)) & ^0x3f) == 0 && ((int64(1)<<(_la-5))&1152921504606847617) != 0 {
		{
			p.SetState(175)
			p.DeclType()
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

// IFuncArgDeclsContext is an interface to support dynamic dispatch.
type IFuncArgDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFuncArgDeclsContext differentiates from other interfaces.
	IsFuncArgDeclsContext()
}

type FuncArgDeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncArgDeclsContext() *FuncArgDeclsContext {
	var p = new(FuncArgDeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_funcArgDecls
	return p
}

func InitEmptyFuncArgDeclsContext(p *FuncArgDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_funcArgDecls
}

func (*FuncArgDeclsContext) IsFuncArgDeclsContext() {}

func NewFuncArgDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncArgDeclsContext {
	var p = new(FuncArgDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_funcArgDecls

	return p
}

func (s *FuncArgDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncArgDeclsContext) CopyAll(ctx *FuncArgDeclsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FuncArgDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncArgDeclsASTContext struct {
	FuncArgDeclsContext
}

func NewFuncArgDeclsASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncArgDeclsASTContext {
	var p = new(FuncArgDeclsASTContext)

	InitEmptyFuncArgDeclsContext(&p.FuncArgDeclsContext)
	p.parser = parser
	p.CopyAll(ctx.(*FuncArgDeclsContext))

	return p
}

func (s *FuncArgDeclsASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgDeclsASTContext) AllSingleVarDeclNoExps() []ISingleVarDeclNoExpsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			len++
		}
	}

	tst := make([]ISingleVarDeclNoExpsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			tst[i] = t.(ISingleVarDeclNoExpsContext)
			i++
		}
	}

	return tst
}

func (s *FuncArgDeclsASTContext) SingleVarDeclNoExps(i int) ISingleVarDeclNoExpsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
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

	return t.(ISingleVarDeclNoExpsContext)
}

func (s *FuncArgDeclsASTContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(MiniGoParserCOMA)
}

func (s *FuncArgDeclsASTContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(MiniGoParserCOMA, i)
}

func (s *FuncArgDeclsASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitFuncArgDeclsAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) FuncArgDecls() (localctx IFuncArgDeclsContext) {
	localctx = NewFuncArgDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MiniGoParserRULE_funcArgDecls)
	var _la int

	localctx = NewFuncArgDeclsASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.SingleVarDeclNoExps()
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MiniGoParserCOMA {
		{
			p.SetState(179)
			p.Match(MiniGoParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
			p.SingleVarDeclNoExps()
		}

		p.SetState(185)
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

// IDeclTypeContext is an interface to support dynamic dispatch.
type IDeclTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclTypeContext differentiates from other interfaces.
	IsDeclTypeContext()
}

type DeclTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclTypeContext() *DeclTypeContext {
	var p = new(DeclTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_declType
	return p
}

func InitEmptyDeclTypeContext(p *DeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_declType
}

func (*DeclTypeContext) IsDeclTypeContext() {}

func NewDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclTypeContext {
	var p = new(DeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_declType

	return p
}

func (s *DeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclTypeContext) CopyAll(ctx *DeclTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructDeclTypeDeckASTContext struct {
	DeclTypeContext
}

func NewStructDeclTypeDeckASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructDeclTypeDeckASTContext {
	var p = new(StructDeclTypeDeckASTContext)

	InitEmptyDeclTypeContext(&p.DeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclTypeContext))

	return p
}

func (s *StructDeclTypeDeckASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclTypeDeckASTContext) StructDeclType() IStructDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDeclTypeContext)
}

func (s *StructDeclTypeDeckASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitStructDeclTypeDeckAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayDeclTypeDeclASTContext struct {
	DeclTypeContext
}

func NewArrayDeclTypeDeclASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayDeclTypeDeclASTContext {
	var p = new(ArrayDeclTypeDeclASTContext)

	InitEmptyDeclTypeContext(&p.DeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclTypeContext))

	return p
}

func (s *ArrayDeclTypeDeclASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayDeclTypeDeclASTContext) ArrayDeclType() IArrayDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayDeclTypeContext)
}

func (s *ArrayDeclTypeDeclASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitArrayDeclTypeDeclAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclTypeParASTContext struct {
	DeclTypeContext
}

func NewDeclTypeParASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclTypeParASTContext {
	var p = new(DeclTypeParASTContext)

	InitEmptyDeclTypeContext(&p.DeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclTypeContext))

	return p
}

func (s *DeclTypeParASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclTypeParASTContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARIZQ, 0)
}

func (s *DeclTypeParASTContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *DeclTypeParASTContext) PARDER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARDER, 0)
}

func (s *DeclTypeParASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitDeclTypeParAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceDeclTypeDeclASTContext struct {
	DeclTypeContext
}

func NewSliceDeclTypeDeclASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceDeclTypeDeclASTContext {
	var p = new(SliceDeclTypeDeclASTContext)

	InitEmptyDeclTypeContext(&p.DeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclTypeContext))

	return p
}

func (s *SliceDeclTypeDeclASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceDeclTypeDeclASTContext) SliceDeclType() ISliceDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceDeclTypeContext)
}

func (s *SliceDeclTypeDeclASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSliceDeclTypeDeclAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclTypeIDASTContext struct {
	DeclTypeContext
}

func NewDeclTypeIDASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclTypeIDASTContext {
	var p = new(DeclTypeIDASTContext)

	InitEmptyDeclTypeContext(&p.DeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclTypeContext))

	return p
}

func (s *DeclTypeIDASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclTypeIDASTContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniGoParserID, 0)
}

func (s *DeclTypeIDASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitDeclTypeIDAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) DeclType() (localctx IDeclTypeContext) {
	localctx = NewDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MiniGoParserRULE_declType)
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDeclTypeParASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(186)
			p.Match(MiniGoParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.DeclType()
		}
		{
			p.SetState(188)
			p.Match(MiniGoParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewDeclTypeIDASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.Match(MiniGoParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewSliceDeclTypeDeclASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(191)
			p.SliceDeclType()
		}

	case 4:
		localctx = NewArrayDeclTypeDeclASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(192)
			p.ArrayDeclType()
		}

	case 5:
		localctx = NewStructDeclTypeDeckASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(193)
			p.StructDeclType()
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

// ISliceDeclTypeContext is an interface to support dynamic dispatch.
type ISliceDeclTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSliceDeclTypeContext differentiates from other interfaces.
	IsSliceDeclTypeContext()
}

type SliceDeclTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceDeclTypeContext() *SliceDeclTypeContext {
	var p = new(SliceDeclTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_sliceDeclType
	return p
}

func InitEmptySliceDeclTypeContext(p *SliceDeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_sliceDeclType
}

func (*SliceDeclTypeContext) IsSliceDeclTypeContext() {}

func NewSliceDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceDeclTypeContext {
	var p = new(SliceDeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_sliceDeclType

	return p
}

func (s *SliceDeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceDeclTypeContext) CopyAll(ctx *SliceDeclTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SliceDeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceDeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SliceDeclTypeASTContext struct {
	SliceDeclTypeContext
}

func NewSliceDeclTypeASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceDeclTypeASTContext {
	var p = new(SliceDeclTypeASTContext)

	InitEmptySliceDeclTypeContext(&p.SliceDeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*SliceDeclTypeContext))

	return p
}

func (s *SliceDeclTypeASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceDeclTypeASTContext) CORCHIZQ() antlr.TerminalNode {
	return s.GetToken(MiniGoParserCORCHIZQ, 0)
}

func (s *SliceDeclTypeASTContext) CORCHDER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserCORCHDER, 0)
}

func (s *SliceDeclTypeASTContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *SliceDeclTypeASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSliceDeclTypeAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) SliceDeclType() (localctx ISliceDeclTypeContext) {
	localctx = NewSliceDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MiniGoParserRULE_sliceDeclType)
	localctx = NewSliceDeclTypeASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(MiniGoParserCORCHIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
		p.Match(MiniGoParserCORCHDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(198)
		p.DeclType()
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

// IArrayDeclTypeContext is an interface to support dynamic dispatch.
type IArrayDeclTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArrayDeclTypeContext differentiates from other interfaces.
	IsArrayDeclTypeContext()
}

type ArrayDeclTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayDeclTypeContext() *ArrayDeclTypeContext {
	var p = new(ArrayDeclTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_arrayDeclType
	return p
}

func InitEmptyArrayDeclTypeContext(p *ArrayDeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_arrayDeclType
}

func (*ArrayDeclTypeContext) IsArrayDeclTypeContext() {}

func NewArrayDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayDeclTypeContext {
	var p = new(ArrayDeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_arrayDeclType

	return p
}

func (s *ArrayDeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayDeclTypeContext) CopyAll(ctx *ArrayDeclTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ArrayDeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayDeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArrayDeclTypeASTContext struct {
	ArrayDeclTypeContext
}

func NewArrayDeclTypeASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayDeclTypeASTContext {
	var p = new(ArrayDeclTypeASTContext)

	InitEmptyArrayDeclTypeContext(&p.ArrayDeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArrayDeclTypeContext))

	return p
}

func (s *ArrayDeclTypeASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayDeclTypeASTContext) CORCHIZQ() antlr.TerminalNode {
	return s.GetToken(MiniGoParserCORCHIZQ, 0)
}

func (s *ArrayDeclTypeASTContext) INT() antlr.TerminalNode {
	return s.GetToken(MiniGoParserINT, 0)
}

func (s *ArrayDeclTypeASTContext) CORCHDER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserCORCHDER, 0)
}

func (s *ArrayDeclTypeASTContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *ArrayDeclTypeASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitArrayDeclTypeAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) ArrayDeclType() (localctx IArrayDeclTypeContext) {
	localctx = NewArrayDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MiniGoParserRULE_arrayDeclType)
	localctx = NewArrayDeclTypeASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(MiniGoParserCORCHIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)
		p.Match(MiniGoParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.Match(MiniGoParserCORCHDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(203)
		p.DeclType()
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

// IStructDeclTypeContext is an interface to support dynamic dispatch.
type IStructDeclTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStructDeclTypeContext differentiates from other interfaces.
	IsStructDeclTypeContext()
}

type StructDeclTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDeclTypeContext() *StructDeclTypeContext {
	var p = new(StructDeclTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_structDeclType
	return p
}

func InitEmptyStructDeclTypeContext(p *StructDeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_structDeclType
}

func (*StructDeclTypeContext) IsStructDeclTypeContext() {}

func NewStructDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclTypeContext {
	var p = new(StructDeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_structDeclType

	return p
}

func (s *StructDeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclTypeContext) CopyAll(ctx *StructDeclTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StructDeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructDeclTypeASTContext struct {
	StructDeclTypeContext
}

func NewStructDeclTypeASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructDeclTypeASTContext {
	var p = new(StructDeclTypeASTContext)

	InitEmptyStructDeclTypeContext(&p.StructDeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*StructDeclTypeContext))

	return p
}

func (s *StructDeclTypeASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclTypeASTContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(MiniGoParserSTRUCT, 0)
}

func (s *StructDeclTypeASTContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLLAVEIZQ, 0)
}

func (s *StructDeclTypeASTContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLLAVEDER, 0)
}

func (s *StructDeclTypeASTContext) StructMemDecls() IStructMemDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructMemDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructMemDeclsContext)
}

func (s *StructDeclTypeASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitStructDeclTypeAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) StructDeclType() (localctx IStructDeclTypeContext) {
	localctx = NewStructDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MiniGoParserRULE_structDeclType)
	var _la int

	localctx = NewStructDeclTypeASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(MiniGoParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.Match(MiniGoParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MiniGoParserID {
		{
			p.SetState(207)
			p.StructMemDecls()
		}

	}
	{
		p.SetState(210)
		p.Match(MiniGoParserLLAVEDER)
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

// IStructMemDeclsContext is an interface to support dynamic dispatch.
type IStructMemDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStructMemDeclsContext differentiates from other interfaces.
	IsStructMemDeclsContext()
}

type StructMemDeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructMemDeclsContext() *StructMemDeclsContext {
	var p = new(StructMemDeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_structMemDecls
	return p
}

func InitEmptyStructMemDeclsContext(p *StructMemDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_structMemDecls
}

func (*StructMemDeclsContext) IsStructMemDeclsContext() {}

func NewStructMemDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructMemDeclsContext {
	var p = new(StructMemDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_structMemDecls

	return p
}

func (s *StructMemDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *StructMemDeclsContext) CopyAll(ctx *StructMemDeclsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StructMemDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructMemDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructMemDeclsASTContext struct {
	StructMemDeclsContext
}

func NewStructMemDeclsASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructMemDeclsASTContext {
	var p = new(StructMemDeclsASTContext)

	InitEmptyStructMemDeclsContext(&p.StructMemDeclsContext)
	p.parser = parser
	p.CopyAll(ctx.(*StructMemDeclsContext))

	return p
}

func (s *StructMemDeclsASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructMemDeclsASTContext) AllSingleVarDeclNoExps() []ISingleVarDeclNoExpsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			len++
		}
	}

	tst := make([]ISingleVarDeclNoExpsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			tst[i] = t.(ISingleVarDeclNoExpsContext)
			i++
		}
	}

	return tst
}

func (s *StructMemDeclsASTContext) SingleVarDeclNoExps(i int) ISingleVarDeclNoExpsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
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

	return t.(ISingleVarDeclNoExpsContext)
}

func (s *StructMemDeclsASTContext) AllPyCOMA() []antlr.TerminalNode {
	return s.GetTokens(MiniGoParserPyCOMA)
}

func (s *StructMemDeclsASTContext) PyCOMA(i int) antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, i)
}

func (s *StructMemDeclsASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitStructMemDeclsAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) StructMemDecls() (localctx IStructMemDeclsContext) {
	localctx = NewStructMemDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MiniGoParserRULE_structMemDecls)
	var _la int

	localctx = NewStructMemDeclsASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.SingleVarDeclNoExps()
	}
	{
		p.SetState(213)
		p.Match(MiniGoParserPyCOMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MiniGoParserID {
		{
			p.SetState(214)
			p.SingleVarDeclNoExps()
		}
		{
			p.SetState(215)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(221)
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

// IIdentifierListContext is an interface to support dynamic dispatch.
type IIdentifierListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIdentifierListContext differentiates from other interfaces.
	IsIdentifierListContext()
}

type IdentifierListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierListContext() *IdentifierListContext {
	var p = new(IdentifierListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_identifierList
	return p
}

func InitEmptyIdentifierListContext(p *IdentifierListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_identifierList
}

func (*IdentifierListContext) IsIdentifierListContext() {}

func NewIdentifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierListContext {
	var p = new(IdentifierListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_identifierList

	return p
}

func (s *IdentifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierListContext) CopyAll(ctx *IdentifierListContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IdentifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IdentifierListASTContext struct {
	IdentifierListContext
}

func NewIdentifierListASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierListASTContext {
	var p = new(IdentifierListASTContext)

	InitEmptyIdentifierListContext(&p.IdentifierListContext)
	p.parser = parser
	p.CopyAll(ctx.(*IdentifierListContext))

	return p
}

func (s *IdentifierListASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListASTContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(MiniGoParserID)
}

func (s *IdentifierListASTContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(MiniGoParserID, i)
}

func (s *IdentifierListASTContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(MiniGoParserCOMA)
}

func (s *IdentifierListASTContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(MiniGoParserCOMA, i)
}

func (s *IdentifierListASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitIdentifierListAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) IdentifierList() (localctx IIdentifierListContext) {
	localctx = NewIdentifierListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MiniGoParserRULE_identifierList)
	var _la int

	localctx = NewIdentifierListASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.Match(MiniGoParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MiniGoParserCOMA {
		{
			p.SetState(223)
			p.Match(MiniGoParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)
			p.Match(MiniGoParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(229)
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
	p.RuleIndex = MiniGoParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_expression

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

type SumaExpresionASTContext struct {
	ExpressionContext
}

func NewSumaExpresionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SumaExpresionASTContext {
	var p = new(SumaExpresionASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *SumaExpresionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SumaExpresionASTContext) AllExpression() []IExpressionContext {
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

func (s *SumaExpresionASTContext) Expression(i int) IExpressionContext {
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

func (s *SumaExpresionASTContext) SUMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserSUMA, 0)
}

func (s *SumaExpresionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSumaExpresionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionASTContext struct {
	ExpressionContext
}

func NewExpressionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionASTContext {
	var p = new(ExpressionASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionASTContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *ExpressionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitExpressionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type MayorExpresionASTContext struct {
	ExpressionContext
}

func NewMayorExpresionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MayorExpresionASTContext {
	var p = new(MayorExpresionASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MayorExpresionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MayorExpresionASTContext) AllExpression() []IExpressionContext {
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

func (s *MayorExpresionASTContext) Expression(i int) IExpressionContext {
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

func (s *MayorExpresionASTContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(MiniGoParserMAYOR, 0)
}

func (s *MayorExpresionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitMayorExpresionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type MenorIgExpresionASTContext struct {
	ExpressionContext
}

func NewMenorIgExpresionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MenorIgExpresionASTContext {
	var p = new(MenorIgExpresionASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MenorIgExpresionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MenorIgExpresionASTContext) AllExpression() []IExpressionContext {
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

func (s *MenorIgExpresionASTContext) Expression(i int) IExpressionContext {
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

func (s *MenorIgExpresionASTContext) MENORIG() antlr.TerminalNode {
	return s.GetToken(MiniGoParserMENORIG, 0)
}

func (s *MenorIgExpresionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitMenorIgExpresionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type RestaExExpresionASTContext struct {
	ExpressionContext
}

func NewRestaExExpresionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RestaExExpresionASTContext {
	var p = new(RestaExExpresionASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *RestaExExpresionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RestaExExpresionASTContext) RESTA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserRESTA, 0)
}

func (s *RestaExExpresionASTContext) Expression() IExpressionContext {
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

func (s *RestaExExpresionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitRestaExExpresionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AnddExpresionASTContext struct {
	ExpressionContext
}

func NewAnddExpresionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AnddExpresionASTContext {
	var p = new(AnddExpresionASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AnddExpresionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnddExpresionASTContext) AllExpression() []IExpressionContext {
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

func (s *AnddExpresionASTContext) Expression(i int) IExpressionContext {
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

func (s *AnddExpresionASTContext) ANDD() antlr.TerminalNode {
	return s.GetToken(MiniGoParserANDD, 0)
}

func (s *AnddExpresionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitAnddExpresionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type MayorMExpresionASTContext struct {
	ExpressionContext
}

func NewMayorMExpresionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MayorMExpresionASTContext {
	var p = new(MayorMExpresionASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MayorMExpresionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MayorMExpresionASTContext) AllExpression() []IExpressionContext {
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

func (s *MayorMExpresionASTContext) Expression(i int) IExpressionContext {
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

func (s *MayorMExpresionASTContext) MAYORM() antlr.TerminalNode {
	return s.GetToken(MiniGoParserMAYORM, 0)
}

func (s *MayorMExpresionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitMayorMExpresionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type MenorExpresionASTContext struct {
	ExpressionContext
}

func NewMenorExpresionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MenorExpresionASTContext {
	var p = new(MenorExpresionASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MenorExpresionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MenorExpresionASTContext) AllExpression() []IExpressionContext {
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

func (s *MenorExpresionASTContext) Expression(i int) IExpressionContext {
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

func (s *MenorExpresionASTContext) MENOR() antlr.TerminalNode {
	return s.GetToken(MiniGoParserMENOR, 0)
}

func (s *MenorExpresionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitMenorExpresionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SumaExExpresionASTContext struct {
	ExpressionContext
}

func NewSumaExExpresionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SumaExExpresionASTContext {
	var p = new(SumaExExpresionASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *SumaExExpresionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SumaExExpresionASTContext) SUMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserSUMA, 0)
}

func (s *SumaExExpresionASTContext) Expression() IExpressionContext {
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

func (s *SumaExExpresionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSumaExExpresionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitExpresionASTContext struct {
	ExpressionContext
}

func NewBitExpresionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitExpresionASTContext {
	var p = new(BitExpresionASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BitExpresionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitExpresionASTContext) AllExpression() []IExpressionContext {
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

func (s *BitExpresionASTContext) Expression(i int) IExpressionContext {
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

func (s *BitExpresionASTContext) BITCLE() antlr.TerminalNode {
	return s.GetToken(MiniGoParserBITCLE, 0)
}

func (s *BitExpresionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitBitExpresionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type DiferenteExpresionASTContext struct {
	ExpressionContext
}

func NewDiferenteExpresionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DiferenteExpresionASTContext {
	var p = new(DiferenteExpresionASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *DiferenteExpresionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DiferenteExpresionASTContext) AllExpression() []IExpressionContext {
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

func (s *DiferenteExpresionASTContext) Expression(i int) IExpressionContext {
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

func (s *DiferenteExpresionASTContext) DIFERENTE() antlr.TerminalNode {
	return s.GetToken(MiniGoParserDIFERENTE, 0)
}

func (s *DiferenteExpresionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitDiferenteExpresionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultExpresionASTContext struct {
	ExpressionContext
}

func NewMultExpresionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultExpresionASTContext {
	var p = new(MultExpresionASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MultExpresionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultExpresionASTContext) AllExpression() []IExpressionContext {
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

func (s *MultExpresionASTContext) Expression(i int) IExpressionContext {
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

func (s *MultExpresionASTContext) MULTI() antlr.TerminalNode {
	return s.GetToken(MiniGoParserMULTI, 0)
}

func (s *MultExpresionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitMultExpresionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type DivExpresionASTContext struct {
	ExpressionContext
}

func NewDivExpresionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivExpresionASTContext {
	var p = new(DivExpresionASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *DivExpresionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivExpresionASTContext) AllExpression() []IExpressionContext {
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

func (s *DivExpresionASTContext) Expression(i int) IExpressionContext {
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

func (s *DivExpresionASTContext) DIV() antlr.TerminalNode {
	return s.GetToken(MiniGoParserDIV, 0)
}

func (s *DivExpresionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitDivExpresionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrExpresionASTContext struct {
	ExpressionContext
}

func NewOrExpresionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExpresionASTContext {
	var p = new(OrExpresionASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *OrExpresionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExpresionASTContext) AllExpression() []IExpressionContext {
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

func (s *OrExpresionASTContext) Expression(i int) IExpressionContext {
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

func (s *OrExpresionASTContext) OR() antlr.TerminalNode {
	return s.GetToken(MiniGoParserOR, 0)
}

func (s *OrExpresionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitOrExpresionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type XorExpresionASTContext struct {
	ExpressionContext
}

func NewXorExpresionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *XorExpresionASTContext {
	var p = new(XorExpresionASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *XorExpresionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XorExpresionASTContext) AllExpression() []IExpressionContext {
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

func (s *XorExpresionASTContext) Expression(i int) IExpressionContext {
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

func (s *XorExpresionASTContext) XOR() antlr.TerminalNode {
	return s.GetToken(MiniGoParserXOR, 0)
}

func (s *XorExpresionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitXorExpresionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdenticoExpresionASTContext struct {
	ExpressionContext
}

func NewIdenticoExpresionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdenticoExpresionASTContext {
	var p = new(IdenticoExpresionASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IdenticoExpresionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdenticoExpresionASTContext) AllExpression() []IExpressionContext {
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

func (s *IdenticoExpresionASTContext) Expression(i int) IExpressionContext {
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

func (s *IdenticoExpresionASTContext) IDENTICO() antlr.TerminalNode {
	return s.GetToken(MiniGoParserIDENTICO, 0)
}

func (s *IdenticoExpresionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitIdenticoExpresionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type MayorIgExpresionASTContext struct {
	ExpressionContext
}

func NewMayorIgExpresionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MayorIgExpresionASTContext {
	var p = new(MayorIgExpresionASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MayorIgExpresionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MayorIgExpresionASTContext) AllExpression() []IExpressionContext {
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

func (s *MayorIgExpresionASTContext) Expression(i int) IExpressionContext {
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

func (s *MayorIgExpresionASTContext) MAYORIG() antlr.TerminalNode {
	return s.GetToken(MiniGoParserMAYORIG, 0)
}

func (s *MayorIgExpresionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitMayorIgExpresionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExclExpresionASTContext struct {
	ExpressionContext
}

func NewExclExpresionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExclExpresionASTContext {
	var p = new(ExclExpresionASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExclExpresionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExclExpresionASTContext) EXCL() antlr.TerminalNode {
	return s.GetToken(MiniGoParserEXCL, 0)
}

func (s *ExclExpresionASTContext) Expression() IExpressionContext {
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

func (s *ExclExpresionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitExclExpresionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndExpresionASTContext struct {
	ExpressionContext
}

func NewAndExpresionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExpresionASTContext {
	var p = new(AndExpresionASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AndExpresionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpresionASTContext) AllExpression() []IExpressionContext {
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

func (s *AndExpresionASTContext) Expression(i int) IExpressionContext {
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

func (s *AndExpresionASTContext) AND() antlr.TerminalNode {
	return s.GetToken(MiniGoParserAND, 0)
}

func (s *AndExpresionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitAndExpresionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrdExpresionASTContext struct {
	ExpressionContext
}

func NewOrdExpresionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrdExpresionASTContext {
	var p = new(OrdExpresionASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *OrdExpresionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrdExpresionASTContext) AllExpression() []IExpressionContext {
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

func (s *OrdExpresionASTContext) Expression(i int) IExpressionContext {
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

func (s *OrdExpresionASTContext) ORD() antlr.TerminalNode {
	return s.GetToken(MiniGoParserORD, 0)
}

func (s *OrdExpresionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitOrdExpresionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type RestaExpresionASTContext struct {
	ExpressionContext
}

func NewRestaExpresionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RestaExpresionASTContext {
	var p = new(RestaExpresionASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *RestaExpresionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RestaExpresionASTContext) AllExpression() []IExpressionContext {
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

func (s *RestaExpresionASTContext) Expression(i int) IExpressionContext {
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

func (s *RestaExpresionASTContext) RESTA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserRESTA, 0)
}

func (s *RestaExpresionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitRestaExpresionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type XorExExpresionASTContext struct {
	ExpressionContext
}

func NewXorExExpresionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *XorExExpresionASTContext {
	var p = new(XorExExpresionASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *XorExExpresionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XorExExpresionASTContext) XOR() antlr.TerminalNode {
	return s.GetToken(MiniGoParserXOR, 0)
}

func (s *XorExExpresionASTContext) Expression() IExpressionContext {
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

func (s *XorExExpresionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitXorExExpresionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type PorcentExpresionASTContext struct {
	ExpressionContext
}

func NewPorcentExpresionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PorcentExpresionASTContext {
	var p = new(PorcentExpresionASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *PorcentExpresionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PorcentExpresionASTContext) AllExpression() []IExpressionContext {
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

func (s *PorcentExpresionASTContext) Expression(i int) IExpressionContext {
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

func (s *PorcentExpresionASTContext) PORCENT() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPORCENT, 0)
}

func (s *PorcentExpresionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitPorcentExpresionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type MenorMExpresionASTContext struct {
	ExpressionContext
}

func NewMenorMExpresionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MenorMExpresionASTContext {
	var p = new(MenorMExpresionASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MenorMExpresionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MenorMExpresionASTContext) AllExpression() []IExpressionContext {
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

func (s *MenorMExpresionASTContext) Expression(i int) IExpressionContext {
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

func (s *MenorMExpresionASTContext) MEMORM() antlr.TerminalNode {
	return s.GetToken(MiniGoParserMEMORM, 0)
}

func (s *MenorMExpresionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitMenorMExpresionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *MiniGoParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, MiniGoParserRULE_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MiniGoParserPARIZQ, MiniGoParserAPPEND, MiniGoParserLEN, MiniGoParserCAP, MiniGoParserID, MiniGoParserINT, MiniGoParserFLOAT, MiniGoParserRUNE, MiniGoParserRAWSTRING, MiniGoParserINTERPRETEDSTRING:
		localctx = NewExpressionASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(231)
			p.primaryExpression(0)
		}

	case MiniGoParserSUMA:
		localctx = NewSumaExExpresionASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(232)
			p.Match(MiniGoParserSUMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)
			p.expression(4)
		}

	case MiniGoParserRESTA:
		localctx = NewRestaExExpresionASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(234)
			p.Match(MiniGoParserRESTA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(235)
			p.expression(3)
		}

	case MiniGoParserEXCL:
		localctx = NewExclExpresionASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(236)
			p.Match(MiniGoParserEXCL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(237)
			p.expression(2)
		}

	case MiniGoParserXOR:
		localctx = NewXorExExpresionASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(238)
			p.Match(MiniGoParserXOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(239)
			p.expression(1)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(301)
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
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(299)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultExpresionASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(242)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
					goto errorExit
				}
				{
					p.SetState(243)
					p.Match(MiniGoParserMULTI)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(244)
					p.expression(24)
				}

			case 2:
				localctx = NewDivExpresionASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(245)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
					goto errorExit
				}
				{
					p.SetState(246)
					p.Match(MiniGoParserDIV)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(247)
					p.expression(23)
				}

			case 3:
				localctx = NewPorcentExpresionASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(248)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
				{
					p.SetState(249)
					p.Match(MiniGoParserPORCENT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(250)
					p.expression(22)
				}

			case 4:
				localctx = NewMayorMExpresionASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(251)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(252)
					p.Match(MiniGoParserMAYORM)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(253)
					p.expression(21)
				}

			case 5:
				localctx = NewMenorMExpresionASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(254)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(255)
					p.Match(MiniGoParserMEMORM)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(256)
					p.expression(20)
				}

			case 6:
				localctx = NewAndExpresionASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(257)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(258)
					p.Match(MiniGoParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(259)
					p.expression(19)
				}

			case 7:
				localctx = NewBitExpresionASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(260)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(261)
					p.Match(MiniGoParserBITCLE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(262)
					p.expression(18)
				}

			case 8:
				localctx = NewSumaExpresionASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(263)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(264)
					p.Match(MiniGoParserSUMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(265)
					p.expression(17)
				}

			case 9:
				localctx = NewRestaExpresionASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(266)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(267)
					p.Match(MiniGoParserRESTA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(268)
					p.expression(16)
				}

			case 10:
				localctx = NewOrExpresionASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(269)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(270)
					p.Match(MiniGoParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(271)
					p.expression(15)
				}

			case 11:
				localctx = NewXorExpresionASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(272)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(273)
					p.Match(MiniGoParserXOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(274)
					p.expression(14)
				}

			case 12:
				localctx = NewIdenticoExpresionASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(275)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(276)
					p.Match(MiniGoParserIDENTICO)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(277)
					p.expression(13)
				}

			case 13:
				localctx = NewDiferenteExpresionASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(278)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(279)
					p.Match(MiniGoParserDIFERENTE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(280)
					p.expression(12)
				}

			case 14:
				localctx = NewMayorExpresionASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(281)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(282)
					p.Match(MiniGoParserMAYOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(283)
					p.expression(11)
				}

			case 15:
				localctx = NewMayorIgExpresionASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(284)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(285)
					p.Match(MiniGoParserMAYORIG)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(286)
					p.expression(10)
				}

			case 16:
				localctx = NewMenorExpresionASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(287)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(288)
					p.Match(MiniGoParserMENOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(289)
					p.expression(9)
				}

			case 17:
				localctx = NewMenorIgExpresionASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(290)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(291)
					p.Match(MiniGoParserMENORIG)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(292)
					p.expression(8)
				}

			case 18:
				localctx = NewAnddExpresionASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(293)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(294)
					p.Match(MiniGoParserANDD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(295)
					p.expression(7)
				}

			case 19:
				localctx = NewOrdExpresionASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(296)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(297)
					p.Match(MiniGoParserORD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(298)
					p.expression(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
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

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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
	p.RuleIndex = MiniGoParserRULE_expressionList
	return p
}

func InitEmptyExpressionListContext(p *ExpressionListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_expressionList
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) CopyAll(ctx *ExpressionListContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionListASTContext struct {
	ExpressionListContext
}

func NewExpressionListASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionListASTContext {
	var p = new(ExpressionListASTContext)

	InitEmptyExpressionListContext(&p.ExpressionListContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionListContext))

	return p
}

func (s *ExpressionListASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListASTContext) AllExpression() []IExpressionContext {
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

func (s *ExpressionListASTContext) Expression(i int) IExpressionContext {
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

func (s *ExpressionListASTContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(MiniGoParserCOMA)
}

func (s *ExpressionListASTContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(MiniGoParserCOMA, i)
}

func (s *ExpressionListASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitExpressionListAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MiniGoParserRULE_expressionList)
	var _la int

	localctx = NewExpressionListASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(304)
		p.expression(0)
	}
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MiniGoParserCOMA {
		{
			p.SetState(305)
			p.Match(MiniGoParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(306)
			p.expression(0)
		}

		p.SetState(311)
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

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_primaryExpression
	return p
}

func InitEmptyPrimaryExpressionContext(p *PrimaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_primaryExpression
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) CopyAll(ctx *PrimaryExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrimaryExpressionInASTContext struct {
	PrimaryExpressionContext
}

func NewPrimaryExpressionInASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExpressionInASTContext {
	var p = new(PrimaryExpressionInASTContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *PrimaryExpressionInASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionInASTContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *PrimaryExpressionInASTContext) Index() IIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *PrimaryExpressionInASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitPrimaryExpressionInAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryExpressionArASTContext struct {
	PrimaryExpressionContext
}

func NewPrimaryExpressionArASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExpressionArASTContext {
	var p = new(PrimaryExpressionArASTContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *PrimaryExpressionArASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionArASTContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *PrimaryExpressionArASTContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *PrimaryExpressionArASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitPrimaryExpressionArAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryExpressionAppendASTContext struct {
	PrimaryExpressionContext
}

func NewPrimaryExpressionAppendASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExpressionAppendASTContext {
	var p = new(PrimaryExpressionAppendASTContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *PrimaryExpressionAppendASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionAppendASTContext) AppendExpression() IAppendExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAppendExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAppendExpressionContext)
}

func (s *PrimaryExpressionAppendASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitPrimaryExpressionAppendAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryExpressionSeASTContext struct {
	PrimaryExpressionContext
}

func NewPrimaryExpressionSeASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExpressionSeASTContext {
	var p = new(PrimaryExpressionSeASTContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *PrimaryExpressionSeASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionSeASTContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *PrimaryExpressionSeASTContext) Selector() ISelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorContext)
}

func (s *PrimaryExpressionSeASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitPrimaryExpressionSeAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryExpressionCapEASTContext struct {
	PrimaryExpressionContext
}

func NewPrimaryExpressionCapEASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExpressionCapEASTContext {
	var p = new(PrimaryExpressionCapEASTContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *PrimaryExpressionCapEASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionCapEASTContext) CapExpression() ICapExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICapExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICapExpressionContext)
}

func (s *PrimaryExpressionCapEASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitPrimaryExpressionCapEAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryExpressionOpASTContext struct {
	PrimaryExpressionContext
}

func NewPrimaryExpressionOpASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExpressionOpASTContext {
	var p = new(PrimaryExpressionOpASTContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *PrimaryExpressionOpASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionOpASTContext) Operand() IOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *PrimaryExpressionOpASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitPrimaryExpressionOpAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryExpressionLengASTContext struct {
	PrimaryExpressionContext
}

func NewPrimaryExpressionLengASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExpressionLengASTContext {
	var p = new(PrimaryExpressionLengASTContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *PrimaryExpressionLengASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionLengASTContext) LengthExpression() ILengthExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILengthExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILengthExpressionContext)
}

func (s *PrimaryExpressionLengASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitPrimaryExpressionLengAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	return p.primaryExpression(0)
}

func (p *MiniGoParser) primaryExpression(_p int) (localctx IPrimaryExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimaryExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 40
	p.EnterRecursionRule(localctx, 40, MiniGoParserRULE_primaryExpression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MiniGoParserPARIZQ, MiniGoParserID, MiniGoParserINT, MiniGoParserFLOAT, MiniGoParserRUNE, MiniGoParserRAWSTRING, MiniGoParserINTERPRETEDSTRING:
		localctx = NewPrimaryExpressionOpASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(313)
			p.Operand()
		}

	case MiniGoParserAPPEND:
		localctx = NewPrimaryExpressionAppendASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(314)
			p.AppendExpression()
		}

	case MiniGoParserLEN:
		localctx = NewPrimaryExpressionLengASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(315)
			p.LengthExpression()
		}

	case MiniGoParserCAP:
		localctx = NewPrimaryExpressionCapEASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(316)
			p.CapExpression()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(325)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPrimaryExpressionSeASTContext(p, NewPrimaryExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_primaryExpression)
				p.SetState(319)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(320)
					p.Selector()
				}

			case 2:
				localctx = NewPrimaryExpressionInASTContext(p, NewPrimaryExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_primaryExpression)
				p.SetState(321)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(322)
					p.Index()
				}

			case 3:
				localctx = NewPrimaryExpressionArASTContext(p, NewPrimaryExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_primaryExpression)
				p.SetState(323)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(324)
					p.Arguments()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
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

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_operand
	return p
}

func InitEmptyOperandContext(p *OperandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_operand
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) CopyAll(ctx *OperandContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type OperandLitASTContext struct {
	OperandContext
}

func NewOperandLitASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OperandLitASTContext {
	var p = new(OperandLitASTContext)

	InitEmptyOperandContext(&p.OperandContext)
	p.parser = parser
	p.CopyAll(ctx.(*OperandContext))

	return p
}

func (s *OperandLitASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandLitASTContext) Literal() ILiteralContext {
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

func (s *OperandLitASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitOperandLitAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type OperandPARASTContext struct {
	OperandContext
}

func NewOperandPARASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OperandPARASTContext {
	var p = new(OperandPARASTContext)

	InitEmptyOperandContext(&p.OperandContext)
	p.parser = parser
	p.CopyAll(ctx.(*OperandContext))

	return p
}

func (s *OperandPARASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandPARASTContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARIZQ, 0)
}

func (s *OperandPARASTContext) Expression() IExpressionContext {
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

func (s *OperandPARASTContext) PARDER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARDER, 0)
}

func (s *OperandPARASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitOperandPARAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type OperandIDASTContext struct {
	OperandContext
}

func NewOperandIDASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OperandIDASTContext {
	var p = new(OperandIDASTContext)

	InitEmptyOperandContext(&p.OperandContext)
	p.parser = parser
	p.CopyAll(ctx.(*OperandContext))

	return p
}

func (s *OperandIDASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandIDASTContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniGoParserID, 0)
}

func (s *OperandIDASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitOperandIDAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, MiniGoParserRULE_operand)
	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MiniGoParserINT, MiniGoParserFLOAT, MiniGoParserRUNE, MiniGoParserRAWSTRING, MiniGoParserINTERPRETEDSTRING:
		localctx = NewOperandLitASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(330)
			p.Literal()
		}

	case MiniGoParserID:
		localctx = NewOperandIDASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(331)
			p.Match(MiniGoParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserPARIZQ:
		localctx = NewOperandPARASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(332)
			p.Match(MiniGoParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(333)
			p.expression(0)
		}
		{
			p.SetState(334)
			p.Match(MiniGoParserPARDER)
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
	p.RuleIndex = MiniGoParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_literal

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

type LiteralIntASTContext struct {
	LiteralContext
}

func NewLiteralIntASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralIntASTContext {
	var p = new(LiteralIntASTContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *LiteralIntASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralIntASTContext) INT() antlr.TerminalNode {
	return s.GetToken(MiniGoParserINT, 0)
}

func (s *LiteralIntASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitLiteralIntAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralFloatASTContext struct {
	LiteralContext
}

func NewLiteralFloatASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralFloatASTContext {
	var p = new(LiteralFloatASTContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *LiteralFloatASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralFloatASTContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(MiniGoParserFLOAT, 0)
}

func (s *LiteralFloatASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitLiteralFloatAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralInterASTContext struct {
	LiteralContext
}

func NewLiteralInterASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralInterASTContext {
	var p = new(LiteralInterASTContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *LiteralInterASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralInterASTContext) INTERPRETEDSTRING() antlr.TerminalNode {
	return s.GetToken(MiniGoParserINTERPRETEDSTRING, 0)
}

func (s *LiteralInterASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitLiteralInterAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralRuneASTContext struct {
	LiteralContext
}

func NewLiteralRuneASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralRuneASTContext {
	var p = new(LiteralRuneASTContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *LiteralRuneASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralRuneASTContext) RUNE() antlr.TerminalNode {
	return s.GetToken(MiniGoParserRUNE, 0)
}

func (s *LiteralRuneASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitLiteralRuneAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralRawsASTContext struct {
	LiteralContext
}

func NewLiteralRawsASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralRawsASTContext {
	var p = new(LiteralRawsASTContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *LiteralRawsASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralRawsASTContext) RAWSTRING() antlr.TerminalNode {
	return s.GetToken(MiniGoParserRAWSTRING, 0)
}

func (s *LiteralRawsASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitLiteralRawsAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, MiniGoParserRULE_literal)
	p.SetState(343)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MiniGoParserINT:
		localctx = NewLiteralIntASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(338)
			p.Match(MiniGoParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserFLOAT:
		localctx = NewLiteralFloatASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(339)
			p.Match(MiniGoParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserRUNE:
		localctx = NewLiteralRuneASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(340)
			p.Match(MiniGoParserRUNE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserRAWSTRING:
		localctx = NewLiteralRawsASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(341)
			p.Match(MiniGoParserRAWSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserINTERPRETEDSTRING:
		localctx = NewLiteralInterASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(342)
			p.Match(MiniGoParserINTERPRETEDSTRING)
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

// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_index
	return p
}

func InitEmptyIndexContext(p *IndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_index
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) CopyAll(ctx *IndexContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IndexASTContext struct {
	IndexContext
}

func NewIndexASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexASTContext {
	var p = new(IndexASTContext)

	InitEmptyIndexContext(&p.IndexContext)
	p.parser = parser
	p.CopyAll(ctx.(*IndexContext))

	return p
}

func (s *IndexASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexASTContext) CORCHIZQ() antlr.TerminalNode {
	return s.GetToken(MiniGoParserCORCHIZQ, 0)
}

func (s *IndexASTContext) Expression() IExpressionContext {
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

func (s *IndexASTContext) CORCHDER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserCORCHDER, 0)
}

func (s *IndexASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitIndexAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) Index() (localctx IIndexContext) {
	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, MiniGoParserRULE_index)
	localctx = NewIndexASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(345)
		p.Match(MiniGoParserCORCHIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(346)
		p.expression(0)
	}
	{
		p.SetState(347)
		p.Match(MiniGoParserCORCHDER)
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

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_arguments
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) CopyAll(ctx *ArgumentsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArgumentsASTContext struct {
	ArgumentsContext
}

func NewArgumentsASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArgumentsASTContext {
	var p = new(ArgumentsASTContext)

	InitEmptyArgumentsContext(&p.ArgumentsContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArgumentsContext))

	return p
}

func (s *ArgumentsASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsASTContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARIZQ, 0)
}

func (s *ArgumentsASTContext) PARDER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARDER, 0)
}

func (s *ArgumentsASTContext) ExpressionList() IExpressionListContext {
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

func (s *ArgumentsASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitArgumentsAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, MiniGoParserRULE_arguments)
	var _la int

	localctx = NewArgumentsASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(349)
		p.Match(MiniGoParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1030976700448) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&63) != 0) {
		{
			p.SetState(350)
			p.ExpressionList()
		}

	}
	{
		p.SetState(353)
		p.Match(MiniGoParserPARDER)
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

// ISelectorContext is an interface to support dynamic dispatch.
type ISelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSelectorContext differentiates from other interfaces.
	IsSelectorContext()
}

type SelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorContext() *SelectorContext {
	var p = new(SelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_selector
	return p
}

func InitEmptySelectorContext(p *SelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_selector
}

func (*SelectorContext) IsSelectorContext() {}

func NewSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorContext {
	var p = new(SelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_selector

	return p
}

func (s *SelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorContext) CopyAll(ctx *SelectorContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SelectorASTContext struct {
	SelectorContext
}

func NewSelectorASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectorASTContext {
	var p = new(SelectorASTContext)

	InitEmptySelectorContext(&p.SelectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*SelectorContext))

	return p
}

func (s *SelectorASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorASTContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPUNTO, 0)
}

func (s *SelectorASTContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniGoParserID, 0)
}

func (s *SelectorASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSelectorAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) Selector() (localctx ISelectorContext) {
	localctx = NewSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, MiniGoParserRULE_selector)
	localctx = NewSelectorASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		p.Match(MiniGoParserPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(356)
		p.Match(MiniGoParserID)
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

// IAppendExpressionContext is an interface to support dynamic dispatch.
type IAppendExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAppendExpressionContext differentiates from other interfaces.
	IsAppendExpressionContext()
}

type AppendExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAppendExpressionContext() *AppendExpressionContext {
	var p = new(AppendExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_appendExpression
	return p
}

func InitEmptyAppendExpressionContext(p *AppendExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_appendExpression
}

func (*AppendExpressionContext) IsAppendExpressionContext() {}

func NewAppendExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AppendExpressionContext {
	var p = new(AppendExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_appendExpression

	return p
}

func (s *AppendExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AppendExpressionContext) CopyAll(ctx *AppendExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AppendExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AppendExpressionASTContext struct {
	AppendExpressionContext
}

func NewAppendExpressionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AppendExpressionASTContext {
	var p = new(AppendExpressionASTContext)

	InitEmptyAppendExpressionContext(&p.AppendExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AppendExpressionContext))

	return p
}

func (s *AppendExpressionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendExpressionASTContext) APPEND() antlr.TerminalNode {
	return s.GetToken(MiniGoParserAPPEND, 0)
}

func (s *AppendExpressionASTContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARIZQ, 0)
}

func (s *AppendExpressionASTContext) AllExpression() []IExpressionContext {
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

func (s *AppendExpressionASTContext) Expression(i int) IExpressionContext {
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

func (s *AppendExpressionASTContext) COMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserCOMA, 0)
}

func (s *AppendExpressionASTContext) PARDER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARDER, 0)
}

func (s *AppendExpressionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitAppendExpressionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) AppendExpression() (localctx IAppendExpressionContext) {
	localctx = NewAppendExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, MiniGoParserRULE_appendExpression)
	localctx = NewAppendExpressionASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(358)
		p.Match(MiniGoParserAPPEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(359)
		p.Match(MiniGoParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(360)
		p.expression(0)
	}
	{
		p.SetState(361)
		p.Match(MiniGoParserCOMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(362)
		p.expression(0)
	}
	{
		p.SetState(363)
		p.Match(MiniGoParserPARDER)
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

// ILengthExpressionContext is an interface to support dynamic dispatch.
type ILengthExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLengthExpressionContext differentiates from other interfaces.
	IsLengthExpressionContext()
}

type LengthExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLengthExpressionContext() *LengthExpressionContext {
	var p = new(LengthExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_lengthExpression
	return p
}

func InitEmptyLengthExpressionContext(p *LengthExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_lengthExpression
}

func (*LengthExpressionContext) IsLengthExpressionContext() {}

func NewLengthExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LengthExpressionContext {
	var p = new(LengthExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_lengthExpression

	return p
}

func (s *LengthExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LengthExpressionContext) CopyAll(ctx *LengthExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LengthExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LengthExpressionASTContext struct {
	LengthExpressionContext
}

func NewLengthExpressionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LengthExpressionASTContext {
	var p = new(LengthExpressionASTContext)

	InitEmptyLengthExpressionContext(&p.LengthExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*LengthExpressionContext))

	return p
}

func (s *LengthExpressionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthExpressionASTContext) LEN() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLEN, 0)
}

func (s *LengthExpressionASTContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARIZQ, 0)
}

func (s *LengthExpressionASTContext) Expression() IExpressionContext {
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

func (s *LengthExpressionASTContext) PARDER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARDER, 0)
}

func (s *LengthExpressionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitLengthExpressionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) LengthExpression() (localctx ILengthExpressionContext) {
	localctx = NewLengthExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, MiniGoParserRULE_lengthExpression)
	localctx = NewLengthExpressionASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(365)
		p.Match(MiniGoParserLEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(366)
		p.Match(MiniGoParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(367)
		p.expression(0)
	}
	{
		p.SetState(368)
		p.Match(MiniGoParserPARDER)
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

// ICapExpressionContext is an interface to support dynamic dispatch.
type ICapExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCapExpressionContext differentiates from other interfaces.
	IsCapExpressionContext()
}

type CapExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCapExpressionContext() *CapExpressionContext {
	var p = new(CapExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_capExpression
	return p
}

func InitEmptyCapExpressionContext(p *CapExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_capExpression
}

func (*CapExpressionContext) IsCapExpressionContext() {}

func NewCapExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CapExpressionContext {
	var p = new(CapExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_capExpression

	return p
}

func (s *CapExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *CapExpressionContext) CopyAll(ctx *CapExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *CapExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CapExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CapExpressionASTContext struct {
	CapExpressionContext
}

func NewCapExpressionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CapExpressionASTContext {
	var p = new(CapExpressionASTContext)

	InitEmptyCapExpressionContext(&p.CapExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*CapExpressionContext))

	return p
}

func (s *CapExpressionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CapExpressionASTContext) CAP() antlr.TerminalNode {
	return s.GetToken(MiniGoParserCAP, 0)
}

func (s *CapExpressionASTContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARIZQ, 0)
}

func (s *CapExpressionASTContext) Expression() IExpressionContext {
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

func (s *CapExpressionASTContext) PARDER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARDER, 0)
}

func (s *CapExpressionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitCapExpressionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) CapExpression() (localctx ICapExpressionContext) {
	localctx = NewCapExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, MiniGoParserRULE_capExpression)
	localctx = NewCapExpressionASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		p.Match(MiniGoParserCAP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(371)
		p.Match(MiniGoParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(372)
		p.expression(0)
	}
	{
		p.SetState(373)
		p.Match(MiniGoParserPARDER)
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

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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
	p.RuleIndex = MiniGoParserRULE_statementList
	return p
}

func InitEmptyStatementListContext(p *StatementListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_statementList
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) CopyAll(ctx *StatementListContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StatementListASTContext struct {
	StatementListContext
}

func NewStatementListASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatementListASTContext {
	var p = new(StatementListASTContext)

	InitEmptyStatementListContext(&p.StatementListContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementListContext))

	return p
}

func (s *StatementListASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListASTContext) AllStatement() []IStatementContext {
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

func (s *StatementListASTContext) Statement(i int) IStatementContext {
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

func (s *StatementListASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitStatementListAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) StatementList() (localctx IStatementListContext) {
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, MiniGoParserRULE_statementList)
	var _la int

	localctx = NewStatementListASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7494024895781699892) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&63) != 0) {
		{
			p.SetState(375)
			p.Statement()
		}

		p.SetState(380)
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) CopyAll(ctx *BlockContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BlockASTContext struct {
	BlockContext
}

func NewBlockASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockASTContext {
	var p = new(BlockASTContext)

	InitEmptyBlockContext(&p.BlockContext)
	p.parser = parser
	p.CopyAll(ctx.(*BlockContext))

	return p
}

func (s *BlockASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockASTContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLLAVEIZQ, 0)
}

func (s *BlockASTContext) StatementList() IStatementListContext {
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

func (s *BlockASTContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLLAVEDER, 0)
}

func (s *BlockASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitBlockAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, MiniGoParserRULE_block)
	localctx = NewBlockASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(381)
		p.Match(MiniGoParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(382)
		p.StatementList()
	}
	{
		p.SetState(383)
		p.Match(MiniGoParserLLAVEDER)
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
	p.RuleIndex = MiniGoParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_statement

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

type ContinueStASTContext struct {
	StatementContext
}

func NewContinueStASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStASTContext {
	var p = new(ContinueStASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ContinueStASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStASTContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(MiniGoParserCONTINUE, 0)
}

func (s *ContinueStASTContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, 0)
}

func (s *ContinueStASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitContinueStAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeDeclStASTContext struct {
	StatementContext
}

func NewTypeDeclStASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeDeclStASTContext {
	var p = new(TypeDeclStASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *TypeDeclStASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclStASTContext) TypeDecl() ITypeDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDeclContext)
}

func (s *TypeDeclStASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitTypeDeclStAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type BlockStASTContext struct {
	StatementContext
}

func NewBlockStASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockStASTContext {
	var p = new(BlockStASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *BlockStASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *BlockStASTContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, 0)
}

func (s *BlockStASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitBlockStAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type VariableDeclStASTContext struct {
	StatementContext
}

func NewVariableDeclStASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableDeclStASTContext {
	var p = new(VariableDeclStASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *VariableDeclStASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclStASTContext) VariableDecl() IVariableDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclContext)
}

func (s *VariableDeclStASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitVariableDeclStAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SimpleStatementStASTContext struct {
	StatementContext
}

func NewSimpleStatementStASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleStatementStASTContext {
	var p = new(SimpleStatementStASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *SimpleStatementStASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStatementStASTContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *SimpleStatementStASTContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, 0)
}

func (s *SimpleStatementStASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSimpleStatementStAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type BreakStASTContext struct {
	StatementContext
}

func NewBreakStASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStASTContext {
	var p = new(BreakStASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *BreakStASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStASTContext) BREAK() antlr.TerminalNode {
	return s.GetToken(MiniGoParserBREAK, 0)
}

func (s *BreakStASTContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, 0)
}

func (s *BreakStASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitBreakStAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type LoopStASTContext struct {
	StatementContext
}

func NewLoopStASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LoopStASTContext {
	var p = new(LoopStASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *LoopStASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopStASTContext) Loop() ILoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopContext)
}

func (s *LoopStASTContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, 0)
}

func (s *LoopStASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitLoopStAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintlStASTContext struct {
	StatementContext
}

func NewPrintlStASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintlStASTContext {
	var p = new(PrintlStASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *PrintlStASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintlStASTContext) PRINTLN() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPRINTLN, 0)
}

func (s *PrintlStASTContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARIZQ, 0)
}

func (s *PrintlStASTContext) PARDER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARDER, 0)
}

func (s *PrintlStASTContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, 0)
}

func (s *PrintlStASTContext) ExpressionList() IExpressionListContext {
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

func (s *PrintlStASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitPrintlStAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatementASTContext struct {
	StatementContext
}

func NewStatementASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatementASTContext {
	var p = new(StatementASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *StatementASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementASTContext) PRINT() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPRINT, 0)
}

func (s *StatementASTContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARIZQ, 0)
}

func (s *StatementASTContext) PARDER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPARDER, 0)
}

func (s *StatementASTContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, 0)
}

func (s *StatementASTContext) ExpressionList() IExpressionListContext {
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

func (s *StatementASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitStatementAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfStatementStASTContext struct {
	StatementContext
}

func NewIfStatementStASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStatementStASTContext {
	var p = new(IfStatementStASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *IfStatementStASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementStASTContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *IfStatementStASTContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, 0)
}

func (s *IfStatementStASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitIfStatementStAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStASTContext struct {
	StatementContext
}

func NewReturnStASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStASTContext {
	var p = new(ReturnStASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ReturnStASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStASTContext) RETURN() antlr.TerminalNode {
	return s.GetToken(MiniGoParserRETURN, 0)
}

func (s *ReturnStASTContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, 0)
}

func (s *ReturnStASTContext) Expression() IExpressionContext {
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

func (s *ReturnStASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitReturnStAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SwitchStASTContext struct {
	StatementContext
}

func NewSwitchStASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchStASTContext {
	var p = new(SwitchStASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *SwitchStASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStASTContext) Switch_() ISwitchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchContext)
}

func (s *SwitchStASTContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, 0)
}

func (s *SwitchStASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSwitchStAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, MiniGoParserRULE_statement)
	var _la int

	p.SetState(425)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MiniGoParserPRINT:
		localctx = NewStatementASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(385)
			p.Match(MiniGoParserPRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(386)
			p.Match(MiniGoParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(388)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1030976700448) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&63) != 0) {
			{
				p.SetState(387)
				p.ExpressionList()
			}

		}
		{
			p.SetState(390)
			p.Match(MiniGoParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(391)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserPRINTLN:
		localctx = NewPrintlStASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(392)
			p.Match(MiniGoParserPRINTLN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(393)
			p.Match(MiniGoParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(395)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1030976700448) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&63) != 0) {
			{
				p.SetState(394)
				p.ExpressionList()
			}

		}
		{
			p.SetState(397)
			p.Match(MiniGoParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(398)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserRETURN:
		localctx = NewReturnStASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(399)
			p.Match(MiniGoParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(401)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1030976700448) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&63) != 0) {
			{
				p.SetState(400)
				p.expression(0)
			}

		}
		{
			p.SetState(403)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserBREAK:
		localctx = NewBreakStASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(404)
			p.Match(MiniGoParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(405)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserCONTINUE:
		localctx = NewContinueStASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(406)
			p.Match(MiniGoParserCONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(407)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserPyCOMA, MiniGoParserPARIZQ, MiniGoParserSUMA, MiniGoParserRESTA, MiniGoParserXOR, MiniGoParserEXCL, MiniGoParserAPPEND, MiniGoParserLEN, MiniGoParserCAP, MiniGoParserID, MiniGoParserINT, MiniGoParserFLOAT, MiniGoParserRUNE, MiniGoParserRAWSTRING, MiniGoParserINTERPRETEDSTRING:
		localctx = NewSimpleStatementStASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(408)
			p.SimpleStatement()
		}
		{
			p.SetState(409)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserLLAVEIZQ:
		localctx = NewBlockStASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(411)
			p.Block()
		}
		{
			p.SetState(412)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserSWITCH:
		localctx = NewSwitchStASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(414)
			p.Switch_()
		}
		{
			p.SetState(415)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserIF:
		localctx = NewIfStatementStASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(417)
			p.IfStatement()
		}
		{
			p.SetState(418)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserFOR:
		localctx = NewLoopStASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(420)
			p.Loop()
		}
		{
			p.SetState(421)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserTYPE:
		localctx = NewTypeDeclStASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(423)
			p.TypeDecl()
		}

	case MiniGoParserVAR:
		localctx = NewVariableDeclStASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(424)
			p.VariableDecl()
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

// ISimpleStatementContext is an interface to support dynamic dispatch.
type ISimpleStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSimpleStatementContext differentiates from other interfaces.
	IsSimpleStatementContext()
}

type SimpleStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleStatementContext() *SimpleStatementContext {
	var p = new(SimpleStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_simpleStatement
	return p
}

func InitEmptySimpleStatementContext(p *SimpleStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_simpleStatement
}

func (*SimpleStatementContext) IsSimpleStatementContext() {}

func NewSimpleStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleStatementContext {
	var p = new(SimpleStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_simpleStatement

	return p
}

func (s *SimpleStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleStatementContext) CopyAll(ctx *SimpleStatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SimpleStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssignSimpStASTContext struct {
	SimpleStatementContext
}

func NewAssignSimpStASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignSimpStASTContext {
	var p = new(AssignSimpStASTContext)

	InitEmptySimpleStatementContext(&p.SimpleStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*SimpleStatementContext))

	return p
}

func (s *AssignSimpStASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignSimpStASTContext) AllExpressionList() []IExpressionListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionListContext); ok {
			len++
		}
	}

	tst := make([]IExpressionListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionListContext); ok {
			tst[i] = t.(IExpressionListContext)
			i++
		}
	}

	return tst
}

func (s *AssignSimpStASTContext) ExpressionList(i int) IExpressionListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
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

	return t.(IExpressionListContext)
}

func (s *AssignSimpStASTContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MiniGoParserASSIGN, 0)
}

func (s *AssignSimpStASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitAssignSimpStAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentStatementSimpStASTContext struct {
	SimpleStatementContext
}

func NewAssignmentStatementSimpStASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentStatementSimpStASTContext {
	var p = new(AssignmentStatementSimpStASTContext)

	InitEmptySimpleStatementContext(&p.SimpleStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*SimpleStatementContext))

	return p
}

func (s *AssignmentStatementSimpStASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementSimpStASTContext) AssignmentStatement() IAssignmentStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentStatementContext)
}

func (s *AssignmentStatementSimpStASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitAssignmentStatementSimpStAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SimpleStatementASTContext struct {
	SimpleStatementContext
}

func NewSimpleStatementASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleStatementASTContext {
	var p = new(SimpleStatementASTContext)

	InitEmptySimpleStatementContext(&p.SimpleStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*SimpleStatementContext))

	return p
}

func (s *SimpleStatementASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStatementASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSimpleStatementAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionSimpStASTContext struct {
	SimpleStatementContext
}

func NewExpressionSimpStASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionSimpStASTContext {
	var p = new(ExpressionSimpStASTContext)

	InitEmptySimpleStatementContext(&p.SimpleStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*SimpleStatementContext))

	return p
}

func (s *ExpressionSimpStASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionSimpStASTContext) Expression() IExpressionContext {
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

func (s *ExpressionSimpStASTContext) SUMAD() antlr.TerminalNode {
	return s.GetToken(MiniGoParserSUMAD, 0)
}

func (s *ExpressionSimpStASTContext) RESTAD() antlr.TerminalNode {
	return s.GetToken(MiniGoParserRESTAD, 0)
}

func (s *ExpressionSimpStASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitExpressionSimpStAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) SimpleStatement() (localctx ISimpleStatementContext) {
	localctx = NewSimpleStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, MiniGoParserRULE_simpleStatement)
	var _la int

	p.SetState(437)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSimpleStatementASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)

	case 2:
		localctx = NewExpressionSimpStASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(428)
			p.expression(0)
		}
		p.SetState(430)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MiniGoParserSUMAD || _la == MiniGoParserRESTAD {
			{
				p.SetState(429)
				_la = p.GetTokenStream().LA(1)

				if !(_la == MiniGoParserSUMAD || _la == MiniGoParserRESTAD) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

	case 3:
		localctx = NewAssignmentStatementSimpStASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(432)
			p.AssignmentStatement()
		}

	case 4:
		localctx = NewAssignSimpStASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(433)
			p.ExpressionList()
		}
		{
			p.SetState(434)
			p.Match(MiniGoParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(435)
			p.ExpressionList()
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

// IAssignmentStatementContext is an interface to support dynamic dispatch.
type IAssignmentStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssignmentStatementContext differentiates from other interfaces.
	IsAssignmentStatementContext()
}

type AssignmentStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentStatementContext() *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_assignmentStatement
	return p
}

func InitEmptyAssignmentStatementContext(p *AssignmentStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_assignmentStatement
}

func (*AssignmentStatementContext) IsAssignmentStatementContext() {}

func NewAssignmentStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_assignmentStatement

	return p
}

func (s *AssignmentStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentStatementContext) CopyAll(ctx *AssignmentStatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AssignmentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IncremAsSTASTContext struct {
	AssignmentStatementContext
}

func NewIncremAsSTASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IncremAsSTASTContext {
	var p = new(IncremAsSTASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *IncremAsSTASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncremAsSTASTContext) AllExpression() []IExpressionContext {
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

func (s *IncremAsSTASTContext) Expression(i int) IExpressionContext {
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

func (s *IncremAsSTASTContext) INCREM() antlr.TerminalNode {
	return s.GetToken(MiniGoParserINCREM, 0)
}

func (s *IncremAsSTASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitIncremAsSTAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignDivAsSTASTContext struct {
	AssignmentStatementContext
}

func NewAssignDivAsSTASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignDivAsSTASTContext {
	var p = new(AssignDivAsSTASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *AssignDivAsSTASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignDivAsSTASTContext) AllExpression() []IExpressionContext {
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

func (s *AssignDivAsSTASTContext) Expression(i int) IExpressionContext {
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

func (s *AssignDivAsSTASTContext) ASSINGDIV() antlr.TerminalNode {
	return s.GetToken(MiniGoParserASSINGDIV, 0)
}

func (s *AssignDivAsSTASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitAssignDivAsSTAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssingnXorAsSTASTContext struct {
	AssignmentStatementContext
}

func NewAssingnXorAsSTASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssingnXorAsSTASTContext {
	var p = new(AssingnXorAsSTASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *AssingnXorAsSTASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssingnXorAsSTASTContext) AllExpression() []IExpressionContext {
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

func (s *AssingnXorAsSTASTContext) Expression(i int) IExpressionContext {
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

func (s *AssingnXorAsSTASTContext) ASSIGNXOR() antlr.TerminalNode {
	return s.GetToken(MiniGoParserASSIGNXOR, 0)
}

func (s *AssingnXorAsSTASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitAssingnXorAsSTAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignMultAsSTASTContext struct {
	AssignmentStatementContext
}

func NewAssignMultAsSTASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignMultAsSTASTContext {
	var p = new(AssignMultAsSTASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *AssignMultAsSTASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignMultAsSTASTContext) AllExpression() []IExpressionContext {
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

func (s *AssignMultAsSTASTContext) Expression(i int) IExpressionContext {
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

func (s *AssignMultAsSTASTContext) ASSIGNMULT() antlr.TerminalNode {
	return s.GetToken(MiniGoParserASSIGNMULT, 0)
}

func (s *AssignMultAsSTASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitAssignMultAsSTAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndCompAsSTASTContext struct {
	AssignmentStatementContext
}

func NewAndCompAsSTASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndCompAsSTASTContext {
	var p = new(AndCompAsSTASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *AndCompAsSTASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndCompAsSTASTContext) AllExpression() []IExpressionContext {
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

func (s *AndCompAsSTASTContext) Expression(i int) IExpressionContext {
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

func (s *AndCompAsSTASTContext) ANDCOMP() antlr.TerminalNode {
	return s.GetToken(MiniGoParserANDCOMP, 0)
}

func (s *AndCompAsSTASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitAndCompAsSTAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignBitAsSTASTContext struct {
	AssignmentStatementContext
}

func NewAssignBitAsSTASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignBitAsSTASTContext {
	var p = new(AssignBitAsSTASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *AssignBitAsSTASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignBitAsSTASTContext) AllExpression() []IExpressionContext {
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

func (s *AssignBitAsSTASTContext) Expression(i int) IExpressionContext {
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

func (s *AssignBitAsSTASTContext) ASSIGNBIT() antlr.TerminalNode {
	return s.GetToken(MiniGoParserASSIGNBIT, 0)
}

func (s *AssignBitAsSTASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitAssignBitAsSTAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignPorAsSTASTContext struct {
	AssignmentStatementContext
}

func NewAssignPorAsSTASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignPorAsSTASTContext {
	var p = new(AssignPorAsSTASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *AssignPorAsSTASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignPorAsSTASTContext) AllExpression() []IExpressionContext {
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

func (s *AssignPorAsSTASTContext) Expression(i int) IExpressionContext {
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

func (s *AssignPorAsSTASTContext) ASSIGNPOR() antlr.TerminalNode {
	return s.GetToken(MiniGoParserASSIGNPOR, 0)
}

func (s *AssignPorAsSTASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitAssignPorAsSTAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentStatementASTContext struct {
	AssignmentStatementContext
}

func NewAssignmentStatementASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentStatementASTContext {
	var p = new(AssignmentStatementASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *AssignmentStatementASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementASTContext) AllExpressionList() []IExpressionListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionListContext); ok {
			len++
		}
	}

	tst := make([]IExpressionListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionListContext); ok {
			tst[i] = t.(IExpressionListContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentStatementASTContext) ExpressionList(i int) IExpressionListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
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

	return t.(IExpressionListContext)
}

func (s *AssignmentStatementASTContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(MiniGoParserIGUAL, 0)
}

func (s *AssignmentStatementASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitAssignmentStatementAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignNorAsSTASTContext struct {
	AssignmentStatementContext
}

func NewAssignNorAsSTASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignNorAsSTASTContext {
	var p = new(AssignNorAsSTASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *AssignNorAsSTASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignNorAsSTASTContext) AllExpression() []IExpressionContext {
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

func (s *AssignNorAsSTASTContext) Expression(i int) IExpressionContext {
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

func (s *AssignNorAsSTASTContext) ASSIGNOR() antlr.TerminalNode {
	return s.GetToken(MiniGoParserASSIGNOR, 0)
}

func (s *AssignNorAsSTASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitAssignNorAsSTAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignMaAsSTASTContext struct {
	AssignmentStatementContext
}

func NewAssignMaAsSTASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignMaAsSTASTContext {
	var p = new(AssignMaAsSTASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *AssignMaAsSTASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignMaAsSTASTContext) AllExpression() []IExpressionContext {
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

func (s *AssignMaAsSTASTContext) Expression(i int) IExpressionContext {
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

func (s *AssignMaAsSTASTContext) ASSIGMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserASSIGMA, 0)
}

func (s *AssignMaAsSTASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitAssignMaAsSTAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type DecremeAsSTASTContext struct {
	AssignmentStatementContext
}

func NewDecremeAsSTASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecremeAsSTASTContext {
	var p = new(DecremeAsSTASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *DecremeAsSTASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecremeAsSTASTContext) AllExpression() []IExpressionContext {
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

func (s *DecremeAsSTASTContext) Expression(i int) IExpressionContext {
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

func (s *DecremeAsSTASTContext) DECREME() antlr.TerminalNode {
	return s.GetToken(MiniGoParserDECREME, 0)
}

func (s *DecremeAsSTASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitDecremeAsSTAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignMeAsSTASTContext struct {
	AssignmentStatementContext
}

func NewAssignMeAsSTASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignMeAsSTASTContext {
	var p = new(AssignMeAsSTASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *AssignMeAsSTASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignMeAsSTASTContext) AllExpression() []IExpressionContext {
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

func (s *AssignMeAsSTASTContext) Expression(i int) IExpressionContext {
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

func (s *AssignMeAsSTASTContext) ASSIGNME() antlr.TerminalNode {
	return s.GetToken(MiniGoParserASSIGNME, 0)
}

func (s *AssignMeAsSTASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitAssignMeAsSTAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) AssignmentStatement() (localctx IAssignmentStatementContext) {
	localctx = NewAssignmentStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, MiniGoParserRULE_assignmentStatement)
	p.SetState(487)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignmentStatementASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(439)
			p.ExpressionList()
		}
		{
			p.SetState(440)
			p.Match(MiniGoParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(441)
			p.ExpressionList()
		}

	case 2:
		localctx = NewIncremAsSTASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(443)
			p.expression(0)
		}
		{
			p.SetState(444)
			p.Match(MiniGoParserINCREM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(445)
			p.expression(0)
		}

	case 3:
		localctx = NewAndCompAsSTASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(447)
			p.expression(0)
		}
		{
			p.SetState(448)
			p.Match(MiniGoParserANDCOMP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(449)
			p.expression(0)
		}

	case 4:
		localctx = NewDecremeAsSTASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(451)
			p.expression(0)
		}
		{
			p.SetState(452)
			p.Match(MiniGoParserDECREME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(453)
			p.expression(0)
		}

	case 5:
		localctx = NewAssignNorAsSTASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(455)
			p.expression(0)
		}
		{
			p.SetState(456)
			p.Match(MiniGoParserASSIGNOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(457)
			p.expression(0)
		}

	case 6:
		localctx = NewAssignMultAsSTASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(459)
			p.expression(0)
		}
		{
			p.SetState(460)
			p.Match(MiniGoParserASSIGNMULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(461)
			p.expression(0)
		}

	case 7:
		localctx = NewAssingnXorAsSTASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(463)
			p.expression(0)
		}
		{
			p.SetState(464)
			p.Match(MiniGoParserASSIGNXOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(465)
			p.expression(0)
		}

	case 8:
		localctx = NewAssignMaAsSTASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(467)
			p.expression(0)
		}
		{
			p.SetState(468)
			p.Match(MiniGoParserASSIGMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(469)
			p.expression(0)
		}

	case 9:
		localctx = NewAssignMeAsSTASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(471)
			p.expression(0)
		}
		{
			p.SetState(472)
			p.Match(MiniGoParserASSIGNME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(473)
			p.expression(0)
		}

	case 10:
		localctx = NewAssignBitAsSTASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(475)
			p.expression(0)
		}
		{
			p.SetState(476)
			p.Match(MiniGoParserASSIGNBIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(477)
			p.expression(0)
		}

	case 11:
		localctx = NewAssignPorAsSTASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(479)
			p.expression(0)
		}
		{
			p.SetState(480)
			p.Match(MiniGoParserASSIGNPOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(481)
			p.expression(0)
		}

	case 12:
		localctx = NewAssignDivAsSTASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(483)
			p.expression(0)
		}
		{
			p.SetState(484)
			p.Match(MiniGoParserASSINGDIV)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(485)
			p.expression(0)
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

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) CopyAll(ctx *IfStatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfElseStASTContext struct {
	IfStatementContext
}

func NewIfElseStASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfElseStASTContext {
	var p = new(IfElseStASTContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IfElseStASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseStASTContext) IF() antlr.TerminalNode {
	return s.GetToken(MiniGoParserIF, 0)
}

func (s *IfElseStASTContext) Expression() IExpressionContext {
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

func (s *IfElseStASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfElseStASTContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MiniGoParserELSE, 0)
}

func (s *IfElseStASTContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *IfElseStASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitIfElseStAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfElseExpIfstASTContext struct {
	IfStatementContext
}

func NewIfElseExpIfstASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfElseExpIfstASTContext {
	var p = new(IfElseExpIfstASTContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IfElseExpIfstASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseExpIfstASTContext) IF() antlr.TerminalNode {
	return s.GetToken(MiniGoParserIF, 0)
}

func (s *IfElseExpIfstASTContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *IfElseExpIfstASTContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, 0)
}

func (s *IfElseExpIfstASTContext) Expression() IExpressionContext {
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

func (s *IfElseExpIfstASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfElseExpIfstASTContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MiniGoParserELSE, 0)
}

func (s *IfElseExpIfstASTContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *IfElseExpIfstASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitIfElseExpIfstAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfStatementASTContext struct {
	IfStatementContext
}

func NewIfStatementASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStatementASTContext {
	var p = new(IfStatementASTContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IfStatementASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementASTContext) IF() antlr.TerminalNode {
	return s.GetToken(MiniGoParserIF, 0)
}

func (s *IfStatementASTContext) Expression() IExpressionContext {
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

func (s *IfStatementASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStatementASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitIfStatementAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfElseBlockStASTContext struct {
	IfStatementContext
}

func NewIfElseBlockStASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfElseBlockStASTContext {
	var p = new(IfElseBlockStASTContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IfElseBlockStASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseBlockStASTContext) IF() antlr.TerminalNode {
	return s.GetToken(MiniGoParserIF, 0)
}

func (s *IfElseBlockStASTContext) Expression() IExpressionContext {
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

func (s *IfElseBlockStASTContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfElseBlockStASTContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
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

	return t.(IBlockContext)
}

func (s *IfElseBlockStASTContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MiniGoParserELSE, 0)
}

func (s *IfElseBlockStASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitIfElseBlockStAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfElseExpressionBlASTContext struct {
	IfStatementContext
}

func NewIfElseExpressionBlASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfElseExpressionBlASTContext {
	var p = new(IfElseExpressionBlASTContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IfElseExpressionBlASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseExpressionBlASTContext) IF() antlr.TerminalNode {
	return s.GetToken(MiniGoParserIF, 0)
}

func (s *IfElseExpressionBlASTContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *IfElseExpressionBlASTContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, 0)
}

func (s *IfElseExpressionBlASTContext) Expression() IExpressionContext {
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

func (s *IfElseExpressionBlASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfElseExpressionBlASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitIfElseExpressionBlAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfElseExpBlockBASTContext struct {
	IfStatementContext
}

func NewIfElseExpBlockBASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfElseExpBlockBASTContext {
	var p = new(IfElseExpBlockBASTContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IfElseExpBlockBASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseExpBlockBASTContext) IF() antlr.TerminalNode {
	return s.GetToken(MiniGoParserIF, 0)
}

func (s *IfElseExpBlockBASTContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *IfElseExpBlockBASTContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, 0)
}

func (s *IfElseExpBlockBASTContext) Expression() IExpressionContext {
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

func (s *IfElseExpBlockBASTContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfElseExpBlockBASTContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
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

	return t.(IBlockContext)
}

func (s *IfElseExpBlockBASTContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MiniGoParserELSE, 0)
}

func (s *IfElseExpBlockBASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitIfElseExpBlockBAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, MiniGoParserRULE_ifStatement)
	p.SetState(527)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIfStatementASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(489)
			p.Match(MiniGoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(490)
			p.expression(0)
		}
		{
			p.SetState(491)
			p.Block()
		}

	case 2:
		localctx = NewIfElseStASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(493)
			p.Match(MiniGoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(494)
			p.expression(0)
		}
		{
			p.SetState(495)
			p.Block()
		}
		{
			p.SetState(496)
			p.Match(MiniGoParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(497)
			p.IfStatement()
		}

	case 3:
		localctx = NewIfElseBlockStASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(499)
			p.Match(MiniGoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(500)
			p.expression(0)
		}
		{
			p.SetState(501)
			p.Block()
		}
		{
			p.SetState(502)
			p.Match(MiniGoParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(503)
			p.Block()
		}

	case 4:
		localctx = NewIfElseExpressionBlASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(505)
			p.Match(MiniGoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(506)
			p.SimpleStatement()
		}
		{
			p.SetState(507)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(508)
			p.expression(0)
		}
		{
			p.SetState(509)
			p.Block()
		}

	case 5:
		localctx = NewIfElseExpIfstASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(511)
			p.Match(MiniGoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(512)
			p.SimpleStatement()
		}
		{
			p.SetState(513)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(514)
			p.expression(0)
		}
		{
			p.SetState(515)
			p.Block()
		}
		{
			p.SetState(516)
			p.Match(MiniGoParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(517)
			p.IfStatement()
		}

	case 6:
		localctx = NewIfElseExpBlockBASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(519)
			p.Match(MiniGoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(520)
			p.SimpleStatement()
		}
		{
			p.SetState(521)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(522)
			p.expression(0)
		}
		{
			p.SetState(523)
			p.Block()
		}
		{
			p.SetState(524)
			p.Match(MiniGoParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(525)
			p.Block()
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

// ILoopContext is an interface to support dynamic dispatch.
type ILoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLoopContext differentiates from other interfaces.
	IsLoopContext()
}

type LoopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopContext() *LoopContext {
	var p = new(LoopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_loop
	return p
}

func InitEmptyLoopContext(p *LoopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_loop
}

func (*LoopContext) IsLoopContext() {}

func NewLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopContext {
	var p = new(LoopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_loop

	return p
}

func (s *LoopContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopContext) CopyAll(ctx *LoopContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForBlLoopASTContext struct {
	LoopContext
}

func NewForBlLoopASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForBlLoopASTContext {
	var p = new(ForBlLoopASTContext)

	InitEmptyLoopContext(&p.LoopContext)
	p.parser = parser
	p.CopyAll(ctx.(*LoopContext))

	return p
}

func (s *ForBlLoopASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForBlLoopASTContext) FOR() antlr.TerminalNode {
	return s.GetToken(MiniGoParserFOR, 0)
}

func (s *ForBlLoopASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForBlLoopASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitForBlLoopAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForSimpStEXLoopASTContext struct {
	LoopContext
}

func NewForSimpStEXLoopASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForSimpStEXLoopASTContext {
	var p = new(ForSimpStEXLoopASTContext)

	InitEmptyLoopContext(&p.LoopContext)
	p.parser = parser
	p.CopyAll(ctx.(*LoopContext))

	return p
}

func (s *ForSimpStEXLoopASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForSimpStEXLoopASTContext) FOR() antlr.TerminalNode {
	return s.GetToken(MiniGoParserFOR, 0)
}

func (s *ForSimpStEXLoopASTContext) AllSimpleStatement() []ISimpleStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			len++
		}
	}

	tst := make([]ISimpleStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimpleStatementContext); ok {
			tst[i] = t.(ISimpleStatementContext)
			i++
		}
	}

	return tst
}

func (s *ForSimpStEXLoopASTContext) SimpleStatement(i int) ISimpleStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
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

	return t.(ISimpleStatementContext)
}

func (s *ForSimpStEXLoopASTContext) AllPyCOMA() []antlr.TerminalNode {
	return s.GetTokens(MiniGoParserPyCOMA)
}

func (s *ForSimpStEXLoopASTContext) PyCOMA(i int) antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, i)
}

func (s *ForSimpStEXLoopASTContext) Expression() IExpressionContext {
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

func (s *ForSimpStEXLoopASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForSimpStEXLoopASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitForSimpStEXLoopAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForExpBlLoopASTContext struct {
	LoopContext
}

func NewForExpBlLoopASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForExpBlLoopASTContext {
	var p = new(ForExpBlLoopASTContext)

	InitEmptyLoopContext(&p.LoopContext)
	p.parser = parser
	p.CopyAll(ctx.(*LoopContext))

	return p
}

func (s *ForExpBlLoopASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForExpBlLoopASTContext) FOR() antlr.TerminalNode {
	return s.GetToken(MiniGoParserFOR, 0)
}

func (s *ForExpBlLoopASTContext) Expression() IExpressionContext {
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

func (s *ForExpBlLoopASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForExpBlLoopASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitForExpBlLoopAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForSimpStSimpStLoopASTContext struct {
	LoopContext
}

func NewForSimpStSimpStLoopASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForSimpStSimpStLoopASTContext {
	var p = new(ForSimpStSimpStLoopASTContext)

	InitEmptyLoopContext(&p.LoopContext)
	p.parser = parser
	p.CopyAll(ctx.(*LoopContext))

	return p
}

func (s *ForSimpStSimpStLoopASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForSimpStSimpStLoopASTContext) FOR() antlr.TerminalNode {
	return s.GetToken(MiniGoParserFOR, 0)
}

func (s *ForSimpStSimpStLoopASTContext) AllSimpleStatement() []ISimpleStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			len++
		}
	}

	tst := make([]ISimpleStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimpleStatementContext); ok {
			tst[i] = t.(ISimpleStatementContext)
			i++
		}
	}

	return tst
}

func (s *ForSimpStSimpStLoopASTContext) SimpleStatement(i int) ISimpleStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
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

	return t.(ISimpleStatementContext)
}

func (s *ForSimpStSimpStLoopASTContext) AllPyCOMA() []antlr.TerminalNode {
	return s.GetTokens(MiniGoParserPyCOMA)
}

func (s *ForSimpStSimpStLoopASTContext) PyCOMA(i int) antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, i)
}

func (s *ForSimpStSimpStLoopASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForSimpStSimpStLoopASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitForSimpStSimpStLoopAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) Loop() (localctx ILoopContext) {
	localctx = NewLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, MiniGoParserRULE_loop)
	p.SetState(550)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForBlLoopASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(529)
			p.Match(MiniGoParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(530)
			p.Block()
		}

	case 2:
		localctx = NewForExpBlLoopASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(531)
			p.Match(MiniGoParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(532)
			p.expression(0)
		}
		{
			p.SetState(533)
			p.Block()
		}

	case 3:
		localctx = NewForSimpStEXLoopASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(535)
			p.Match(MiniGoParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(536)
			p.SimpleStatement()
		}
		{
			p.SetState(537)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(538)
			p.expression(0)
		}
		{
			p.SetState(539)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(540)
			p.SimpleStatement()
		}
		{
			p.SetState(541)
			p.Block()
		}

	case 4:
		localctx = NewForSimpStSimpStLoopASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(543)
			p.Match(MiniGoParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(544)
			p.SimpleStatement()
		}
		{
			p.SetState(545)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(546)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(547)
			p.SimpleStatement()
		}
		{
			p.SetState(548)
			p.Block()
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

// ISwitchContext is an interface to support dynamic dispatch.
type ISwitchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSwitchContext differentiates from other interfaces.
	IsSwitchContext()
}

type SwitchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchContext() *SwitchContext {
	var p = new(SwitchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_switch
	return p
}

func InitEmptySwitchContext(p *SwitchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_switch
}

func (*SwitchContext) IsSwitchContext() {}

func NewSwitchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchContext {
	var p = new(SwitchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_switch

	return p
}

func (s *SwitchContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchContext) CopyAll(ctx *SwitchContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SwitchExprASTContext struct {
	SwitchContext
}

func NewSwitchExprASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchExprASTContext {
	var p = new(SwitchExprASTContext)

	InitEmptySwitchContext(&p.SwitchContext)
	p.parser = parser
	p.CopyAll(ctx.(*SwitchContext))

	return p
}

func (s *SwitchExprASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchExprASTContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(MiniGoParserSWITCH, 0)
}

func (s *SwitchExprASTContext) Expression() IExpressionContext {
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

func (s *SwitchExprASTContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLLAVEIZQ, 0)
}

func (s *SwitchExprASTContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseListContext)
}

func (s *SwitchExprASTContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLLAVEDER, 0)
}

func (s *SwitchExprASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSwitchExprAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SwitchSimpStASTContext struct {
	SwitchContext
}

func NewSwitchSimpStASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchSimpStASTContext {
	var p = new(SwitchSimpStASTContext)

	InitEmptySwitchContext(&p.SwitchContext)
	p.parser = parser
	p.CopyAll(ctx.(*SwitchContext))

	return p
}

func (s *SwitchSimpStASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchSimpStASTContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(MiniGoParserSWITCH, 0)
}

func (s *SwitchSimpStASTContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *SwitchSimpStASTContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, 0)
}

func (s *SwitchSimpStASTContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLLAVEIZQ, 0)
}

func (s *SwitchSimpStASTContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseListContext)
}

func (s *SwitchSimpStASTContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLLAVEDER, 0)
}

func (s *SwitchSimpStASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSwitchSimpStAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SwitchASTContext struct {
	SwitchContext
}

func NewSwitchASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchASTContext {
	var p = new(SwitchASTContext)

	InitEmptySwitchContext(&p.SwitchContext)
	p.parser = parser
	p.CopyAll(ctx.(*SwitchContext))

	return p
}

func (s *SwitchASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchASTContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(MiniGoParserSWITCH, 0)
}

func (s *SwitchASTContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *SwitchASTContext) PyCOMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPyCOMA, 0)
}

func (s *SwitchASTContext) Expression() IExpressionContext {
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

func (s *SwitchASTContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLLAVEIZQ, 0)
}

func (s *SwitchASTContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseListContext)
}

func (s *SwitchASTContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLLAVEDER, 0)
}

func (s *SwitchASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSwitchAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SwitchExCaseASTContext struct {
	SwitchContext
}

func NewSwitchExCaseASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchExCaseASTContext {
	var p = new(SwitchExCaseASTContext)

	InitEmptySwitchContext(&p.SwitchContext)
	p.parser = parser
	p.CopyAll(ctx.(*SwitchContext))

	return p
}

func (s *SwitchExCaseASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchExCaseASTContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(MiniGoParserSWITCH, 0)
}

func (s *SwitchExCaseASTContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLLAVEIZQ, 0)
}

func (s *SwitchExCaseASTContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseListContext)
}

func (s *SwitchExCaseASTContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLLAVEDER, 0)
}

func (s *SwitchExCaseASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSwitchExCaseAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) Switch_() (localctx ISwitchContext) {
	localctx = NewSwitchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, MiniGoParserRULE_switch)
	p.SetState(578)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSwitchASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(552)
			p.Match(MiniGoParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(553)
			p.SimpleStatement()
		}
		{
			p.SetState(554)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(555)
			p.expression(0)
		}
		{
			p.SetState(556)
			p.Match(MiniGoParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(557)
			p.ExpressionCaseClauseList()
		}
		{
			p.SetState(558)
			p.Match(MiniGoParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewSwitchExprASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(560)
			p.Match(MiniGoParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(561)
			p.expression(0)
		}
		{
			p.SetState(562)
			p.Match(MiniGoParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(563)
			p.ExpressionCaseClauseList()
		}
		{
			p.SetState(564)
			p.Match(MiniGoParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewSwitchSimpStASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(566)
			p.Match(MiniGoParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(567)
			p.SimpleStatement()
		}
		{
			p.SetState(568)
			p.Match(MiniGoParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(569)
			p.Match(MiniGoParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(570)
			p.ExpressionCaseClauseList()
		}
		{
			p.SetState(571)
			p.Match(MiniGoParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewSwitchExCaseASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(573)
			p.Match(MiniGoParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(574)
			p.Match(MiniGoParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(575)
			p.ExpressionCaseClauseList()
		}
		{
			p.SetState(576)
			p.Match(MiniGoParserLLAVEDER)
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

// IExpressionCaseClauseListContext is an interface to support dynamic dispatch.
type IExpressionCaseClauseListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionCaseClauseListContext differentiates from other interfaces.
	IsExpressionCaseClauseListContext()
}

type ExpressionCaseClauseListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionCaseClauseListContext() *ExpressionCaseClauseListContext {
	var p = new(ExpressionCaseClauseListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_expressionCaseClauseList
	return p
}

func InitEmptyExpressionCaseClauseListContext(p *ExpressionCaseClauseListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_expressionCaseClauseList
}

func (*ExpressionCaseClauseListContext) IsExpressionCaseClauseListContext() {}

func NewExpressionCaseClauseListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionCaseClauseListContext {
	var p = new(ExpressionCaseClauseListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_expressionCaseClauseList

	return p
}

func (s *ExpressionCaseClauseListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionCaseClauseListContext) CopyAll(ctx *ExpressionCaseClauseListContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionCaseClauseListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionCaseClauseListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionCaseClauseListASTContext struct {
	ExpressionCaseClauseListContext
}

func NewExpressionCaseClauseListASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionCaseClauseListASTContext {
	var p = new(ExpressionCaseClauseListASTContext)

	InitEmptyExpressionCaseClauseListContext(&p.ExpressionCaseClauseListContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionCaseClauseListContext))

	return p
}

func (s *ExpressionCaseClauseListASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionCaseClauseListASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitExpressionCaseClauseListAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionCaseClauseExpASTContext struct {
	ExpressionCaseClauseListContext
}

func NewExpressionCaseClauseExpASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionCaseClauseExpASTContext {
	var p = new(ExpressionCaseClauseExpASTContext)

	InitEmptyExpressionCaseClauseListContext(&p.ExpressionCaseClauseListContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionCaseClauseListContext))

	return p
}

func (s *ExpressionCaseClauseExpASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionCaseClauseExpASTContext) ExpressionCaseClause() IExpressionCaseClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseContext)
}

func (s *ExpressionCaseClauseExpASTContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseListContext)
}

func (s *ExpressionCaseClauseExpASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitExpressionCaseClauseExpAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) ExpressionCaseClauseList() (localctx IExpressionCaseClauseListContext) {
	localctx = NewExpressionCaseClauseListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, MiniGoParserRULE_expressionCaseClauseList)
	p.SetState(584)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MiniGoParserLLAVEDER:
		localctx = NewExpressionCaseClauseListASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)

	case MiniGoParserCASE, MiniGoParserDEFAULT:
		localctx = NewExpressionCaseClauseExpASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(581)
			p.ExpressionCaseClause()
		}
		{
			p.SetState(582)
			p.ExpressionCaseClauseList()
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

// IExpressionCaseClauseContext is an interface to support dynamic dispatch.
type IExpressionCaseClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionCaseClauseContext differentiates from other interfaces.
	IsExpressionCaseClauseContext()
}

type ExpressionCaseClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionCaseClauseContext() *ExpressionCaseClauseContext {
	var p = new(ExpressionCaseClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_expressionCaseClause
	return p
}

func InitEmptyExpressionCaseClauseContext(p *ExpressionCaseClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_expressionCaseClause
}

func (*ExpressionCaseClauseContext) IsExpressionCaseClauseContext() {}

func NewExpressionCaseClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionCaseClauseContext {
	var p = new(ExpressionCaseClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_expressionCaseClause

	return p
}

func (s *ExpressionCaseClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionCaseClauseContext) CopyAll(ctx *ExpressionCaseClauseContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionCaseClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionCaseClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionCaseClauseASTContext struct {
	ExpressionCaseClauseContext
}

func NewExpressionCaseClauseASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionCaseClauseASTContext {
	var p = new(ExpressionCaseClauseASTContext)

	InitEmptyExpressionCaseClauseContext(&p.ExpressionCaseClauseContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionCaseClauseContext))

	return p
}

func (s *ExpressionCaseClauseASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionCaseClauseASTContext) ExpressionSwitchCase() IExpressionSwitchCaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSwitchCaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSwitchCaseContext)
}

func (s *ExpressionCaseClauseASTContext) DOSP() antlr.TerminalNode {
	return s.GetToken(MiniGoParserDOSP, 0)
}

func (s *ExpressionCaseClauseASTContext) StatementList() IStatementListContext {
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

func (s *ExpressionCaseClauseASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitExpressionCaseClauseAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) ExpressionCaseClause() (localctx IExpressionCaseClauseContext) {
	localctx = NewExpressionCaseClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, MiniGoParserRULE_expressionCaseClause)
	localctx = NewExpressionCaseClauseASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(586)
		p.ExpressionSwitchCase()
	}
	{
		p.SetState(587)
		p.Match(MiniGoParserDOSP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(588)
		p.StatementList()
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

// IExpressionSwitchCaseContext is an interface to support dynamic dispatch.
type IExpressionSwitchCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionSwitchCaseContext differentiates from other interfaces.
	IsExpressionSwitchCaseContext()
}

type ExpressionSwitchCaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionSwitchCaseContext() *ExpressionSwitchCaseContext {
	var p = new(ExpressionSwitchCaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_expressionSwitchCase
	return p
}

func InitEmptyExpressionSwitchCaseContext(p *ExpressionSwitchCaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_expressionSwitchCase
}

func (*ExpressionSwitchCaseContext) IsExpressionSwitchCaseContext() {}

func NewExpressionSwitchCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionSwitchCaseContext {
	var p = new(ExpressionSwitchCaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_expressionSwitchCase

	return p
}

func (s *ExpressionSwitchCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionSwitchCaseContext) CopyAll(ctx *ExpressionSwitchCaseContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionSwitchCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionSwitchCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DefaultExpressionSASTContext struct {
	ExpressionSwitchCaseContext
}

func NewDefaultExpressionSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefaultExpressionSASTContext {
	var p = new(DefaultExpressionSASTContext)

	InitEmptyExpressionSwitchCaseContext(&p.ExpressionSwitchCaseContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionSwitchCaseContext))

	return p
}

func (s *DefaultExpressionSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultExpressionSASTContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(MiniGoParserDEFAULT, 0)
}

func (s *DefaultExpressionSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitDefaultExpressionSAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionSwitchCaseASTContext struct {
	ExpressionSwitchCaseContext
}

func NewExpressionSwitchCaseASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionSwitchCaseASTContext {
	var p = new(ExpressionSwitchCaseASTContext)

	InitEmptyExpressionSwitchCaseContext(&p.ExpressionSwitchCaseContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionSwitchCaseContext))

	return p
}

func (s *ExpressionSwitchCaseASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionSwitchCaseASTContext) CASE() antlr.TerminalNode {
	return s.GetToken(MiniGoParserCASE, 0)
}

func (s *ExpressionSwitchCaseASTContext) ExpressionList() IExpressionListContext {
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

func (s *ExpressionSwitchCaseASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitExpressionSwitchCaseAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) ExpressionSwitchCase() (localctx IExpressionSwitchCaseContext) {
	localctx = NewExpressionSwitchCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, MiniGoParserRULE_expressionSwitchCase)
	p.SetState(593)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MiniGoParserCASE:
		localctx = NewExpressionSwitchCaseASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(590)
			p.Match(MiniGoParserCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(591)
			p.ExpressionList()
		}

	case MiniGoParserDEFAULT:
		localctx = NewDefaultExpressionSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(592)
			p.Match(MiniGoParserDEFAULT)
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

func (p *MiniGoParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 18:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 20:
		var t *PrimaryExpressionContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryExpressionContext)
		}
		return p.PrimaryExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MiniGoParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 23)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 16:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 17:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 18:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *MiniGoParser) PrimaryExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 19:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 20:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 21:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
