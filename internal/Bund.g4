//
//
//

grammar Bund;

expressions
  : term*
;

term
  : ( ns
  )
;

ns
  : '[' name=NAME ':' (body+=term)* ';;'
  ;

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
