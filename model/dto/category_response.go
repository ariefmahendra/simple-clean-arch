package dto

type CategoryResponse struct {
	Id   int    `validate:"required"`
	Name string `validate:"required"`
}
