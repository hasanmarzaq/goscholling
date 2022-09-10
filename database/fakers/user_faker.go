package fakers

import (
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"

	"github.com/hasanmarzaq87/goscholling/api/models"
	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) *models.User {

	return &models.User{
		// ID:       uuid.New().String(),
		Username: faker.Username(),
		// LastName:  faker.LastName(),
		Email:     faker.Email(),
		Password:  "$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi", // password
		Phone:     faker.Phonenumber(),
		Gender:    faker.Gender(),
		Religion:  "islam",
		StatusID:  1,
		RoleID:    1,
		IsActive:  1,
		Uuid:      uuid.New().String(),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
}
