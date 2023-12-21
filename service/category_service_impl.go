package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"simple-rest-api-clean-arch/helper"
	"simple-rest-api-clean-arch/model/domain"
	"simple-rest-api-clean-arch/model/dto"
	"simple-rest-api-clean-arch/repository"
)

type CategoryServiceImpl struct {
	categoryRepository repository.CategoryRepository
	DB                 *sql.DB
	validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{categoryRepository: categoryRepository, DB: DB, validate: validate}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request dto.CategoryCreateRequest) dto.CategoryResponse {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{Name: request.Name}

	category = service.categoryRepository.Insert(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request dto.CategoryUpdateRequest) dto.CategoryResponse {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.categoryRepository.FindByID(ctx, tx, request.Id)
	helper.PanicIfError(err)

	category.Name = request.Name

	category = service.categoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.categoryRepository.FindByID(ctx, tx, id)
	helper.PanicIfError(err)

	category = service.categoryRepository.Update(ctx, tx, category)

}

func (service *CategoryServiceImpl) FindByID(ctx context.Context, id int) dto.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.categoryRepository.FindByID(ctx, tx, id)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []dto.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories, err := service.categoryRepository.FindALl(ctx, tx)
	helper.PanicIfError(err)

	return helper.ToCategoryResponses(categories)
}
