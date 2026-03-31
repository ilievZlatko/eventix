seed:
	go run ./apps/api/cmd/seed/main.go

run:
	cd apps/api && air

migrate:
	migrate -path apps/api/migrations -database "postgres://postgres:Password123%21@localhost:5432/eventix?sslmode=disable" up
