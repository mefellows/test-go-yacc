
//line lexer_go.rl:1

//line lexer_go.rl:56


package main
import "fmt"


//line lexer_go.go:12
var _simple_lexer_actions []byte = []byte{
	0, 1, 1, 1, 2, 1, 3, 1, 9, 
	1, 10, 1, 11, 1, 12, 1, 13, 
	1, 14, 1, 15, 1, 16, 2, 3, 
	0, 2, 3, 4, 2, 3, 5, 2, 
	3, 6, 2, 3, 7, 2, 3, 8, 
}

var _simple_lexer_key_offsets []byte = []byte{
	0, 0, 2, 4, 9, 15, 22, 29, 
	44, 47, 49, 54, 60, 66, 72, 78, 
	84, 90, 96, 102, 108, 114, 120, 126, 
	132, 138, 144, 150, 156, 
}

var _simple_lexer_trans_keys []byte = []byte{
	48, 57, 48, 57, 95, 65, 90, 97, 
	122, 95, 101, 65, 90, 97, 122, 95, 
	102, 110, 65, 90, 97, 122, 95, 97, 
	101, 65, 90, 98, 122, 32, 43, 45, 
	61, 104, 111, 116, 9, 13, 48, 57, 
	65, 90, 97, 122, 46, 48, 57, 48, 
	57, 95, 65, 90, 97, 122, 95, 97, 
	65, 90, 98, 122, 95, 116, 65, 90, 
	97, 122, 95, 101, 65, 90, 97, 122, 
	95, 114, 65, 90, 97, 122, 95, 102, 
	65, 90, 97, 122, 95, 114, 65, 90, 
	97, 122, 95, 103, 65, 90, 97, 122, 
	95, 101, 65, 90, 97, 122, 95, 116, 
	65, 90, 97, 122, 95, 109, 65, 90, 
	97, 122, 95, 112, 65, 90, 97, 122, 
	95, 101, 65, 90, 97, 122, 95, 114, 
	65, 90, 97, 122, 95, 97, 65, 90, 
	98, 122, 95, 116, 65, 90, 97, 122, 
	95, 117, 65, 90, 97, 122, 95, 114, 
	65, 90, 97, 122, 95, 101, 65, 90, 
	97, 122, 
}

var _simple_lexer_single_lengths []byte = []byte{
	0, 0, 0, 1, 2, 3, 3, 7, 
	1, 0, 1, 2, 2, 2, 2, 2, 
	2, 2, 2, 2, 2, 2, 2, 2, 
	2, 2, 2, 2, 2, 
}

var _simple_lexer_range_lengths []byte = []byte{
	0, 1, 1, 2, 2, 2, 2, 4, 
	1, 1, 2, 2, 2, 2, 2, 2, 
	2, 2, 2, 2, 2, 2, 2, 2, 
	2, 2, 2, 2, 2, 
}

var _simple_lexer_index_offsets []byte = []byte{
	0, 0, 2, 4, 8, 13, 19, 25, 
	37, 40, 42, 46, 51, 56, 61, 66, 
	71, 76, 81, 86, 91, 96, 101, 106, 
	111, 116, 121, 126, 131, 
}

var _simple_lexer_indicies []byte = []byte{
	0, 1, 3, 2, 4, 4, 4, 1, 
	4, 5, 4, 4, 1, 4, 6, 7, 
	4, 4, 1, 4, 8, 9, 4, 4, 
	1, 10, 11, 11, 12, 14, 15, 16, 
	10, 0, 13, 13, 1, 18, 19, 17, 
	3, 20, 4, 4, 4, 21, 4, 23, 
	4, 4, 22, 4, 24, 4, 4, 22, 
	4, 26, 4, 4, 25, 4, 27, 4, 
	4, 22, 4, 7, 4, 4, 22, 4, 
	28, 4, 4, 22, 4, 29, 4, 4, 
	22, 4, 30, 4, 4, 22, 4, 31, 
	4, 4, 22, 4, 32, 4, 4, 22, 
	4, 33, 4, 4, 22, 4, 34, 4, 
	4, 22, 4, 35, 4, 4, 22, 4, 
	36, 4, 4, 22, 4, 37, 4, 4, 
	22, 4, 38, 4, 4, 22, 4, 39, 
	4, 4, 22, 4, 40, 4, 4, 22, 
	
}

var _simple_lexer_trans_targs []byte = []byte{
	8, 0, 7, 9, 10, 11, 15, 10, 
	16, 20, 7, 1, 7, 3, 4, 5, 
	6, 7, 2, 8, 7, 7, 7, 12, 
	13, 7, 14, 10, 17, 18, 19, 10, 
	21, 22, 23, 24, 25, 26, 27, 28, 
	10, 
}

var _simple_lexer_trans_actions []byte = []byte{
	23, 0, 19, 0, 38, 0, 0, 29, 
	0, 0, 9, 0, 7, 0, 0, 0, 
	0, 13, 0, 5, 15, 21, 17, 0, 
	0, 11, 0, 32, 0, 0, 0, 26, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	35, 
}

