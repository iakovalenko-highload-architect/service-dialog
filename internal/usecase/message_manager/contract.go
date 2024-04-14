package message_manager

import (
	"context"

	"service-dialog/internal/models"
)

type cache interface {
	Set(dialog models.Dialog)
	Get(userID1 string, userID2 string) *models.Dialog
}

type storage interface {
	Insert(ctx context.Context, message models.Message) (*models.Message, error)
	FinMessagesByDialogID(ctx context.Context, dialogID string) ([]models.Message, error)
	InsertDialog(ctx context.Context, dialog models.Dialog) (*models.Dialog, error)
}
