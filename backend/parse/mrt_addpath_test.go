package parse

import (
	"io"
	"os"
	"testing"
)

func TestMrtAddPath(t *testing.T) {
	mrt := MrtAddPath{}
	file, err := os.Open("./testfile/06-09-2023-08-29.mrt")
	if err != nil {
		t.Error(err)
		return
	}
	data, err := io.ReadAll(file)
	if err != nil {
		t.Error(err)
		return
	}
	mrt.Init(data)
	var drawing Drawing
	err = mrt.ParseAndMerge(&drawing)
	if err != nil {
		t.Error(err)
		return
	}
	for _, v := range drawing.BGP.AS {
		t.Log(*v)
	}
	t.Log(drawing.BGP.Link)
}
