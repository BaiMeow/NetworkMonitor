package fetch

type Fetcher interface {
	GetData() ([]byte, error)
}

var Spawn = make(map[string]func(map[string]any) (Fetcher, error))

func Register(name string, spawnFunc func(map[string]any) (Fetcher, error)) {
	Spawn[name] = spawnFunc
}
