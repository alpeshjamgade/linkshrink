APP_BINARY=shrinklinkApp
DATABASE_URL="postgresql://postgres:postgres@localhost:5432/shrinklink?sslmode=disable"
MIGRATION_PATH="./internal/repo/migrations/"

build:
	@echo "Building shrinklink..."
	CGO_ENABLED=0 go build -o _build/${APP_BINARY} ./cmd/main.go
	@echo "Done!"

run: build
	@echo "Starting shrinklink..."
	./_build/${APP_BINARY}

docker-build:
	@echo "Building docker image..."
	docker build -f Dockerfile -t alpeshjamgade/shrinklink:${TAG} .
	@echo "Done!!, Image: alpeshjamgade/shrinklink:${TAG}"

migration_create:
	migrate create -ext sql -dir ${MIGRATION_PATH} -seq ${name}

migration_up:
	migrate -path ${MIGRATION_PATH} -database ${DATABASE_URL} -verbose up

migration_down:
	migrate -path ${MIGRATION_PATH} -database ${DATABASE_URL} -verbose down

migration_fix:
	migrate -path ${MIGRATION_PATH} -database ${DATABASE_URL} force VERSION