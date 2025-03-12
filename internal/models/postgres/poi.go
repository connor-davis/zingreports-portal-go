package postgres

type Poi struct {
	Base
	Name string `json:"name" gorm:"type:text;not null;" validate:"required"`
	Key  string `json:"key" gorm:"type:text;uniqueIndex;not null;" validate:"required"`
}
