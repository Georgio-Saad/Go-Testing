package controllers

import (
	"inoutgo/data/request"
	"inoutgo/data/response"
	"inoutgo/helpers"
	"inoutgo/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagsController struct {
	tagsService service.TagsService
}

// Create CONTROLLER
func (controller *TagsController) Create(ctx *gin.Context) {
	createTagsRequest := request.CreateTagsRequest{}

	err := ctx.ShouldBindJSON(&createTagsRequest)

	helpers.ErrorPanic(err)

	controller.tagsService.Create(createTagsRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Update CONTROLLER
func (controller *TagsController) Update(ctx *gin.Context) {
	updateTagsRequest := request.UpdateTagsRequest{}

	err := ctx.ShouldBindJSON(&updateTagsRequest)

	helpers.ErrorPanic(err)

	tagId := ctx.Param("tag_id")

	id, err := strconv.Atoi(tagId)

	helpers.ErrorPanic(err)

	updateTagsRequest.Id = id

	controller.tagsService.Update(updateTagsRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Delete CONTROLLER
func (controller *TagsController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tag_id")

	id, err := strconv.Atoi(tagId)

	helpers.ErrorPanic(err)

	controller.tagsService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// FindById CONTROLLER
func (controller *TagsController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tag_id")

	id, err := strconv.Atoi(tagId)

	helpers.ErrorPanic(err)

	tagResponse := controller.tagsService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAll CONTROLLER
func (controller *TagsController) FindAll(ctx *gin.Context) {
	tagResponse := controller.tagsService.FindAll()

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func NewTagsController(service service.TagsService) *TagsController {
	return &TagsController{
		tagsService: service,
	}
}
