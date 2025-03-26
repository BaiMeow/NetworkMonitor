package rosospf

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"github.com/BaiMeow/NetworkMonitor/graph/entity"
	"github.com/go-routeros/routeros/v3/proto"
	"testing"

	"github.com/stretchr/testify/assert"
)

// todo fix test
func TestROSOSPF(t *testing.T) {
	testcases := []struct {
		name        string
		feterOutput string
		ospfData    string
	}{
		{
			name:        "ROS6",
			feterOutput: "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",
			ospfData:    "[{\"area_id\":\"atomInner\",\"router\":[{\"router_id\":\"10.28.0.1\"},{\"router_id\":\"10.28.0.2\"},{\"router_id\":\"10.28.0.5\"},{\"router_id\":\"10.28.0.7\"},{\"router_id\":\"10.28.0.8\"}],\"links\":[{\"src\":\"10.28.0.1\",\"dst\":\"10.28.0.2\",\"cost\":1},{\"src\":\"10.28.0.1\",\"dst\":\"10.28.0.5\",\"cost\":1},{\"src\":\"10.28.0.1\",\"dst\":\"10.28.0.7\",\"cost\":1},{\"src\":\"10.28.0.2\",\"dst\":\"10.28.0.1\",\"cost\":20},{\"src\":\"10.28.0.2\",\"dst\":\"10.28.0.5\",\"cost\":20},{\"src\":\"10.28.0.5\",\"dst\":\"10.28.0.1\",\"cost\":1},{\"src\":\"10.28.0.5\",\"dst\":\"10.28.0.2\",\"cost\":1},{\"src\":\"10.28.0.5\",\"dst\":\"10.28.0.7\",\"cost\":1},{\"src\":\"10.28.0.5\",\"dst\":\"10.28.0.8\",\"cost\":1},{\"src\":\"10.28.0.7\",\"dst\":\"10.28.0.5\",\"cost\":10},{\"src\":\"10.28.0.7\",\"dst\":\"10.28.0.1\",\"cost\":10},{\"src\":\"10.28.0.7\",\"dst\":\"10.28.0.8\",\"cost\":10},{\"src\":\"10.28.0.8\",\"dst\":\"10.28.0.5\",\"cost\":1},{\"src\":\"10.28.0.8\",\"dst\":\"10.28.0.7\",\"cost\":1}]}]",
		},
		{
			name:        "ROS7",
			feterOutput: "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",
			ospfData:    "[{\"area_id\":\"atomInner\",\"router\":[{\"router_id\":\"10.28.0.1\"},{\"router_id\":\"10.28.0.2\"},{\"router_id\":\"10.28.0.5\"},{\"router_id\":\"10.28.0.7\"},{\"router_id\":\"10.28.0.8\"}],\"links\":[{\"src\":\"10.28.0.1\",\"dst\":\"10.28.0.2\",\"cost\":1},{\"src\":\"10.28.0.1\",\"dst\":\"10.28.0.5\",\"cost\":1},{\"src\":\"10.28.0.1\",\"dst\":\"10.28.0.7\",\"cost\":1},{\"src\":\"10.28.0.2\",\"dst\":\"10.28.0.1\",\"cost\":20},{\"src\":\"10.28.0.2\",\"dst\":\"10.28.0.5\",\"cost\":20},{\"src\":\"10.28.0.5\",\"dst\":\"10.28.0.1\",\"cost\":1},{\"src\":\"10.28.0.5\",\"dst\":\"10.28.0.2\",\"cost\":1},{\"src\":\"10.28.0.5\",\"dst\":\"10.28.0.7\",\"cost\":1},{\"src\":\"10.28.0.5\",\"dst\":\"10.28.0.8\",\"cost\":1},{\"src\":\"10.28.0.7\",\"dst\":\"10.28.0.5\",\"cost\":10},{\"src\":\"10.28.0.7\",\"dst\":\"10.28.0.1\",\"cost\":10},{\"src\":\"10.28.0.7\",\"dst\":\"10.28.0.8\",\"cost\":10},{\"src\":\"10.28.0.8\",\"dst\":\"10.28.0.5\",\"cost\":1},{\"src\":\"10.28.0.8\",\"dst\":\"10.28.0.7\",\"cost\":1}]}]",
		},
	}

	for _, v := range testcases {
		t.Run(v.name, func(t *testing.T) {
			gobOutput, err := base64.StdEncoding.DecodeString(v.feterOutput)
			var sentences []*proto.Sentence
			gob.NewDecoder(bytes.NewReader(gobOutput)).Decode(&sentences)
			p := RosOSPF{}
			res, err := p.Parse(context.Background(), sentences)
			assert.NoError(t, err)
			o := &entity.OSPF{}
			_ = json.Unmarshal([]byte(v.ospfData), o)
			assert.Equal(t, o, res)
		})
	}

}
