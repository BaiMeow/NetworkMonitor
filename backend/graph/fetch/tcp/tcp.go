package tcp

import (
	"context"
	"fmt"
	"io"
	"net"
	"strconv"
	"time"

	"github.com/BaiMeow/NetworkMonitor/conf"
	"github.com/BaiMeow/NetworkMonitor/graph/fetch"
	"github.com/BaiMeow/NetworkMonitor/trace"
)

func init() {
	fetch.Register("tcp", func(config map[string]any) (fetch.Fetcher, error) {
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

type Tcp struct {
	fetch.Base
	host string
	port int
}

func (t *Tcp) GetData(ctx context.Context) (any, error) {
	ctx, span := trace.Tracer.Start(ctx,
		"fetch/tcp/Tcp.GetData",
	)
	defer span.End()
	now := time.Now()
	dialer := &net.Dialer{}
	conn, err := dialer.DialContext(ctx, "tcp", net.JoinHostPort(t.host, strconv.Itoa(t.port)))
	if err != nil {
		return nil, fmt.Errorf("fail to dial tcp: %v", err)
	}
	defer conn.Close()
	_ = conn.SetReadDeadline(now.Add(conf.ProbeTimeout * 4 / 5))
	return io.ReadAll(conn)
}
