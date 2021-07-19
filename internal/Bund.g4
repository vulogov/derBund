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
    | true_term
    | false_term
    | call_term
    | begin
    | end
  )
;

ns
  : '[' name=NAME ':' (body+=term)* ';;'
  ;

block
  : '(' (body+=term)* ')'
  ;

//
// Various terms
//
true_term:    value=TRUE ;
false_term:   value=FALSE ;

call_term:    value=NAME ;
begin:        value=TOBEGIN;
end:          value=TOEND ;



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

TOBEGIN: ':' ;
TOEND:   ';' ;
SLASH:   '/' ;

NAME
  : ID_START ID_CONTINUE*
  ;

SKIP_
  : ( SPACES | COMMENT | LINE_JOINING | CCOMMENT | BLOCK_COMMENT ) -> skip
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

fragment LINE_JOINING
  : '\\' SPACES? ( '\r'? '\n' | '\r' | '\f' )
  ;

fragment COMMENT
  : '##' ~[\r\n\f]*
  ;

fragment BLOCK_COMMENT
  :   '/*' .*? '*/'
  ;

fragment CCOMMENT
  :   '//' ~[\r\n]*
  ;

fragment SPACES
  : [ \t]+
  ;
