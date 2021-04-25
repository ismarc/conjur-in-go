package model

import (
	"net"
	"time"
)

type Credential struct {
	RoleID        string `gorm:"not null;primaryKey"`
	ClientID      string
	ApiKey        []byte
	EncryptedHash []byte
	Expiration    time.Time   `gorm:"type:timestamp without time zone"`
	RestrictedTo  []net.IPNet `gorm:"type:cidr[];default:'{}'::cidr[];not null"`
}

// CREATE TABLE public.credentials (
// 	role_id text NOT NULL,
// 	client_id text,
// 	api_key bytea,
// 	encrypted_hash bytea,
// 	expiration timestamp without time zone,
// 	restricted_to cidr[] DEFAULT '{}'::cidr[] NOT NULL
// );
