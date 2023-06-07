package parse

type Parser interface {
	Init(input []byte)
	Parse() (Graph, error)
}

var Spawn = make(map[string]func(map[string]any) (Parser, error))

func Register(name string, spawnFunc func(map[string]any) (Parser, error)) {
	Spawn[name] = spawnFunc
}
