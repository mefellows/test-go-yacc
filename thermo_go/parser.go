//go:generate go tool yacc -o thermo.go -p "Thermo" thermo.y
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Program struct {
	statements []Statement
}

func NewProgram() *Program {
	return &Program{statements: make([]Statement, 0)}
}

func (p *Program) append(s Statement) {
	p.statements = append(p.statements, s)
}

type Token int

const (
	NULL Token = iota
	INTEGER
	FLOAT
	ASSIGNMENT
	IDENTIFIER

	//	WORD
	//	STATE
	HEATER
	TARGET
	TEMP
)

type Statement struct {
	Name  Token
	Value string
}

type Parser struct{}

func (p *Parser) Parse(input string) {
	fmt.Printf("Parsing input: %s", input)
}

type ThermoLex struct {
	s        string
	done     bool
	program  *Program
	location int
}

func (l *ThermoLex) Lex(lval *ThermoSymType) int {
	// Ragel has digested the entire input string.
	// Each time Lex is invoked, we just need to pop the last statement from the array
	// set the value and return the token
	if len(l.program.statements) > l.location && len(l.program.statements) > 0 {
		//	fmt.Printf("Nr statements:  %d", len(l.program.statements))
		//	fmt.Printf("Got statement: %v, at location: %d", l.program.statements[l.location], l.location)
		statement := l.program.statements[l.location]
		lval.str = statement.Value
		l.location = l.location + 1
		return int(statement.Name)
	}

	lval.str = ""
	l.location = 0
	return 0
}

func (l *ThermoLex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}

type Thermoy struct {
}

func (c *Thermoy) PrintResult(do string) {
	fmt.Printf("Doing something: %s", do)
}

func main() {
	fi := bufio.NewReader(os.NewFile(0, "stdin"))

	for {
		var eqn string
		var ok bool

		fmt.Printf("> ")
		if eqn, ok = readline(fi); ok {
			program = run_machine(eqn)
			ThermoParse(&ThermoLex{program: program})
		} else {
			break
		}
	}
}

func readline(fi *bufio.Reader) (string, bool) {
	s, err := fi.ReadString('\n')
	if err != nil {
		return "", false
	}
	return s, true
}
