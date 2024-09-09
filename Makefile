APP_BINARY=urlShortnerApp
DATABASE_URL="postgresql://postgres:postgres@localhost:5432/urlshortner?sslmode=disable"
MIGRATION_PATH="./migrations/"

build:
	@echo "Building urlshortner..."
	CGO_ENABLED=0 go build -o _build/${APP_BINARY} main.go
	@echo "Done!"

run: build
	@echo "Starting urlshortner..."
	./_build/${APP_BINARY}

docker-build:
	@echo "Building docker image..."
	docker build -f Dockerfile -t alpeshjamgade/url-shortner:${TAG} .
	@echo "Done!!, Image: registry.tradelab.in/url-shortner:${TAG}"

migration_create:
	migrate create -ext sql -dir ${MIGRATION_PATH} -seq ${name}
migration_up:
	migrate -path ${MIGRATION_PATH} -database ${DATABASE_URL} -verbose up

migration_down:
	migrate -path ${MIGRATION_PATH} -database ${DATABASE_URL} -verbose down

migration_fix:
	migrate -path ${MIGRATION_PATH} -database ${DATABASE_URL} force VERSION