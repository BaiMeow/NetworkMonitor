package fetch

import (
	"encoding/base64"
	"testing"
)

func TestROS_GetData(t *testing.T) {
	ros6 := &ROS{
		Address:  "10.28.0.7:8728",
		Username: "networkMonitor",
		Password: "7ezvrpDYbGpyaWhwesjH",
	}
	resp, err := ros6.GetData()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(base64.StdEncoding.EncodeToString(resp))

	ros7 := &ROS{
		Address:  "10.28.0.1:8728",
		Username: "networkMonitor",
		Password: "7ezvrpDYbGpyaWhwesjH",
	}
	resp, err = ros7.GetData()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(base64.StdEncoding.EncodeToString(resp))
}
