package message_manager

import (
	"context"

	"service-dialog/internal/models"
)

func (m *MessageManager) Create(ctx context.Context, message models.Message) (*models.Message, error) {
	return m.storage.Insert(ctx, message)
}
