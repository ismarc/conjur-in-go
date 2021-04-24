package main

//go:generate go run ../../scripts/embed_migrations.go
func main() {
	Execute()
}
