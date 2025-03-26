package ros

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/gob"
	"github.com/go-routeros/routeros/v3/proto"
	"testing"
)

func TestROS6_GetData(t *testing.T) {
	ros6 := &ROS{
		Address:  "10.28.0.7:8728",
		Username: "networkMonitor",
		Password: "7ezvrpDYbGpyaWhwesjH",
	}
	resp, err := ros6.GetData(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	sentences := resp.([2][]*proto.Sentence)
	gob.Register(sentences)
	var buf bytes.Buffer
	gob.NewEncoder(&buf).Encode(sentences)
	t.Log(base64.StdEncoding.EncodeToString(buf.Bytes()))
}

func TestROS7_GetData(t *testing.T) {
	ros7 := &ROS{
		Address:  "10.28.0.1:8728",
		Username: "networkMonitor",
		Password: "7ezvrpDYbGpyaWhwesjH",
	}
	resp, err := ros7.GetData(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	sentences := resp.([2][]*proto.Sentence)
	gob.Register(sentences)
	var buf bytes.Buffer
	gob.NewEncoder(&buf).Encode(sentences)
	t.Log(base64.StdEncoding.EncodeToString(buf.Bytes()))
}
