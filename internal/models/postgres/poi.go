package postgres

type Poi struct {
	Base
	Name string `json:"name" gorm:"type:text;"`
	Key  string `json:"key" gorm:"type:text;uniqueIndex;"`
}
