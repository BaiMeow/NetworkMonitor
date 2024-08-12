package utils

func Map[T any, V any](t []T, f func(T) V) []V {
	v := make([]V, len(t))
	for i := range t {
		v[i] = f(t[i])
	}
	return v
}
