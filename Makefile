build:
	@echo "Building urlshortner..."
	CGO_ENABLED=0 go build -o _build/app.build main.go
	@echo "Done!"

run: build
	./_build/app.build