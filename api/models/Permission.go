package models

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	ID         uint32         `gorm:"primary_key;auto_increment" json:"id"`
	Name       string         `gorm:"size:255;not null;" json:"name"`
	ActionName string         `gorm:"size:255;not null;" json:"action_name"`
	Path       string         `gorm:"size:255;not null;" json:"path"`
	FullPath   string         `gorm:"size:255;not null;" json:"full_path"`
	Methods    string         `gorm:"size:255;not null;" json:"methods"`
	RoutePath  string         `gorm:"size:255;not null;" json:"route_path"`
	Uuid       string         `gorm:"size:50;not null" json:"uuid"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy  uint32         `json:"created_by"`
	UpdatedBy  uint32         `json:"updated_by"`
	DeletedBy  uint32         `json:"deleted_by"`
}
type Permissions []Permission
