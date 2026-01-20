run: migrate-up
	go run ./cmd/server

test: test-unit test-intergration

test-unit:
	go test -v ./...

test-intergration:
	go test -v -tags=integration ./...

migrate-up:
	migrate -path migrations -database $${DB_URL} up
migrate-down:
	migrate -path migrations -database $${DB_URL} down