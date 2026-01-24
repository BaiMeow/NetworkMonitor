package analysis

type IdReleaser int

func (r *IdReleaser) Next() int {
	*r++
	return int(*r - 1)
}
