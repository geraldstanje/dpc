//line parser.y:2
package main

import __yyfmt__ "fmt"

//line parser.y:3
//line parser.y:7
type yySymType struct {
	yys int
	f   float64
	i   int
	s   string
}

const tFNUMBER = 57346
const tNUMBER = 57347
const tID = 57348
const tQSTRING = 57349
const tAND = 57350
const tARRAY = 57351
const tASSIGN = 57352
const tBEGIN = 57353
const tBOOLEAN = 57354
const tBREAK = 57355
const tCHAR = 57356
const tCONTINUE = 57357
const tDIV = 57358
const tDO = 57359
const tDOTDOT = 57360
const tDOWNTO = 57361
const tELSE = 57362
const tEND = 57363
const tFALSE = 57364
const tFOR = 57365
const tFUNCTION = 57366
const tGE = 57367
const tGOTO = 57368
const tIF = 57369
const tINTEGER = 57370
const tLE = 57371
const tMOD = 57372
const tNE = 57373
const tNOT = 57374
const tOF = 57375
const tOR = 57376
const tPROCEDURE = 57377
const tPROGRAM = 57378
const tREAL = 57379
const tRECORD = 57380
const tREPEAT = 57381
const tSTRING = 57382
const tTHEN = 57383
const tTO = 57384
const tTRUE = 57385
const tTYPE = 57386
const tUNTIL = 57387
const tVAR = 57388
const tWHILE = 57389
const UMINUS = 57390

var yyToknames = []string{
	"tFNUMBER",
	"tNUMBER",
	"tID",
	"tQSTRING",
	"tAND",
	"tARRAY",
	"tASSIGN",
	"tBEGIN",
	"tBOOLEAN",
	"tBREAK",
	"tCHAR",
	"tCONTINUE",
	"tDIV",
	"tDO",
	"tDOTDOT",
	"tDOWNTO",
	"tELSE",
	"tEND",
	"tFALSE",
	"tFOR",
	"tFUNCTION",
	"tGE",
	"tGOTO",
	"tIF",
	"tINTEGER",
	"tLE",
	"tMOD",
	"tNE",
	"tNOT",
	"tOF",
	"tOR",
	"tPROCEDURE",
	"tPROGRAM",
	"tREAL",
	"tRECORD",
	"tREPEAT",
	"tSTRING",
	"tTHEN",
	"tTO",
	"tTRUE",
	"tTYPE",
	"tUNTIL",
	"tVAR",
	"tWHILE",
	"'<'",
	"'>'",
	"'='",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"UMINUS",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 25,
	20, 39,
	56, 39,
	-2, 74,
	-1, 131,
	25, 0,
	29, 0,
	31, 0,
	48, 0,
	49, 0,
	50, 0,
	-2, 55,
	-1, 132,
	25, 0,
	29, 0,
	31, 0,
	48, 0,
	49, 0,
	50, 0,
	-2, 56,
	-1, 133,
	25, 0,
	29, 0,
	31, 0,
	48, 0,
	49, 0,
	50, 0,
	-2, 57,
	-1, 138,
	25, 0,
	29, 0,
	31, 0,
	48, 0,
	49, 0,
	50, 0,
	-2, 62,
	-1, 139,
	25, 0,
	29, 0,
	31, 0,
	48, 0,
	49, 0,
	50, 0,
	-2, 63,
	-1, 141,
	25, 0,
	29, 0,
	31, 0,
	48, 0,
	49, 0,
	50, 0,
	-2, 65,
}

const yyNprod = 78
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 444

