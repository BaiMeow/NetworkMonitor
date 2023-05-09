package fetch

import (
	"fmt"
	"os/exec"
)

func init() {
	Register("command", func(m map[string]any) (Fetcher, error) {
		cmd, ok := m["cmd"].(string)
		if !ok {
			return nil, fmt.Errorf("cmd is not string")
		}
		return &Command{Command: cmd}, nil
	})
}

type Command struct {
	Command string
}

func (c *Command) GetData() (string, error) {
	output, err := exec.Command(c.Command).CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
