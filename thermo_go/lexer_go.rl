%%{

  machine simple_lexer;

  action print { /*fmt.Printf("Current mark: %d, fpc: %d", mark, fpc) */ }
  
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
      emit(TOKHEAT, data, program, ts, te) 
    };
    
    heater => { 
      emit(TOKHEATER, data, program, ts, te) 
    };
    
    temperature => { 
      emit(TOKTEMPERATURE, data, program, ts, te) 
    };

    integer => { 
      emit(NUMBER, data, program, ts, te) 
    };
    
    word => { 
      emit(WORD, data, program, ts, te) 
    };
    
    space;
    
  *|;

}%%

package main
//import "fmt"

%% write data;
var cs, p, pe, ts, te, act, eof, mark int
var program *Program

func emit(token_name int, data []byte, program *Program, ts int, te int) {
  value := string(data[ts:te])
  program.append(Statement{Name: token_name, Value: value})
}

func reset() {
  cs, p, pe, ts, te, act, eof, mark = 0,0,0,0,0,0,0,0
}

func run_machine(input string) *Program {
  program = NewProgram()
  reset()
  data := []byte(input)
  pe = len(data)
  //eof = -1
  eof = len(data)
  %% write init;
  %% write exec;
  
  return program
}
