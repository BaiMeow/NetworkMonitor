package tcp

import (
	"context"
	"testing"
	"time"

	"github.com/BaiMeow/NetworkMonitor/conf"
	"github.com/BaiMeow/NetworkMonitor/graph/fetch"
	"github.com/BaiMeow/NetworkMonitor/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

func init() {
	trace.Tracer = noop.Tracer{}
	conf.ProbeTimeout = time.Second * 10
}

func TestTcp_GetData(t *testing.T) {
	fetcher, err := fetch.Spawn["tcp"](map[string]any{
		"addr": "network.srv.csmantle.top:30050",
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
	t.Log(string(data.([]byte)))
}
