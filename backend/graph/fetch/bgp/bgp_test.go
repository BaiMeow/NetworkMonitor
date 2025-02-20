package bgp

import (
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/graph/fetch"
	apipb "github.com/osrg/gobgp/v3/api"
	"testing"
	"time"
)

func TestBGPDial(t *testing.T) {
	bgp, err := fetch.Spawn["bgp"](map[string]any{
		"mode":          "dial",
		"asn":           2,
		"router-id":     "172.16.4.102",
		"peer-asn":      4211110101,
		"neighbor-addr": "172.16.4.4",
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
	fmt.Println(countPath(data.([]*apipb.Destination)))
}

func TestBGPListen(t *testing.T) {
	bgp, err := fetch.Spawn["bgp"](map[string]any{
		"mode":      "listen",
		"asn":       2,
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
	fmt.Println(countPath(data.([]*apipb.Destination)))
}

func countPath(dess []*apipb.Destination) int {
	var sum int
	for _, des := range dess {
		sum += len(des.GetPaths())
	}
	return sum
}
