package birdospf

import (
	"context"
	_ "embed"
	"github.com/BaiMeow/NetworkMonitor/trace"
	"go.opentelemetry.io/otel/trace/noop"
	"os"
	"path/filepath"
	"testing"
)

func init() {
	trace.Tracer = noop.Tracer{}
}

func TestMUSTParse(t *testing.T) {
	entries, err := os.ReadDir("./testdata")
	if err != nil {
		t.Error(err)
		return
	}

	for _, entry := range entries {
		t.Run(entry.Name(), func(t *testing.T) {
			path := filepath.Join("./testdata", entry.Name())
			data, err := os.ReadFile(path)
			if err != nil {
				t.Errorf("read %s fail: %v", path, err)
				return
			}
			var p BirdOSPF
			gr, err := p.Parse(context.Background(), data)
			if err != nil {
				t.Errorf("parse %s fail: %v", path, err)
				return
			}
			t.Log(gr)
		})
	}
}
