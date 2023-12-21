package router

import (
	"github.com/gin-gonic/gin"
	"simple-rest-api-clean-arch/controller"
)

func NewRouter(categoryController controller.CategoryController) *gin.Engine {

	router := gin.Default()

	router.GET("/category", categoryController.FindAll)
	router.POST("/category", categoryController.Create)
	router.PUT("/category", categoryController.Update)
	router.GET("/category/:id", categoryController.FindByID)
	router.DELETE("/category/:id", categoryController.Delete)

	return router
}
