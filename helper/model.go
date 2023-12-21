package helper

import (
	"simple-rest-api-clean-arch/model/domain"
	"simple-rest-api-clean-arch/model/dto"
)

func ToCategoryResponse(category domain.Category) dto.CategoryResponse {
	return dto.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []dto.CategoryResponse {
	var categoriesResponses []dto.CategoryResponse
	for _, item := range categories {
		categoriesResponses = append(categoriesResponses, ToCategoryResponse(item))
	}

	return categoriesResponses
}
