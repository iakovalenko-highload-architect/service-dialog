package server

import (
	"context"

	dto "service-dialog/internal/generated"
)

type createHandler interface {
	Handle(ctx context.Context, req *dto.CreateRequest) (*dto.CreateResponse, error)
}

type getHandler interface {
	Handle(ctx context.Context, req *dto.GetRequest) (*dto.GetResponse, error)
}
