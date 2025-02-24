package ssh

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/graph/fetch"
	"github.com/BaiMeow/NetworkMonitor/utils"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"strconv"
	"time"
)

func init() {
	fetch.Register("ssh", func(config map[string]any) (fetch.Fetcher, error) {
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
		command, ok := config["command"].(string)
		if !ok {
			return nil, fmt.Errorf("command is not string")
		}
		fetcher := &SshWithPassword{
			Host:     host,
			Port:     port,
			User:     user,
			Password: password,
			Command:  command,
		}
		b64pubkey, ok := config["public-key"].(string)
		if !ok {
			log.Println("missing pubkey field in ssh, may not be safe")
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

type SshWithPassword struct {
	fetch.Base
	Host      string
	Port      int
	User      string
	Password  string
	Command   string
	PublicKey ssh.PublicKey
}

func (s *SshWithPassword) GetData(ctx context.Context) (any, error) {
	cfg := &ssh.ClientConfig{
		User: s.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(s.Password),
		},
		Timeout: 30 * time.Second,
	}
	if s.PublicKey != nil {
		cfg.HostKeyCallback = ssh.FixedHostKey(s.PublicKey)
	} else {
		cfg.HostKeyCallback = ssh.InsecureIgnoreHostKey()
	}

	var output []byte

	if err := utils.CtxWarp(ctx, func() error {
		c, err := ssh.Dial("tcp", net.JoinHostPort(s.Host, strconv.Itoa(s.Port)), cfg)
		if err != nil {
			return fmt.Errorf("fail to dial ssh: %v", err)
		}
		defer c.Close()

		utils.CtxCheckDone(ctx)

		ss, err := c.NewSession()
		if err != nil {
			return fmt.Errorf("fail to dial ssh: %v", err)
		}
		defer ss.Close()
		utils.CtxCheckDone(ctx)

		output, err = ss.CombinedOutput(s.Command)
		if err != nil {
			return fmt.Errorf("fail to run command: %v", err)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return output, nil
}
