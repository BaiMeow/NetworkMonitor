package bgp

import (
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/graph/fetch"
	"testing"
	"time"
)

func TestBGPDial(t *testing.T) {
	bgp, err := fetch.Spawn["bgp"](map[string]any{
		"mode":          "dial",
		"asn":           4211110102,
		"router-id":     "172.16.4.100",
		"peer-asn":      4211110101,
		"neighbor-addr": "172.16.4.6",
	})
	if err != nil {
		t.Error(err)
		return
	}
	data, err := bgp.GetData()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(data)
}

func TestBGPListen(t *testing.T) {
	bgp, err := fetch.Spawn["bgp"](map[string]any{
		"mode":      "listen",
		"asn":       4211110102,
		"router-id": "172.16.4.100",
	})
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(time.Second * 10)
	data, err := bgp.GetData()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(data)
}
