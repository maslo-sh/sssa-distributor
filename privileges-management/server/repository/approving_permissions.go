package repository

import (
	"gorm.io/gorm"
	"privileges-management/model"
)

type ApprovingPermissionsRepository interface {
	Read(uint) interface{}
	Create(interface{})
	Delete(uint)
	ReadByResourceId(id int) []model.ApprovingPermission
}

type ApprovingPermissionsRepositoryImpl struct {
	db *gorm.DB
}

func NewApprovingPermissionsRepository(db *gorm.DB) ApprovingPermissionsRepository {
	return &ApprovingPermissionsRepositoryImpl{db}
}

func (apr *ApprovingPermissionsRepositoryImpl) Read(id uint) interface{} {
	var fetchedApprovingPermission model.ApprovingPermission
	apr.db.First(&fetchedApprovingPermission, id)
	return fetchedApprovingPermission
}

func (apr *ApprovingPermissionsRepositoryImpl) Create(approvingPermission interface{}) {
	castedApprovingPermission := approvingPermission.(model.ApprovingPermission)
	apr.db.Create(&castedApprovingPermission)
}

func (apr *ApprovingPermissionsRepositoryImpl) Delete(id uint) {
	var fetchedApprovingPermission model.ApprovingPermission
	apr.db.Delete(&fetchedApprovingPermission, id)
}

func (apr *ApprovingPermissionsRepositoryImpl) ReadByResourceId(id int) []model.ApprovingPermission {
	var permissions []model.ApprovingPermission
	apr.db.Find(&permissions, "id = ?", id)

	return permissions
}
