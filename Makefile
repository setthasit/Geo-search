up:
	docker compose up --build -d

down:
	docker compose down -v

migrate:
	docker compose -f docker-compose.yaml run --rm geo-search-migrator migrate:elasticsearch

seed: migrate
	docker compose -f docker-compose.yaml run --rm geo-search-migrator seed
