grammar BoardGameLang;

// Parser rules
expression
    : '(' expression ')'
    | expression argList ('{' statementList '}')?
    | expression '[' expressionList ']'
    | expression '^' expression
    | expression ('*' | '/' | '%') expression
    | expression ('+' | '-') expression
    | expression ('..=' | '..') expression
    | expression ('=='|'!='|'<='|'>='|'<'|'>') expression
    | member
    | literal
    ;

member
    : member '.' NAME
    | NAME
    ;

literal
    : list
    | BOOL
    | NUMBER
    | STRING
    ;

list
    : '[' (expression (',' expression)* ','?)? ']'
    ;

argList
    : '(' ')'
    | '(' expressionList (',' namedArgs)? ','? ')'
    | '(' namedArgs ','? ')'
    ;

expressionList
    : expression (',' expression)*
    ;

namedArgs
    : namedArg (',' namedArg)*
    ;


namedArg
    : NAME '=' expression
    ;

statementList
    : statement*
    ;

statement
    : apply
    | while
    | if
    | functionDecl
    | return ';'
    | expression ';'
    | assignment ';'
    ;

apply
    : 'apply' expression '{' statementList '}'
    ;

while
    : 'while' expression '{' statementList '}'
    ;

if
    : 'if' expression '{' statementList '}'
        ('else' 'if' expression '{' statementList '}' )*
        ('else' '{' statementList '}')?
    ;

functionDecl
    : 'fn' NAME functionDeclArgs '{' statementList '}'
    ;

functionDeclArgs
    : '(' ')'
    | '(' nameList (',' namedArgsDeclList)? ','? ')'
    | '(' namedArgsDeclList ','? ')'
    ;

nameList
    : NAME (',' NAME)*
    ;

namedArgsDeclList
    : namedArgDecl (',' namedArgDecl)*
    ;

namedArgDecl
    : NAME '=' expression
    ;

return
    : 'return' expression
    ;

assignment
    : NAME '=' expression
    ;


// Lexer rules
LINE_COMMENT : '#' .*? '\r'? '\n' -> skip;
MULTILINE_COMMENT : '/#' .*? '#/' -> skip;
BOOL : 'true' | 'false';
NAME : [a-zA-Z_][a-zA-Z0-9_]*;
NUMBER : '-'?[0-9]+('.'[0-9]+)?; // Float or int
STRING : '"' ('\\'[\\"] | .)*? '"';
WS : [ \t\r\n]+ -> skip;