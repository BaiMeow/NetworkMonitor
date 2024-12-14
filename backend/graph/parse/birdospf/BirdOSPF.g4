grammar BirdOSPF;

state            : area+;
area             : 'area' IP ('distance' INT)? router*;
router           : 'router' IP (('distance' INT) | routerEntry)+;
routerEntry      : 'router' IP ('metric' INT)
                  | 'stubnet' Prefix ('metric' INT)
                  | 'xnetwork' Prefix ('metric' INT)
                  | 'external' Prefix ('metric2' INT)
                  | 'xrouter' IP ('metric' INT);


// Lexer Rules
Prefix  : IP '/' INT;
VERSION : 'BIRD' WS+ INT '.' INT '.' INT WS+ 'ready.' -> skip;
IP      : [0-9]+ '.' [0-9]+ '.' [0-9]+ '.' [0-9]+;
STRING  : '"' .*? '"'
        | [a-zA-Z][a-zA-Z0-9_-]*
        ;
INT  : [0-9]+;
WS      : [ \t\r\n]+ -> skip;