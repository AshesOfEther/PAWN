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
    | expression op='&&' expression                             # And
    | expression op='||' expression                             # Or
    | member                                                    # MemberExpr
    | literal                                                   # LiteralExpr
    ;

member
    : member '.' NAME    # Property
    | NAME               # Name
    ;

literal
    : '[' (expression (',' expression)* ','?)? ']'    # List
    | BOOL                                            # Bool
    | NUMBER                                          # Number
    | STRING                                          # String
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
    : 'apply' expression '{' statementList '}'            # Apply
    | 'while' expression '{' statementList '}'            # While
    | 'if' expression '{' statementList '}'
        ('else' 'if' expression '{' statementList '}' )*
        ('else' '{' statementList '}')?                   # If
    | 'fn' NAME functionDeclArgs '{' statementList '}'    # FunctionDecl
    | 'return' expression                                 # Return
    | expression ';'                                      # ExpressionStatement
    | NAME '=' expression                                 # Assignment
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
LESS_EQUAL : '<=';
GREATER_EQUAL : '>=';
WS : [ \t\r\n]+ -> skip;
