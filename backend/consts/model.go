package consts

import "time"

type LinkTime struct {
	Time  time.Time `json:"time"`
	Links int       `json:"links"`
}
