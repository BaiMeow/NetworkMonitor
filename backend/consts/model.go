package consts

import "time"

type LinkTime struct {
	Time  time.Time `json:"time"`
	Links int       `json:"links"`
}

type DirectedLinkTime struct {
	Time      time.Time `json:"time"`
	InDegree  int       `json:"in_degree"`
	OutDegree int       `json:"out_degree"`
}