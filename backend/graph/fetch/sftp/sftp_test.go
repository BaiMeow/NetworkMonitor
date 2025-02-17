package sftp

import (
	"encoding/base64"
	"github.com/BaiMeow/NetworkMonitor/graph/fetch"
	"golang.org/x/crypto/ssh"
	"net"
	"testing"
)

func TestSftpWithPassword_GetData(t *testing.T) {
	fetcher, err := fetch.Spawn["sftp"](map[string]any{
		"host":     "172.16.4.6",
		"port":     22,
		"user":     "ubuntu",
		"password": "111111",
		"filepath": "/var/log/bird/{{- (Now.Add (Second -60)).Format \"01-02-2006-15-04\"}}.mrt",
	})
	if err != nil {
		t.Error(err)
		return
	}
	data, err := fetcher.GetData()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(data)
}

func TestReadPublicKey(t *testing.T) {
	_, err := ssh.Dial("tcp", "172.16.4.5:22", &ssh.ClientConfig{
		User: "root",
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			t.Log(base64.StdEncoding.EncodeToString(key.Marshal()))
			return nil
		},
	})
	if err != nil {
		t.Error(err)
	}
}
