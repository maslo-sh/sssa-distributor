package repository

import (
	"gorm.io/gorm"
	"privileges-management/model"
)

type AccessRequestsRepository interface {
	Read(uint) interface{}
	Create(interface{})
	Delete(uint)
}

type AccessRequestsRepositoryImpl struct {
	db *gorm.DB
}

func NewAccessRequestRepository(db *gorm.DB) AccessRequestsRepository {
	return &AccessRequestsRepositoryImpl{db}
}

func (ar *AccessRequestsRepositoryImpl) Read(id uint) interface{} {
	var fetchedAccessRequest *model.AccessRequest
	ar.db.First(&fetchedAccessRequest, id)
	return fetchedAccessRequest
}

func (ar *AccessRequestsRepositoryImpl) Create(accessRequest interface{}) {
	castedAccessRequest := accessRequest.(*model.AccessRequest)
	ar.db.Create(&castedAccessRequest)
}

func (ar *AccessRequestsRepositoryImpl) Delete(id uint) {
	var fetchedAccessRequest *model.AccessRequest
	ar.db.Delete(&fetchedAccessRequest, id)
}
