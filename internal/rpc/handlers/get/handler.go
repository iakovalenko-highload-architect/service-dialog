package get

import (
	"context"
	"errors"
	"fmt"

	"github.com/labstack/gommon/log"
	"google.golang.org/grpc/metadata"

	dto "service-dialog/internal/generated"
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

func (h *Handler) Handle(ctx context.Context, req *dto.GetRequest) (*dto.GetResponse, error) {
	var requestID string
	meta, _ := metadata.FromIncomingContext(ctx)
	if vals := meta.Get("request_id"); len(vals) == 1 {
		requestID = vals[0]
	}

	if !validator.IsValidUUID(req.FromUserID) || !validator.IsValidUUID(req.ToUserID) {
		return nil, errors.New("not valid uuid")
	}

	res, err := h.messageManager.Get(ctx, req.FromUserID, req.ToUserID)
	if err != nil {
		log.Error(err, fmt.Sprintf("X-Request-Id: %s", requestID))
		return nil, err
	}

	messages := make([]*dto.Message, 0, len(res))
	for _, msg := range res {
		messages = append(messages, &dto.Message{
			FromUserID: msg.FromID,
			ToUserID:   msg.ToID,
			Text:       msg.Text,
		})
	}
	return &dto.GetResponse{Messages: messages}, nil
}
