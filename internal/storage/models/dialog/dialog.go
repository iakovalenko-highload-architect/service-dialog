package dialog

import (
	"service-dialog/internal/models"
)

type Dialog struct {
	ID      string `db:"id"`
	UserID1 string `db:"user_id_1"`
	UserID2 string `db:"user_id_2"`
}

func Exported(internal Dialog) models.Dialog {
	return models.Dialog{
		ID:      internal.ID,
		UserID1: internal.UserID1,
		UserID2: internal.UserID2,
	}
}

func Imported(external models.Dialog) Dialog {
	return Dialog{
		ID:      external.ID,
		UserID1: external.UserID1,
		UserID2: external.UserID2,
	}
}
