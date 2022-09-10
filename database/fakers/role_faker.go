package fakers

import (
	"time"

	"github.com/google/uuid"

	"github.com/hasanmarzaq87/goscholling/api/models"
	"gorm.io/gorm"
)

func RoleFaker(db *gorm.DB) *models.Roles {
	return &models.Roles{
		{
			Name:      "Super Admin",
			Alias:     "superadmin",
			Uuid:      uuid.New().String(),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		{
			Name:      "Teacher",
			Alias:     "teacher",
			Uuid:      uuid.New().String(),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		{
			Name:      "Student",
			Alias:     "student",
			Uuid:      uuid.New().String(),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		{
			Name:      "Parent",
			Alias:     "parent",
			Uuid:      uuid.New().String(),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
	}

}
