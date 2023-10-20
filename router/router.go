package router

import (
	"inoutgo/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(tagsController *controllers.TagsController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome Home")
	})

	baseRouter := router.Group("/api")
	tagsRouter := baseRouter.Group("/tags")
	tagsRouter.GET("", tagsController.FindAll)
	tagsRouter.GET("/:tag_id", tagsController.FindById)
	tagsRouter.POST("", tagsController.Create)
	tagsRouter.PATCH("/:tag_id", tagsController.Update)
	tagsRouter.DELETE("/:tag_id", tagsController.Delete)

	return router
}
