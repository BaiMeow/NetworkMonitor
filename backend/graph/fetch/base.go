package fetch

import "context"

type Base struct {
}

func (b Base) GetData(ctx context.Context) (any, error) {
	panic("implement me")
}

func (b Base) CleanUp() error {
	return nil
}

var _ Fetcher = (*Base)(nil)
