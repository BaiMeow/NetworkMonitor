package parse

import "sync"

type Drawing struct {
	OSPF map[uint32]*OSPF
	BGP  BGP
	sync.Mutex
}
