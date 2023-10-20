package repositories

import (
	"errors"
	"inoutgo/data/request"
	"inoutgo/helpers"
	"inoutgo/models"

	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

// Delete implements TagsRepository.
func (t *TagsRepositoryImpl) Delete(tagsId int) {
	var tags models.Tags

	result := t.Db.Where("id = ?", tagsId).Delete(&tags)

	helpers.ErrorPanic(result.Error)
}

// FindAll implements TagsRepository.
func (t *TagsRepositoryImpl) FindAll() []models.Tags {
	var tags []models.Tags

	result := t.Db.Find(&tags)

	helpers.ErrorPanic(result.Error)

	return tags
}

// FindById implements TagsRepository.
func (t *TagsRepositoryImpl) FindById(tagsId int) (tags models.Tags, err error) {
	var tag models.Tags

	result := t.Db.Find(&tags, tagsId)

	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("Tag not found")
	}
}

// Update implements TagsRepository.
func (t *TagsRepositoryImpl) Update(tag models.Tags) {
	var updateTag = request.UpdateTagsRequest{Id: tag.Id, Name: tag.Name}

	result := t.Db.Model(&tag).Updates(updateTag)

	helpers.ErrorPanic(result.Error)
}

// save implements TagsRepository.
func (t *TagsRepositoryImpl) Save(tags models.Tags) {
	result := t.Db.Create(&tags)

	helpers.ErrorPanic(result.Error)
}

func NewTagsRepositoryImpl(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}
