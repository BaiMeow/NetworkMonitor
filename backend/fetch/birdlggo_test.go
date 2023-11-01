package fetch

import (
	"fmt"
	"testing"
)

func TestBirdLgGo_GetData(t *testing.T) {
	fetcher, err := Spawn["bird-lg-go"](map[string]any{
		"api":    "https://lg.internal.potat0.cc/api/",
		"server": "she",
		"type":   "bird",
		"args":   "show ospf state all OSPF_Potat0_v4",
	})
	if err != nil {
		t.Error(err)
		return
	}
	data, err := fetcher.GetData()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(data))
}
