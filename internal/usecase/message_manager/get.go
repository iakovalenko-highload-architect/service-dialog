package message_manager

import (
	"context"

	"service-dialog/internal/models"
)

func (m *MessageManager) Get(ctx context.Context, fromID string, toID string) ([]models.Message, error) {
	dialog := m.cache.Get(fromID, toID)
	if dialog == nil {
		return nil, nil
	}
	return m.storage.FinMessagesByDialogID(ctx, dialog.ID)
}
