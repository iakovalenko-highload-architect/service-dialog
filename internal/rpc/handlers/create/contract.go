package create

import (
	"context"

	"service-dialog/internal/models"
)

type messageManager interface {
	Create(ctx context.Context, message models.Message) (*models.Message, error)
}
