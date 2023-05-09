package fetch

import (
	"fmt"
	"github.com/BaiMeow/script-tools/ssh"
	"net"
	"strconv"
)

func init() {
	Register("ssh", func(config map[string]any) (Fetcher, error) {
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
		return &SshWithPassword{
			Host:     host,
			Port:     port,
			User:     user,
			Password: password,
			Command:  command,
		}, nil
	})
}

type SshWithPassword struct {
	Host     string
	Port     int
	User     string
	Password string
	Command  string
}

func (s *SshWithPassword) GetData() (string, error) {
	session, err := ssh.ConnectWithPassword(net.JoinHostPort(s.Host, strconv.Itoa(s.Port)), s.User, s.Password)
	if err != nil {
		return "", err
	}
	output, err := session.CombinedOutput(s.Command)
	if err != nil {
		return "", err
	}
	return string(output), nil
}
