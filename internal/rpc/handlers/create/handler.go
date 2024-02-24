package create

import (
	"context"
	"errors"
	"fmt"

	"github.com/labstack/gommon/log"
	"google.golang.org/grpc/metadata"

	dto "service-dialog/internal/generated"
	"service-dialog/internal/models"
	"service-dialog/internal/utils/validator"
)

type Handler struct {
	messageManager messageManager
}

func New(messageManager messageManager) *Handler {
	return &Handler{
		messageManager: messageManager,
	}
}

func (h *Handler) Handle(ctx context.Context, req *dto.CreateRequest) (*dto.CreateResponse, error) {
	var requestID string
	meta, _ := metadata.FromIncomingContext(ctx)
	if vals := meta.Get("request_id"); len(vals) == 1 {
		requestID = vals[0]
	}

	if !validator.IsValidUUID(req.FromUserID) || !validator.IsValidUUID(req.ToUserID) {
		return nil, errors.New("not valid uuid")
	}

	res, err := h.messageManager.Create(ctx, models.Message{
		FromID: req.FromUserID,
		ToID:   req.ToUserID,
		Text:   req.Text,
	})
	if err != nil {
		log.Error(err, fmt.Sprintf("X-Request-Id: %s", requestID))
		return nil, err
	}

	return &dto.CreateResponse{Message: &dto.Message{
		FromUserID: res.FromID,
		ToUserID:   res.ToID,
		Text:       res.Text,
	}}, nil
}
