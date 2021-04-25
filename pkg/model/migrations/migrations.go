package migrations

import (
	"sort"

	"github.com/go-gormigrate/gormigrate/v2"
)

func GetMigrations() []*gormigrate.Migration {
	migrations := []*gormigrate.Migration{}
	fileMigrations := GetFileMigrations()
	migrations = append(migrations, fileMigrations...)
	sort.Slice(migrations, func(i, j int) bool {
		return fileMigrations[i].ID < fileMigrations[j].ID
	})

	return migrations
}
