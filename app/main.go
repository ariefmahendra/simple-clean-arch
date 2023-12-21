package main

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"simple-rest-api-clean-arch/config"
	"simple-rest-api-clean-arch/controller"
	"simple-rest-api-clean-arch/helper"
	"simple-rest-api-clean-arch/repository"
	"simple-rest-api-clean-arch/router"
	"simple-rest-api-clean-arch/service"
	"time"
)

func main() {

	db := config.DatabaseConnection()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	routes := router.NewRouter(categoryController)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
