// Code generated by goyacc - DO NOT EDIT.



import __yyfmt__ "fmt"

#include <stdio.h>
#include <stdlib.h>

type yySymType 

type yyXError struct {
	state, xsym int
}

const (
	yyDefault   = 57385
	yyEofCode   = 57344
	ADDOP       = 57375
	AND         = 57368
	BOOL        = 57350
	BREAK       = 57383
	CONST       = 57352
	CONTINUE    = 57384
	COOK        = 57359
	DEC         = 57371
	DIVOP       = 57378
	ECHO        = 57374
	ELSE        = 57366
	ELSE_IF     = 57367
	EQ          = 57379
	FALSE       = 57373
	FLAG        = 57353
	FLOOP       = 57357
	FOR         = 57381
	FUNC        = 57346
	GE          = 57360
	GT          = 57363
	ID          = 57347
	IF          = 57365
	INC         = 57370
	LE          = 57361
	LT          = 57364
	MAIN_TOKEN  = 57351
	MULOP       = 57377
	NEQ         = 57362
	NUM         = 57354
	NUM_FLAG    = 57348
	OR          = 57369
	RETURN      = 57382
	STR_FLAG    = 57349
	SUBOP       = 57376
	TRUE        = 57372
	VAL1        = 57355
	WHILE       = 57380
	WLOOP       = 57358
	_digit      = 57356
	yyErrCode   = 57345

	yyMaxDepth = 200
	yyTabOfs   = -90
)

