%%{

  machine simple_lexer;

  action print { fmt.Printf("Current mark: %d, fpc: %d", mark, fpc) }
  
  integer     = ('+'|'-')?[0-9]+ >print;
  float       = ('+'|'-')?[0-9]+'.'[0-9]+;
  assignment  = '=';
  heater	  = ("heater");
  heat	      = ("heat");
  state       = ("on"|"off");     
  target      = ("target"|"set");
  temperature = "temperature";
  newline     = '\n';
  whitespace  = [\t]+;
  word        = [a-zA-Z][a-zA-Z_]+; 
  
  gosh := |*
    
    target => { 
      emit(TOKTARGET, data, program, ts, te) 
    };
    
    state=> { 
      emit(STATE, data, program, ts, te) 
    };
    
    heat => { 
	  fmt.Printf("Heat!")
      emit(TOKHEAT, data, program, ts, te) 
    };
    
    heater => { 
	  fmt.Printf("Heater!")
      emit(TOKHEATER, data, program, ts, te) 
    };
    
    temperature => { 
      emit(TOKTEMPERATURE, data, program, ts, te) 
    };

    integer => { 
      emit(NUMBER, data, program, ts, te) 
    };
    
    word => { 
	  fmt.Printf("word")
      emit(WORD, data, program, ts, te) 
    };
    
    space;
    
  *|;

}%%

package main
import "fmt"

%% write data;
var cs, p, pe, ts, te, act, eof, mark int
var program *Program

func emit(token_name Token, data []byte, program *Program, ts int, te int) {
  value := string(data[ts:te])
  //fmt.Printf("Inserting statement %v,%v onto array", token_name, value)
  program.append(Statement{Name: token_name, Value: value})
  //fmt.Printf("Program %v", program)
}

func reset() {
  cs, p, pe, ts, te, act, eof, mark = 0,0,0,0,0,0,0,0
}

func run_machine(input string) *Program {
  program = NewProgram()
  reset()
  data := []byte(input)
  fmt.Printf("Running the state machine with input %v...\n", data)
  pe = len(data)
  //eof = -1
  eof = len(data)
  %% write init;
  %% write exec;
  
  fmt.Printf("Finished. The state of the machine is: %s\n", cs)
  fmt.Printf("p: %s pe: %s\n", p, pe)
  return program
}
