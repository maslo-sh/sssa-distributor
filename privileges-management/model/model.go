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
	Username        string     `json:"username"`
	ValidityInHours int        `json:"validityInHours,omitempty"`
	GivenApproves   int        `json:"givenApproves,omitempty"`
	ExpiryTimestamp *time.Time `json:"expiryTimestamp,omitempty"`
	Status          string     `json:"status"`
	Reasoning       string     `json:"justification,omitempty"`
	ResourceID      uint       `json:"-"`
	Resource        Resource   `json:"resource"`
}

type Resource struct {
	gorm.Model
	SharesCreated     int             `json:"sharesCreated"`
	MinSharesRequired int             `json:"minSharesRequired"`
	ResourceDN        string          `json:"resourceDN"`
	AccessRequests    []AccessRequest `gorm:"foreignKey:ResourceID" json:"-"`
}

type ApprovingPermission struct {
	gorm.Model
	ResourceID uint     `json:"resourceId"`
	Resource   Resource `json:"-"`
	Username   string   `json:"username"`
}
