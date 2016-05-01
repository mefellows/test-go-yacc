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

func main() {
	yyparse()
}

var heater = "default"
type stringy []byte
%}

%token TOKHEATER TOKHEAT TOKTARGET TOKTEMPERATURE

%union 
{
	number int
	str stringy
}

%token <number> STATE
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
