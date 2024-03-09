package message

import (
	"time"

	"service-dialog/internal/models"
)

type Message struct {
	ID        int64     `db:"id"`
	FromID    string    `db:"from_id"`
	ToID      string    `db:"to_id"`
	Text      string    `db:"text_"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func Exported(internal Message) models.Message {
	return models.Message{
		FromID: internal.FromID,
		ToID:   internal.ToID,
		Text:   internal.Text,
	}
}

func Imported(external models.Message) Message {
	return Message{
		FromID: external.FromID,
		ToID:   external.ToID,
		Text:   external.Text,
	}
}
