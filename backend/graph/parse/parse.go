package parse

import (
	"context"
	"github.com/BaiMeow/NetworkMonitor/graph/entity"
)

type Parser[T entity.DrawType] interface {
	Parse(ctx context.Context, input any) (T, error)
	CleanUp() error
}

type ParserSpawner[T entity.DrawType] = func(map[string]any) (Parser[T], error)

var registry = make(map[string]any)

func Register[T entity.DrawType](name string, spawnFunc ParserSpawner[T]) {
	registry[name] = spawnFunc
}

func GetSpawner[T entity.DrawType](name string) ParserSpawner[T] {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	return registry[name].(ParserSpawner[T])
}
