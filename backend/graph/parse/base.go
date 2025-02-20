package parse

import (
	"github.com/BaiMeow/NetworkMonitor/graph/entity"
)

type Base[T entity.DrawType] struct {
}

func (b *Base[T]) CleanUp() error {
	return nil
}

func (b *Base[T]) Parse(input any) (T, error) {
	panic("implement me")
}

var _ Parser[*entity.BGP] = (*Base[*entity.BGP])(nil)
