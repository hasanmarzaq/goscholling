package fakers

import (
	"time"

	"github.com/google/uuid"
	"github.com/hasanmarzaq87/goscholling/api/models"
	"gorm.io/gorm"
)

func PermissionFaker(db *gorm.DB) *models.Permissions {
	return &models.Permissions{
		{
			Name:       "Super Admin",
			ActionName: "superadmin",
			Path:       "superadmin",
			FullPath:   "superadmin",
			Methods:    "superadmin",
			RoutePath:  "superadmin",
			Uuid:       uuid.New().String(),
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
			DeletedAt:  gorm.DeletedAt{},
		},
	}

}
