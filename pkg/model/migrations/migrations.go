package migrations

import (
	"conjur-in-go/pkg/model"
	"fmt"
	"sort"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func GetMigrations() []*gormigrate.Migration {
	migrations := GetLocalMigrations()
	fileMigrations := GetFileMigrations()
	migrations = append(migrations, fileMigrations...)
	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].ID < migrations[j].ID
	})

	fmt.Printf("Migrations to apply:\n")
	for _, v := range migrations {
		fmt.Printf(" - %s\n", v.ID)
	}
	return migrations
}

func GetLocalMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "20210424_001_base_v5_models",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(
					&model.Secret{},
					&model.Resource{},
					&model.Role{},
					&model.PolicyVersion{},
					&model.Annotation{},
					&model.AuthenticatorConfig{},
					&model.Credential{},
				)
			},
		},
	}
}
