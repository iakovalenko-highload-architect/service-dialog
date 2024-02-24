package server

import (
	"context"

	dto "service-dialog/internal/generated"
)

type ServiceDialogServer struct {
	createHandler createHandler
	getHandler    getHandler
}

func New(
	createHandler createHandler,
	getHandler getHandler,
) *ServiceDialogServer {
	return &ServiceDialogServer{
		createHandler: createHandler,
		getHandler:    getHandler,
	}
}

func (s *ServiceDialogServer) Create(ctx context.Context, req *dto.CreateRequest) (*dto.CreateResponse, error) {
	if req == nil {
		return nil, errEmptyRequest
	}
	return s.createHandler.Handle(ctx, req)
}

func (s *ServiceDialogServer) Get(ctx context.Context, req *dto.GetRequest) (*dto.GetResponse, error) {
	if req == nil {
		return nil, errEmptyRequest
	}
	return s.getHandler.Handle(ctx, req)
}
