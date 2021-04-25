package model

type AuthenticatorConfig struct {
	ID         int    `gorm:"type:serial;primaryKey;not null"`
	ResourceID string `gorm:"not null;unique"`
	Enabled    bool   `gorm:"not null;default:false"`
}
