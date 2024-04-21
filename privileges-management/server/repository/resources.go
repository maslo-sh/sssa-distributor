package repository

import (
	"gorm.io/gorm"
	"privileges-management/model"
)

type ResourcesRepository interface {
	Read(uint) interface{}
	Create(interface{})
	Delete(uint)
}

type ResourcesRepositoryImpl struct {
	db *gorm.DB
}

func NewResourcesRepository(db *gorm.DB) ResourcesRepository {
	return &ResourcesRepositoryImpl{db}
}

func (ar *ResourcesRepositoryImpl) Read(id uint) interface{} {
	var fetchedResource *model.Resource
	ar.db.First(&fetchedResource, id)
	return fetchedResource
}

func (ar *ResourcesRepositoryImpl) Create(resource interface{}) {
	castedResource := resource.(*model.Resource)
	ar.db.Create(&castedResource)
}

func (ar *ResourcesRepositoryImpl) Delete(id uint) {
	var fetchedResource *model.Resource
	ar.db.Delete(&fetchedResource, id)
}
