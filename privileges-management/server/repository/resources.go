package repository

import (
	"gorm.io/gorm"
	"privileges-management/model"
)

type ResourcesRepository interface {
	ReadAll() []model.Resource
	Read(uint) *model.Resource
	Create(resource *model.Resource)
	Delete(uint)
}

type ResourcesRepositoryImpl struct {
	db *gorm.DB
}

func NewResourcesRepository(db *gorm.DB) ResourcesRepository {
	return &ResourcesRepositoryImpl{db}
}

func (ar *ResourcesRepositoryImpl) ReadAll() []model.Resource {
	var fetchedResource []model.Resource
	ar.db.Find(&fetchedResource)
	return fetchedResource
}

func (ar *ResourcesRepositoryImpl) Read(id uint) *model.Resource {
	var fetchedResource *model.Resource
	ar.db.First(&fetchedResource, id)
	return fetchedResource
}

func (ar *ResourcesRepositoryImpl) Create(resource *model.Resource) {
	ar.db.Create(&resource)
}

func (ar *ResourcesRepositoryImpl) Delete(id uint) {
	var fetchedResource *model.Resource
	ar.db.Delete(&fetchedResource, id)
}
