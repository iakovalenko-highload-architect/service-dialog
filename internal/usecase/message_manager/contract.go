package message_manager

import (
	"context"

	"service-dialog/internal/models"
)

type storage interface {
	Insert(ctx context.Context, message models.Message) (*models.Message, error)
	FinByUserIDs(ctx context.Context, fromID string, toID string) ([]models.Message, error)
}
