package model

import (
	"gorm.io/gorm"
	"time"
)

type Credentials struct {
	Username, Password string
}

type AccessRequest struct {
	ID              string    `gorm:"primaryKey"`
	Username        string    `json:"username"`
	ResourceID      uint      `json:"resourceId"`
	ValidityInHours int       `json:"validityInHours,omitempty"`
	GivenApproves   int       `json:"givenApproves,omitempty"`
	ExpiryTimestamp time.Time `json:"expiryTimestamp,omitempty"`
	Status          string    `json:"status"`
	Reasoning       string    `json:"justification,omitempty"`
}

type Resource struct {
	gorm.Model
	SharesCreated     int    `json:"sharesCreated"`
	MinSharesRequired int    `json:"minSharesRequired"`
	ResourceDN        string `json:"resourceDN"`
}

type ApprovingPermission struct {
	gorm.Model
	ResourceID uint   `json:"resourceId"`
	Username   string `json:"username"`
}
