build:
	docker compose build url-shortener

launch:
	docker compose up url-shortener

migrate-up:
	migrate -path ./migrations -database 'postgres://user:password@localhost:5432/dbname?sslmode=disable' up

migrate-down:
	migrate -path ./migrations -database 'postgres://user:password@localhost:5432/dbname?sslmode=disable' down