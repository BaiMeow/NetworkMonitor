package conf

import (
	"fmt"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	err := Init()
	if err != nil {
		t.Error(err)
	}
	t.Log(Probes)
}

func Test111(t *testing.T) {
	err := Init()
	if err != nil {
		t.Error(err)
	}

	os.MkdirAll("tmp", 0777)
	os.Chdir("tmp")

	for k, v := range Metas {
		f, _ := os.Create(k + ".json")
		f.Write([]byte(fmt.Sprintf(`{
    "display":"%s"
}`, v["name"])))
	}
}
