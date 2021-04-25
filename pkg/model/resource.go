package model

import "time"

type Resource struct {
	ResourceId string    `gorm:"not null;primaryKey;check:has_account,((public.account(resource_id) IS NOT NULL))"`
	OwnerId    string    `gorm:"not null;check:has_kind,((public.kind(resource_id) IS NOT NULL))"`
	CreatedAt  time.Time `gorm:"type:timestamp without time zone;default:transaction_timestamp()"`
	PolicyID   string    `gorm:"check:verify_policy_kind,((public.kind(policy_id) = 'policy'::text))"`
}
