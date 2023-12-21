package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"simple-rest-api-clean-arch/helper"
	model "simple-rest-api-clean-arch/model/domain"
)

type CategoryRepositoryImpl struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &CategoryRepositoryImpl{db: db}
}

func (repo *CategoryRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, category model.Category) model.Category {
	var SQL = "insert into category(name) values($1) returning id"
	var id int
	err := tx.QueryRowContext(ctx, SQL, category.Name).Scan(&id)
	helper.PanicIfError(err)

	category.Id = id
	return category
}

func (repo *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category model.Category) model.Category {
	SQL := "update category set name = $1 where id = $2"

	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repo *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) {
	SQL := "delete from category where id = $1"

	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfError(err)
}

func (repo *CategoryRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, id int) (model.Category, error) {
	SQL := "select id, name from category where id = $1"

	category := model.Category{}
	err := tx.QueryRowContext(ctx, SQL, id).Scan(&category.Id, &category.Name)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Category{}, errors.New("category not found")
	}

	return category, nil
}

func (repo *CategoryRepositoryImpl) FindALl(ctx context.Context, tx *sql.Tx) ([]model.Category, error) {
	SQL := "select id, name from category"

	rowsContext, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var categoryList []model.Category

	for rowsContext.Next() {
		category := model.Category{}
		err := rowsContext.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categoryList = append(categoryList, category)
	}

	if len(categoryList) == 0 {
		return categoryList, errors.New("category not found")
	}

	fmt.Println(categoryList)

	return categoryList, nil
}
