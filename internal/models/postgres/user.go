package postgres

type User struct {
	Base
	Image       string `json:"image" gorm:"type:text;"`
	Name        string `json:"name" gorm:"type:text;not null;" validate:"required"`
	Email       string `json:"email" gorm:"type:text;not null;" validate:"required"`
	Password    string `json:"-" gorm:"type:text;not null;" validate:"required"`
	Role        string `json:"role" gorm:"type:text;not null;" validate:"required"`
	MfaSecret   string `json:"-" gorm:"type:text"`
	MfaEnabled  bool   `json:"mfaEnabled" gorm:"default:false;not null;"`
	MfaVerified bool   `json:"mfaVerified" gorm:"default:false;not null;"`
}
