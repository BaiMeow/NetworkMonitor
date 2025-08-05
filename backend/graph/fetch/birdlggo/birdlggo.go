package birdlggo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/graph/fetch"
	"github.com/BaiMeow/NetworkMonitor/trace"
	"github.com/BaiMeow/NetworkMonitor/utils"
	"github.com/BaiMeow/NetworkMonitor/utils/ctxex"
	"net/http"
)

func init() {
	fetch.Register("bird-lg-go", func(config map[string]any) (fetch.Fetcher, error) {
		api, ok := config["api"].(string)
		if !ok {
			return nil, fmt.Errorf("host is not string")
		}
		server, ok := config["server"].(string)
		if !ok {
			return nil, fmt.Errorf("server is not string")
		}
		Type, ok := config["req-type"].(string)
		if !ok {
			return nil, fmt.Errorf("type is not string")
		}
		args, ok := config["args"].(string)
		if !ok {
			return nil, fmt.Errorf("args is not string")
		}
		return &BirdLgGo{
			API:    api,
			Server: server,
			Type:   Type,
			Args:   args,
		}, nil
	})
}

type BirdLgGo struct {
	fetch.Base
	API    string
	Server string
	Type   string
	Args   string
}

type birdLgGoPayload struct {
	Servers []string `json:"servers"`
	Type    string   `json:"type"`
	Args    string   `json:"args"`
}

type birdLgGoResp struct {
	Error  string           `json:"error"`
	Result []birdLgGoResult `json:"result"`
}

type birdLgGoResult struct {
	Server string `json:"server"`
	Data   string `json:"data"`
}

func (b *BirdLgGo) GetData(ctx context.Context) (any, error) {
	ctx, span := trace.Tracer.Start(ctx,
		"fetch/birdlggo/BirdLgGo.GetData",
	)
	defer span.End()

	payload, err := json.Marshal(birdLgGoPayload{
		Servers: []string{b.Server},
		Type:    b.Type,
		Args:    b.Args,
	})
	if err != nil {
		return nil, fmt.Errorf("fail to marshal payload:%v", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, b.API, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("new request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fail to post request:%v", err)
	}
	defer resp.Body.Close()

	rbody, err := ctxex.IoReadAll(ctx, resp.Body)
	if err != nil {
		return nil, fmt.Errorf("fail to read body:%v", err)
	}

	var res birdLgGoResp
	if err := json.Unmarshal(rbody, &res); err != nil {
		return nil, fmt.Errorf("fail to unmarshal body:%v", err)
	}

	if res.Error != "" {
		return nil, fmt.Errorf("bird-lg-go err:%v", res.Error)
	}

	if len(res.Result) <= 0 {
		return nil, fmt.Errorf("empty bird-lg-go result:%v", res.Result)
	}
	data, ok := utils.FindFunc(res.Result, func(s birdLgGoResult) bool {
		return s.Server == b.Server
	})
	if !ok {
		return nil, fmt.Errorf("target bird-lg-go result not found")
	}

	return []byte(data.Data), nil
}

func (b *BirdLgGo) CleanUp() error {
	return nil
}
