package api

import (
	"context"
)

type BarRequest struct {
	BarField int
}

func (s *Skull) Bar(ctx context.Context, req *BarRequest) (*FooResponse, error) {
	return nil, nil
}
