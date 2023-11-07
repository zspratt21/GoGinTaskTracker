.PHONY: build-containers start-containers stop-containers restart-containers run-migrations build-windows build-linux build-mac setup-envs tidy

setup-envs:
	@cp example.env .env
	@cp example.migration.env .migration.env

build-containers:
	@docker-compose build

build-containers-no-cache:
	@docker-compose build --no-cache

start-containers: build-containers
	@docker-compose up -d

stop-containers:
	@docker-compose down

restart-containers:
	@docker-compose restart

run-migrations: start-containers
	@go run ./cmd/migrations/main.go

run-tests:
	@go test ./src/...

tidy:
	@go mod tidy

build-windows:
	@GOOS=windows GOARCH=amd64 go build -o ./build/win/app-amd64-win.exe ./src
	@cp .env ./build/win/.env
	@mkdir -p ./build/win/templates
	@cp ./templates/welcome.html ./build/win/templates/welcome.html

build-linux:
	@GOOS=linux GOARCH=amd64 go build -o ./build/linux/app-amd64-linux ./src
	@cp .env ./build/linux/.env
	@mkdir -p ./build/linux/templates
	@cp ./templates/welcome.html ./build/linux/templates/welcome.html

build-mac:
	@GOOS=darwin GOARCH=amd64 go build -o ./build/mac/app-amd64-darwin ./src
	@cp .env ./build/mac/.env
	@mkdir -p ./build/mac/templates
	@cp ./templates/welcome.html ./build/mac/templates/welcome.html
