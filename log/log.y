%{
#include <stdio.h>
#include <string.h>
 
void yyerror(const char *str)
{
        fprintf(stderr,"error: %s\n",str);
}
 
int yywrap()
{
        return 1;
} 
  
int main()
{
        yyparse();
} 

%}

%token SEMICOLON ZONETOK FILETOK WORD FILENAME QUOTE OBRACE EBRACE

%%
statement: 
		'h'
        {
                printf("A statement of '%s' was encountered\n", $1);
        }
		;
%%
