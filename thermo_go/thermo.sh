lex thermo.l
yacc -d thermo.y
cc lex.yy.c y.tab.c -o thermo 
