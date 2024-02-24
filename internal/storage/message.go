package storage

import (
	"context"

	"github.com/AlekSi/pointer"
	"github.com/go-faster/errors"
	"github.com/jmoiron/sqlx"

	"service-dialog/internal/models"
	dto "service-dialog/internal/storage/models/message"
)

func (p *Postgres) Insert(ctx context.Context, message models.Message) (*models.Message, error) {
	query := `
		insert into messages (from_id, to_id, text_)
		values(:from_id, :to_id, :text_)
		returning id, from_id, to_id, text_
	`

	rows, err := p.conn.NamedQueryContext(ctx, query, dto.Imported(message))
	if err != nil {
		return nil, errors.Wrap(err, "insert message error")
	}

	var inserted dto.Message
	for rows.Next() {
		if err := rows.StructScan(&inserted); err != nil {
			return nil, errors.Wrap(err, "scan message insert result error")
		}
	}

	return pointer.To(dto.Exported(inserted)), nil
}

func (p *Postgres) FinByUserIDs(ctx context.Context, fromID string, toID string) ([]models.Message, error) {
	query := `
		select from_id, to_id, text_
		from messages
		where (from_id = $1 and to_id = $2) or (from_id = $2 and to_id = $1)
		order by id
	`

	query, args, err := sqlx.In(query, fromID, toID)
	if err != nil {
		return nil, errors.Wrap(err, "create message select error")
	}
	dbQuery := p.conn.Rebind(query)

	var messages []dto.Message
	err = p.conn.SelectContext(ctx, &messages, dbQuery, args...)
	if err != nil {
		return nil, errors.Wrap(err, "message select error")
	}
	if len(messages) == 0 {
		return nil, nil
	}

	res := make([]models.Message, 0, len(messages))
	for _, message := range messages {
		res = append(res, dto.Exported(message))
	}
	return res, nil
}
