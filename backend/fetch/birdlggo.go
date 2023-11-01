package fetch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/utils"
	"io"
	"net/http"
)

func init() {
	Register("bird-lg-go", func(config map[string]any) (Fetcher, error) {
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

func (b *BirdLgGo) GetData() ([]byte, error) {
	payload, err := json.Marshal(birdLgGoPayload{
		Servers: []string{b.Server},
		Type:    b.Type,
		Args:    b.Args,
	})
	if err != nil {
		return nil, fmt.Errorf("fail to marshal payload:%v", err)
	}

	resp, err := http.Post(b.API, "application/json", bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("fail to post request:%v", err)
	}

	rbody, err := io.ReadAll(resp.Body)
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
