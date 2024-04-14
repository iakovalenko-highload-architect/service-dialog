package message_manager

import (
	"context"

	"github.com/go-faster/errors"

	"service-dialog/internal/models"
)

func (m *MessageManager) Create(ctx context.Context, message models.Message) (*models.Message, error) {
	var err error

	dialog := m.cache.Get(message.FromID, message.ToID)
	if dialog == nil {
		dialog, err = m.storage.InsertDialog(ctx, models.Dialog{
			UserID1: message.FromID,
			UserID2: message.ToID,
		})
		if err != nil {
			return nil, errors.Wrap(err, "insert dialog error")
		}

		m.cache.Set(*dialog)
	}

	message.DialogID = dialog.ID

	return m.storage.Insert(ctx, message)
}
