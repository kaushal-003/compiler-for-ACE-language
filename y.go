// Code generated by goyacc grammar.y. DO NOT EDIT.

//line grammar.y:1

package main

import __yyfmt__ "fmt"

//line grammar.y:3

//line grammar.y:4
type yySymType struct {
	yys        int
	FUNC       string
	ID         string
	NUM_FLAG   string
	STR_FLAG   string
	MAIN_TOKEN string
	CONST      string
	FLAG       string
	NUM        string
	VAL1       string
	FLOOP      string
	WLOOP      string
	COOK       string
	GE         string
	LE         string
	NEQ        string
	GT         string
	LT         string
	IF         string
	ELSE       string
	ELSE_IF    string
	AND        string
	OR         string
	INC        string
	DEC        string
	TRUE       string
	FALSE      string
	ECHO       string
	_digit     string
	CHAR       string
	CONTINUE   string
	BREAK      string
	RETURN     string
	VAL9       string
	NUM1       string
	HEADOF     string
	TAILOF     string
	LEN        string
	COLON      string
	SEMI_COL   string
	L_PAREN    string
	R_PAREN    string
	L_CB       string
	R_CB       string
	L_B        string
	R_B        string
	COMMA      string
	ASSIGN_EQ  string
	EQ         string
	NOT        string
	PLUS       string
	MINUS      string
	MULT       string
	DIV        string
	ZERO       string
}

const FUNC = 57346
const ID = 57347
const NUM_FLAG = 57348
const STR_FLAG = 57349
const MAIN_TOKEN = 57350
const CONST = 57351
const FLAG = 57352
const NUM = 57353
const VAL1 = 57354
const _digit = 57355
const CHAR = 57356
const CONTINUE = 57357
const BREAK = 57358
const RETURN = 57359
const VAL9 = 57360
const NUM1 = 57361
const HEADOF = 57362
const LEN = 57363
const TAILOF = 57364
const FLOOP = 57365
const WLOOP = 57366
const COOK = 57367
const GE = 57368
const LE = 57369
const NEQ = 57370
const GT = 57371
const LT = 57372
const IF = 57373
const ELSE = 57374
const ELSE_IF = 57375
const AND = 57376
const OR = 57377
const INC = 57378
const DEC = 57379
const TRUE = 57380
const FALSE = 57381
const ECHO = 57382
const COLON = 57383
const SEMI_COL = 57384
const L_PAREN = 57385
const R_PAREN = 57386
const L_CB = 57387
const R_CB = 57388
const L_B = 57389
const R_B = 57390
const COMMA = 57391
const ASSIGN_EQ = 57392
const EQ = 57393
const NOT = 57394
const PLUS = 57395
const MINUS = 57396
const MULT = 57397
const DIV = 57398
const ZERO = 57399
const ADDOP = 57400
const SUBOP = 57401
const MULOP = 57402
const DIVOP = 57403
const WHILE = 57404
const FOR = 57405

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"FUNC",
	"ID",
	"NUM_FLAG",
	"STR_FLAG",
	"MAIN_TOKEN",
	"CONST",
	"FLAG",
	"NUM",
	"VAL1",
	"_digit",
	"CHAR",
	"CONTINUE",
	"BREAK",
	"RETURN",
	"VAL9",
	"NUM1",
	"HEADOF",
	"LEN",
	"TAILOF",
	"FLOOP",
	"WLOOP",
	"COOK",
	"GE",
	"LE",
	"NEQ",
	"GT",
	"LT",
	"IF",
	"ELSE",
	"ELSE_IF",
	"AND",
	"OR",
	"INC",
	"DEC",
	"TRUE",
	"FALSE",
	"ECHO",
	"COLON",
	"SEMI_COL",
	"L_PAREN",
	"R_PAREN",
	"L_CB",
	"R_CB",
	"L_B",
	"R_B",
	"COMMA",
	"ASSIGN_EQ",
	"EQ",
	"NOT",
	"PLUS",
	"MINUS",
	"MULT",
	"DIV",
	"ZERO",
	"ADDOP",
	"SUBOP",
	"MULOP",
	"DIVOP",
	"WHILE",
	"FOR",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line grammar.y:324

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 27,
	46, 26,
	-2, 75,
	-1, 28,
	46, 26,
	-2, 75,
	-1, 37,
	46, 26,
	-2, 75,
	-1, 38,
	46, 26,
	-2, 75,
	-1, 39,
	46, 26,
	-2, 75,
	-1, 57,
	46, 26,
	-2, 75,
	-1, 177,
	46, 26,
	-2, 75,
	-1, 189,
	46, 26,
	-2, 75,
	-1, 300,
	46, 26,
	-2, 75,
	-1, 307,
	46, 26,
	-2, 75,
	-1, 312,
	46, 26,
	-2, 75,
}

