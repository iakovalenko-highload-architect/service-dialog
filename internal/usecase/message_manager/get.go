package message_manager

import (
	"context"

	"service-dialog/internal/models"
)

func (m *MessageManager) Get(ctx context.Context, fromID string, toID string) ([]models.Message, error) {
	return m.storage.FinByUserIDs(ctx, fromID, toID)
}
