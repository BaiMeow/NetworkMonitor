package fetch

import (
	"encoding/base64"
	"testing"
)

func TestROS6_GetData(t *testing.T) {
	ros6 := &ROS6{
		Address:  "10.28.0.7:8728",
		Username: "networkMonitor",
		Password: "7ezvrpDYbGpyaWhwesjH",
	}
	resp, err := ros6.GetData()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(base64.StdEncoding.EncodeToString(resp))
}
