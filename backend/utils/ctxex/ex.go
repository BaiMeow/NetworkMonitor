package ctxex

import (
	"context"
	"io"
)

func Warp(ctx context.Context, f func() error) error {
	err := make(chan error)
	go func() {
		err <- f()
	}()
	select {
	case <-err:
		return nil
	case <-ctx.Done():
		return context.Cause(ctx)
	}
}

func CheckDone(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func IoReadAll(ctx context.Context, r io.Reader) ([]byte, error) {
	b := make([]byte, 0, 512)
	err := Warp(ctx, func() error {
		for {
			n, err := r.Read(b[len(b):cap(b)])
			if CheckDone(ctx) {
				return context.Cause(ctx)
			}
			b = b[:len(b)+n]
			if err != nil {
				if err == io.EOF {
					err = nil
				}
				return err
			}

			if len(b) == cap(b) {
				// Add more capacity (let append pick how much).
				b = append(b, 0)[:len(b)]
			}
		}
	})
	if err != nil {
		return nil, err
	}
	return b, nil
}
