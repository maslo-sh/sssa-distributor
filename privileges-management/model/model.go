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
	Username        string    `json:"username"`
	ResourceID      uint      `json:"resourceId"`
	ValidityInHours int       `json:"validityInHours,omitempty"`
	expiryTimestamp time.Time `json:"expiryTimestamp,omitempty"`
	status          string    `json:"status"`
	reasoning       string    `json:"justification"`
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
