package postgres

import (
	"log"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Base struct {
	Id        string    `json:"id" gorm:"type:text;primaryKey;"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	log.Printf("Base")

	if b.Id == "" {
		id, err := gonanoid.Generate("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789", 32)

		if err != nil {
			return err
		}

		b.Id = id
	}

	return nil
}
