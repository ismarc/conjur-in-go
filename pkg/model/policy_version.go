package model

import "time"

type PolicyVersion struct {
	ResourceID   string    `gorm:"not null;primaryKey"`
	RoleID       string    `gorm:"not null"`
	Version      int       `gorm:"type:integer; not null;primaryKey"`
	CreatedAt    time.Time `gorm:"not null;type:timestamp with time zone;default:transaction_timestamp()"`
	PolicyText   string    `gorm:"not null"`
	PolicySHA256 string    `gorm:"not null"`
	FinishedAt   time.Time `gorm:"check:created_before_finish,((created_at <= finished_at))"`
	ClientIP     string
}
