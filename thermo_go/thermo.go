//line thermo.y:2
package main

import __yyfmt__ "fmt"

//line thermo.y:2
import (
	"fmt"
)

func yyerror(str []byte) {
	fmt.Errorf("error: %s, passing through as is\n", str)
}

func yywrap() int {
	return 1
}

var heater = "default"

//line thermo.y:21
type ThermoSymType struct {
	yys    int
	number int
	str    string
}

const NUMBER = 57346
const TOKTEMPERATURE = 57347
const WORD = 57348
const TOKHEAT = 57349
const TOKHEATER = 57350
const TOKTARGET = 57351
const STATE = 57352
const SPECIAL = 57353

var ThermoToknames = []string{
	"NUMBER",
	"TOKTEMPERATURE",
	"WORD",
	"TOKHEAT",
	"TOKHEATER",
	"TOKTARGET",
	"STATE",
	"SPECIAL",
}
var ThermoStatenames = []string{}

const ThermoEofCode = 1
const ThermoErrCode = 2
const ThermoMaxDepth = 200

//line yacctab:1
var ThermoExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const ThermoNprod = 13
const ThermoPrivate = 57344

var ThermoTokenNames []string
var ThermoStates []string

const ThermoLast = 17

var ThermoAct = []int{

	9, 11, 10, 16, 12, 13, 15, 14, 17, 8,
	7, 6, 4, 3, 2, 1, 5,
}
var ThermoPact = []int{

	-1000, -7, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -5,
	2, 0, -3, -1000, 4, -1000, -1000, -1000,
}
var ThermoPgo = []int{

	0, 16, 15, 14, 13, 12, 11, 10, 9,
}
var ThermoR1 = []int{

	0, 2, 2, 3, 3, 4, 4, 5, 5, 1,
	6, 7, 8,
}
var ThermoR2 = []int{

	0, 0, 2, 1, 1, 1, 1, 1, 1, 2,
	3, 2, 2,
}
var ThermoChk = []int{

	-1000, -2, -3, -4, -5, -1, -6, -7, -8, 7,
	9, 8, 11, 10, 5, 6, 6, 4,
}
var ThermoDef = []int{

	1, -2, 2, 3, 4, 5, 6, 7, 8, 0,
	0, 0, 0, 9, 0, 11, 12, 10,
}
var ThermoTok1 = []int{

	1,
}
var ThermoTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
}
var ThermoTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var ThermoDebug = 0

type ThermoLexer interface {
	Lex(lval *ThermoSymType) int
	Error(s string)
}

const ThermoFlag = -1000

func ThermoTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(ThermoToknames) {
		if ThermoToknames[c-4] != "" {
			return ThermoToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func ThermoStatname(s int) string {
	if s >= 0 && s < len(ThermoStatenames) {
		if ThermoStatenames[s] != "" {
			return ThermoStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func Thermolex1(lex ThermoLexer, lval *ThermoSymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = ThermoTok1[0]
		goto out
	}
	if char < len(ThermoTok1) {
		c = ThermoTok1[char]
		goto out
	}
	if char >= ThermoPrivate {
		if char < ThermoPrivate+len(ThermoTok2) {
			c = ThermoTok2[char-ThermoPrivate]
			goto out
		}
	}
	for i := 0; i < len(ThermoTok3); i += 2 {
		c = ThermoTok3[i+0]
		if c == char {
			c = ThermoTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = ThermoTok2[1] /* unknown char */
	}
	if ThermoDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", ThermoTokname(c), uint(char))
	}
	return c
}

func ThermoParse(Thermolex ThermoLexer) int {
	var Thermon int
	var Thermolval ThermoSymType
	var ThermoVAL ThermoSymType
	ThermoS := make([]ThermoSymType, ThermoMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Thermostate := 0
	Thermochar := -1
	Thermop := -1
	goto Thermostack

ret0:
	return 0

ret1:
	return 1

Thermostack:
	/* put a state and value onto the stack */
	if ThermoDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", ThermoTokname(Thermochar), ThermoStatname(Thermostate))
	}

	Thermop++
	if Thermop >= len(ThermoS) {
		nyys := make([]ThermoSymType, len(ThermoS)*2)
		copy(nyys, ThermoS)
		ThermoS = nyys
	}
	ThermoS[Thermop] = ThermoVAL
	ThermoS[Thermop].yys = Thermostate

Thermonewstate:
	Thermon = ThermoPact[Thermostate]
	if Thermon <= ThermoFlag {
		goto Thermodefault /* simple state */
	}
	if Thermochar < 0 {
		Thermochar = Thermolex1(Thermolex, &Thermolval)
	}
	Thermon += Thermochar
	if Thermon < 0 || Thermon >= ThermoLast {
		goto Thermodefault
	}
	Thermon = ThermoAct[Thermon]
	if ThermoChk[Thermon] == Thermochar { /* valid shift */
		Thermochar = -1
		ThermoVAL = Thermolval
		Thermostate = Thermon
		if Errflag > 0 {
			Errflag--
		}
		goto Thermostack
	}

Thermodefault:
	/* default state action */
	Thermon = ThermoDef[Thermostate]
	if Thermon == -2 {
		if Thermochar < 0 {
			Thermochar = Thermolex1(Thermolex, &Thermolval)
		}

		/* look through exception table */
		xi := 0
		for {
			if ThermoExca[xi+0] == -1 && ThermoExca[xi+1] == Thermostate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Thermon = ThermoExca[xi+0]
			if Thermon < 0 || Thermon == Thermochar {
				break
			}
		}
		Thermon = ThermoExca[xi+1]
		if Thermon < 0 {
			goto ret0
		}
	}
	if Thermon == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Thermolex.Error("syntax error")
			Nerrs++
			if ThermoDebug >= 1 {
				__yyfmt__.Printf("%s", ThermoStatname(Thermostate))
				__yyfmt__.Printf(" saw %s\n", ThermoTokname(Thermochar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Thermop >= 0 {
				Thermon = ThermoPact[ThermoS[Thermop].yys] + ThermoErrCode
				if Thermon >= 0 && Thermon < ThermoLast {
					Thermostate = ThermoAct[Thermon] /* simulate a shift of "error" */
					if ThermoChk[Thermostate] == ThermoErrCode {
						goto Thermostack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if ThermoDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", ThermoS[Thermop].yys)
				}
				Thermop--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if ThermoDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", ThermoTokname(Thermochar))
			}
			if Thermochar == ThermoEofCode {
				goto ret1
			}
			Thermochar = -1
			goto Thermonewstate /* try again in the same state */
		}
	}

	/* reduction by production Thermon */
	if ThermoDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Thermon, ThermoStatname(Thermostate))
	}

	Thermont := Thermon
	Thermopt := Thermop
	_ = Thermopt // guard against "declared and not used"

	Thermop -= ThermoR2[Thermon]
	ThermoVAL = ThermoS[Thermop+1]

	/* consult goto table to find next state */
	Thermon = ThermoR1[Thermon]
	Thermog := ThermoPgo[Thermon]
	Thermoj := Thermog + ThermoS[Thermop].yys + 1

	if Thermoj >= ThermoLast {
		Thermostate = ThermoAct[Thermog]
	} else {
		Thermostate = ThermoAct[Thermoj]
		if ThermoChk[Thermostate] != -Thermon {
			Thermostate = ThermoAct[Thermog]
		}
	}
	// dummy call; replaced with literal code
	switch Thermont {

	case 4:
		//line thermo.y:41
		{
			fmt.Printf("Running a command\n")
		}
	case 9:
		//line thermo.y:55
		{
			if ThermoS[Thermopt-0].str == "on" {
				fmt.Printf("\tHeater '%s' turned on\n", heater)
				ThermoVAL.str = ThermoS[Thermopt-0].str
			} else {
				fmt.Printf("\tHeat '%s' turned off\n", heater)
			}
		}
	case 10:
		//line thermo.y:67
		{
			fmt.Printf("\tHeater '%s' temperature set to %d\n", heater, ThermoS[Thermopt-0].number)
		}
	case 11:
		//line thermo.y:74
		{
			fmt.Printf("\tSelected heater '%s'\n", ThermoS[Thermopt-0].str)
			heater = ThermoS[Thermopt-0].str
		}
	case 12:
		//line thermo.y:82
		{
			fmt.Printf("\tspecial word: '%s'\n", ThermoS[Thermopt-1].str)
		}
	}
	goto Thermostack /* stack new state and value */
}
