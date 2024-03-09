package get

import (
	"context"

	"service-dialog/internal/models"
)

type messageManager interface {
	Get(ctx context.Context, fromID string, toID string) ([]models.Message, error)
}
