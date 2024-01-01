package fetch

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/go-routeros/routeros"
	"github.com/go-routeros/routeros/proto"
)

var _ Fetcher = (*ROS)(nil)

func init() {
	gob.Register(&proto.Sentence{})

	Register("ros", func(config map[string]any) (Fetcher, error) {
		addr, ok := config["address"].(string)
		if !ok {
			return nil, fmt.Errorf("host is not string")
		}
		username, ok := config["username"].(string)
		if !ok {
			return nil, fmt.Errorf("server is not string")
		}
		passwd, ok := config["password"].(string)
		if !ok {
			return nil, fmt.Errorf("type is not string")
		}
		return &ROS{
			Address:  addr,
			Username: username,
			Password: passwd,
		}, nil
	})
}

type ROS struct {
	Address  string //<IP or domain name>:port
	Username string
	Password string
}

func (R *ROS) GetData() ([]byte, error) {
	client, err := routeros.Dial(R.Address, R.Username, R.Password)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	reply, err := client.Run("/routing/ospf/lsa/print", "detail", "?type=router")
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	gob.NewEncoder(buf).Encode(reply.Re)
	return buf.Bytes(), nil
}
