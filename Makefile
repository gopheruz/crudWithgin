include .env
migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up
migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down