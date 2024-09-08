APP_BINARY=urlShortnerApp
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