package service

import (
	"context"
	"simple-rest-api-clean-arch/model/dto"
)

type CategoryService interface {
	Create(ctx context.Context, request dto.CategoryCreateRequest) dto.CategoryResponse
	Update(ctx context.Context, request dto.CategoryUpdateRequest) dto.CategoryResponse
	Delete(ctx context.Context, id int)
	FindByID(ctx context.Context, id int) dto.CategoryResponse
	FindAll(ctx context.Context) []dto.CategoryResponse
}