const yyPrivate = 57344

const yyLast = 372

var yyAct = [...]int{
	35, 90, 94, 254, 250, 115, 245, 114, 240, 112,
	167, 92, 161, 16, 110, 135, 136, 137, 138, 139,
	194, 135, 136, 137, 138, 139, 193, 157, 117, 51,
	156, 186, 170, 181, 117, 191, 190, 192, 58, 59,
	60, 169, 140, 141, 142, 143, 108, 111, 140, 141,
	142, 143, 282, 116, 113, 30, 116, 113, 77, 155,
	196, 117, 154, 71, 117, 181, 118, 103, 86, 279,
	153, 275, 118, 152, 140, 141, 142, 143, 101, 23,
	88, 89, 271, 88, 89, 140, 141, 142, 143, 166,
	285, 286, 165, 164, 147, 148, 132, 126, 162, 118,
	31, 98, 118, 128, 284, 131, 95, 99, 146, 74,
	75, 76, 151, 61, 249, 150, 62, 24, 22, 26,
	297, 25, 312, 295, 293, 163, 258, 257, 256, 253,
	252, 248, 247, 243, 96, 97, 182, 184, 242, 93,
	228, 178, 179, 176, 244, 159, 185, 84, 91, 158,
	239, 204, 120, 83, 85, 119, 33, 125, 195, 314,
	205, 309, 81, 80, 82, 98, 305, 280, 213, 214,
	95, 99, 220, 222, 216, 217, 218, 215, 229, 219,
	88, 89, 44, 277, 88, 89, 230, 273, 199, 87,
	234, 206, 41, 42, 43, 200, 269, 263, 96, 97,
	47, 46, 49, 93, 241, 260, 209, 98, 48, 98,
	129, 130, 251, 99, 73, 99, 223, 50, 221, 56,
	180, 34, 302, 307, 98, 300, 189, 177, 201, 28,
	99, 27, 237, 210, 306, 291, 290, 289, 283, 259,
	226, 225, 224, 160, 149, 183, 236, 183, 255, 127,
	32, 21, 20, 235, 123, 122, 121, 72, 65, 64,
	63, 10, 183, 9, 231, 144, 57, 267, 238, 174,
	124, 88, 89, 129, 130, 288, 311, 14, 168, 241,
	294, 246, 296, 12, 298, 251, 299, 301, 70, 66,
	67, 276, 272, 68, 102, 17, 19, 69, 268, 18,
	175, 303, 7, 233, 304, 8, 29, 232, 308, 52,
	15, 17, 19, 313, 211, 18, 207, 202, 197, 78,
	212, 292, 208, 203, 198, 79, 17, 19, 188, 266,
	18, 265, 187, 264, 262, 227, 173, 172, 171, 145,
	109, 107, 106, 105, 104, 55, 54, 53, 1, 5,
	281, 278, 274, 6, 270, 45, 40, 310, 287, 134,
	133, 261, 100, 39, 38, 37, 36, 13, 11, 4,
	3, 2,
}

