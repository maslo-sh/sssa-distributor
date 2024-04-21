package repository

import (
	"gorm.io/gorm"
	"privileges-management/model"
)

type AccessRequestsRepository interface {
	ReadAll() []model.AccessRequest
	Read(uint) *model.AccessRequest
	Create(*model.AccessRequest)
	Update(*model.AccessRequest)
	Delete(uint)
}

type AccessRequestsRepositoryImpl struct {
	db *gorm.DB
}

func NewAccessRequestRepository(db *gorm.DB) AccessRequestsRepository {
	return &AccessRequestsRepositoryImpl{db}
}

func (ar *AccessRequestsRepositoryImpl) ReadAll() []model.AccessRequest {
	var requests []model.AccessRequest
	ar.db.Find(&requests)
	return requests
}

func (ar *AccessRequestsRepositoryImpl) Read(id uint) *model.AccessRequest {
	var fetchedAccessRequest *model.AccessRequest
	ar.db.First(&fetchedAccessRequest, id)
	return fetchedAccessRequest
}

func (ar *AccessRequestsRepositoryImpl) Create(accessRequest *model.AccessRequest) {
	ar.db.Create(&accessRequest)
}

func (ar *AccessRequestsRepositoryImpl) Update(accessRequest *model.AccessRequest) {
	ar.db.Save(&accessRequest)
}

func (ar *AccessRequestsRepositoryImpl) Delete(id uint) {
	var fetchedAccessRequest *model.AccessRequest
	ar.db.Delete(&fetchedAccessRequest, id)
}
