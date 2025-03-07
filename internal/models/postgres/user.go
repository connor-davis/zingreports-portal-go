package postgres

type User struct {
	Base
	Image    string `json:"image" gorm:"type:text;"`
	Name     string `json:"name" gorm:"type:text;not null;"`
	Email    string `json:"email" gorm:"type:text;"`
	Password string `json:"password" gorm:"type:text;"`
	Role     string `json:"role" gorm:"type:text;"`
}
