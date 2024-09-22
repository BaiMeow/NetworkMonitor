package fetch

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
)

func init() {
	Register("tcp", func(config map[string]any) (Fetcher, error) {
		addr, ok := config["addr"].(string)
		if !ok {
			return nil, fmt.Errorf("addr is not string")
		}
		host, port, err := net.SplitHostPort(addr)
		if err != nil {
			return nil, fmt.Errorf("fail to split host port: %v", err)
		}

		portInt, err := strconv.Atoi(port)
		if err != nil {
			return nil, fmt.Errorf("fail to convert port to int: %v", err)
		}

		return &Tcp{
			host: host,
			port: portInt,
		}, nil
	})
}

var _ Fetcher = (*Tcp)(nil)

type Tcp struct {
	host string
	port int
}

func (t *Tcp) GetData() ([]byte, error) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(t.host, strconv.Itoa(t.port)), time.Second*30)
	if err != nil {
		return nil, fmt.Errorf("fail to dial tcp: %v", err)
	}
	defer conn.Close()
	return io.ReadAll(conn)
}
