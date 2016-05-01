%{
#include <stdio.h>
#include "y.tab.h"
%}

%%

zone                    printf("zone token"); return ZONETOK;
file                    return FILETOK;
[a-zA-Z][a-zA-Z0-9]*    yylval.str=strdup(yytext); printf("word: %s", yytext); return WORD;
[a-zA-Z0-9\/.-]+        yylval.str=strdup(yytext); printf("filename: %s", yytext); return FILENAME;
\"                      return QUOTE;
\{                      return OBRACE;
\}                      return EBRACE;
;                       return SEMICOLON;
\n                      /* ignore EOL */;
[ \t]+                  /* ignore whitespace */;
%%

