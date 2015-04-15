%{
package main
import (
	"fmt"
)

func yyerror(str []byte) {
	fmt.Errorf("error: %s\n",str)
}

func yywrap() int {
	return 1
}

var heater = "default"
%}

%token TOKHEATER TOKHEAT TOKTARGET TOKTEMPERATURE

// Will be ${PREFIX}SymType
%union 
{
	number int
	str string
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%type <str> heat_switch 

%token <str> STATE
%token <number> NUMBER
%token <str> WORD

%%

commands:
	| commands command
	;


command:
	heat_switch | target_set | heater_select

heat_switch:
	TOKHEAT STATE 
	{
		//if $2 {
			fmt.Printf("\tHeater '%s' turned on\n", heater)
			$$ = $2
		//} else
		//	fmt.Printf("\tHeat '%s' turned off\n", heater)
        //}
	}
	;

target_set:
	TOKTARGET TOKTEMPERATURE NUMBER
	{
		fmt.Printf("\tHeater '%s' temperature set to %d\n",heater, $3)
	}
	;

heater_select:
	TOKHEATER WORD
	{
		fmt.Printf("\tSelected heater '%s'\n",$2)
		heater = $2
	}
	;

heater_select:
	WORD
	{
		fmt.Printf("\tword: '%s'\n",$1)
	}
	;
