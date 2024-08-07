package model

import (
	"gorm.io/gorm"
	"time"
)

type Credentials struct {
	Username, Password string
}

type AccessRequest struct {
	gorm.Model
	Username        string
	ValidityInHours int        `json:"omitempty"`
	GivenApproves   int        `json:"omitempty"`
	ExpiryTimestamp *time.Time `json:"omitempty"`
	Status          string
	Reasoning       string   `json:"omitempty"`
	ResourceID      uint     `json:"-"`
	Resource        Resource `json:"resource" gorm:"foreignKey:ResourceID"`
}

type Resource struct {
	gorm.Model
	SharesCreated     int
	MinSharesRequired int
	ResourceDN        string
}

type ApprovingPermission struct {
	gorm.Model
	ResourceID uint     `json:"resourceId"`
	Resource   Resource `json:"-"`
	Username   string   `json:"username"`
}
