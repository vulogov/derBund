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
    | lambda
    | alambda
    | datablock
    | floatblock
    | intblock
    | uintblock
    | trueblock
    | falseblock
    | ignoreblock
    | true_term
    | false_term
    | string_term
    | integer
    | uinteger
    | float
    | ufloat
    | complex_term
    | call_term
    | call_sys
    | cmd_term
    | cmd_sys
    | begin
    | end
    | execute_term
    | return_term
    | drop
    | duplicate
  )
;

data
  : ( true_term
    | false_term
    | string_term
    | integer
    | uinteger
    | float
    | ufloat
    | complex_term
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

uintblock
  : '(uint' (body+=uinteger)* ')'
  ;

trueblock
  : '(true' (body+=term)* ')'
  ;
falseblock
  : '(false' (body+=term)* ')'
  ;

ignoreblock
  : '(ignore' (body+=term)* ')'
  ;

lambda
  : '[' name=NAME ']' (body+=term)* '.'
  ;

alambda
  : '[]' (body+=term)* '.'
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
uinteger:     value=UINTEGER ;
float:        value=FLOAT_NUMBER ;
ufloat:       value=UFLOAT_NUMBER ;
allfloat:     value=(FLOAT_NUMBER|UFLOAT_NUMBER) ;
complex_term: value=COMPLEX_NUMBER ;
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
execute_term: value=EXECUTE ;
return_term:  value=RETURN ;



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

UINTEGER
  :  ('U'|'u') DECIMAL_INTEGER
  ;


DECIMAL_INTEGER
  : NON_ZERO_DIGIT DIGIT*
  | '0'+
  ;

FLOAT_NUMBER
  : EXPONENT_OR_POINT_FLOAT
  ;

UFLOAT_NUMBER
  : EXPONENT_OR_POINT_UFLOAT
  ;

COMPLEX_NUMBER
  : '(' FLOAT_NUMBER FLOAT_NUMBER 'i)'
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
  : ('≉'|'≈'|'∐'|'∏'|'∇'|'∆'|'∪'|'∩'|'∉'|'∈'|'⊉'|'⊈'|'⊇'|'⊆'|'⊅'|'⊄'|'⊃'|'⊂'|'÷'|'\\'|'+'|'-'|'&'|'='|'<'|'>')+
  ;

SYS
  : ('_'|'@'|'%'|'?'|'*')+
  ;

EXECUTE
  : ('!')+
  ;

RETURN
  : ('$')+
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

fragment EXPONENT_OR_POINT_UFLOAT
  : ('U'|'u') ([0-9]+ | POINT_FLOAT) [eE] [+-]? [0-9]+
  | ('U'|'u') POINT_FLOAT
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
