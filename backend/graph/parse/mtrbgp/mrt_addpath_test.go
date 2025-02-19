package mtrbgp

import (
	"github.com/BaiMeow/NetworkMonitor/graph/parse"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestMrtAddPath(t *testing.T) {
	dir, err := os.ReadDir("./testfile")
	if err != nil {
		t.Error(err)
		return
	}
	for _, entry := range dir {
		t.Run(entry.Name(), func(t *testing.T) {
			mrt := MrtAddPath{}
			file, err := os.Open(filepath.Join("./testfile", entry.Name()))
			if err != nil {
				t.Error(err)
				return
			}
			data, err := io.ReadAll(file)
			if err != nil {
				t.Error(err)
				return
			}
			var drawing parse.Drawing
			drawing.BGP = new(parse.BGP)
			err = mrt.ParseAndMerge(data, &drawing)
			if err != nil {
				t.Error(err)
				return
			}
			for _, v := range drawing.BGP.AS {
				t.Log(*v)
			}
			t.Log(drawing.BGP.Link)
		})
	}
}