var yyAct = []int{

	85, 67, 66, 114, 116, 84, 62, 127, 127, 113,
	150, 63, 15, 158, 126, 50, 21, 40, 149, 42,
	101, 44, 41, 119, 43, 172, 68, 22, 45, 69,
	59, 155, 77, 112, 76, 18, 22, 38, 160, 79,
	64, 81, 82, 154, 72, 124, 22, 118, 74, 104,
	105, 117, 106, 107, 111, 73, 80, 75, 71, 2,
	78, 34, 162, 102, 42, 39, 4, 41, 36, 43,
	34, 152, 16, 120, 34, 35, 22, 38, 174, 122,
	123, 121, 70, 11, 16, 37, 164, 165, 129, 130,
	131, 132, 133, 134, 135, 136, 137, 138, 139, 140,
	141, 142, 22, 128, 8, 87, 7, 143, 26, 145,
	157, 147, 115, 88, 148, 9, 8, 11, 7, 77,
	151, 76, 20, 5, 115, 22, 146, 98, 156, 108,
	13, 54, 53, 46, 57, 74, 31, 16, 103, 83,
	61, 14, 65, 58, 75, 33, 32, 78, 17, 56,
	94, 93, 3, 60, 161, 163, 170, 166, 161, 52,
	168, 153, 19, 87, 12, 169, 10, 6, 1, 0,
	55, 88, 0, 22, 167, 0, 175, 0, 48, 49,
	96, 0, 0, 0, 97, 98, 99, 22, 173, 100,
	51, 0, 47, 0, 87, 0, 0, 0, 0, 0,
	0, 0, 88, 89, 91, 90, 95, 92, 94, 93,
	0, 96, 0, 0, 0, 97, 98, 99, 0, 144,
	100, 0, 87, 0, 0, 0, 0, 0, 0, 0,
	88, 0, 0, 87, 89, 91, 90, 95, 92, 94,
	93, 88, 171, 0, 98, 0, 0, 125, 100, 0,
	96, 0, 0, 0, 97, 98, 99, 0, 0, 100,
	0, 0, 0, 87, 0, 95, 92, 94, 93, 0,
	0, 88, 0, 89, 91, 90, 95, 92, 94, 93,
	96, 0, 0, 0, 97, 98, 99, 0, 0, 100,
	0, 0, 0, 87, 0, 0, 0, 159, 0, 0,
	0, 88, 109, 89, 91, 90, 95, 92, 94, 93,
	96, 0, 0, 0, 97, 98, 99, 0, 0, 100,
	0, 0, 0, 87, 0, 0, 0, 0, 0, 0,
	0, 88, 0, 89, 91, 90, 95, 92, 94, 93,
	96, 0, 0, 0, 97, 98, 99, 0, 0, 100,
	0, 0, 0, 87, 0, 0, 86, 0, 0, 0,
	0, 88, 0, 89, 91, 90, 95, 92, 94, 93,
	96, 0, 0, 0, 97, 98, 99, 0, 0, 100,
	0, 0, 0, 25, 0, 0, 0, 0, 11, 0,
	23, 0, 24, 89, 91, 90, 95, 92, 94, 93,
	28, 0, 25, 0, 27, 0, 0, 11, 0, 23,
	0, 24, 0, 0, 0, 0, 30, 0, 0, 28,
	0, 0, 110, 27, 29, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 30, 0, 0, 0, 0,
	0, 0, 0, 29,
}
var yyPact = []int{

	23, -1000, 146, 10, -1000, 60, 106, 131, 142, -22,
	-1000, 396, -1000, 140, 139, 16, -1000, 18, -1000, 64,
	396, 9, 7, -1000, -1000, -42, -1000, 127, 137, 127,
	396, 72, -52, -52, 136, 20, 20, -1000, 0, -1000,
	127, 127, 133, -1000, 127, 315, -43, 132, 127, 127,
	-38, 127, 127, -1000, -1000, -1000, -1000, -1000, 119, 285,
	377, -2, -26, 78, -5, -1000, -9, -1000, -1000, -37,
	20, 78, -52, -52, -1000, -1000, -1000, -1000, -1000, -11,
	-1000, 345, 186, -1000, -50, 345, 396, 127, 127, 127,
	127, 127, 127, 127, 127, 127, 127, 127, 127, 127,
	127, 127, -38, -1000, -1000, -1000, 155, -1000, 127, 396,
	127, -1000, 107, -46, -1000, 131, 12, -1000, -1000, 156,
	-1000, -13, -28, -1000, -1000, -1000, -1000, 127, 90, -1000,
	-1000, 214, 214, 214, 97, -1000, -1000, 97, 214, 214,
	-1000, 214, 97, -51, -1000, 255, -1000, 345, -18, -1000,
	78, 3, 20, 68, 66, 107, 345, 396, -1000, 127,
	-1000, -1000, 20, -1000, 151, -1000, -1000, -1000, 225, -1000,
	-36, 396, 45, -1000, 107, -1000,
}
var yyPgo = []int{

	0, 168, 123, 167, 108, 4, 2, 1, 9, 6,
	166, 164, 3, 162, 122, 16, 15, 0, 5,
}
var yyR1 = []int{

	0, 1, 5, 5, 2, 2, 2, 6, 6, 6,
	6, 6, 6, 6, 7, 7, 7, 7, 7, 3,
	3, 10, 11, 11, 9, 9, 8, 8, 12, 12,
	4, 13, 13, 14, 14, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 18, 18, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 16, 16, 16, 16,
}
var yyR2 = []int{

	0, 7, 3, 1, 6, 6, 0, 1, 1, 8,
	2, 4, 4, 2, 1, 1, 1, 1, 1, 2,
	0, 4, 6, 4, 3, 0, 3, 1, 4, 3,
	3, 1, 0, 3, 2, 3, 1, 1, 4, 1,
	1, 4, 6, 8, 4, 4, 3, 1, 4, 2,
	2, 2, 1, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 1,
	1, 1, 1, 1, 1, 4, 3, 2,
}
var yyChk = []int{

	-1000, -1, 36, 6, 56, -2, -3, 46, 44, -4,
	-10, 11, -11, 24, 35, -5, 6, 6, 57, -13,
	-14, -15, -16, 13, 15, 6, -4, 27, 23, 47,
	39, -2, 6, 6, 58, 59, 50, 21, -15, 56,
	10, 60, 57, 62, 63, -17, 6, 65, 51, 52,
	-16, 63, 32, 5, 4, 43, 22, 7, 6, -17,
	-14, -4, -9, 63, -9, 6, -6, -7, 6, 9,
	62, 38, 24, 35, 28, 37, 14, 12, 40, -6,
	56, -17, -17, 6, -18, -17, 41, 8, 16, 48,
	50, 49, 52, 54, 53, 51, 25, 29, 30, 31,
	34, 63, -16, 6, -17, -17, -17, -17, 10, 17,
	45, 56, 59, -8, -12, 46, -5, 56, 56, 60,
	-6, -8, -9, -9, 56, 61, 64, 58, -15, -17,
	-17, -17, -17, -17, -17, -17, -17, -17, -17, -17,
	-17, -17, -17, -18, 64, -17, -15, -17, -7, 64,
	56, -5, 59, 5, 56, 59, -17, 20, 64, 42,
	56, -12, 59, -6, 18, 21, -7, -15, -17, -6,
	5, 17, 61, -15, 33, -7,
}
var yyDef = []int{

	0, -2, 0, 0, 6, 20, 0, 0, 0, 0,
	19, 32, 6, 0, 0, 0, 3, 0, 1, 0,
	31, 0, 0, 36, 37, -2, 40, 0, 0, 0,
	0, 0, 25, 25, 0, 0, 0, 30, 0, 34,
	0, 0, 0, 77, 0, 0, 74, 0, 0, 0,
	52, 0, 0, 69, 70, 71, 72, 73, 0, 0,
	0, 0, 0, 0, 0, 2, 0, 7, 8, 0,
	0, 0, 25, 25, 14, 15, 16, 17, 18, 0,
	33, 35, 0, 76, 0, 47, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 49, 74, 50, 51, 0, 68, 0, 0,
	0, 21, 0, 0, 27, 0, 0, 23, 4, 0,
	10, 0, 0, 13, 5, 75, 38, 0, 41, 53,
	54, -2, -2, -2, 58, 59, 60, 61, -2, -2,
	64, -2, 66, 0, 67, 0, 44, 45, 0, 24,
	0, 0, 0, 0, 0, 0, 46, 0, 48, 0,
	22, 26, 0, 29, 0, 11, 12, 42, 0, 28,
	0, 0, 0, 43, 0, 9,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	63, 64, 53, 51, 58, 52, 57, 54, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 59, 56,
	48, 50, 49, 3, 65, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 60, 3, 61, 62,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 55,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
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

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
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
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
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
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
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
			if yyn < 0 || yyn == yychar {
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
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
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
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
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
