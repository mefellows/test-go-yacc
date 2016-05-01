%{
package main
import (
	"fmt"
)

func yyerror(str []byte) {
	fmt.Errorf("error: %s, passing through as is\n",str)
}

func yywrap() int {
	return 1
}

var heater = "default"
%}

//%token TOKHEATER TOKHEAT TOKTARGET TOKTEMPERATURE

// Will be ${PREFIX}SymType
%union 
{
	number int
	str string
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%type <str> heat_switch 
%token <number> NUMBER TOKTEMPERATURE 
%token <str> WORD TOKHEAT TOKHEATER TOKTARGET STATE SPECIAL

%%

commands:
	| commands command
	;

command:
	other_special_command | special_command	
    {
		fmt.Printf("Running a command\n")
    }

other_special_command:
	heat_switch | target_set
    ;

special_command:
	heater_select | specific_command
    ;

heat_switch:
	TOKHEAT STATE 
	{
		if $2 == "on" {
			fmt.Printf("\tHeater '%s' turned on\n", heater);
			$$ = $2;
		} else {
			fmt.Printf("\tHeat '%s' turned off\n", heater)
        }
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

specific_command:
	SPECIAL WORD
	{
		fmt.Printf("\tspecial word: '%s'\n",$1)
	}
	;
