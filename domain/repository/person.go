package repository

import (
	"context"

	"my-project/domain/model"
)

type IPerson interface {
	GetByName(ctx context.Context, name string) (model.Person, error)
}
