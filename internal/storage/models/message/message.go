package message

import (
	"time"

	"service-dialog/internal/models"
)

type Message struct {
	ID        string    `db:"id"`
	DialogID  string    `db:"dialog_id"`
	FromID    string    `db:"from_id"`
	ToID      string    `db:"to_id"`
	Text      string    `db:"text_"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func Exported(internal Message) models.Message {
	return models.Message{
		ID:       internal.ID,
		DialogID: internal.DialogID,
		FromID:   internal.FromID,
		ToID:     internal.ToID,
		Text:     internal.Text,
	}
}

func Imported(external models.Message) Message {
	return Message{
		ID:       external.ID,
		DialogID: external.DialogID,
		FromID:   external.FromID,
		ToID:     external.ToID,
		Text:     external.Text,
	}
}