var _simple_lexer_to_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 1, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 
}

var _simple_lexer_from_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 3, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 
}

var _simple_lexer_eof_trans []byte = []byte{
	0, 0, 3, 0, 0, 0, 0, 0, 
	18, 21, 22, 23, 23, 26, 23, 23, 
	23, 23, 23, 23, 23, 23, 23, 23, 
	23, 23, 23, 23, 23, 
}

const simple_lexer_start int = 7
const simple_lexer_first_final int = 7
const simple_lexer_error int = 0

const simple_lexer_en_gosh int = 7


//line lexer_go.rl:62
var cs, p, pe, ts, te, act, eof, mark int
var program *Program

func emit(token_name Token, data []byte, program *Program, ts int, te int) {
  value := string(data[ts:te])
  fmt.Printf("Inserting statement %v,%v onto array", token_name, value)
  program.append(Statement{Name: token_name, Value: value})
  fmt.Printf("Program %v", program)
}

func run_machine(input string) *Program {
  program = NewProgram()
  data := []byte(input)
  fmt.Printf("Running the state machine with input %v...\n", data)
  pe = len(data)
  //eof = -1
  eof = len(data)
  
//line lexer_go.go:159
	{
	cs = simple_lexer_start
	ts = 0
	te = 0
	act = 0
	}

//line lexer_go.rl:80
  
//line lexer_go.go:169
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	_acts = int(_simple_lexer_from_state_actions[cs])
	_nacts = uint(_simple_lexer_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		 _acts++
		switch _simple_lexer_actions[_acts - 1] {
		case 2:
//line NONE:1
ts = p

//line lexer_go.go:192
		}
	}

	_keys = int(_simple_lexer_key_offsets[cs])
	_trans = int(_simple_lexer_index_offsets[cs])

	_klen = int(_simple_lexer_single_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case data[p] < _simple_lexer_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _simple_lexer_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_simple_lexer_range_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case data[p] < _simple_lexer_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _simple_lexer_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	_trans = int(_simple_lexer_indicies[_trans])
_eof_trans:
	cs = int(_simple_lexer_trans_targs[_trans])

	if _simple_lexer_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_simple_lexer_trans_actions[_trans])
	_nacts = uint(_simple_lexer_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _simple_lexer_actions[_acts-1] {
		case 0:
//line lexer_go.rl:5
 fmt.Printf("Current mark: %d, fpc: %d", mark, p) 
		case 3:
//line NONE:1
te = p+1

		case 4:
//line lexer_go.rl:20
act = 1;
		case 5:
//line lexer_go.rl:24
act = 2;
		case 6:
//line lexer_go.rl:28
act = 3;
		case 7:
//line lexer_go.rl:32
act = 4;
		case 8:
//line lexer_go.rl:48
act = 8;
		case 9:
//line lexer_go.rl:44
te = p+1
{ 
      emit(ASSIGNMENT, data, program, ts, te) 
    }
		case 10:
//line lexer_go.rl:52
te = p+1

		case 11:
//line lexer_go.rl:28
te = p
p--
{ 
      emit(HEATER, data, program, ts, te) 
    }
		case 12:
//line lexer_go.rl:36
te = p
p--
{ 
      emit(INTEGER, data, program, ts, te) 
    }
		case 13:
//line lexer_go.rl:40
te = p
p--
{ 
      emit(FLOAT, data, program, ts, te) 
    }
		case 14:
//line lexer_go.rl:48
te = p
p--
{ 
      emit(WORD, data, program, ts, te) 
    }
		case 15:
//line lexer_go.rl:36
p = (te) - 1
{ 
      emit(INTEGER, data, program, ts, te) 
    }
		case 16:
//line NONE:1
	switch act {
	case 1:
	{p = (te) - 1
 
      emit(TARGET, data, program, ts, te) 
    }
	case 2:
	{p = (te) - 1
 
      emit(STATE, data, program, ts, te) 
    }
	case 3:
	{p = (te) - 1
 
      emit(HEATER, data, program, ts, te) 
    }
	case 4:
	{p = (te) - 1
 
      emit(TEMP, data, program, ts, te) 
    }
	case 8:
	{p = (te) - 1
 
      emit(WORD, data, program, ts, te) 
    }
	}
	
//line lexer_go.go:358
		}
	}

_again:
	_acts = int(_simple_lexer_to_state_actions[cs])
	_nacts = uint(_simple_lexer_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _simple_lexer_actions[_acts-1] {
		case 1:
//line NONE:1
ts = 0

//line lexer_go.go:372
		}
	}

	if cs == 0 {
		goto _out
	}
	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	if p == eof {
		if _simple_lexer_eof_trans[cs] > 0 {
			_trans = int(_simple_lexer_eof_trans[cs] - 1)
			goto _eof_trans
		}
	}

	_out: {}
	}

//line lexer_go.rl:81
  
  fmt.Printf("Finished. The state of the machine is: %s\n", cs)
  fmt.Printf("p: %s pe: %s\n", p, pe)
  return program
}
