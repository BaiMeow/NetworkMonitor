package fetch

import "testing"

func TestTcp_GetData(t *testing.T) {
	fetcher, err := Spawn["tcp"](map[string]any{
		"addr": "1.1.1.1:11111",
	})
	if err != nil {
		t.Error(err)
		return
	}
	data, err := fetcher.GetData()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(data))
}
