package model

import (
	"gorm.io/gorm"
)

type Credentials struct {
	Username, Password string
}

type AccessRequest struct {
	gorm.Model
	Username        string
	ResourceID      uint
	ValidityInHours int
}

type Resource struct {
	gorm.Model
	SharesCreated     int
	MinSharesRequired int
	ResourceDN        string
}

type ApprovingPermission struct {
	gorm.Model
	ResourceID uint
	Username   string
}
