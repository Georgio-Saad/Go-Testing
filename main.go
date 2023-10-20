package main

import (
	"inoutgo/config"
	"inoutgo/controllers"
	"inoutgo/helpers"
	"inoutgo/models"
	"inoutgo/repositories"
	"inoutgo/router"
	"inoutgo/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&models.Tags{})

	// Repositories
	tagsRepositories := repositories.NewTagsRepositoryImpl(db)

	// Services
	tagsService := service.NewTagsServiceImpl(tagsRepositories, validate)

	// Controller
	tagsController := controllers.NewTagsController(tagsService)

	// Router
	routes := router.NewRouter(tagsController)

	routes.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome")
	})

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()

	helpers.ErrorPanic(err)

}
