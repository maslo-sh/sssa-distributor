package repository

import (
	"gorm.io/gorm"
	"privileges-management/model"
)

type ApprovingPermissionsRepository interface {
	Read(uint) *model.ApprovingPermission
	Create(*model.ApprovingPermission)
	Delete(uint)
	ReadByResourceId(id uint) []model.ApprovingPermission
}

type ApprovingPermissionsRepositoryImpl struct {
	db *gorm.DB
}

func NewApprovingPermissionsRepository(db *gorm.DB) ApprovingPermissionsRepository {
	return &ApprovingPermissionsRepositoryImpl{db}
}

func (apr *ApprovingPermissionsRepositoryImpl) Read(id uint) *model.ApprovingPermission {
	var fetchedApprovingPermission model.ApprovingPermission
	apr.db.First(&fetchedApprovingPermission, id)
	return &fetchedApprovingPermission
}

func (apr *ApprovingPermissionsRepositoryImpl) Create(approvingPermission *model.ApprovingPermission) {
	apr.db.Create(approvingPermission)
}

func (apr *ApprovingPermissionsRepositoryImpl) Delete(id uint) {
	var fetchedApprovingPermission model.ApprovingPermission
	apr.db.Delete(&fetchedApprovingPermission, id)
}

func (apr *ApprovingPermissionsRepositoryImpl) ReadByResourceId(id uint) []model.ApprovingPermission {
	var permissions []model.ApprovingPermission
	apr.db.Find(&permissions, "resource_id = ?", id)
	return permissions
}
