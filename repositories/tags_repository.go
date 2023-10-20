package repositories

import "inoutgo/models"

type TagsRepository interface {
	Save(tags models.Tags)
	Update(tag models.Tags)
	Delete(tagsId int)
	FindById(tagsId int) (tags models.Tags, err error)
	FindAll() []models.Tags
}
