package parse

type Base struct {
}

func (b *Base) Stop() error {
	return nil
}

func (b *Base) ParseAndMerge(input any, drawing *Drawing) error {
	panic("implement me")
}

var _ Parser = (*Base)(nil)
