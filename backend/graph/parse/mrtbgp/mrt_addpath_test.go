package mrtbgp

import (
	"context"
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
			res, err := mrt.Parse(context.Background(), data)
			if err != nil {
				t.Error(err)
				return
			}
			for _, v := range res.AS {
				t.Log(*v)
			}
			t.Log(res.Link)
		})
	}
}