var yyPact = [...]int{
	345, -1000, -1000, -1000, 345, 297, -1000, 220, 218, 305,
	208, 207, -1000, 69, -1000, -1000, 74, -1000, -1000, -1000,
	186, 184, 305, 51, 206, 108, 175, 177, 177, -1000,
	-1000, 320, 342, 341, 340, 173, 224, 177, 177, 177,
	-1000, -1000, -1000, -1000, 66, -1000, 217, 216, 215, 283,
	214, 168, -1000, 51, 51, 51, -1000, 177, -1000, -1000,
	-1000, 314, 142, 96, 289, 96, 339, 338, 337, 336,
	-4, 335, 42, -1000, -1000, -1000, -1000, -1000, 107, 104,
	213, 212, 211, 229, 110, -1000, -1000, 45, -1000, -1000,
	205, 160, 239, 160, -5, -1000, -1000, -1000, -1000, -1000,
	223, 334, 58, 200, 65, 23, 12, -20, 102, -1000,
	199, 49, 49, 44, 43, 40, -1000, 265, -1000, -9,
	-18, 333, 332, 331, 228, 288, 95, 182, 239, 160,
	160, 176, -11, 219, 219, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 96, -19, 323, -1000, -1000, 181,
	15, 313, 183, 312, 146, 311, 188, 309, 45, 45,
	-1000, -1000, 42, -1000, 45, 45, 45, -1000, 265, 204,
	202, 198, 197, 196, 330, 92, -1000, 177, -1000, -1000,
	-1000, -1000, 32, 219, 32, 222, 298, -1000, -1000, 177,
	210, 203, 189, 227, 103, -1000, 9, 90, 85, 97,
	-1000, 269, 84, 83, 67, -1000, 233, 82, 81, -1000,
	230, 80, 79, 78, 195, -1000, -1000, -1000, -1000, -1000,
	32, -1000, 32, -1000, -1000, -1000, -1000, -1000, -1000, 159,
	21, 329, -1000, -1000, 151, 328, 326, 324, 226, 286,
	150, 33, -1000, -1000, 280, 141, 22, -1000, -1000, 279,
	137, 20, -1000, -1000, 121, 3, -1000, -1000, -1000, -1000,
	-1000, 194, 54, 242, 193, 192, 191, 316, 76, -1000,
	-1000, 9, 75, -1000, -1000, 269, 72, -1000, -1000, 233,
	-1000, -1000, 230, 180, 219, -1000, -1000, -1000, 179, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	177, 32, 96, 120, 190, -1000, 178, 177, 115, 244,
	-1000, 77, 177, 113, -1000,
}

var yyPgo = [...]int{
	0, 348, 371, 370, 369, 368, 0, 283, 367, 277,
	13, 55, 366, 365, 364, 363, 362, 1, 361, 11,
	2, 360, 359, 358, 357, 356, 355, 14, 12, 9,
	7, 10, 5, 8, 6, 4, 3, 354, 352, 351,
	350,
}

var yyR1 = [...]int{
	0, 1, 1, 2, 2, 4, 3, 5, 5, 7,
	7, 7, 8, 9, 9, 9, 9, 11, 11, 10,
	10, 10, 6, 6, 6, 6, 6, 14, 13, 16,
	16, 16, 16, 16, 16, 17, 17, 19, 19, 19,
	19, 19, 19, 19, 21, 21, 21, 21, 21, 20,
	20, 20, 20, 20, 22, 22, 22, 22, 15, 23,
	23, 24, 24, 18, 18, 18, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 27, 27, 28, 28,
	30, 30, 31, 31, 26, 26, 26, 26, 26, 26,
	26, 26, 26, 26, 26, 26, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 32, 32, 33, 33, 37,
	37, 34, 34, 38, 38, 35, 35, 39, 39, 29,
	29, 29, 29, 29, 29, 36, 36, 40, 40,
}

var yyR2 = [...]int{
	0, 1, 1, 2, 0, 8, 7, 1, 0, 3,
	1, 1, 1, 3, 5, 5, 5, 2, 0, 1,
	1, 1, 3, 2, 2, 2, 0, 11, 7, 4,
	4, 3, 3, 2, 2, 2, 1, 3, 3, 3,
	3, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 3, 0, 1, 1, 1, 1, 8, 8,
	0, 4, 0, 3, 2, 2, 1, 1, 1, 1,
	6, 6, 6, 6, 1, 0, 2, 2, 2, 0,
	2, 1, 2, 0, 8, 8, 8, 8, 6, 6,
	6, 6, 8, 8, 8, 6, 5, 5, 5, 5,
	7, 7, 7, 7, 6, 6, 3, 3, 3, 3,
	6, 6, 6, 6, 6, 6, 6, 6, 3, 2,
	3, 3, 3, 5, 4, 1, 1, 2, 0, 2,
	0, 2, 0, 2, 0, 2, 0, 2, 0, 3,
	3, 3, 1, 1, 1, 2, 0, 2, 0,
}

