package model

type Annotation struct {
	ResourceID string `gorm:"not null;primaryKey"`
	Name       string `gorm:"not null;primaryKey"`
	Value      string `gorm:"not null"`
	PolicyID   string `gorm:"check:verify_policy_kind,((public.kind(policy_id) = 'policy'::text))"`
}
