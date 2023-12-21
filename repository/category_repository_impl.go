package repository

import (
	"context"
	"database/sql"
	"errors"
	"simple-rest-api-clean-arch/helper"
	model "simple-rest-api-clean-arch/model/domain"
)

type CategoryRepositoryImpl struct {
	db *sql.DB
}

func (repo *CategoryRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, category model.Category) model.Category {
	SQL := "insert into category(id, name) values(?, ?)"

	result, err := tx.ExecContext(ctx, SQL, category.Id, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repo *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category model.Category) model.Category {
	SQL := "update category set name = ? where id = ?"

	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repo *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) {
	SQL := "delete from category where id = ?"

	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfError(err)
}

func (repo *CategoryRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, id int) (model.Category, error) {
	SQL := "select id, name from category where id = ?"

	category := model.Category{}
	err := tx.QueryRowContext(ctx, SQL, id).Scan(&category.Id, &category.Name)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Category{}, errors.New("category not found")
	} else {
		panic(err)
	}

	return category, nil
}

func (repo *CategoryRepositoryImpl) FindALl(ctx context.Context, tx *sql.Tx) ([]model.Category, error) {
	SQL := "select id, name from category"

	rowsContext, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var categoryList []model.Category
	category := model.Category{}
	if rowsContext.Next() {
		err := rowsContext.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categoryList = append(categoryList, category)
	} else {
		return categoryList, errors.New("category not found")
	}

	return categoryList, nil

}
