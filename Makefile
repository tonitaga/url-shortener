build:
	docker compose build url-shortener

launch:
	docker compose up url-shortener

migrate-up:
	migrate -path ./migrations -database 'postgres://tonitaga:password@0.0.0.0:5432/url-shortener?sslmode=disable' up

migrate-down:
	migrate -path ./migrations -database 'postgres://tonitaga:password@0.0.0.0:5432/url-shortener?sslmode=disable' down