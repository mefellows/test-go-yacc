package main

import (
	"fmt"
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
	WORD
	STATE
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

func main() {
	program := run_machine("test = -100")
	fmt.Printf("Program: %v", program)
	program = run_machine("heat on")
	fmt.Printf("Program: %v", program)
	program = run_machine("heat on")
}
