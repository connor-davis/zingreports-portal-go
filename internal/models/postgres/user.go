package postgres

import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type User struct {
	Id        string    `json:"id" gorm:"type:text;primaryKey;"`
	Image     string    `json:"image" gorm:"type:text;"`
	Name      string    `json:"name" gorm:"type:text;not null;"`
	Email     string    `json:"email" gorm:"type:text;"`
	Password  string    `json:"password" gorm:"type:text;"`
	Role      string    `json:"role" gorm:"type:text;"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	id, err := gonanoid.New()

	if err != nil {
		return err
	}

	u.Id = id

	return nil
}
