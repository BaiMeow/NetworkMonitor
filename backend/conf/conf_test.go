package conf

import (
	"testing"
)

func TestInit(t *testing.T) {
	err := Init()
	if err != nil {
		t.Error(err)
	}
	t.Log(Probes)
}
