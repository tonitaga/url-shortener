build:
	docker compose build url-shortener

launch:
	docker compose up url-shortener

migrate-up:
	migrate -path ./migrations -database 'postgres://postgres:password@localhost:5432/postgres?sslmode=disable' up

migrate-down:
	migrate -path ./migrations -database 'postgres://postgres:password@localhost:5432/postgres?sslmode=disable' down