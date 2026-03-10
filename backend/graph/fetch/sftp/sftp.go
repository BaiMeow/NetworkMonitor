package sftp

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"

	"github.com/BaiMeow/NetworkMonitor/conf"
	"github.com/BaiMeow/NetworkMonitor/graph/fetch"
	"github.com/BaiMeow/NetworkMonitor/template"
	"github.com/BaiMeow/NetworkMonitor/trace"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func init() {
	fetch.Register("sftp", func(config map[string]any) (fetch.Fetcher, error) {
		host, ok := config["host"].(string)
		if !ok {
			return nil, fmt.Errorf("host is not string")
		}
		port, ok := config["port"].(int)
		if !ok {
			return nil, fmt.Errorf("port is not int")
		}
		user, ok := config["user"].(string)
		if !ok {
			return nil, fmt.Errorf("user is not string")
		}
		password, ok := config["password"].(string)
		if !ok {
			return nil, fmt.Errorf("password is not string")
		}
		filepathTplstr, ok := config["filepath"].(string)
		if !ok {
			return nil, fmt.Errorf("filepath is not string")
		}
		filepathTpl, err := template.Parse(filepathTplstr)
		if err != nil {
			return nil, fmt.Errorf("fail to parse filepath as template:%v", err)
		}

		fetcher := &WithPassword{
			Host:     host,
			Port:     port,
			User:     user,
			Password: password,
			filepath: filepathTpl,
		}

		b64pubkey, ok := config["public-key"].(string)
		if !ok {
			log.Println("missing pubkey field in sftp, may not be safe")
			return fetcher, nil
		}

		pubkeyBytes, err := base64.StdEncoding.DecodeString(b64pubkey)
		if err != nil {
			return nil, fmt.Errorf("fail to decode pubkey: %v", err)
		}

		fetcher.PublicKey, err = ssh.ParsePublicKey(pubkeyBytes)
		if err != nil {
			return nil, fmt.Errorf("fail to parse pubkey: %v", err)
		}

		return fetcher, nil
	})
}

type WithPassword struct {
	fetch.Base
	Host      string
	Port      int
	User      string
	Password  string
	PublicKey ssh.PublicKey
	// filepath use template to parse filepath, it's for reading rotate log file (like .mrt)
	filepath *template.Template
}

func (s *WithPassword) GetData(ctx context.Context) (any, error) {
	ctx, span := trace.Tracer.Start(ctx,
		"fetch/sftp/WithPassword.GetData",
	)
	defer span.End()

	cfg := &ssh.ClientConfig{
		User: s.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(s.Password),
		},
		Timeout: conf.ProbeTimeout * 4 / 5,
	}
	if s.PublicKey != nil {
		cfg.HostKeyCallback = ssh.FixedHostKey(s.PublicKey)
	} else {
		cfg.HostKeyCallback = ssh.InsecureIgnoreHostKey()
	}

	var content []byte

	c, err := ssh.Dial("tcp", net.JoinHostPort(s.Host, strconv.Itoa(s.Port)), cfg)
	if err != nil {
		return nil, fmt.Errorf("fail to dial ssh: %v", err)
	}
	defer c.Close()
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	sftpc, err := sftp.NewClient(c)
	if err != nil {
		return nil, fmt.Errorf("fail to dial sftp: %v", err)
	}
	defer sftpc.Close()
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	fp, err := s.filepath.ExecuteString()
	if err != nil {
		return nil, fmt.Errorf("fail to get filepath: %v", err)
	}
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	file, err := sftpc.Open(fp)
	if err != nil {
		return nil, fmt.Errorf("fail to open file from sftp: %v", err)
	}
	defer file.Close()
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	content, err = io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("fail to read file from sftp: %v", err)
	}

	return content, nil
}
