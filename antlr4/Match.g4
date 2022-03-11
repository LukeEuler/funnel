grammar Match;

sentence    : clause;

clause      : condition
            | NOT clause
            | clause AND clause
            | clause OR clause
            | '(' clause ')'
            ;

condition   : key EXIST
            | key EQUAL VALUE
            | key CONTAINS VALUE
            ;

key         : KI ('.' KI)*;

// Tokens
AND         : '&';
OR          : '|';
NOT         : '!';
EXIST       : '+';
EQUAL       : '=';
CONTAINS    : '>';
KI          : ([a-zA-Z] | '0'..'9' | '_')+;
VALUE       : '\'' .*? '\'';
WHITESPACE  : [ \r\n\t]+ -> skip;
