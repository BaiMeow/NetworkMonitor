package fetch

import (
	"fmt"
	"os/exec"
	"strings"
)

func init() {
	Register("cmd", func(m map[string]any) (Fetcher, error) {
		cmd, ok := m["command"].(string)
		if !ok {
			return nil, fmt.Errorf("cmd is not string")
		}
		return &Command{Command: cmd}, nil
	})
}

var _ Fetcher = (*Command)(nil)

type Command struct {
	Command string
}

func (c *Command) GetData() ([]byte, error) {
	ss := strings.SplitN(c.Command, " ", 2)
	var cmd, arg string
	switch len(ss) {
	case 0:
		return nil, fmt.Errorf("command is empty")
	case 1:
		cmd = ss[0]
	case 2:
		cmd, arg = ss[0], ss[1]
	}
	output, err := exec.Command(cmd, arg).Output()
	if err != nil {
		return nil, err
	}
	return output, nil
}
