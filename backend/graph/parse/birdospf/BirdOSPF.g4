grammar BirdOSPF;

state            : area+ EOF;
area             : 'area' IP
                  ( router | network )*;
router           : 'router' IP
                    distance
                    routerEntry*;
network          : 'network' Prefix
                    DR IP
                    distance
                    ('router' IP)*
                  ;
routerEntry      : ('router' IP
                  | 'stubnet' Prefix
                  | 'xnetwork' Prefix
                  | 'network' Prefix
                  | 'external' Prefix
                  | 'xrouter' IP )
                  ( Metric | Metric2 ) INT ( 'via' IP )?
                  ;
distance         : Distance INT
                  | Unreachable;

// Lexer Rules
DR          : 'dr';
Distance    : 'distance';
Unreachable : 'unreachable';
Metric      : 'metric';
Metric2     : 'metric2';
Prefix      : IP '/' INT;
VERSION     : 'BIRD' WS+ INT '.' INT ('.' INT)? WS+ 'ready.' -> skip;
IP          : [0-9]+ '.' [0-9]+ '.' [0-9]+ '.' [0-9]+;
STRING      : '"' .*? '"'
            | [a-zA-Z][a-zA-Z0-9_-]*
            ;
INT         : [0-9]+;
WS          : [ \t\r\n]+ -> skip;