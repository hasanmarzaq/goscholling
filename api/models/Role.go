package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID         uint32         `gorm:"primary_key;auto_increment" json:"id"`
	Name       string         `gorm:"size:255;not null;" json:"name"`
	Alias      string         `gorm:"size:255;not null;" json:"alias"`
	Uuid       string         `gorm:"size:50;not null" json:"uuid"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy  uint32         `json:"created_by"`
	UpdatedBy  uint32         `json:"updated_by"`
	DeletedBy  uint32         `json:"deleted_by"`
	Permission []*Permission  `gorm:"many2many:role_permissions;"`
}
type Roles []Role
