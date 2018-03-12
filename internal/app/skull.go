package api

import (
	"context"
	u "net/url"
)

// Skull will be used by skulleton:skullernator to generate code.
// The tool determind this struct by the "skulleton:skullernator" keyword.
type Skull struct {
}

type FooResponse struct {
}

func (s *Skull) Foo(ctx context.Context, req *u.URL) (*FooResponse, error) {
	return nil, nil
}
