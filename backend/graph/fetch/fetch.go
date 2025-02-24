package fetch

import "context"

type Fetcher interface {
	GetData(ctx context.Context) (any, error)
	CleanUp() error
}

var Spawn = make(map[string]func(map[string]any) (Fetcher, error))

func Register(name string, spawnFunc func(map[string]any) (Fetcher, error)) {
	Spawn[name] = spawnFunc
}
