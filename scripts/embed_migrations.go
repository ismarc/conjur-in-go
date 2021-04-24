package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type Entry struct {
	ID   string
	Up   string
	Down string
}

type MigrationData struct {
	Timestamp  time.Time
	Migrations []*Entry
}

func main() {
	var migrations map[string]*Entry
	migrations = make(map[string]*Entry)

	baseDir := filepath.Join("..", "..", "assets", "db", "migrations")
	fs, _ := ioutil.ReadDir(baseDir)
	for _, f := range fs {
		if strings.HasSuffix(f.Name(), ".sql") {
			base := strings.TrimSuffix(f.Name(), ".sql")
			fmt.Printf("Processing file: %s\n", filepath.Join("assets", "db", "migrations", f.Name()))
			if strings.HasSuffix(base, "_up") {
				key := strings.TrimSuffix(base, "_up")
				bytes, err := ioutil.ReadFile(filepath.Join(baseDir, f.Name()))
				if err != nil {
					panic(err)
				}
				if val, ok := migrations[key]; ok {
					val.Up = string(bytes)
				} else {
					migrations[key] = &Entry{ID: key, Up: string(bytes)}
				}
			} else if strings.HasSuffix(base, "_down") {
				key := strings.TrimSuffix(base, "_down")
				bytes, err := ioutil.ReadFile(filepath.Join(baseDir, f.Name()))
				if err != nil {
					panic(err)
				}
				if val, ok := migrations[key]; ok {
					val.Down = string(bytes)
				} else {
					migrations[key] = &Entry{ID: key, Down: string(bytes)}
				}
			}
		}
	}

	migrationData := MigrationData{
		Timestamp: time.Now(),
	}

	for _, v := range migrations {
		migrationData.Migrations = append(migrationData.Migrations, v)
	}

	migrationTemplate, err := template.ParseFiles(filepath.Join("..", "..", "assets", "db", "migrations", "migration.go.tmpl"))
	if err != nil {
		panic(err)
	}

	out, err := os.Create(filepath.Join("..", "..", "pkg", "model", "migrations", "file_migrations.go"))
	if err != nil {
		panic(err)
	}

	err = migrationTemplate.Execute(out, migrationData)
	if err != nil {
		panic(err)
	}
}
