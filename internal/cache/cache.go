package cache

import (
	"context"
	"fmt"

	"github.com/go-faster/errors"

	"service-dialog/internal/models"
)

type DialogsCache struct {
	data    map[string]string
	storage storage
}

func Must(ctx context.Context, storage storage) *DialogsCache {
	c := &DialogsCache{
		storage: storage,
		data:    make(map[string]string),
	}
	if err := c.init(ctx); err != nil {
		panic(errors.Wrap(err, "init cache error"))
	}
	return c
}

func formKey(userID1 string, userID2 string) string {
	return fmt.Sprintf("%s:%s", userID1, userID2)
}

func (d *DialogsCache) init(ctx context.Context) error {
	dialogs, err := d.storage.GetAllDialogs(ctx)
	if err != nil {
		return errors.Wrap(err, "get all dialogs error")
	}

	for _, dialog := range dialogs {
		d.Set(dialog)
	}

	return nil
}

func (d *DialogsCache) Set(dialog models.Dialog) {
	d.data[formKey(dialog.UserID1, dialog.UserID2)] = dialog.ID
	d.data[formKey(dialog.UserID2, dialog.UserID1)] = dialog.ID
}

func (d *DialogsCache) Get(userID1 string, userID2 string) *models.Dialog {
	if dialogID, ok := d.data[formKey(userID1, userID2)]; ok {
		return &models.Dialog{
			ID:      dialogID,
			UserID1: userID1,
			UserID2: userID2,
		}
	}

	return nil
}
