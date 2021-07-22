//
//
//

grammar Bund;

expressions
  : root_term*
;

root_term
   : ( ns
     | block
   )
;

term
  : ( ns
    | block
    | datablock
    | floatblock
    | intblock
    | trueblock
    | falseblock
    | true_term
    | false_term
    | string_term
    | integer
    | float
    | call_term
    | call_sys
    | cmd_term
    | cmd_sys
    | begin
    | end
    | execute_term
    | drop
    | duplicate
  )
;

data
  : ( true_term
    | false_term
    | string_term
    | integer
    | float
    | call_term
    | call_sys
    | cmd_term
    | cmd_sys
  )
;

ns
  : '[' name=NAME ':' (body+=term)* ';;'
  ;

block
  : '(' (body+=term)* ')'
  ;

datablock
  : '(*' (body+=data)* ')'
  ;

floatblock
  : '(float' (body+=float)* ')'
  ;

intblock
  : '(int' (body+=integer)* ')'
  ;

trueblock
  : '(true' (body+=term)* ')'
  ;
falseblock
  : '(false' (body+=term)* ')'
  ;

lambda
  : '[' name=NAME ']' (body+=term)* '.'
  ;
lambda_cmd
  : '[[' name=CMD ']]' (body+=term)* '.'
  ;

//
// Various terms
//
true_term:    value=TRUE ;
false_term:   value=FALSE ;
string_term:  value=STRING ;
integer:      value=INTEGER ;
float:        value=FLOAT_NUMBER ;
call_term:    value=NAME ;
call_sys
  : value=NAME '.(' syscmd=(SYS|NAME) ')'
  ;
cmd_term:     value=CMD ;
cmd_sys
  : value=CMD '.(' syscmd=(SYS|NAME) ')'
  ;
begin:        value=TOBEGIN;
end:          value=TOEND ;
drop:         value=DROP ;
duplicate:    value=DUPLICATE ;
execute_term: value=EXECUTE;



TRUE
  : 'true'
  | 'True'
  | 'TRUE'
  | 'T'
  | 'yes'
  | 'Yes'
  | 'YES'
  ;

FALSE
  : 'false'
  | 'False'
  | 'FALSE'
  | 'F'
  | 'no'
  | 'No'
  | 'NO'
  ;

INTEGER
  :  (SIGN)? DECIMAL_INTEGER
  ;

DECIMAL_INTEGER
  : NON_ZERO_DIGIT DIGIT*
  | '0'+
  ;

FLOAT_NUMBER
  : EXPONENT_OR_POINT_FLOAT
  ;

STRING
  : SHORT_STRING
  | LONG_STRING
  ;

TOBEGIN: ':' ;
TOEND:   ';' ;
SLASH:   '/' ;
DROP:    ',' ;
DUPLICATE: '^';

NAME
  : ID_START ID_CONTINUE*
  ;

CMD
  : ('+'|'-'|'&'|'='|'<'|'>')+
  ;

SYS
  : ('@'|'$'|'`'|'?'|'*')+
  ;

EXECUTE
  : ('!')+
  ;

COMMENT
  : '##' ~[\r\n]* -> skip
  ;
BLOCK_COMMENT
  :   '/*' .*? '*/' -> skip
  ;
WS
  : [ \r\n\t]+ -> skip
  ;
SHEBANG
  : '#' '!' ~('\n'|'\r')* -> channel(HIDDEN)
  ;
//
// fragments
//

fragment NON_ZERO_DIGIT
  : [1-9]
  ;

fragment DIGIT
  : [0-9]
  ;

fragment EXPONENT_OR_POINT_FLOAT
  : (SIGN)? ([0-9]+ | POINT_FLOAT) [eE] [+-]? [0-9]+
  | (SIGN)? POINT_FLOAT
  ;

fragment POINT_FLOAT
  : [0-9]* '.' [0-9]+
  | [0-9]+ '.'
  ;

fragment SIGN
  : ('+' | '-')
  ;

fragment RN
  : '\r'? '\n'
  ;

fragment SHORT_STRING
  : '\'' ('\\' (RN | .) | ~[\\\r\n'])* '\''
  | '"'  ('\\' (RN | .) | ~[\\\r\n"])* '"'
  ;

fragment LONG_STRING
  : '\'\'\'' LONG_STRING_ITEM*? '\'\'\''
  | '"""' LONG_STRING_ITEM*? '"""'
  ;

fragment LONG_STRING_ITEM
  : ~'\\'
  | '\\' (RN | .)
  ;

fragment ID_START
 : ([A-Z]|[a-z]|SLASH)
 | [a-z]
 ;

fragment ID_CONTINUE
 : ID_START
 | [0-9]
 | SLASH
 ;
