package birdospf

import (
	_ "embed"
	"os"
	"path/filepath"
	"testing"

	"github.com/BaiMeow/NetworkMonitor/graph/parse"
)

func TestMUSTParse(t *testing.T) {
	entries, err := os.ReadDir("testdata")
	if err != nil {
		t.Error(err)
		return
	}

	for _, entry := range entries {
		path := filepath.Join("testdata", entry.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			t.Errorf("read %s fail: %v", path, err)
			return
		}
		str := string(data)
		var p BirdOSPF
		p.asn = 4242424242
		p.Init([]byte(str))
		var drawing parse.Drawing
		drawing.OSPF = make(map[uint32]*parse.OSPF)
		err = p.ParseAndMerge(&drawing)
		if err != nil {
			t.Errorf("parse %s fail: %v", path, err)
			return
		}
		t.Log(drawing.OSPF[4242424242])
	}
}
