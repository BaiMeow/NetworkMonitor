package birdlggo

import (
	"context"
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/graph/fetch"
	"github.com/BaiMeow/NetworkMonitor/trace"
	"go.opentelemetry.io/otel/trace/noop"
	"testing"
)

func init() {
	trace.Tracer = noop.Tracer{}
}

func TestBirdLgGo_GetData(t *testing.T) {
	fetcher, err := fetch.Spawn["bird-lg-go"](map[string]any{
		"api":      "https://lg.internal.potat0.cc/api/",
		"server":   "she",
		"req-type": "bird",
		"args":     "show ospf state all OSPF_Potat0_v4",
	})
	if err != nil {
		t.Error(err)
		return
	}
	data, err := fetcher.GetData(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(data.([]byte)))
}
