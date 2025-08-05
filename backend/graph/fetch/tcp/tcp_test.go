package tcp

import (
	"context"
	"github.com/BaiMeow/NetworkMonitor/graph/fetch"
	"github.com/BaiMeow/NetworkMonitor/trace"
	"go.opentelemetry.io/otel/trace/noop"
	"testing"
)

func init() {
	trace.Tracer = noop.Tracer{}
}

func TestTcp_GetData(t *testing.T) {
	fetcher, err := fetch.Spawn["tcp"](map[string]any{
		"addr": "1.1.1.1:11111",
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
