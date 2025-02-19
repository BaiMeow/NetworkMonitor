package fetch

type Base struct {
}

func (b Base) GetData() (any, error) {
	panic("implement me")
}

func (b Base) Stop() error {
	return nil
}

var _ Fetcher = (*Base)(nil)
