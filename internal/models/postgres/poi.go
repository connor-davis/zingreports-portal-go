package postgres

import (
	"github.com/connor-davis/zingreports-portal-go/internal/helpers"
	"gorm.io/gorm"
)

type Poi struct {
	Base
	Name string `json:"name" gorm:"type:text;not null;" validate:"required"`
	Key  string `json:"key" gorm:"type:text;uniqueIndex;not null;" validate:"required"`
}

func (p *Poi) BeforeCreate(tx *gorm.DB) error {
	err := helpers.Validate(p)

	if err != nil {
		return err
	}

	return nil
}
