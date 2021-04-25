package model

import "time"

type Role struct {
	RoleID    string    `gorm:"not null;primaryKey;check:has_account,((public.account(role_id) IS NOT NULL))"`
	CreatedAt time.Time `gorm:"not null;type:timestamp without time zone;default:transaction_timestamp();check:has_kind,((public.kind(role_id) IS NOT NULL))"`
	PolicyID  string    `gorm:"check:verify_policy_kind,((public.kind(policy_id) = 'policy'::text))"`
}