var yyChk = [...]int{
	-1000, -1, -2, -3, -4, 4, -1, 5, 8, 43,
	43, -5, -7, -8, -9, 5, -10, 6, 10, 7,
	44, 44, 49, 5, 43, 47, 45, 45, 45, -7,
	-11, 49, 44, 48, 46, -6, -12, -13, -14, -15,
	-25, 15, 16, 17, 5, -26, 24, 23, 31, 25,
	40, -6, -9, 5, 5, 5, 46, 42, -6, -6,
	-6, 47, 50, 43, 43, 43, 6, 7, 10, 14,
	5, -10, 43, 46, -11, -11, -11, -6, 5, 11,
	21, 20, 22, 11, 5, 12, -32, 47, 38, 39,
	-17, 52, -19, 43, -20, 10, 38, 39, 5, 11,
	-16, -10, 5, -17, 5, 5, 5, 5, 50, 5,
	-27, 5, -29, 12, -30, -32, 11, 19, 57, 48,
	48, 43, 43, 43, 41, 47, -29, 44, -19, 34,
	35, -19, -20, -21, -22, 26, 27, 28, 29, 30,
	53, 54, 55, 56, 42, 5, 50, 36, 37, 44,
	50, 47, 50, 47, 50, 47, 50, 47, 47, 43,
	44, -28, 49, -28, 49, 49, 49, -31, 13, 50,
	50, 5, 5, 5, 41, 12, 48, 45, -19, -19,
	44, 44, -20, 43, -20, -17, 50, 9, 5, 45,
	21, 20, 22, 11, 5, -30, 45, 5, 11, 5,
	12, 45, 5, 11, 5, -32, 45, 5, 11, 18,
	45, 5, 11, -29, -29, -27, -29, -29, -29, -31,
	-20, 14, -20, 14, 44, 44, 44, 5, 48, -6,
	-20, 42, 9, 5, -6, 43, 43, 43, 41, 47,
	-33, -30, 48, 48, 47, -34, 12, 48, 48, 47,
	-35, -32, 48, 48, -36, 18, 48, 48, 48, 44,
	46, -18, 5, 46, 5, 5, 5, 41, 12, 46,
	-37, 49, 12, 46, -38, 49, 12, 46, -39, 49,
	46, -40, 49, 44, 50, 36, 37, -23, 33, 44,
	44, 44, 5, 48, -33, 48, -34, 48, -35, -36,
	45, -20, 43, -6, -17, 46, 44, 45, -6, 46,
	-24, 32, 45, -6, 46,
}

var yyDef = [...]int{
	4, -2, 1, 2, 4, 0, 3, 0, 0, 8,
	0, 0, 7, 10, 11, 12, 0, 19, 20, 21,
	0, 0, 0, 18, 0, 0, 0, -2, -2, 9,
	13, 0, 0, 0, 0, 0, 0, -2, -2, -2,
	66, 67, 68, 69, 0, 74, 0, 0, 0, 0,
	0, 0, 17, 18, 18, 18, 6, -2, 23, 24,
	25, 0, 0, 53, 0, 53, 0, 0, 0, 0,
	119, 0, 0, 5, 14, 15, 16, 22, 0, 0,
	0, 0, 0, 121, 0, 120, 122, 0, 125, 126,
	0, 53, 36, 53, 0, 41, 42, 43, 49, 50,
	0, 0, 0, 0, 109, 106, 107, 108, 0, 118,
	0, 79, 79, 142, 0, 144, 143, 83, 81, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 35, 53,
	53, 0, 0, 53, 53, 44, 45, 46, 47, 48,
	54, 55, 56, 57, 53, 0, 0, 33, 34, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	124, 76, 0, 77, 0, 0, 0, 80, 83, 53,
	53, 0, 0, 0, 0, 0, 123, -2, 39, 40,
	37, 52, 38, 53, 51, 0, 0, 31, 32, -2,
	0, 0, 0, 0, 0, 97, 128, 0, 0, 0,
	96, 132, 0, 0, 0, 98, 136, 0, 0, 99,
	146, 0, 0, 0, 0, 78, 139, 140, 141, 82,
	70, 73, 71, 72, 88, 89, 90, 91, 95, 0,
	0, 0, 29, 30, 0, 0, 0, 0, 0, 0,
	0, 130, 113, 117, 0, 0, 134, 110, 114, 0,
	0, 138, 111, 115, 0, 148, 112, 116, 104, 105,
	28, 0, 0, 60, 0, 0, 0, 0, 0, 100,
	127, 128, 0, 101, 131, 132, 0, 102, 135, 136,
	103, 145, 146, 0, 53, 64, 65, 58, 0, 84,
	85, 86, 87, 93, 129, 92, 133, 94, 137, 147,
	-2, 63, 53, 0, 0, 27, 0, -2, 0, 62,
	59, 0, -2, 0, 61,
}

var yyTok1 = [...]int{
	1,
}

var yyTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63,
}

var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
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
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	}
	goto yystack /* stack new state and value */
}
