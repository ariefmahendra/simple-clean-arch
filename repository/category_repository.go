package repository

import (
	"context"
	"database/sql"
	model "simple-rest-api-clean-arch/model/domain"
)

type CategoryRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, category model.Category) model.Category
	Update(ctx context.Context, tx *sql.Tx, category model.Category) model.Category
	Delete(ctx context.Context, tx *sql.Tx, id int)
	FindByID(ctx context.Context, tx *sql.Tx, id int) (model.Category, error)
	FindALl(ctx context.Context, tx *sql.Tx) ([]model.Category, error)
}
