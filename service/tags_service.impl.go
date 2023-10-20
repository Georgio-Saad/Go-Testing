package service

import (
	"inoutgo/data/request"
	"inoutgo/data/response"
	"inoutgo/helpers"
	"inoutgo/models"
	"inoutgo/repositories"

	"github.com/go-playground/validator/v10"
)

type TagsServiceImpl struct {
	TagsRepository repositories.TagsRepository
	validate       *validator.Validate
}

// Create implements TagsService.
func (t *TagsServiceImpl) Create(tags request.CreateTagsRequest) {
	err := t.validate.Struct(tags)
	helpers.ErrorPanic(err)
	tagModel := models.Tags{
		Name: tags.Name,
	}

	t.TagsRepository.Save(tagModel)
}

// Delete implements TagsService.
func (t *TagsServiceImpl) Delete(tagsId int) {
	t.TagsRepository.Delete(tagsId)
}

// FindAll implements TagsService.
func (t *TagsServiceImpl) FindAll() []response.TagsResponse {
	result := t.TagsRepository.FindAll()

	var tags []response.TagsResponse

	for _, value := range result {
		tag := response.TagsResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}

	return tags
}

// FindById implements TagsService.
func (t *TagsServiceImpl) FindById(tagsId int) response.TagsResponse {
	tagData, err := t.TagsRepository.FindById(tagsId)

	helpers.ErrorPanic(err)

	tagResponse := response.TagsResponse{
		Id:   tagData.Id,
		Name: tagData.Name,
	}

	return tagResponse
}

// Update implements TagsService.
func (t *TagsServiceImpl) Update(tags request.UpdateTagsRequest) {
	tagData, err := t.TagsRepository.FindById(tags.Id)

	helpers.ErrorPanic(err)

	tagData.Name = tags.Name
	t.TagsRepository.Update(tagData)
}

func NewTagsServiceImpl(tagRepository repositories.TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{
		TagsRepository: tagRepository,
		validate:       validate,
	}
}
