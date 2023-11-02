package fetch

import (
	"encoding/base64"
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/template"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

func init() {
	Register("sftp", func(config map[string]any) (Fetcher, error) {
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

		fetcher := &SftpWithPassword{
			Host:     host,
			Port:     port,
			User:     user,
			Password: password,
			filepath: filepathTpl.ExecuteString,
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

type SftpWithPassword struct {
	Host      string
	Port      int
	User      string
	Password  string
	PublicKey ssh.PublicKey
	// filepath use template to parse filepath, it's for reading rotate log file (like .mrt)
	filepath func() (string, error)
}

func (s *SftpWithPassword) GetData() ([]byte, error) {
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

	c, err := ssh.Dial("tcp", net.JoinHostPort(s.Host, strconv.Itoa(s.Port)), cfg)
	if err != nil {
		return nil, fmt.Errorf("fail to dial ssh: %v", err)
	}
	defer c.Close()

	sftpc, err := sftp.NewClient(c)
	if err != nil {
		return nil, fmt.Errorf("fail to dial sftp: %v", err)
	}
	defer sftpc.Close()

	fp, err := s.filepath()
	if err != nil {
		return nil, fmt.Errorf("fail to get filepath: %v", err)
	}

	file, err := sftpc.Open(fp)
	if err != nil {
		return nil, fmt.Errorf("fail to open file from sftp: %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("fail to read file from sftp: %v", err)
	}

	return content, nil
}
