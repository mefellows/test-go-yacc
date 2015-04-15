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
	s    string
	done bool
}

func (l *ThermoLex) Lex(lval *ThermoSymType) int {
	//lval.str = []byte("hello")
	if !l.done {
		lval.str = l.s
		l.done = true
		return int(WORD)
	}
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

		fmt.Printf("equation: ")
		if eqn, ok = readline(fi); ok {
			ThermoParse(&ThermoLex{s: eqn})
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
