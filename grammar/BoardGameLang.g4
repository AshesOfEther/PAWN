grammar BoardGameLang;

type
    : type '[]'

    | TYPE
    ;

// Parser rules
expression
    : equalities
    ;

member
    : member '.' NAME
    | NAME
    ;

equalities
    : range ('=='|'!='|'<='|'>='|'<'|'>') equalities
    | range
    ;

range
    : sum ('..='|'..') sum
    | sum
    ;

atom
    : '(' expression ')'
    | atom argList ('{' statementList '}')? // call
    | atom '[' expressionList ']' // indexing
    | member
    | literal
    ;

literal
    : list
    | BOOL
    | INT
    | FLOAT
    | STRING
    ;

power
    : atom '^' power
    | atom
    ;

sum
    : product ('+'|'-') sum
    | product
    ;

product
    : power ('*'|'/'|'%') product
    | power
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
    | declaration ';'
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
    : 'fn' type? NAME functionDeclArgs '{' statementList '}'
    ;

functionDeclArgs
    : '(' ')'
    | '(' positionalArgsDeclList (',' namedArgsDeclList)? ','? ')'
    | '(' namedArgsDeclList ','? ')'
    ;

positionalArgsDeclList
    : positionalArgDecl (',' positionalArgDecl)*
    ;

positionalArgDecl
    : type NAME
    ;

namedArgsDeclList
    : namedArgDecl (',' namedArgDecl)*
    ;

namedArgDecl
    : type NAME '=' expression
    ;

return
    : 'return' expression
    ;

declaration
    : type NAME '=' expression
    ;

assignment
    : NAME '=' expression
    ;


// Lexer rules
LINE_COMMENT : '#' .*? '\r'? '\n' -> skip;
MULTILINE_COMMENT : '/#' .*? '#/' -> skip;
TYPE : 'Int' | 'Float' | 'String' | 'Boolean';
BOOL : 'true' | 'false';
NAME : [a-zA-Z_][a-zA-Z0-9_]*;
FLOAT : '-'?[0-9]+'.'[0-9]+;
INT : '-'?[0-9]+;
STRING : '"' ('\\'[\\"] | .)*? '"';
WS : [ \t\r\n]+ -> skip;
