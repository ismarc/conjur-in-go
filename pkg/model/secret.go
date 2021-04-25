package model

import (
	"time"
)

type Secret struct {
	ResourceId string    `gorm:"primaryKey;not null"`
	Version    int       `gorm:"primaryKey;type:integer;not null"`
	Value      []byte    `gorm:"not null"`
	ExpiresAt  time.Time `gorm:"type:timestamp without time zone"`
}

// Unneeded for normal table name format
// func (s Secret) TableName() string {
// 	return "secrets"
// }
