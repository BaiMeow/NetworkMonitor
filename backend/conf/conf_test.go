package conf

import (
	"github.com/stretchr/testify/assert"
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

func TestEnvConf(t *testing.T) {
	os.Setenv("NETM_INFLUXDB_TOKEN", "66666")
	err := Init()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "66666", Influxdb.Token)
}