var (

	yyPrec = map[int]int{
		ADDOP: 0,
		SUBOP: 0,
		MULOP: 1,
		DIVOP: 1,
		AND: 2,
		OR: 2,
		EQ: 3,
		NEQ: 3,
		LT: 3,
		LE: 3,
		GT: 3,
		GE: 3,
		'(': 4,
		')': 4,
		',': 5,
		'=': 6,
		IF: 7,
		ELSE: 7,
		WHILE: 7,
		FOR: 7,
		RETURN: 7,
		BREAK: 7,
		CONTINUE: 7,
		'{': 7,
		'}': 7,
		}

	yyXLAT = map[int]int{
		    59:   0, // ';' (63x)
		 57347:   1, // ID (55x)
		    41:   2, // ')' (38x)
		   125:   3, // '}' (33x)
		 57353:   4, // FLAG (32x)
		    40:   5, // '(' (30x)
		 57348:   6, // NUM_FLAG (24x)
		 57349:   7, // STR_FLAG (24x)
		 57354:   8, // NUM (21x)
		 57368:   9, // AND (18x)
		 57359:  10, // COOK (18x)
		 57374:  11, // ECHO (18x)
		 57357:  12, // FLOOP (18x)
		 57365:  13, // IF (18x)
		 57369:  14, // OR (18x)
		 57358:  15, // WLOOP (18x)
		 57397:  16, // dt (17x)
		    42:  17, // '*' (13x)
		    43:  18, // '+' (13x)
		    45:  19, // '-' (13x)
		    47:  20, // '/' (13x)
		 57401:  21, // exp (12x)
		 57355:  22, // VAL1 (12x)
		   123:  23, // '{' (11x)
		 57403:  24, // floop (11x)
		 57405:  25, // if_cond (11x)
		 57407:  26, // l (11x)
		 57411:  27, // stmt (11x)
		 57412:  28, // stmts (11x)
		 57413:  29, // wloop (11x)
		    61:  30, // '=' (10x)
		 57360:  31, // GE (10x)
		 57363:  32, // GT (10x)
		 57361:  33, // LE (10x)
		 57364:  34, // LT (10x)
		 57362:  35, // NEQ (10x)
		 57394:  36, // c (8x)
		 57373:  37, // FALSE (8x)
		 57372:  38, // TRUE (8x)
		 57344:  39, // $end (6x)
		 57402:  40, // f1 (6x)
		 57409:  41, // op (6x)
		    33:  42, // '!' (4x)
		    44:  43, // ',' (4x)
		 57396:  44, // cond (4x)
		 57398:  45, // e1 (4x)
		    91:  46, // '[' (2x)
		    93:  47, // ']' (2x)
		 57395:  48, // comp (2x)
		 57352:  49, // CONST (2x)
		 57371:  50, // DEC (2x)
		 57370:  51, // INC (2x)
		 57393:  52, // assign (1x)
		 57366:  53, // ELSE (1x)
		 57399:  54, // else_cond (1x)
		 57367:  55, // ELSE_IF (1x)
		 57400:  56, // else_if_cond (1x)
		 57346:  57, // FUNC (1x)
		 57404:  58, // func_decl (1x)
		 57406:  59, // iter (1x)
		 57408:  60, // main_func (1x)
		 57351:  61, // MAIN_TOKEN (1x)
		 57410:  62, // program (1x)
		 57385:  63, // $default (0x)
		    48:  64, // '0' (0x)
		 57356:  65, // _digit (0x)
		 57375:  66, // ADDOP (0x)
		 57350:  67, // BOOL (0x)
		 57383:  68, // BREAK (0x)
		 57384:  69, // CONTINUE (0x)
		 57378:  70, // DIVOP (0x)
		 57379:  71, // EQ (0x)
		 57345:  72, // error (0x)
		 57381:  73, // FOR (0x)
		 57377:  74, // MULOP (0x)
		 57382:  75, // RETURN (0x)
		 57376:  76, // SUBOP (0x)
		 57386:  77, // VAL2 (0x)
		 57387:  78, // VAL2_T (0x)
		 57388:  79, // VAL3 (0x)
		 57389:  80, // VAL4 (0x)
		 57390:  81, // VAL6 (0x)
		 57391:  82, // VAL6_T (0x)
		 57392:  83, // VAL7 (0x)
		 57380:  84, // WHILE (0x)
	}

	yySymNames = []string{
		"';'",
		"ID",
		"')'",
		"'}'",
		"FLAG",
		"'('",
		"NUM_FLAG",
		"STR_FLAG",
		"NUM",
		"AND",
		"COOK",
		"ECHO",
		"FLOOP",
		"IF",
		"OR",
		"WLOOP",
		"dt",
		"'*'",
		"'+'",
		"'-'",
		"'/'",
		"exp",
		"VAL1",
		"'{'",
		"floop",
		"if_cond",
		"l",
		"stmt",
		"stmts",
		"wloop",
		"'='",
		"GE",
		"GT",
		"LE",
		"LT",
		"NEQ",
		"c",
		"FALSE",
		"TRUE",
		"$end",
		"f1",
		"op",
		"'!'",
		"','",
		"cond",
		"e1",
		"'['",
		"']'",
		"comp",
		"CONST",
		"DEC",
		"INC",
		"assign",
		"ELSE",
		"else_cond",
		"ELSE_IF",
		"else_if_cond",
		"FUNC",
		"func_decl",
		"iter",
		"main_func",
		"MAIN_TOKEN",
		"program",
		"$default",
		"'0'",
		"_digit",
		"ADDOP",
		"BOOL",
		"BREAK",
		"CONTINUE",
		"DIVOP",
		"EQ",
		"error",
		"FOR",
		"MULOP",
		"RETURN",
		"SUBOP",
		"VAL2",
		"VAL2_T",
		"VAL3",
		"VAL4",
		"VAL6",
		"VAL6_T",
		"VAL7",
		"WHILE",
	}

	yyTokenLiteralStrings = map[int]string{
	}

	yyReductions = map[int]struct{xsym, components int}{
		0: {0, 1},
		1: {62, 0},
		2: {62, 1},
		3: {62, 1},
		4: {58, 8},
		5: {60, 8},
		6: {45, 0},
		7: {45, 3},
		8: {45, 2},
		9: {45, 0},
		10: {16, 1},
		11: {16, 1},
		12: {16, 1},
		13: {28, 3},
		14: {28, 2},
		15: {28, 2},
		16: {28, 2},
		17: {28, 0},
		18: {24, 11},
		19: {29, 7},
		20: {52, 4},
		21: {52, 4},
		22: {52, 3},
		23: {52, 3},
		24: {52, 2},
		25: {52, 2},
		26: {44, 2},
		27: {44, 1},
		28: {36, 3},
		29: {36, 3},
		30: {36, 3},
		31: {36, 3},
		32: {36, 1},
		33: {36, 1},
		34: {36, 1},
		35: {36, 1},
		36: {36, 1},
		37: {48, 1},
		38: {48, 1},
		39: {48, 1},
		40: {48, 1},
		41: {48, 1},
		42: {21, 1},
		43: {21, 1},
		44: {21, 2},
		45: {21, 3},
		46: {40, 3},
		47: {40, 0},
		48: {41, 1},
		49: {41, 1},
		50: {41, 1},
		51: {41, 1},
		52: {25, 8},
		53: {56, 8},
		54: {56, 0},
		55: {54, 4},
		56: {54, 0},
		57: {59, 3},
		58: {59, 2},
		59: {59, 2},
		60: {27, 1},
		61: {27, 1},
		62: {27, 1},
		63: {27, 1},
		64: {27, 0},
		65: {26, 5},
		66: {26, 5},
		67: {26, 5},
		68: {26, 7},
		69: {26, 7},
		70: {26, 7},
		71: {26, 7},
		72: {26, 6},
		73: {26, 3},
		74: {26, 3},
		75: {26, 3},
		76: {26, 3},
		77: {26, 3},
		78: {26, 3},
		79: {26, 3},
		80: {26, 2},
		81: {26, 3},
		82: {26, 3},
		83: {26, 3},
		84: {26, 5},
		85: {26, 5},
		86: {26, 5},
		87: {26, 6},
		88: {26, 5},
		89: {26, 3},
		90: {77, 2},
		91: {77, 1},
		92: {78, 2},
		93: {78, 0},
		94: {79, 1},
		95: {79, 1},
		96: {80, 3},
		97: {80, 3},
		98: {80, 3},
		99: {80, 1},
		100: {80, 1},
		101: {80, 1},
		102: {81, 2},
		103: {81, 0},
		104: {82, 2},
		105: {82, 0},
		106: {83, 3},
		107: {83, 3},
		108: {83, 3},
		109: {83, 1},
		110: {83, 1},
		111: {83, 1},
	}

	yyXErrors = map[yyXError]string{
	}

	yyParseTab = [179][]uint16{
		// 0
		{39: 89, 57: 94, 92, 60: 93, 62: 91},
		{39: 90},
		{39: 88},
		{39: 87},
		{1: 95, 61: 96},
		// 5
		{5: 263},
		{5: 97},
		{2: 84, 4: 102, 6: 101, 103, 16: 99, 43: 100, 45: 98},
		{2: 107},
		{1: 105},
		// 10
		{2: 84, 4: 102, 6: 101, 103, 16: 99, 43: 100, 45: 104},
		{1: 80},
		{1: 79},
		{1: 78},
		{2: 82},
		// 15
		{2: 84, 4: 102, 6: 101, 103, 16: 99, 43: 100, 45: 106},
		{2: 83},
		{23: 108},
		{26, 119, 3: 73, 102, 6: 101, 103, 10: 118, 121, 114, 116, 15: 115, 120, 24: 112, 113, 117, 110, 109, 111},
		{3: 262},
		// 20
		{260},
		{29, 119, 3: 73, 102, 6: 101, 103, 10: 118, 121, 114, 116, 15: 115, 120, 24: 112, 113, 117, 110, 259, 111},
		{28, 119, 3: 73, 102, 6: 101, 103, 10: 118, 121, 114, 116, 15: 115, 120, 24: 112, 113, 117, 110, 258, 111},
		{27, 119, 3: 73, 102, 6: 101, 103, 10: 118, 121, 114, 116, 15: 115, 120, 24: 112, 113, 117, 110, 257, 111},
		{5: 231},
		// 25
		{5: 225},
		{5: 169},
		{30},
		{1: 141, 4: 138, 6: 137, 139, 16: 140},
		{30: 129},
		// 30
		{1: 124},
		{5: 122},
		{2: 123},
		{1},
		{30: 125},
		// 35
		{46: 126},
		{22: 127},
		{47: 128},
		{3},
		{5: 132, 22: 130, 131},
		// 40
		{9},
		{22: 135},
		{22: 133},
		{2: 134},
		{2},
		// 45
		{3: 136},
		{6},
		{1: 163},
		{1: 157},
		{1: 151},
		// 50
		{1: 146},
		{10, 30: 142},
		{5: 143},
		{22: 144},
		{2: 145},
		// 55
		{18},
		{11, 30: 147},
		{46: 148},
		{22: 149},
		{47: 150},
		// 60
		{19},
		{17, 30: 152},
		{22: 153, 154},
		{25},
		{22: 155},
		// 65
		{3: 156},
		{21},
		{15, 30: 158},
		{22: 159, 160},
		{23},
		// 70
		{22: 161},
		{3: 162},
		{20},
		{16, 30: 164},
		{22: 165, 166},
		// 75
		{24},
		{22: 167},
		{3: 168},
		{22},
		{1: 174, 4: 175, 172, 8: 176, 21: 173, 36: 171, 178, 177, 42: 170, 44: 179},
		// 80
		{1: 174, 4: 175, 172, 8: 176, 21: 173, 36: 224, 178, 177},
		{63, 2: 63, 9: 220, 14: 221},
		{1: 174, 4: 175, 172, 8: 176, 21: 218, 36: 217, 178, 177},
		{17: 207, 205, 206, 208, 31: 198, 201, 199, 202, 200, 40: 203, 204, 48: 197},
		{58, 2: 58, 9: 58, 14: 58, 17: 48, 48, 48, 48, 31: 48, 48, 48, 48, 48},
		// 85
		{57, 2: 57, 9: 57, 14: 57},
		{56, 2: 56, 9: 56, 14: 56, 17: 47, 47, 47, 47, 31: 47, 47, 47, 47, 47},
		{55, 2: 55, 9: 55, 14: 55},
		{54, 2: 54, 9: 54, 14: 54},
		{2: 180},
		// 90
		{23: 181},
		{26, 119, 3: 73, 102, 6: 101, 103, 10: 118, 121, 114, 116, 15: 115, 120, 24: 112, 113, 117, 110, 182, 111},
		{3: 183},
		{36, 36, 3: 36, 36, 6: 36, 36, 10: 36, 36, 36, 36, 15: 36, 55: 185, 184},
		{38, 38, 3: 38, 38, 6: 38, 38, 10: 38, 38, 38, 38, 15: 38},
		// 95
		{5: 186},
		{1: 174, 4: 175, 172, 8: 176, 21: 173, 36: 171, 178, 177, 42: 170, 44: 187},
		{2: 188},
		{23: 189},
		{26, 119, 3: 73, 102, 6: 101, 103, 10: 118, 121, 114, 116, 15: 115, 120, 24: 112, 113, 117, 110, 190, 111},
		// 100
		{3: 191},
		{34, 34, 3: 34, 34, 6: 34, 34, 10: 34, 34, 34, 34, 15: 34, 53: 193, 192},
		{37, 37, 3: 37, 37, 6: 37, 37, 10: 37, 37, 37, 37, 15: 37},
		{23: 194},
		{26, 119, 3: 73, 102, 6: 101, 103, 10: 118, 121, 114, 116, 15: 115, 120, 24: 112, 113, 117, 110, 195, 111},
		// 105
		{3: 196},
		{35, 35, 3: 35, 35, 6: 35, 35, 10: 35, 35, 35, 35, 15: 35},
		{1: 209, 5: 212, 8: 210, 21: 216},
		{1: 53, 5: 53, 8: 53},
		{1: 52, 5: 52, 8: 52},
		// 110
		{1: 51, 5: 51, 8: 51},
		{1: 50, 5: 50, 8: 50},
		{1: 49, 5: 49, 8: 49},
		{46, 2: 46, 9: 46, 14: 46, 17: 46, 46, 46, 46, 31: 46, 46, 46, 46, 46},
		{1: 209, 5: 212, 8: 210, 21: 211},
		// 115
		{1: 42, 5: 42, 8: 42},
		{1: 41, 5: 41, 8: 41},
		{1: 40, 5: 40, 8: 40},
		{1: 39, 5: 39, 8: 39},
		{48, 2: 48, 9: 48, 14: 48, 17: 48, 48, 48, 48, 31: 48, 48, 48, 48, 48},
		// 120
		{47, 2: 47, 9: 47, 14: 47, 17: 47, 47, 47, 47, 31: 47, 47, 47, 47, 47},
		{43, 2: 43, 9: 43, 14: 43, 17: 207, 205, 206, 208, 31: 43, 43, 43, 43, 43, 40: 215, 204},
		{1: 209, 5: 212, 8: 210, 21: 213},
		{2: 214, 17: 207, 205, 206, 208, 40: 203, 204},
		{45, 2: 45, 9: 45, 14: 45, 17: 45, 45, 45, 45, 31: 45, 45, 45, 45, 45},
		// 125
		{46, 2: 46, 9: 46, 14: 46, 17: 46, 46, 46, 46, 31: 46, 46, 46, 46, 46},
		{61, 2: 61, 9: 61, 14: 61, 17: 207, 205, 206, 208, 40: 203, 204},
		{2: 219, 9: 220, 14: 221},
		{2: 214, 17: 207, 205, 206, 208, 31: 198, 201, 199, 202, 200, 40: 203, 204, 48: 197},
		{62, 2: 62, 9: 62, 14: 62},
		// 130
		{1: 174, 4: 175, 172, 8: 176, 21: 173, 36: 223, 178, 177},
		{1: 174, 4: 175, 172, 8: 176, 21: 173, 36: 222, 178, 177},
		{59, 2: 59, 9: 59, 14: 59},
		{60, 2: 60, 9: 60, 14: 60},
		{64, 2: 64, 9: 220, 14: 221},
		// 135
		{1: 174, 4: 175, 172, 8: 176, 21: 173, 36: 171, 178, 177, 42: 170, 44: 226},
		{2: 227},
		{23: 228},
		{26, 119, 3: 73, 102, 6: 101, 103, 10: 118, 121, 114, 116, 15: 115, 120, 24: 112, 113, 117, 110, 229, 111},
		{3: 230},
		// 140
		{71, 71, 3: 71, 71, 6: 71, 71, 10: 71, 71, 71, 71, 15: 71},
		{1: 234, 4: 102, 6: 101, 103, 16: 233, 52: 232},
		{244},
		{1: 240},
		{30: 235, 50: 237, 236},
		// 145
		{1: 239, 49: 238},
		{66},
		{65},
		{68},
		{67},
		// 150
		{30: 241},
		{1: 243, 49: 242},
		{70},
		{69},
		{1: 174, 4: 175, 172, 8: 176, 21: 173, 36: 171, 178, 177, 42: 170, 44: 245},
		// 155
		{246},
		{1: 248, 59: 247},
		{2: 253},
		{30: 249, 50: 251, 250},
		{1: 209, 5: 212, 8: 210, 21: 252},
		// 160
		{2: 32},
		{2: 31},
		{2: 43, 17: 207, 205, 206, 208, 40: 203, 204},
		{23: 254},
		{26, 119, 3: 73, 102, 6: 101, 103, 10: 118, 121, 114, 116, 15: 115, 120, 24: 112, 113, 117, 110, 255, 111},
		// 165
		{3: 256},
		{72, 72, 3: 72, 72, 6: 72, 72, 10: 72, 72, 72, 72, 15: 72},
		{3: 74},
		{3: 75},
		{3: 76},
		// 170
		{26, 119, 3: 73, 102, 6: 101, 103, 10: 118, 121, 114, 116, 15: 115, 120, 24: 112, 113, 117, 110, 261, 111},
		{3: 77},
		{39: 85},
		{2: 84, 4: 102, 6: 101, 103, 16: 99, 43: 100, 45: 264},
		{2: 265},
		// 175
		{23: 266},
		{26, 119, 3: 73, 102, 6: 101, 103, 10: 118, 121, 114, 116, 15: 115, 120, 24: 112, 113, 117, 110, 267, 111},
		{3: 268},
		{39: 86},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	if c < 0x7f {
		return __yyfmt__.Sprintf("%q", c)
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), lval: %+v\n", yySymName(n), n, n, lval)
	}
	return n
}
	
func yyParse(yylex yyLexer) int {
	const yyError = 72

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)


	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() { 
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yylval.yys = yystate
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := yyTokenLiteralStrings[yychar]
				if ls == "" {
					ls = yySymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError])+yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x])+yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}



// Lexical part
_digit : [0-9];
VAL1: '\"' [^'\"] '\"'
    ;
NUM_FLAG:   [1-9]
    ;
