package ros

import (
	"context"
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/graph/fetch"
	"github.com/go-routeros/routeros/v3"
	"github.com/go-routeros/routeros/v3/proto"
	"net"
	"time"
)

func init() {
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
	fetch.Base
	Address  string //<IP or domain name>:port
	Username string
	Password string
}

func (R *ROS) GetData(ctx context.Context) (any, error) {
	dialer := &net.Dialer{
		Timeout: time.Second * 10,
	}
	conn, err := dialer.DialContext(ctx, "tcp", R.Address)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client, _ := routeros.NewClient(conn)
	defer client.Close()

	err = client.LoginContext(ctx, R.Username, R.Password)
	if err != nil {
		return nil, err
	}
	reply1, err := client.RunContext(ctx, "/routing/ospf/lsa/print", "detail", "?type=router")
	if err != nil {
		return nil, err
	}
	reply2, err := client.RunContext(ctx, "/routing/ospf/area/print", "detail")
	if err != nil {
		return nil, err
	}
	return [2][]*proto.Sentence{reply1.Re, reply2.Re}, nil
}
