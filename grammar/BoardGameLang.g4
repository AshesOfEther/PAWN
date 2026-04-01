grammar BoardGameLang;

// Parser rules
expression
    : '(' expression ')'                                        # Parens
    | expression argList ('{' statementList '}')?               # Call
    | expression '[' expressionList ']'                         # Index
    | expression '^' expression                                 # Power
    | expression op=('*' | '/' | '%') expression                # MulDivMod
    | expression op=('+' | '-') expression                      # AddSub
    | expression op=('..=' | '..') expression                   # Range
    | expression op=('=='|'!='|'<='|'>='|'<'|'>') expression    # Comparison
    | member                                                    # MemberExpr
    | literal                                                   # LiteralExpr
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
MUL : '*';
DIV : '/';
MOD : '%';
ADD : '+';
SUB : '-';
RANGE : '..';
RANGE_INCLUSIVE : '..=';
EQUAL : '==';
NOT_EQUAL : '!=';
LESS_THAN : '<';
GREATER_THAN : '>';
LESS_EUQAL : '<=';
GREATER_EQUAL : '>=';
WS : [ \t\r\n]+ -> skip;