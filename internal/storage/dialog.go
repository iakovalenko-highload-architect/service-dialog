package storage

import (
	"context"

	"github.com/AlekSi/pointer"
	"github.com/go-faster/errors"
	"github.com/jmoiron/sqlx"

	"service-dialog/internal/models"
	dto "service-dialog/internal/storage/models/dialog"
)

func (p *Postgres) FinDialogByUserIDs(ctx context.Context, userID1 string, userID2 string) (*models.Dialog, error) {
	query := `
		select id
		from dialogs
		where (user_id_1 = $1 and user_id_2 = $2) or (user_id_1 = $2 and user_id_2 = $1)
		order by id
	`

	query, args, err := sqlx.In(query, userID1, userID2)
	if err != nil {
		return nil, errors.Wrap(err, "create dialogs select error")
	}
	dbQuery := p.conn.Rebind(query)

	var dialogs []dto.Dialog
	err = p.conn.SelectContext(ctx, &dialogs, dbQuery, args...)
	if err != nil {
		return nil, errors.Wrap(err, "dialog select error")
	}
	if len(dialogs) == 0 {
		return nil, nil
	}

	return pointer.To(dto.Exported(dialogs[0])), nil
}

func (p *Postgres) GetAllDialogs(ctx context.Context) ([]models.Dialog, error) {
	query := `
		select id, user_id_1, user_id_2
		from dialogs
	`

	dbQuery := p.conn.Rebind(query)

	var dialogs []dto.Dialog
	err := p.conn.SelectContext(ctx, &dialogs, dbQuery)
	if err != nil {
		return nil, errors.Wrap(err, "dialog select error")
	}
	if len(dialogs) == 0 {
		return nil, nil
	}

	res := make([]models.Dialog, 0, len(dialogs))
	for _, dialog := range dialogs {
		res = append(res, models.Dialog{
			ID:      dialog.ID,
			UserID1: dialog.UserID1,
			UserID2: dialog.UserID2,
		})
	}
	return res, nil
}

func (p *Postgres) InsertDialog(ctx context.Context, dialog models.Dialog) (*models.Dialog, error) {
	query := `
		insert into dialogs (user_id_1, user_id_2)
		values(:user_id_1, :user_id_2)
		returning id, user_id_1, user_id_2
	`

	rows, err := p.conn.NamedQueryContext(ctx, query, dto.Imported(dialog))
	if err != nil {
		return nil, errors.Wrap(err, "insert dialog error")
	}

	var inserted dto.Dialog
	for rows.Next() {
		if err := rows.StructScan(&inserted); err != nil {
			return nil, errors.Wrap(err, "scan dialog insert result error")
		}
	}

	return pointer.To(dto.Exported(inserted)), nil
}
