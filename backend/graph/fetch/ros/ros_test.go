package ros

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"github.com/go-routeros/routeros/proto"
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
	sentences := resp.([]*proto.Sentence)
	gob.Register(sentences)
	var buf bytes.Buffer
	gob.NewEncoder(&buf).Encode(sentences)
	t.Log(base64.StdEncoding.EncodeToString(buf.Bytes()))

	ros7 := &ROS{
		Address:  "10.28.0.1:8728",
		Username: "networkMonitor",
		Password: "7ezvrpDYbGpyaWhwesjH",
	}
	resp, err = ros7.GetData()
	if err != nil {
		t.Fatal(err)
	}
	sentences = resp.([]*proto.Sentence)
	gob.Register(sentences)
	gob.NewEncoder(&buf).Encode(sentences)
	t.Log(base64.StdEncoding.EncodeToString(buf.Bytes()))
}
