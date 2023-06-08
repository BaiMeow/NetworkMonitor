package fetch

import (
	"bytes"
	"fmt"
	"github.com/BaiMeow/OSPF-monitor/template"
	"github.com/BaiMeow/script-tools/ssh"
	"net"
	"strconv"
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
		return &SftpWithPassword{
			Host:     host,
			Port:     port,
			User:     user,
			Password: password,
			filepath: filepathTpl.ExecuteString,
		}, nil
	})
}

type SftpWithPassword struct {
	Host     string
	Port     int
	User     string
	Password string
	// filepath use template to parse filepath, it's for reading rotate log file (like .mrt)
	filepath func() (string, error)
}

func (s *SftpWithPassword) GetData() ([]byte, error) {
	c, err := ssh.ConnectWithPassword(net.JoinHostPort(s.Host, strconv.Itoa(s.Port)), s.User, s.Password)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	sftp, err := ssh.NewSFTP(c)
	defer sftp.Close()
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	fp, err := s.filepath()
	if err != nil {
		return nil, err
	}
	if err = sftp.Download(fp, &buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
