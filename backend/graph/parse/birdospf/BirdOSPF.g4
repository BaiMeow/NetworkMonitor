grammar BirdOSPF;

state            : area+;
area             : 'area' IP ('distance' INT)? router*;
router           : 'router' IP (('distance' INT) | routerEntry)+;
routerEntry      : ('router' IP
                  | 'stubnet' Prefix
                  | 'xnetwork' Prefix
                  | 'external' Prefix
                  | 'xrouter' IP )
                  ( Metric | Metric2 ) INT;


// Lexer Rules
Metric  : 'metric';
Metric2 : 'metric2';
Prefix  : IP '/' INT;
VERSION : 'BIRD' WS+ INT '.' INT ('.' INT)? WS+ 'ready.' -> skip;
IP      : [0-9]+ '.' [0-9]+ '.' [0-9]+ '.' [0-9]+;
STRING  : '"' .*? '"'
        | [a-zA-Z][a-zA-Z0-9_-]*
        ;
INT  : [0-9]+;
WS      : [ \t\r\n]+ -> skip;