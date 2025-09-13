package bgp

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/BaiMeow/NetworkMonitor/graph/fetch"
	"github.com/BaiMeow/NetworkMonitor/trace"
	apipb "github.com/osrg/gobgp/v3/api"
	"go.opentelemetry.io/otel/trace/noop"
)

func init() {
	trace.Tracer = noop.Tracer{}
}

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
	for {
		time.Sleep(time.Second)
		start := time.Now()
		data, err := bgp.GetData(context.Background())
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println(time.Since(start).String())
		fmt.Println(countPath(data.([]*apipb.Destination)))
	}
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
	data, err := bgp.GetData(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(countPath(data.([]*apipb.Destination)))
}
