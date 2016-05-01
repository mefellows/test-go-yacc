#!/bin/bash
lex log.l
yacc -d log.y --debug --verbose
cc lex.yy.c y.tab.c -o log
