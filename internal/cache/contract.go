package cache

import (
	"context"

	"service-dialog/internal/models"
)

type storage interface {
	GetAllDialogs(ctx context.Context) ([]models.Dialog, error)
}
