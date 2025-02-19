package ros

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/graph/fetch"
	"github.com/go-routeros/routeros"
	"github.com/go-routeros/routeros/proto"
	"net"
	"time"
)

func init() {
	gob.Register(&proto.Sentence{})

	fetch.Register("ros", func(config map[string]any) (fetch.Fetcher, error) {
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

func (R *ROS) GetData() (any, error) {
	conn, err := net.DialTimeout("tcp", R.Address, time.Second*10)
	if err != nil {
		return nil, err
	}
	client, _ := routeros.NewClient(conn)
	defer client.Close()
	err = client.Login(R.Username, R.Password)
	if err != nil {
		return nil, err
	}
	reply, err := client.Run("/routing/ospf/lsa/print", "detail", "?type=router")
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	gob.NewEncoder(buf).Encode(reply.Re)
	return buf.Bytes(), nil
}
